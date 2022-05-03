// THIS FILE IS GENERATED CODE, MANUAL EDITS WILL BE OVERWRITTEN

package bitcoindclient

import (
	"context"
	"encoding/json"
)

// EnumerateSignersResp holds the response to the EnumerateSigners call.
//  {                               (json object)
//    "signers" : [                 (json array)
//      {                           (json object)
//        "fingerprint" : "hex",    (string) Master key fingerprint
//        "name" : "str"            (string) Device name
//      },
//      ...
//    ]
//  }
type EnumerateSignersResp struct {
	Signers []EnumerateSignersRespSigners `json:"signers"`
}

type EnumerateSignersRespSigners struct {
	// Master key fingerprint
	Fingerprint string `json:"fingerprint"`

	// Device name
	Name string `json:"name"`
}

// EnumerateSigners RPC method.
// Returns a list of external signers from -signer.
func (bc *BitcoindClient) EnumerateSigners(ctx context.Context) (result EnumerateSignersResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "enumeratesigners", nil); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}
