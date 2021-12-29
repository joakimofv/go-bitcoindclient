package bitcoindclient

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type readyContext struct {
	versionCheck         func(bitcoindVersion int) bool
	dontUseRpc           bool
	subscribeEventWithin *time.Duration
}

// ReadyOption is a modifier of the Ready function.
type ReadyOption func(*readyContext)

// WithVersionCheck makes Ready check that the bitcoind major version (from "getnetworkinfo") matches that of this package.
func WithVersionCheck() ReadyOption {
	return func(rctx *readyContext) {
		rctx.versionCheck = func(bitcoindVersion int) bool { return bitcoindVersion/10000 == MAJOR_VERSION }
	}
}

// WithVersionCheckFn makes Ready check that the bitcoind version (from "getnetworkinfo") passes the supplied checkFn.
// bitcoindVersion is a number with the format MMmmpp (M=Major, m=minor, p=patch).
// E.g. 220000 for v22.0.0.
func WithVersionCheckFn(checkFn func(bitcoindVersion int) bool) ReadyOption {
	return func(rctx *readyContext) {
		rctx.versionCheck = checkFn
	}
}

// WithoutRpc makes Ready not do an RPC call. Can be used in conjunction with WithZmqMessageWithin
// to only check the ZMQ connection.
func WithoutRpc() ReadyOption {
	return func(rctx *readyContext) {
		rctx.dontUseRpc = true
	}
}

// WithZmqMessageWithin makes Ready check if there has arrived any subscribe event (ZMQ message) in the preceeding period of time.
//
// ZMQ messages will only arrive if there are active subscriptions.
// To confirm that the connection is healthy you can for example do SubscribeSequence, then do Ready(WithZmqMessageWithin(time.Hour))
// once per second until successful.
//
// There is no way to detect a ZMQ connection error (other than the absence of success), this is an inherent drawback of the ZMQ protocol.
func WithZmqMessageWithin(period time.Duration) ReadyOption {
	return func(rctx *readyContext) {
		rctx.subscribeEventWithin = &period
	}
}

// Ready checks the connection health. By default it does a RPC call ("getnetworkinfo") and returns a nil error
// if there was a sane response.
//
// Does not retry on connection failure, in contrast to the RPC methods that retry until their context is cancelled.
//
// What is checked on can be changed by giving ReadyOption(s) as parameters.
func (bc *BitcoindClient) Ready(opts ...ReadyOption) (err error) {
	rctx := readyContext{}
	for _, opt := range opts {
		opt(&rctx)
	}

	// RPC
	if !rctx.dontUseRpc {
		// Modified bc.GetNetworkInfo to skip retries.
		var result GetNetworkInfoResp
		var resultRaw json.RawMessage
		if resultRaw, err = bc.sendRequest(context.Background(), "getnetworkinfo", nil, true); err != nil {
			return
		}
		if err = json.Unmarshal(resultRaw, &result); err != nil {
			return
		}

		if rctx.versionCheck != nil {
			if !rctx.versionCheck(int(result.Version)) {
				err = fmt.Errorf("Bad bitcoind version: %v (bitcoindclient MAJOR_VERSION: %v)", int(result.Version), MAJOR_VERSION)
				return
			}
		}
	}

	// ZMQ
	if rctx.subscribeEventWithin != nil {
		if time.Since(bc.subs.latestEvent) > *rctx.subscribeEventWithin {
			err = fmt.Errorf("Latest subscribe event at %v not within %v of the present.", bc.subs.latestEvent, *rctx.subscribeEventWithin)
			return
		}
	}

	return
}
