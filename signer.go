// THIS FILE IS GENERATED CODE, MANUAL EDITS WILL BE OVERWRITTEN

package bitcoindclient

import (
	"context"
	"encoding/json"
)

// EnumerateSignersResp holds the response to the EnumerateSigners call.
//  {                  (json object)
//    "signers" : [    (json array)
//      "hex",         (string) Master key fingerprint
//      "str",         (string) Device name
//      ...
//    ]
//  }
type EnumerateSignersResp struct {
	// Element: hex    Master key fingerprint
	// Element: str    Device name
	Signers []string `json:"signers"`
}

// EnumerateSigners RPC method.
// Returns a list of external signers from -signer.
func (bc *BitcoindClient) EnumerateSigners(ctx context.Context) (result EnumerateSignersResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "enumeratesigners", nil, false); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}
