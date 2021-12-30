package bitcoindclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"time"
)

type reqWrapper struct {
	Method string          `json:"method"`
	Params json.RawMessage `json:"params"`
	Id     float64         `json:"id"`
}

type respWrapper struct {
	Result json.RawMessage `json:"result"`
	Error  *BitcoindError  `json:"error"`
	Id     float64         `json:"id"`
}

// BitcoindError represents an error that originates from bitcoind.
type BitcoindError struct {
	JsonRPCError
}

// JsonRPCError is an error on the format of the JSON-RPC 2.0 standard.
type JsonRPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	//Data  interface{} `json:"data"` // Omitted
}

func (err *JsonRPCError) Error() string {
	return fmt.Sprintf("(error code %d) %s", err.Code, err.Message)
}

// UnmarshalError indicates a problem with unmarshaling JSON data.
// The data is provided in the member B, can be used for debugging.
type UnmarshalError struct {
	B          []byte
	structName string
}

func (err *UnmarshalError) Error() string {
	return fmt.Sprintf("Could not unmarshal %q into any member of %s", err.B, err.structName)
}

type uriPathKey struct{}
type retriesKey struct{}

// UseUriPath changes the RpcUriPath for a single request, when applied to that requests context.
func UseUriPath(ctx context.Context, newPath string) context.Context {
	return context.WithValue(ctx, uriPathKey{}, newPath)
}

// UseConnectionRetries enables retries on connection failure, as many as the given number.
// Negative number for infinite retries, 0 for no retries (default).
//
// If the context is canceled or expired the latest connection error will be returned instead of context.Canceled/context.DeadlineExceeded.
func UseConnectionRetries(ctx context.Context, retries int) context.Context {
	return context.WithValue(ctx, retriesKey{}, retries)
}

func (bc *BitcoindClient) sendRequest(ctx context.Context, method string, paramsStruct interface{}) (result json.RawMessage, err error) {
	if bc.Cfg.RpcAddress == "" {
		err = ErrRpcDisabled
		return
	}
	// Create the request.
	params := []byte("[]")
	if paramsStruct != nil {
		paramsStruct = allocateNilObjects(paramsStruct)
		params, err = json.Marshal(paramsStruct)
		if err != nil {
			return
		}
	}
	var body []byte
	body, err = json.Marshal(reqWrapper{
		Method: method,
		Params: json.RawMessage(params),
		Id:     1,
	})
	if err != nil {
		return
	}

	var resp *http.Response
	var nonCtxErr error
	buf := new(bytes.Buffer)
	backoff := 250 * time.Millisecond // The time it takes to start up bitcoind in regtest mode.
	retries := 0
	if r, ok := ctx.Value(retriesKey{}).(int); ok {
		retries = r
	}
	for i := 0; true; i++ {
		buf.Reset()
		buf.Write(body)
		var req *http.Request
		uriPath := bc.Cfg.RpcUriPath
		if u, ok := ctx.Value(uriPathKey{}).(string); ok {
			uriPath = u
		}
		req, err = http.NewRequestWithContext(ctx, "POST", "http://"+bc.Cfg.RpcAddress+uriPath, buf)
		if err != nil {
			return
		}
		req.Header["Content-Type"] = []string{"application/json"}
		req.Header["Connection"] = []string{"close"}
		req.SetBasicAuth(bc.Cfg.RpcUser, bc.Cfg.RpcPassword)

		// Send the request, get a response.
		resp, err = bc.httpClient.Do(req)
		if err != nil {
			if ctx.Err() != nil {
				break
			}
			nonCtxErr = err
			if i == retries {
				break
			}
			select {
			case <-time.After(backoff):
			case <-ctx.Done():
				break
			}
			backoff = time.Second
			continue
		}
		defer resp.Body.Close()
		break
	}
	if err != nil {
		if nonCtxErr != nil {
			err = nonCtxErr
		}
		return
	}

	// Extract the result from the response.
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var respWrapper respWrapper
	err = json.Unmarshal(body, &respWrapper)
	if err != nil {
		err = fmt.Errorf("Unexpected response format: %s", body)
		return
	}
	if respWrapper.Error != nil {
		err = respWrapper.Error
		return
	}
	result = respWrapper.Result
	return
}

// allocateNilObjects makes empty slices and maps be represented by
// [] and {} instead of null when marshalled into JSON, because that
// is what bitcoind expects.
func allocateNilObjects(paramsStruct interface{}) interface{} {
	v := reflect.ValueOf(paramsStruct)
	// v isn't settable, make a copy that is.
	vCopy := reflect.New(v.Type()).Elem()
	vCopy.Set(v)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Kind() == reflect.Map && v.Field(i).IsNil() {
			vCopy.Field(i).Set(reflect.MakeMap(v.Field(i).Type()))
		}
		if v.Field(i).Kind() == reflect.Slice && v.Field(i).IsNil() {
			vCopy.Field(i).Set(reflect.MakeSlice(v.Field(i).Type(), 0, 0))
		}
	}
	return vCopy.Interface()
}
