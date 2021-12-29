// THIS FILE IS GENERATED CODE, MANUAL EDITS WILL BE OVERWRITTEN

package bitcoindclient

import (
	"bytes"
	"context"
	"encoding/json"
)

// GenerateBlockReq holds the arguments for the GenerateBlock call.
//  1. output               (string, required) The address or descriptor to send the newly generated bitcoin to.
//  2. transactions         (json array, required) An array of hex strings which are either txids or raw transactions.
//                          Txids must reference transactions currently in the mempool.
//                          All transactions must be valid and in valid order, otherwise the block will be rejected.
//       [
//         "rawtx/txid",    (string)
//         ...
//       ]
type GenerateBlockReq struct {
	// The address or descriptor to send the newly generated bitcoin to.
	Output string `json:"output"`

	// An array of hex strings which are either txids or raw transactions.
	// Txids must reference transactions currently in the mempool.
	// All transactions must be valid and in valid order, otherwise the block will be rejected.
	// Element: RawTxOrTxID
	Transactions []string `json:"transactions"`
}

// GenerateBlockResp holds the response to the GenerateBlock call.
//  {                    (json object)
//    "hash" : "hex"     (string) hash of generated block
//  }
type GenerateBlockResp struct {
	// hash of generated block
	Hash string `json:"hash"`
}

// GenerateBlock RPC method.
// Mine a block with a set of ordered transactions immediately to a specified address or descriptor (before the RPC call returns)
func (bc *BitcoindClient) GenerateBlock(ctx context.Context, args GenerateBlockReq) (result GenerateBlockResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "generateblock", args, false); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// GenerateToAddressReq holds the arguments for the GenerateToAddress call.
//  1. nblocks     (numeric, required) How many blocks are generated immediately.
//  2. address     (string, required) The address to send the newly generated bitcoin to.
//  3. maxtries    (numeric, optional, default=1000000) How many iterations to try.
type GenerateToAddressReq struct {
	// How many blocks are generated immediately.
	NBlocks float64 `json:"nblocks"`

	// The address to send the newly generated bitcoin to.
	Address string `json:"address"`

	// How many iterations to try.
	// Default: 1000000
	MaxTries *float64 `json:"maxtries,omitempty"`
}

// GenerateToAddressResp holds the response to the GenerateToAddress call.
//  [           (json array) hashes of blocks generated
//    "hex",    (string) blockhash
//    ...
//  ]
type GenerateToAddressResp struct {
	// hashes of blocks generated
	// Element: Hex    blockhash
	Hex []string
}

func (alts GenerateToAddressResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.Hex)
}

func (alts *GenerateToAddressResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Hex) == nil {
		return nil
	}
	alts.Hex = reset.Hex
	return &UnmarshalError{B: b, structName: "GenerateToAddressResp"}
}

// GenerateToAddress RPC method.
// Mine blocks immediately to a specified address (before the RPC call returns)
func (bc *BitcoindClient) GenerateToAddress(ctx context.Context, args GenerateToAddressReq) (result GenerateToAddressResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "generatetoaddress", args, false); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// GenerateToDescriptorReq holds the arguments for the GenerateToDescriptor call.
//  1. num_blocks    (numeric, required) How many blocks are generated immediately.
//  2. descriptor    (string, required) The descriptor to send the newly generated bitcoin to.
//  3. maxtries      (numeric, optional, default=1000000) How many iterations to try.
type GenerateToDescriptorReq struct {
	// How many blocks are generated immediately.
	NumBlocks float64 `json:"num_blocks"`

	// The descriptor to send the newly generated bitcoin to.
	Descriptor string `json:"descriptor"`

	// How many iterations to try.
	// Default: 1000000
	MaxTries *float64 `json:"maxtries,omitempty"`
}

// GenerateToDescriptorResp holds the response to the GenerateToDescriptor call.
//  [           (json array) hashes of blocks generated
//    "hex",    (string) blockhash
//    ...
//  ]
type GenerateToDescriptorResp struct {
	// hashes of blocks generated
	// Element: Hex    blockhash
	Hex []string
}

func (alts GenerateToDescriptorResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.Hex)
}

func (alts *GenerateToDescriptorResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Hex) == nil {
		return nil
	}
	alts.Hex = reset.Hex
	return &UnmarshalError{B: b, structName: "GenerateToDescriptorResp"}
}

// GenerateToDescriptor RPC method.
// Mine blocks immediately to a specified descriptor (before the RPC call returns)
func (bc *BitcoindClient) GenerateToDescriptor(ctx context.Context, args GenerateToDescriptorReq) (result GenerateToDescriptorResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "generatetodescriptor", args, false); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}
