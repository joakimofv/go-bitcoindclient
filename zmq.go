// THIS FILE IS GENERATED CODE, MANUAL EDITS WILL BE OVERWRITTEN

package bitcoindclient

import (
	"bytes"
	"context"
	"encoding/json"
)

// GetZmqNotificationsResp holds the response to the GetZmqNotifications call.
//  [                         (json array)
//    {                       (json object)
//      "type" : "str",       (string) Type of notification
//      "address" : "str",    (string) Address of the publisher
//      "hwm" : n             (numeric) Outbound message high water mark
//    },
//    ...
//  ]
type GetZmqNotificationsResp struct {
	Array []GetZmqNotificationsRespElement
}

func (alts GetZmqNotificationsResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.Array)
}

func (alts *GetZmqNotificationsResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Array) == nil {
		return nil
	}
	alts.Array = reset.Array
	return &UnmarshalError{B: b, structName: "GetZmqNotificationsResp"}
}

type GetZmqNotificationsRespElement struct {
	// Type of notification
	Type string `json:"type"`

	// Address of the publisher
	Address string `json:"address"`

	// Outbound message high water mark
	Hwm float64 `json:"hwm"`
}

// GetZmqNotifications RPC method.
// Returns information about the active ZeroMQ notifications.
func (bc *BitcoindClient) GetZmqNotifications(ctx context.Context) (result GetZmqNotificationsResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "getzmqnotifications", nil, false); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}
