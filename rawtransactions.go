// THIS FILE IS GENERATED CODE, MANUAL EDITS WILL BE OVERWRITTEN

package bitcoindclient

import (
	"bytes"
	"context"
	"encoding/json"
	"reflect"
)

// AnalyzePsbtReq holds the arguments for the AnalyzePsbt call.
//  1. psbt    (string, required) A base64 string of a PSBT
type AnalyzePsbtReq struct {
	// A base64 string of a PSBT
	Psbt string `json:"psbt"`
}

// AnalyzePsbtResp holds the response to the AnalyzePsbt call.
//  {                                   (json object)
//    "inputs" : [                      (json array)
//      {                               (json object)
//        "has_utxo" : true|false,      (boolean) Whether a UTXO is provided
//        "is_final" : true|false,      (boolean) Whether the input is finalized
//        "missing" : {                 (json object, optional) Things that are missing that are required to complete this input
//          "pubkeys" : [               (json array, optional)
//            "hex",                    (string) Public key ID, hash160 of the public key, of a public key whose BIP 32 derivation path is missing
//            ...
//          ],
//          "signatures" : [            (json array, optional)
//            "hex",                    (string) Public key ID, hash160 of the public key, of a public key whose signature is missing
//            ...
//          ],
//          "redeemscript" : "hex",     (string, optional) Hash160 of the redeemScript that is missing
//          "witnessscript" : "hex"     (string, optional) SHA256 of the witnessScript that is missing
//        },
//        "next" : "str"                (string, optional) Role of the next person that this input needs to go to
//      },
//      ...
//    ],
//    "estimated_vsize" : n,            (numeric, optional) Estimated vsize of the final signed transaction
//    "estimated_feerate" : n,          (numeric, optional) Estimated feerate of the final signed transaction in BTC/kvB. Shown only if all UTXO slots in the PSBT have been filled
//    "fee" : n,                        (numeric, optional) The transaction fee paid. Shown only if all UTXO slots in the PSBT have been filled
//    "next" : "str",                   (string) Role of the next person that this psbt needs to go to
//    "error" : "str"                   (string, optional) Error message (if there is one)
//  }
type AnalyzePsbtResp struct {
	Inputs []AnalyzePsbtRespInputs `json:"inputs"`

	// Estimated vsize of the final signed transaction
	EstimatedVSize *float64 `json:"estimated_vsize,omitempty"`

	// Estimated feerate of the final signed transaction in BTC/kvB. Shown only if all UTXO slots in the PSBT have been filled
	EstimatedFeeRate *float64 `json:"estimated_feerate,omitempty"`

	// The transaction fee paid. Shown only if all UTXO slots in the PSBT have been filled
	Fee *float64 `json:"fee,omitempty"`

	// Role of the next person that this psbt needs to go to
	Next string `json:"next"`

	// Error message (if there is one)
	Error string `json:"error,omitempty"`
}

type AnalyzePsbtRespInputs struct {
	// Whether a UTXO is provided
	HasUtxo bool `json:"has_utxo"`

	// Whether the input is finalized
	IsFinal bool `json:"is_final"`

	// Things that are missing that are required to complete this input
	Missing *AnalyzePsbtRespInputsMissing `json:"missing,omitempty"`

	// Role of the next person that this input needs to go to
	Next string `json:"next,omitempty"`
}

type AnalyzePsbtRespInputsMissing struct {
	// Element: Hex    Public key ID, hash160 of the public key, of a public key whose BIP 32 derivation path is missing
	Pubkeys []string `json:"pubkeys,omitempty"`

	// Element: Hex    Public key ID, hash160 of the public key, of a public key whose signature is missing
	Signatures []string `json:"signatures,omitempty"`

	// Hash160 of the redeemScript that is missing
	RedeemScript string `json:"redeemscript,omitempty"`

	// SHA256 of the witnessScript that is missing
	WitnessScript string `json:"witnessscript,omitempty"`
}

// AnalyzePsbt RPC method.
// Analyzes and provides information about the current status of a PSBT and its inputs
func (bc *BitcoindClient) AnalyzePsbt(ctx context.Context, args AnalyzePsbtReq) (result AnalyzePsbtResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "analyzepsbt", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// CombinePsbtReq holds the arguments for the CombinePsbt call.
//  1. txs            (json array, required) The base64 strings of partially signed transactions
//       [
//         "psbt",    (string) A base64 string of a PSBT
//         ...
//       ]
type CombinePsbtReq struct {
	// The base64 strings of partially signed transactions
	// Element: Psbt    A base64 string of a PSBT
	Txs []string `json:"txs"`
}

// CombinePsbtResp holds the response to the CombinePsbt call.
//  "str"    (string) The base64-encoded partially signed transaction
type CombinePsbtResp struct {
	// The base64-encoded partially signed transaction
	Str string
}

func (alts CombinePsbtResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.Str)
}

func (alts *CombinePsbtResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Str) == nil {
		return nil
	}
	alts.Str = reset.Str
	return &UnmarshalError{B: b, structName: "CombinePsbtResp"}
}

// CombinePsbt RPC method.
// Combine multiple partially signed Bitcoin transactions into one transaction.
// Implements the Combiner role.
func (bc *BitcoindClient) CombinePsbt(ctx context.Context, args CombinePsbtReq) (result CombinePsbtResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "combinepsbt", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// CombineRawTransactionReq holds the arguments for the CombineRawTransaction call.
//  1. txs                 (json array, required) The hex strings of partially signed transactions
//       [
//         "hexstring",    (string) A hex-encoded raw transaction
//         ...
//       ]
type CombineRawTransactionReq struct {
	// The hex strings of partially signed transactions
	// Element: HexString    A hex-encoded raw transaction
	Txs []string `json:"txs"`
}

// CombineRawTransactionResp holds the response to the CombineRawTransaction call.
//  "str"    (string) The hex-encoded raw transaction with signature(s)
type CombineRawTransactionResp struct {
	// The hex-encoded raw transaction with signature(s)
	Str string
}

func (alts CombineRawTransactionResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.Str)
}

func (alts *CombineRawTransactionResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Str) == nil {
		return nil
	}
	alts.Str = reset.Str
	return &UnmarshalError{B: b, structName: "CombineRawTransactionResp"}
}

// CombineRawTransaction RPC method.
// Combine multiple partially signed transactions into one transaction.
// The combined transaction may be another partially signed transaction or a
// fully signed transaction.
func (bc *BitcoindClient) CombineRawTransaction(ctx context.Context, args CombineRawTransactionReq) (result CombineRawTransactionResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "combinerawtransaction", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// ConvertToPsbtReq holds the arguments for the ConvertToPsbt call.
//  1. hexstring        (string, required) The hex string of a raw transaction
//  2. permitsigdata    (boolean, optional, default=false) If true, any signatures in the input will be discarded and conversion
//                      will continue. If false, RPC will fail if any signatures are present.
//  3. iswitness        (boolean, optional, default=depends on heuristic tests) Whether the transaction hex is a serialized witness transaction.
//                      If iswitness is not present, heuristic tests will be used in decoding.
//                      If true, only witness deserialization will be tried.
//                      If false, only non-witness deserialization will be tried.
//                      This boolean should reflect whether the transaction has inputs
//                      (e.g. fully valid, or on-chain transactions), if known by the caller.
type ConvertToPsbtReq struct {
	// The hex string of a raw transaction
	HexString string `json:"hexstring"`

	// If true, any signatures in the input will be discarded and conversion
	// will continue. If false, RPC will fail if any signatures are present.
	// Default: false
	PermitSigData bool `json:"permitsigdata,omitempty"`

	// Whether the transaction hex is a serialized witness transaction.
	// If iswitness is not present, heuristic tests will be used in decoding.
	// If true, only witness deserialization will be tried.
	// If false, only non-witness deserialization will be tried.
	// This boolean should reflect whether the transaction has inputs
	// (e.g. fully valid, or on-chain transactions), if known by the caller.
	// Default: depends on heuristic tests
	IsWitness *bool `json:"iswitness,omitempty"`
}

// ConvertToPsbtResp holds the response to the ConvertToPsbt call.
//  "str"    (string) The resulting raw transaction (base64-encoded string)
type ConvertToPsbtResp struct {
	// The resulting raw transaction (base64-encoded string)
	Str string
}

func (alts ConvertToPsbtResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.Str)
}

func (alts *ConvertToPsbtResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Str) == nil {
		return nil
	}
	alts.Str = reset.Str
	return &UnmarshalError{B: b, structName: "ConvertToPsbtResp"}
}

// ConvertToPsbt RPC method.
// Converts a network serialized transaction to a PSBT. This should be used only with createrawtransaction and fundrawtransaction
// createpsbt and walletcreatefundedpsbt should be used for new applications.
func (bc *BitcoindClient) ConvertToPsbt(ctx context.Context, args ConvertToPsbtReq) (result ConvertToPsbtResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "converttopsbt", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// CreatePsbtReq holds the arguments for the CreatePsbt call.
//  1. inputs                      (json array, required) The json objects
//       [
//         {                       (json object)
//           "txid": "hex",        (string, required) The transaction id
//           "vout": n,            (numeric, required) The output number
//           "sequence": n,        (numeric, optional, default=depends on the value of the 'replaceable' and 'locktime' arguments) The sequence number
//         },
//         ...
//       ]
//  2. outputs                     (json array, required) The outputs (key-value pairs), where none of the keys are duplicated.
//                                 That is, each address can only appear once and there can only be one 'data' object.
//                                 For compatibility reasons, a dictionary, which holds the key-value pairs directly, is also
//                                 accepted as second parameter.
//       [
//         {                       (json object)
//           "address": amount,    (numeric or string, required) A key-value pair. The key (string) is the bitcoin address, the value (float or string) is the amount in BTC
//           ...
//         },
//         {                       (json object)
//           "data": "hex",        (string, required) A key-value pair. The key must be "data", the value is hex-encoded data
//         },
//         ...
//       ]
//  3. locktime                    (numeric, optional, default=0) Raw locktime. Non-0 value also locktime-activates inputs
//  4. replaceable                 (boolean, optional, default=false) Marks this transaction as BIP125 replaceable.
//                                 Allows this transaction to be replaced by a transaction with higher fees. If provided, it is an error if explicit sequence numbers are incompatible.
type CreatePsbtReq struct {
	// The json objects
	Inputs []CreatePsbtReqInputs `json:"inputs"`

	// The outputs (key-value pairs), where none of the keys are duplicated.
	// That is, each address can only appear once and there can only be one 'data' object.
	// For compatibility reasons, a dictionary, which holds the key-value pairs directly, is also
	// accepted as second parameter.
	Outputs []CreatePsbtReqOutputs `json:"outputs"`

	// Raw locktime. Non-0 value also locktime-activates inputs
	// Default: 0
	LockTime float64 `json:"locktime,omitempty"`

	// Marks this transaction as BIP125 replaceable.
	// Allows this transaction to be replaced by a transaction with higher fees. If provided, it is an error if explicit sequence numbers are incompatible.
	// Default: false
	Replaceable bool `json:"replaceable,omitempty"`
}

type CreatePsbtReqInputs struct {
	// The transaction id
	TxID string `json:"txid"`

	// The output number
	Vout float64 `json:"vout"`

	// The sequence number
	// Default: depends on the value of the 'replaceable' and 'locktime' arguments
	Sequence *float64 `json:"sequence,omitempty"`
}

// Holder of alternative parameter formats, only one will be used, the first that is non-zero.
type CreatePsbtReqOutputs struct {
	// A key-value pair. The key (string) is the bitcoin address, the value (float or string) is the amount in BTC
	A map[string]float64

	B struct {
		// A key-value pair. The key must be "data", the value is hex-encoded data
		Data string `json:"data"`
	}
}

func (alts CreatePsbtReqOutputs) MarshalJSON() ([]byte, error) {
	if !reflect.ValueOf(alts.A).IsZero() {
		return json.Marshal(alts.A)
	}
	return json.Marshal(alts.B)
}

func (alts *CreatePsbtReqOutputs) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.A) == nil {
		return nil
	}
	alts.A = reset.A
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.B) == nil {
		return nil
	}
	alts.B = reset.B
	return &UnmarshalError{B: b, structName: "CreatePsbtReqOutputs"}
}

// CreatePsbtResp holds the response to the CreatePsbt call.
//  "str"    (string) The resulting raw transaction (base64-encoded string)
type CreatePsbtResp struct {
	// The resulting raw transaction (base64-encoded string)
	Str string
}

func (alts CreatePsbtResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.Str)
}

func (alts *CreatePsbtResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Str) == nil {
		return nil
	}
	alts.Str = reset.Str
	return &UnmarshalError{B: b, structName: "CreatePsbtResp"}
}

// CreatePsbt RPC method.
// Creates a transaction in the Partially Signed Transaction format.
// Implements the Creator role.
func (bc *BitcoindClient) CreatePsbt(ctx context.Context, args CreatePsbtReq) (result CreatePsbtResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "createpsbt", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// CreateRawTransactionReq holds the arguments for the CreateRawTransaction call.
//  1. inputs                      (json array, required) The inputs
//       [
//         {                       (json object)
//           "txid": "hex",        (string, required) The transaction id
//           "vout": n,            (numeric, required) The output number
//           "sequence": n,        (numeric, optional, default=depends on the value of the 'replaceable' and 'locktime' arguments) The sequence number
//         },
//         ...
//       ]
//  2. outputs                     (json array, required) The outputs (key-value pairs), where none of the keys are duplicated.
//                                 That is, each address can only appear once and there can only be one 'data' object.
//                                 For compatibility reasons, a dictionary, which holds the key-value pairs directly, is also
//                                 accepted as second parameter.
//       [
//         {                       (json object)
//           "address": amount,    (numeric or string, required) A key-value pair. The key (string) is the bitcoin address, the value (float or string) is the amount in BTC
//           ...
//         },
//         {                       (json object)
//           "data": "hex",        (string, required) A key-value pair. The key must be "data", the value is hex-encoded data
//         },
//         ...
//       ]
//  3. locktime                    (numeric, optional, default=0) Raw locktime. Non-0 value also locktime-activates inputs
//  4. replaceable                 (boolean, optional, default=false) Marks this transaction as BIP125-replaceable.
//                                 Allows this transaction to be replaced by a transaction with higher fees. If provided, it is an error if explicit sequence numbers are incompatible.
type CreateRawTransactionReq struct {
	// The inputs
	Inputs []CreateRawTransactionReqInputs `json:"inputs"`

	// The outputs (key-value pairs), where none of the keys are duplicated.
	// That is, each address can only appear once and there can only be one 'data' object.
	// For compatibility reasons, a dictionary, which holds the key-value pairs directly, is also
	// accepted as second parameter.
	Outputs []CreateRawTransactionReqOutputs `json:"outputs"`

	// Raw locktime. Non-0 value also locktime-activates inputs
	// Default: 0
	LockTime float64 `json:"locktime,omitempty"`

	// Marks this transaction as BIP125-replaceable.
	// Allows this transaction to be replaced by a transaction with higher fees. If provided, it is an error if explicit sequence numbers are incompatible.
	// Default: false
	Replaceable bool `json:"replaceable,omitempty"`
}

type CreateRawTransactionReqInputs struct {
	// The transaction id
	TxID string `json:"txid"`

	// The output number
	Vout float64 `json:"vout"`

	// The sequence number
	// Default: depends on the value of the 'replaceable' and 'locktime' arguments
	Sequence *float64 `json:"sequence,omitempty"`
}

// Holder of alternative parameter formats, only one will be used, the first that is non-zero.
type CreateRawTransactionReqOutputs struct {
	// A key-value pair. The key (string) is the bitcoin address, the value (float or string) is the amount in BTC
	A map[string]float64

	B struct {
		// A key-value pair. The key must be "data", the value is hex-encoded data
		Data string `json:"data"`
	}
}

func (alts CreateRawTransactionReqOutputs) MarshalJSON() ([]byte, error) {
	if !reflect.ValueOf(alts.A).IsZero() {
		return json.Marshal(alts.A)
	}
	return json.Marshal(alts.B)
}

func (alts *CreateRawTransactionReqOutputs) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.A) == nil {
		return nil
	}
	alts.A = reset.A
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.B) == nil {
		return nil
	}
	alts.B = reset.B
	return &UnmarshalError{B: b, structName: "CreateRawTransactionReqOutputs"}
}

// CreateRawTransactionResp holds the response to the CreateRawTransaction call.
//  "hex"    (string) hex string of the transaction
type CreateRawTransactionResp struct {
	// hex string of the transaction
	Hex string
}

func (alts CreateRawTransactionResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.Hex)
}

func (alts *CreateRawTransactionResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Hex) == nil {
		return nil
	}
	alts.Hex = reset.Hex
	return &UnmarshalError{B: b, structName: "CreateRawTransactionResp"}
}

// CreateRawTransaction RPC method.
// Create a transaction spending the given inputs and creating new outputs.
// Outputs can be addresses or data.
// Returns hex-encoded raw transaction.
// Note that the transaction's inputs are not signed, and
// it is not stored in the wallet or transmitted to the network.
func (bc *BitcoindClient) CreateRawTransaction(ctx context.Context, args CreateRawTransactionReq) (result CreateRawTransactionResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "createrawtransaction", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// DecodePsbtReq holds the arguments for the DecodePsbt call.
//  1. psbt    (string, required) The PSBT base64 string
type DecodePsbtReq struct {
	// The PSBT base64 string
	Psbt string `json:"psbt"`
}

// DecodePsbtResp holds the response to the DecodePsbt call.
//  {                                          (json object)
//    "tx" : {                                 (json object) The decoded network-serialized unsigned transaction.
//      ...                                    The layout is the same as the output of decoderawtransaction.
//    },
//    "unknown" : {                            (json object) The unknown global fields
//      "key" : "hex",                         (string) (key-value pair) An unknown key-value pair
//      ...
//    },
//    "inputs" : [                             (json array)
//      {                                      (json object)
//        "non_witness_utxo" : {               (json object, optional) Decoded network transaction for non-witness UTXOs
//          ...
//        },
//        "witness_utxo" : {                   (json object, optional) Transaction output for witness UTXOs
//          "amount" : n,                      (numeric) The value in BTC
//          "scriptPubKey" : {                 (json object)
//            "asm" : "str",                   (string) The asm
//            "hex" : "hex",                   (string) The hex
//            "type" : "str",                  (string) The type, eg 'pubkeyhash'
//            "address" : "str"                (string)  Bitcoin address if there is one
//          }
//        },
//        "partial_signatures" : {             (json object, optional)
//          "pubkey" : "str",                  (string) The public key and signature that corresponds to it.
//          ...
//        },
//        "sighash" : "str",                   (string, optional) The sighash type to be used
//        "redeem_script" : {                  (json object, optional)
//          "asm" : "str",                     (string) The asm
//          "hex" : "hex",                     (string) The hex
//          "type" : "str"                     (string) The type, eg 'pubkeyhash'
//        },
//        "witness_script" : {                 (json object, optional)
//          "asm" : "str",                     (string) The asm
//          "hex" : "hex",                     (string) The hex
//          "type" : "str"                     (string) The type, eg 'pubkeyhash'
//        },
//        "bip32_derivs" : [                   (json array, optional)
//          {                                  (json object, optional) The public key with the derivation path as the value.
//            "master_fingerprint" : "str",    (string) The fingerprint of the master key
//            "path" : "str"                   (string) The path
//          },
//          ...
//        ],
//        "final_scriptsig" : {                (json object, optional)
//          "asm" : "str",                     (string) The asm
//          "hex" : "str"                      (string) The hex
//        },
//        "final_scriptwitness" : [            (json array)
//          "hex",                             (string) hex-encoded witness data (if any)
//          ...
//        ],
//        "unknown" : {                        (json object) The unknown global fields
//          "key" : "hex",                     (string) (key-value pair) An unknown key-value pair
//          ...
//        }
//      },
//      ...
//    ],
//    "outputs" : [                            (json array)
//      {                                      (json object)
//        "redeem_script" : {                  (json object, optional)
//          "asm" : "str",                     (string) The asm
//          "hex" : "hex",                     (string) The hex
//          "type" : "str"                     (string) The type, eg 'pubkeyhash'
//        },
//        "witness_script" : {                 (json object, optional)
//          "asm" : "str",                     (string) The asm
//          "hex" : "hex",                     (string) The hex
//          "type" : "str"                     (string) The type, eg 'pubkeyhash'
//        },
//        "bip32_derivs" : [                   (json array, optional)
//          {                                  (json object)
//            "pubkey" : "str",                (string) The public key this path corresponds to
//            "master_fingerprint" : "str",    (string) The fingerprint of the master key
//            "path" : "str"                   (string) The path
//          },
//          ...
//        ],
//        "unknown" : {                        (json object) The unknown global fields
//          "key" : "hex",                     (string) (key-value pair) An unknown key-value pair
//          ...
//        }
//      },
//      ...
//    ],
//    "fee" : n                                (numeric, optional) The transaction fee paid if all UTXOs slots in the PSBT have been filled.
//  }
type DecodePsbtResp struct {
	// The decoded network-serialized unsigned transaction.
	// The layout is the same as the output of decoderawtransaction.
	Tx DecodeRawTransactionResp `json:"tx"`

	// The unknown global fields
	Unknown map[string]string `json:"unknown"`

	Inputs []DecodePsbtRespInputs `json:"inputs"`

	Outputs []DecodePsbtRespOutputs `json:"outputs"`

	// The transaction fee paid if all UTXOs slots in the PSBT have been filled.
	Fee *float64 `json:"fee,omitempty"`
}

type DecodePsbtRespInputs struct {
	// Decoded network transaction for non-witness UTXOs
	NonWitnessUtxo *DecodePsbtRespInputsWitnessUtxo `json:"non_witness_utxo,omitempty"`

	// Transaction output for witness UTXOs
	WitnessUtxo *DecodePsbtRespInputsWitnessUtxo `json:"witness_utxo,omitempty"`

	PartialSignatures map[string]string `json:"partial_signatures,omitempty"`

	// The sighash type to be used
	SigHash string `json:"sighash,omitempty"`

	RedeemScript *DecodePsbtRespInputsRedeemScript `json:"redeem_script,omitempty"`

	WitnessScript *DecodePsbtRespInputsWitnessScript `json:"witness_script,omitempty"`

	BIP32Derivs []DecodePsbtRespInputsBIP32Derivs `json:"bip32_derivs,omitempty"`

	FinalScriptSig *DecodePsbtRespInputsFinalScriptSig `json:"final_scriptsig,omitempty"`

	// Element: Hex    hex-encoded witness data (if any)
	FinalScriptWitness []string `json:"final_scriptwitness"`

	// The unknown global fields
	Unknown map[string]string `json:"unknown"`
}

type DecodePsbtRespInputsWitnessUtxo struct {
	// The value in BTC
	Amount float64 `json:"amount"`

	ScriptPubkey struct {
		// The asm
		Asm string `json:"asm"`

		// The hex
		Hex string `json:"hex"`

		// The type, eg 'pubkeyhash'
		Type string `json:"type"`

		// Bitcoin address if there is one
		Address string `json:"address"`
	} `json:"scriptPubKey"`
}

type DecodePsbtRespInputsRedeemScript struct {
	// The asm
	Asm string `json:"asm"`

	// The hex
	Hex string `json:"hex"`

	// The type, eg 'pubkeyhash'
	Type string `json:"type"`
}

type DecodePsbtRespInputsWitnessScript struct {
	// The asm
	Asm string `json:"asm"`

	// The hex
	Hex string `json:"hex"`

	// The type, eg 'pubkeyhash'
	Type string `json:"type"`
}

// The public key with the derivation path as the value.
type DecodePsbtRespInputsBIP32Derivs struct {
	// The fingerprint of the master key
	MasterFingerprint string `json:"master_fingerprint"`

	// The path
	Path string `json:"path"`
}

type DecodePsbtRespInputsFinalScriptSig struct {
	// The asm
	Asm string `json:"asm"`

	// The hex
	Hex string `json:"hex"`
}

type DecodePsbtRespOutputs struct {
	RedeemScript *DecodePsbtRespOutputsRedeemScript `json:"redeem_script,omitempty"`

	WitnessScript *DecodePsbtRespOutputsWitnessScript `json:"witness_script,omitempty"`

	BIP32Derivs []DecodePsbtRespOutputsBIP32Derivs `json:"bip32_derivs,omitempty"`

	// The unknown global fields
	Unknown map[string]string `json:"unknown"`
}

type DecodePsbtRespOutputsRedeemScript struct {
	// The asm
	Asm string `json:"asm"`

	// The hex
	Hex string `json:"hex"`

	// The type, eg 'pubkeyhash'
	Type string `json:"type"`
}

type DecodePsbtRespOutputsWitnessScript struct {
	// The asm
	Asm string `json:"asm"`

	// The hex
	Hex string `json:"hex"`

	// The type, eg 'pubkeyhash'
	Type string `json:"type"`
}

type DecodePsbtRespOutputsBIP32Derivs struct {
	// The public key this path corresponds to
	Pubkey string `json:"pubkey"`

	// The fingerprint of the master key
	MasterFingerprint string `json:"master_fingerprint"`

	// The path
	Path string `json:"path"`
}

// DecodePsbt RPC method.
// Return a JSON object representing the serialized, base64-encoded partially signed Bitcoin transaction.
func (bc *BitcoindClient) DecodePsbt(ctx context.Context, args DecodePsbtReq) (result DecodePsbtResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "decodepsbt", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// DecodeRawTransactionReq holds the arguments for the DecodeRawTransaction call.
//  1. hexstring    (string, required) The transaction hex string
//  2. iswitness    (boolean, optional, default=depends on heuristic tests) Whether the transaction hex is a serialized witness transaction.
//                  If iswitness is not present, heuristic tests will be used in decoding.
//                  If true, only witness deserialization will be tried.
//                  If false, only non-witness deserialization will be tried.
//                  This boolean should reflect whether the transaction has inputs
//                  (e.g. fully valid, or on-chain transactions), if known by the caller.
type DecodeRawTransactionReq struct {
	// The transaction hex string
	HexString string `json:"hexstring"`

	// Whether the transaction hex is a serialized witness transaction.
	// If iswitness is not present, heuristic tests will be used in decoding.
	// If true, only witness deserialization will be tried.
	// If false, only non-witness deserialization will be tried.
	// This boolean should reflect whether the transaction has inputs
	// (e.g. fully valid, or on-chain transactions), if known by the caller.
	// Default: depends on heuristic tests
	IsWitness *bool `json:"iswitness,omitempty"`
}

// DecodeRawTransactionResp holds the response to the DecodeRawTransaction call.
//  {                             (json object)
//    "txid" : "hex",             (string) The transaction id
//    "hash" : "hex",             (string) The transaction hash (differs from txid for witness transactions)
//    "size" : n,                 (numeric) The transaction size
//    "vsize" : n,                (numeric) The virtual transaction size (differs from size for witness transactions)
//    "weight" : n,               (numeric) The transaction's weight (between vsize*4 - 3 and vsize*4)
//    "version" : n,              (numeric) The version
//    "locktime" : xxx,           (numeric) The lock time
//    "vin" : [                   (json array)
//      {                         (json object)
//        "txid" : "hex",         (string) The transaction id
//        "vout" : n,             (numeric) The output number
//        "scriptSig" : {         (json object) The script
//          "asm" : "str",        (string) asm
//          "hex" : "hex"         (string) hex
//        },
//        "txinwitness" : [       (json array)
//          "hex",                (string) hex-encoded witness data (if any)
//          ...
//        ],
//        "sequence" : n          (numeric) The script sequence number
//      },
//      ...
//    ],
//    "vout" : [                  (json array)
//      {                         (json object)
//        "value" : n,            (numeric) The value in BTC
//        "n" : n,                (numeric) index
//        "scriptPubKey" : {      (json object)
//          "asm" : "str",        (string) the asm
//          "hex" : "hex",        (string) the hex
//          "reqSigs" : n,        (numeric, optional) (DEPRECATED, returned only if config option -deprecatedrpc=addresses is passed) Number of required signatures
//          "type" : "str",       (string) The type, eg 'pubkeyhash'
//          "address" : "str",    (string, optional) bitcoin address (only if a well-defined address exists)
//          "addresses" : [       (json array, optional) (DEPRECATED, returned only if config option -deprecatedrpc=addresses is passed) Array of bitcoin addresses
//            "str",              (string) bitcoin address
//            ...
//          ]
//        }
//      },
//      ...
//    ]
//  }
type DecodeRawTransactionResp struct {
	// The transaction id
	TxID string `json:"txid"`

	// The transaction hash (differs from txid for witness transactions)
	Hash string `json:"hash"`

	// The transaction size
	Size float64 `json:"size"`

	// The virtual transaction size (differs from size for witness transactions)
	VSize float64 `json:"vsize"`

	// The transaction's weight (between vsize*4 - 3 and vsize*4)
	Weight float64 `json:"weight"`

	// The version
	Version float64 `json:"version"`

	// The lock time
	LockTime float64 `json:"locktime"`

	Vin []DecodeRawTransactionRespVin `json:"vin"`

	Vout []DecodeRawTransactionRespVout `json:"vout"`
}

type DecodeRawTransactionRespVin struct {
	// The transaction id
	TxID string `json:"txid"`

	// The output number
	Vout float64 `json:"vout"`

	// The script
	ScriptSig struct {
		// asm
		Asm string `json:"asm"`

		// hex
		Hex string `json:"hex"`
	} `json:"scriptSig"`

	// Element: Hex    hex-encoded witness data (if any)
	TxInWitness []string `json:"txinwitness"`

	// The script sequence number
	Sequence float64 `json:"sequence"`
}

type DecodeRawTransactionRespVout struct {
	// The value in BTC
	Value float64 `json:"value"`

	// index
	N float64 `json:"n"`

	ScriptPubkey struct {
		// the asm
		Asm string `json:"asm"`

		// the hex
		Hex string `json:"hex"`

		// (DEPRECATED, returned only if config option -deprecatedrpc=addresses is passed) Number of required signatures
		ReqSigs *float64 `json:"reqSigs,omitempty"`

		// The type, eg 'pubkeyhash'
		Type string `json:"type"`

		// bitcoin address (only if a well-defined address exists)
		Address string `json:"address,omitempty"`

		// (DEPRECATED, returned only if config option -deprecatedrpc=addresses is passed) Array of bitcoin addresses
		// Element: Str    bitcoin address
		Addresses []string `json:"addresses,omitempty"`
	} `json:"scriptPubKey"`
}

// DecodeRawTransaction RPC method.
// Return a JSON object representing the serialized, hex-encoded transaction.
func (bc *BitcoindClient) DecodeRawTransaction(ctx context.Context, args DecodeRawTransactionReq) (result DecodeRawTransactionResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "decoderawtransaction", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// DecodeScriptReq holds the arguments for the DecodeScript call.
//  1. hexstring    (string, required) the hex-encoded script
type DecodeScriptReq struct {
	// the hex-encoded script
	HexString string `json:"hexstring"`
}

// DecodeScriptResp holds the response to the DecodeScript call.
//  {                             (json object)
//    "asm" : "str",              (string) Script public key
//    "type" : "str",             (string) The output type (e.g. nonstandard, pubkey, pubkeyhash, scripthash, multisig, nulldata, witness_v0_scripthash, witness_v0_keyhash, witness_v1_taproot, witness_unknown)
//    "address" : "str",          (string, optional) bitcoin address (only if a well-defined address exists)
//    "reqSigs" : n,              (numeric, optional) (DEPRECATED, returned only if config option -deprecatedrpc=addresses is passed) Number of required signatures
//    "addresses" : [             (json array, optional) (DEPRECATED, returned only if config option -deprecatedrpc=addresses is passed) Array of bitcoin addresses
//      "str",                    (string) bitcoin address
//      ...
//    ],
//    "p2sh" : "str",             (string) address of P2SH script wrapping this redeem script (not returned if the script is already a P2SH)
//    "segwit" : {                (json object) Result of a witness script public key wrapping this redeem script (not returned if the script is a P2SH or witness)
//      "asm" : "str",            (string) String representation of the script public key
//      "hex" : "hex",            (string) Hex string of the script public key
//      "type" : "str",           (string) The type of the script public key (e.g. witness_v0_keyhash or witness_v0_scripthash)
//      "address" : "str",        (string, optional) bitcoin address (only if a well-defined address exists)
//      "reqSigs" : n,            (numeric, optional) (DEPRECATED, returned only if config option -deprecatedrpc=addresses is passed) Number of required signatures
//      "addresses" : [           (json array, optional) (DEPRECATED, returned only if config option -deprecatedrpc=addresses is passed) Array of bitcoin addresses
//        "str",                  (string) segwit address
//        ...
//      ],
//      "p2sh-segwit" : "str"     (string) address of the P2SH script wrapping this witness redeem script
//    }
//  }
type DecodeScriptResp struct {
	// Script public key
	Asm string `json:"asm"`

	// The output type (e.g. nonstandard, pubkey, pubkeyhash, scripthash, multisig, nulldata, witness_v0_scripthash, witness_v0_keyhash, witness_v1_taproot, witness_unknown)
	Type string `json:"type"`

	// bitcoin address (only if a well-defined address exists)
	Address string `json:"address,omitempty"`

	// (DEPRECATED, returned only if config option -deprecatedrpc=addresses is passed) Number of required signatures
	ReqSigs *float64 `json:"reqSigs,omitempty"`

	// (DEPRECATED, returned only if config option -deprecatedrpc=addresses is passed) Array of bitcoin addresses
	// Element: Str    bitcoin address
	Addresses []string `json:"addresses,omitempty"`

	// address of P2SH script wrapping this redeem script (not returned if the script is already a P2SH)
	P2SH string `json:"p2sh"`

	// Result of a witness script public key wrapping this redeem script (not returned if the script is a P2SH or witness)
	Segwit struct {
		// String representation of the script public key
		Asm string `json:"asm"`

		// Hex string of the script public key
		Hex string `json:"hex"`

		// The type of the script public key (e.g. witness_v0_keyhash or witness_v0_scripthash)
		Type string `json:"type"`

		// bitcoin address (only if a well-defined address exists)
		Address string `json:"address,omitempty"`

		// (DEPRECATED, returned only if config option -deprecatedrpc=addresses is passed) Number of required signatures
		ReqSigs *float64 `json:"reqSigs,omitempty"`

		// (DEPRECATED, returned only if config option -deprecatedrpc=addresses is passed) Array of bitcoin addresses
		// Element: Str    segwit address
		Addresses []string `json:"addresses,omitempty"`

		// address of the P2SH script wrapping this witness redeem script
		P2SHSegwit string `json:"p2sh-segwit"`
	} `json:"segwit"`
}

// DecodeScript RPC method.
// Decode a hex-encoded script.
func (bc *BitcoindClient) DecodeScript(ctx context.Context, args DecodeScriptReq) (result DecodeScriptResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "decodescript", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// FinalizePsbtReq holds the arguments for the FinalizePsbt call.
//  1. psbt       (string, required) A base64 string of a PSBT
//  2. extract    (boolean, optional, default=true) If true and the transaction is complete,
//                extract and return the complete transaction in normal network serialization instead of the PSBT.
type FinalizePsbtReq struct {
	// A base64 string of a PSBT
	Psbt string `json:"psbt"`

	// If true and the transaction is complete,
	// extract and return the complete transaction in normal network serialization instead of the PSBT.
	// Default: true
	Extract *bool `json:"extract,omitempty"`
}

// FinalizePsbtResp holds the response to the FinalizePsbt call.
//  {                             (json object)
//    "psbt" : "str",             (string) The base64-encoded partially signed transaction if not extracted
//    "hex" : "hex",              (string) The hex-encoded network transaction if extracted
//    "complete" : true|false     (boolean) If the transaction has a complete set of signatures
//  }
type FinalizePsbtResp struct {
	// The base64-encoded partially signed transaction if not extracted
	Psbt string `json:"psbt"`

	// The hex-encoded network transaction if extracted
	Hex string `json:"hex"`

	// If the transaction has a complete set of signatures
	Complete bool `json:"complete"`
}

// FinalizePsbt RPC method.
// Finalize the inputs of a PSBT. If the transaction is fully signed, it will produce a
// network serialized transaction which can be broadcast with sendrawtransaction. Otherwise a PSBT will be
// created which has the final_scriptSig and final_scriptWitness fields filled for inputs that are complete.
// Implements the Finalizer and Extractor roles.
func (bc *BitcoindClient) FinalizePsbt(ctx context.Context, args FinalizePsbtReq) (result FinalizePsbtResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "finalizepsbt", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// FundRawTransactionReq holds the arguments for the FundRawTransaction call.
//  1. hexstring                          (string, required) The hex string of the raw transaction
//  2. options                            (json object, optional) for backward compatibility: passing in a true instead of an object will result in {"includeWatching":true}
//       {
//         "add_inputs": bool,            (boolean, optional, default=true) For a transaction with existing inputs, automatically include more if they are not enough.
//         "include_unsafe": bool,        (boolean, optional, default=false) Include inputs that are not safe to spend (unconfirmed transactions from outside keys and unconfirmed replacement transactions).
//                                        Warning: the resulting transaction may become invalid if one of the unsafe inputs disappears.
//                                        If that happens, you will need to fund the transaction with different inputs and republish it.
//         "changeAddress": "str",        (string, optional, default=pool address) The bitcoin address to receive the change
//         "changePosition": n,           (numeric, optional, default=random) The index of the change output
//         "change_type": "str",          (string, optional, default=set by -changetype) The output type to use. Only valid if changeAddress is not specified. Options are "legacy", "p2sh-segwit", and "bech32".
//         "includeWatching": bool,       (boolean, optional, default=true for watch-only wallets, otherwise false) Also select inputs which are watch only.
//                                        Only solvable inputs can be used. Watch-only destinations are solvable if the public key and/or output script was imported,
//                                        e.g. with 'importpubkey' or 'importmulti' with the 'pubkeys' or 'desc' field.
//         "lockUnspents": bool,          (boolean, optional, default=false) Lock selected unspent outputs
//         "fee_rate": amount,            (numeric or string, optional, default=not set, fall back to wallet fee estimation) Specify a fee rate in sat/vB.
//         "feeRate": amount,             (numeric or string, optional, default=not set, fall back to wallet fee estimation) Specify a fee rate in BTC/kvB.
//         "subtractFeeFromOutputs": [    (json array, optional, default=[]) The integers.
//                                        The fee will be equally deducted from the amount of each specified output.
//                                        Those recipients will receive less bitcoins than you enter in their corresponding amount field.
//                                        If no outputs are specified here, the sender pays the fee.
//           vout_index,                  (numeric) The zero-based output index, before a change output is added.
//           ...
//         ],
//         "replaceable": bool,           (boolean, optional, default=wallet default) Marks this transaction as BIP125 replaceable.
//                                        Allows this transaction to be replaced by a transaction with higher fees
//         "conf_target": n,              (numeric, optional, default=wallet -txconfirmtarget) Confirmation target in blocks
//         "estimate_mode": "str",        (string, optional, default="unset") The fee estimate mode, must be one of (case insensitive):
//                                        "unset"
//                                        "economical"
//                                        "conservative"
//       }
//  3. iswitness                          (boolean, optional, default=depends on heuristic tests) Whether the transaction hex is a serialized witness transaction.
//                                        If iswitness is not present, heuristic tests will be used in decoding.
//                                        If true, only witness deserialization will be tried.
//                                        If false, only non-witness deserialization will be tried.
//                                        This boolean should reflect whether the transaction has inputs
//                                        (e.g. fully valid, or on-chain transactions), if known by the caller.
type FundRawTransactionReq struct {
	// The hex string of the raw transaction
	HexString string `json:"hexstring"`

	// for backward compatibility: passing in a true instead of an object will result in {"includeWatching":true}
	Options *FundRawTransactionReqOptions `json:"options,omitempty"`

	// Whether the transaction hex is a serialized witness transaction.
	// If iswitness is not present, heuristic tests will be used in decoding.
	// If true, only witness deserialization will be tried.
	// If false, only non-witness deserialization will be tried.
	// This boolean should reflect whether the transaction has inputs
	// (e.g. fully valid, or on-chain transactions), if known by the caller.
	// Default: depends on heuristic tests
	IsWitness *bool `json:"iswitness,omitempty"`
}

type FundRawTransactionReqOptions struct {
	// For a transaction with existing inputs, automatically include more if they are not enough.
	// Default: true
	AddInputs *bool `json:"add_inputs,omitempty"`

	// Include inputs that are not safe to spend (unconfirmed transactions from outside keys and unconfirmed replacement transactions).
	// Warning: the resulting transaction may become invalid if one of the unsafe inputs disappears.
	// If that happens, you will need to fund the transaction with different inputs and republish it.
	// Default: false
	IncludeUnsafe bool `json:"include_unsafe,omitempty"`

	// The bitcoin address to receive the change
	// Default: pool address
	Changeaddress string `json:"changeAddress,omitempty"`

	// The index of the change output
	// Default: random
	ChangePosition *float64 `json:"changePosition,omitempty"`

	// The output type to use. Only valid if changeAddress is not specified. Options are "legacy", "p2sh-segwit", and "bech32".
	// Default: set by -changetype
	ChangeType string `json:"change_type,omitempty"`

	// Also select inputs which are watch only.
	// Only solvable inputs can be used. Watch-only destinations are solvable if the public key and/or output script was imported,
	// e.g. with 'importpubkey' or 'importmulti' with the 'pubkeys' or 'desc' field.
	// Default: true for watch-only wallets, otherwise false
	IncludeWatching *bool `json:"includeWatching,omitempty"`

	// Lock selected unspent outputs
	// Default: false
	LockUnspents bool `json:"lockUnspents,omitempty"`

	// Specify a fee rate in sat/vB.
	// Default: not set, fall back to wallet fee estimation
	FeeRate *float64 `json:"fee_rate,omitempty"`

	// Specify a fee rate in BTC/kvB.
	// Default: not set, fall back to wallet fee estimation
	FeeRate2 *float64 `json:"feeRate,omitempty"`

	// The integers.
	// The fee will be equally deducted from the amount of each specified output.
	// Those recipients will receive less bitcoins than you enter in their corresponding amount field.
	// If no outputs are specified here, the sender pays the fee.
	// Element: VoutIndex    The zero-based output index, before a change output is added.
	SubtractFeeFromOutputs []float64 `json:"subtractFeeFromOutputs,omitempty"`

	// Marks this transaction as BIP125 replaceable.
	// Allows this transaction to be replaced by a transaction with higher fees
	// Default: wallet default
	Replaceable *bool `json:"replaceable,omitempty"`

	// Confirmation target in blocks
	// Default: wallet -txconfirmtarget
	ConfTarget *float64 `json:"conf_target,omitempty"`

	// The fee estimate mode, must be one of (case insensitive):
	// "unset"
	// "economical"
	// "conservative"
	// Default: "unset"
	EstimateMode string `json:"estimate_mode,omitempty"`
}

// FundRawTransactionResp holds the response to the FundRawTransaction call.
//  {                     (json object)
//    "hex" : "hex",      (string) The resulting raw transaction (hex-encoded string)
//    "fee" : n,          (numeric) Fee in BTC the resulting transaction pays
//    "changepos" : n     (numeric) The position of the added change output, or -1
//  }
type FundRawTransactionResp struct {
	// The resulting raw transaction (hex-encoded string)
	Hex string `json:"hex"`

	// Fee in BTC the resulting transaction pays
	Fee float64 `json:"fee"`

	// The position of the added change output, or -1
	ChangePos float64 `json:"changepos"`
}

// FundRawTransaction RPC method.
// If the transaction has no inputs, they will be automatically selected to meet its out value.
// It will add at most one change output to the outputs.
// No existing outputs will be modified unless "subtractFeeFromOutputs" is specified.
// Note that inputs which were signed may need to be resigned after completion since in/outputs have been added.
// The inputs added will not be signed, use signrawtransactionwithkey
//  or signrawtransactionwithwallet for that.
// Note that all existing inputs must have their previous output transaction be in the wallet.
// Note that all inputs selected must be of standard form and P2SH scripts must be
// in the wallet using importaddress or addmultisigaddress (to calculate fees).
// You can see whether this is the case by checking the "solvable" field in the listunspent output.
// Only pay-to-pubkey, multisig, and P2SH versions thereof are currently supported for watch-only
func (bc *BitcoindClient) FundRawTransaction(ctx context.Context, args FundRawTransactionReq) (result FundRawTransactionResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "fundrawtransaction", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// GetRawTransactionReq holds the arguments for the GetRawTransaction call.
//  1. txid         (string, required) The transaction id
//  2. verbose      (boolean, optional, default=false) If false, return a string, otherwise return a json object
//  3. blockhash    (string, optional) The block in which to look for the transaction
type GetRawTransactionReq struct {
	// The transaction id
	TxID string `json:"txid"`

	// If false, return a string, otherwise return a json object
	// Default: false
	Verbose bool `json:"verbose,omitempty"`

	// The block in which to look for the transaction
	Blockhash string `json:"blockhash,omitempty"`
}

// GetRawTransactionResp holds the response to the GetRawTransaction call.
//
// ALTERNATIVE (if verbose is not set or set to false)
//  "str"    (string) The serialized, hex-encoded data for 'txid'
//
// ALTERNATIVE (if verbose is set to true)
//  {                                    (json object)
//    "in_active_chain" : true|false,    (boolean) Whether specified block is in the active chain or not (only present with explicit "blockhash" argument)
//    "hex" : "hex",                     (string) The serialized, hex-encoded data for 'txid'
//    "txid" : "hex",                    (string) The transaction id (same as provided)
//    "hash" : "hex",                    (string) The transaction hash (differs from txid for witness transactions)
//    "size" : n,                        (numeric) The serialized transaction size
//    "vsize" : n,                       (numeric) The virtual transaction size (differs from size for witness transactions)
//    "weight" : n,                      (numeric) The transaction's weight (between vsize*4-3 and vsize*4)
//    "version" : n,                     (numeric) The version
//    "locktime" : xxx,                  (numeric) The lock time
//    "vin" : [                          (json array)
//      {                                (json object)
//        "txid" : "hex",                (string) The transaction id
//        "vout" : n,                    (numeric) The output number
//        "scriptSig" : {                (json object) The script
//          "asm" : "str",               (string) asm
//          "hex" : "hex"                (string) hex
//        },
//        "sequence" : n,                (numeric) The script sequence number
//        "txinwitness" : [              (json array)
//          "hex",                       (string) hex-encoded witness data (if any)
//          ...
//        ]
//      },
//      ...
//    ],
//    "vout" : [                         (json array)
//      {                                (json object)
//        "value" : n,                   (numeric) The value in BTC
//        "n" : n,                       (numeric) index
//        "scriptPubKey" : {             (json object)
//          "asm" : "str",               (string) the asm
//          "hex" : "str",               (string) the hex
//          "reqSigs" : n,               (numeric, optional) (DEPRECATED, returned only if config option -deprecatedrpc=addresses is passed) Number of required signatures
//          "type" : "str",              (string) The type, eg 'pubkeyhash'
//          "address" : "str",           (string, optional) bitcoin address (only if a well-defined address exists)
//          "addresses" : [              (json array, optional) (DEPRECATED, returned only if config option -deprecatedrpc=addresses is passed) Array of bitcoin addresses
//            "str",                     (string) bitcoin address
//            ...
//          ]
//        }
//      },
//      ...
//    ],
//    "blockhash" : "hex",               (string) the block hash
//    "confirmations" : n,               (numeric) The confirmations
//    "blocktime" : xxx,                 (numeric) The block time expressed in UNIX epoch time
//    "time" : n                         (numeric) Same as "blocktime"
//  }
type GetRawTransactionResp struct {
	// The serialized, hex-encoded data for 'txid'
	Str string

	IfVerboseIsSetToTrue GetRawTransactionRespIfVerboseIsSetToTrue
}

func (alts GetRawTransactionResp) MarshalJSON() ([]byte, error) {
	if !reflect.ValueOf(alts.Str).IsZero() {
		return json.Marshal(alts.Str)
	}
	return json.Marshal(alts.IfVerboseIsSetToTrue)
}

func (alts *GetRawTransactionResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Str) == nil {
		return nil
	}
	alts.Str = reset.Str
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.IfVerboseIsSetToTrue) == nil {
		return nil
	}
	alts.IfVerboseIsSetToTrue = reset.IfVerboseIsSetToTrue
	return &UnmarshalError{B: b, structName: "GetRawTransactionResp"}
}

type GetRawTransactionRespIfVerboseIsSetToTrue struct {
	// Whether specified block is in the active chain or not (only present with explicit "blockhash" argument)
	InActiveChain bool `json:"in_active_chain"`

	// The serialized, hex-encoded data for 'txid'
	Hex string `json:"hex"`

	// The transaction id (same as provided)
	TxID string `json:"txid"`

	// The transaction hash (differs from txid for witness transactions)
	Hash string `json:"hash"`

	// The serialized transaction size
	Size float64 `json:"size"`

	// The virtual transaction size (differs from size for witness transactions)
	VSize float64 `json:"vsize"`

	// The transaction's weight (between vsize*4-3 and vsize*4)
	Weight float64 `json:"weight"`

	// The version
	Version float64 `json:"version"`

	// The lock time
	LockTime float64 `json:"locktime"`

	Vin []GetRawTransactionRespIfVerboseIsSetToTrueVin `json:"vin"`

	Vout []GetRawTransactionRespIfVerboseIsSetToTrueVout `json:"vout"`

	// the block hash
	Blockhash string `json:"blockhash"`

	// The confirmations
	Confirmations float64 `json:"confirmations"`

	// The block time expressed in UNIX epoch time
	BlockTime float64 `json:"blocktime"`

	// Same as "blocktime"
	Time float64 `json:"time"`
}

type GetRawTransactionRespIfVerboseIsSetToTrueVin struct {
	// The transaction id
	TxID string `json:"txid"`

	// The output number
	Vout float64 `json:"vout"`

	// The script
	ScriptSig struct {
		// asm
		Asm string `json:"asm"`

		// hex
		Hex string `json:"hex"`
	} `json:"scriptSig"`

	// The script sequence number
	Sequence float64 `json:"sequence"`

	// Element: Hex    hex-encoded witness data (if any)
	TxInWitness []string `json:"txinwitness"`
}

type GetRawTransactionRespIfVerboseIsSetToTrueVout struct {
	// The value in BTC
	Value float64 `json:"value"`

	// index
	N float64 `json:"n"`

	ScriptPubkey struct {
		// the asm
		Asm string `json:"asm"`

		// the hex
		Hex string `json:"hex"`

		// (DEPRECATED, returned only if config option -deprecatedrpc=addresses is passed) Number of required signatures
		ReqSigs *float64 `json:"reqSigs,omitempty"`

		// The type, eg 'pubkeyhash'
		Type string `json:"type"`

		// bitcoin address (only if a well-defined address exists)
		Address string `json:"address,omitempty"`

		// (DEPRECATED, returned only if config option -deprecatedrpc=addresses is passed) Array of bitcoin addresses
		// Element: Str    bitcoin address
		Addresses []string `json:"addresses,omitempty"`
	} `json:"scriptPubKey"`
}

// GetRawTransaction RPC method.
// Return the raw transaction data.
// By default this function only works for mempool transactions. When called with a blockhash
// argument, getrawtransaction will return the transaction if the specified block is available and
// the transaction is found in that block. When called without a blockhash argument, getrawtransaction
// will return the transaction if it is in the mempool, or if -txindex is enabled and the transaction
// is in a block in the blockchain.
// Hint: Use gettransaction for wallet transactions.
// If verbose is 'true', returns an Object with information about 'txid'.
// If verbose is 'false' or omitted, returns a string that is serialized, hex-encoded data for 'txid'.
func (bc *BitcoindClient) GetRawTransaction(ctx context.Context, args GetRawTransactionReq) (result GetRawTransactionResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "getrawtransaction", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// JoinPsbtsReq holds the arguments for the JoinPsbts call.
//  1. txs            (json array, required) The base64 strings of partially signed transactions
//       [
//         "psbt",    (string, required) A base64 string of a PSBT
//         ...
//       ]
type JoinPsbtsReq struct {
	// The base64 strings of partially signed transactions
	// Element: Psbt    A base64 string of a PSBT
	Txs []string `json:"txs"`
}

// JoinPsbtsResp holds the response to the JoinPsbts call.
//  "str"    (string) The base64-encoded partially signed transaction
type JoinPsbtsResp struct {
	// The base64-encoded partially signed transaction
	Str string
}

func (alts JoinPsbtsResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.Str)
}

func (alts *JoinPsbtsResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Str) == nil {
		return nil
	}
	alts.Str = reset.Str
	return &UnmarshalError{B: b, structName: "JoinPsbtsResp"}
}

// JoinPsbts RPC method.
// Joins multiple distinct PSBTs with different inputs and outputs into one PSBT with inputs and outputs from all of the PSBTs
// No input in any of the PSBTs can be in more than one of the PSBTs.
func (bc *BitcoindClient) JoinPsbts(ctx context.Context, args JoinPsbtsReq) (result JoinPsbtsResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "joinpsbts", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// SendRawTransactionReq holds the arguments for the SendRawTransaction call.
//  1. hexstring     (string, required) The hex string of the raw transaction
//  2. maxfeerate    (numeric or string, optional, default="0.10") Reject transactions whose fee rate is higher than the specified value, expressed in BTC/kvB.
//                   Set to 0 to accept any fee rate.
type SendRawTransactionReq struct {
	// The hex string of the raw transaction
	HexString string `json:"hexstring"`

	// Reject transactions whose fee rate is higher than the specified value, expressed in BTC/kvB.
	// Set to 0 to accept any fee rate.
	// Default: "0.10"
	MaxFeeRate *float64 `json:"maxfeerate,omitempty"`
}

// SendRawTransactionResp holds the response to the SendRawTransaction call.
//  "hex"    (string) The transaction hash in hex
type SendRawTransactionResp struct {
	// The transaction hash in hex
	Hex string
}

func (alts SendRawTransactionResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.Hex)
}

func (alts *SendRawTransactionResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Hex) == nil {
		return nil
	}
	alts.Hex = reset.Hex
	return &UnmarshalError{B: b, structName: "SendRawTransactionResp"}
}

// SendRawTransaction RPC method.
// Submit a raw transaction (serialized, hex-encoded) to local node and network.
// The transaction will be sent unconditionally to all peers, so using sendrawtransaction
// for manual rebroadcast may degrade privacy by leaking the transaction's origin, as
// nodes will normally not rebroadcast non-wallet transactions already in their mempool.
// A specific exception, RPC_TRANSACTION_ALREADY_IN_CHAIN, may throw if the transaction cannot be added to the mempool.
// Related RPCs: createrawtransaction, signrawtransactionwithkey
func (bc *BitcoindClient) SendRawTransaction(ctx context.Context, args SendRawTransactionReq) (result SendRawTransactionResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "sendrawtransaction", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// SignRawTransactionWithKeyReq holds the arguments for the SignRawTransactionWithKey call.
//  1. hexstring                        (string, required) The transaction hex string
//  2. privkeys                         (json array, required) The base58-encoded private keys for signing
//       [
//         "privatekey",                (string) private key in base58-encoding
//         ...
//       ]
//  3. prevtxs                          (json array, optional) The previous dependent transaction outputs
//       [
//         {                            (json object)
//           "txid": "hex",             (string, required) The transaction id
//           "vout": n,                 (numeric, required) The output number
//           "scriptPubKey": "hex",     (string, required) script key
//           "redeemScript": "hex",     (string) (required for P2SH) redeem script
//           "witnessScript": "hex",    (string) (required for P2WSH or P2SH-P2WSH) witness script
//           "amount": amount,          (numeric or string) (required for Segwit inputs) the amount spent
//         },
//         ...
//       ]
//  4. sighashtype                      (string, optional, default="DEFAULT") The signature hash type. Must be one of:
//                                      "DEFAULT"
//                                      "ALL"
//                                      "NONE"
//                                      "SINGLE"
//                                      "ALL|ANYONECANPAY"
//                                      "NONE|ANYONECANPAY"
//                                      "SINGLE|ANYONECANPAY"
type SignRawTransactionWithKeyReq struct {
	// The transaction hex string
	HexString string `json:"hexstring"`

	// The base58-encoded private keys for signing
	// Element: PrivateKey    private key in base58-encoding
	Privkeys []string `json:"privkeys"`

	// The previous dependent transaction outputs
	PrevTxs []SignRawTransactionWithKeyReqPrevTxs `json:"prevtxs,omitempty"`

	// The signature hash type. Must be one of:
	// "DEFAULT"
	// "ALL"
	// "NONE"
	// "SINGLE"
	// "ALL|ANYONECANPAY"
	// "NONE|ANYONECANPAY"
	// "SINGLE|ANYONECANPAY"
	// Default: "DEFAULT"
	SigHashType string `json:"sighashtype,omitempty"`
}

type SignRawTransactionWithKeyReqPrevTxs struct {
	// The transaction id
	TxID string `json:"txid"`

	// The output number
	Vout float64 `json:"vout"`

	// script key
	ScriptPubkey string `json:"scriptPubKey"`

	// (required for P2SH) redeem script
	RedeemScript string `json:"redeemScript"`

	// (required for P2WSH or P2SH-P2WSH) witness script
	WitnessScript string `json:"witnessScript"`

	// (required for Segwit inputs) the amount spent
	Amount float64 `json:"amount"`
}

// SignRawTransactionWithKeyResp holds the response to the SignRawTransactionWithKey call.
//  {                             (json object)
//    "hex" : "hex",              (string) The hex-encoded raw transaction with signature(s)
//    "complete" : true|false,    (boolean) If the transaction has a complete set of signatures
//    "errors" : [                (json array, optional) Script verification errors (if there are any)
//      {                         (json object)
//        "txid" : "hex",         (string) The hash of the referenced, previous transaction
//        "vout" : n,             (numeric) The index of the output to spent and used as input
//        "scriptSig" : "hex",    (string) The hex-encoded signature script
//        "sequence" : n,         (numeric) Script sequence number
//        "error" : "str"         (string) Verification or signing error related to the input
//      },
//      ...
//    ]
//  }
type SignRawTransactionWithKeyResp struct {
	// The hex-encoded raw transaction with signature(s)
	Hex string `json:"hex"`

	// If the transaction has a complete set of signatures
	Complete bool `json:"complete"`

	// Script verification errors (if there are any)
	Errors []SignRawTransactionWithKeyRespErrors `json:"errors,omitempty"`
}

type SignRawTransactionWithKeyRespErrors struct {
	// The hash of the referenced, previous transaction
	TxID string `json:"txid"`

	// The index of the output to spent and used as input
	Vout float64 `json:"vout"`

	// The hex-encoded signature script
	ScriptSig string `json:"scriptSig"`

	// Script sequence number
	Sequence float64 `json:"sequence"`

	// Verification or signing error related to the input
	Error string `json:"error"`
}

// SignRawTransactionWithKey RPC method.
// Sign inputs for raw transaction (serialized, hex-encoded).
// The second argument is an array of base58-encoded private
// keys that will be the only keys used to sign the transaction.
// The third optional argument (may be null) is an array of previous transaction outputs that
// this transaction depends on but may not yet be in the block chain.
func (bc *BitcoindClient) SignRawTransactionWithKey(ctx context.Context, args SignRawTransactionWithKeyReq) (result SignRawTransactionWithKeyResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "signrawtransactionwithkey", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// TestMempoolAcceptReq holds the arguments for the TestMempoolAccept call.
//  1. rawtxs          (json array, required) An array of hex strings of raw transactions.
//       [
//         "rawtx",    (string)
//         ...
//       ]
//  2. maxfeerate      (numeric or string, optional, default="0.10") Reject transactions whose fee rate is higher than the specified value, expressed in BTC/kvB
type TestMempoolAcceptReq struct {
	// An array of hex strings of raw transactions.
	// Element: RawTx
	RawTxs []string `json:"rawtxs"`

	// Reject transactions whose fee rate is higher than the specified value, expressed in BTC/kvB
	// Default: "0.10"
	MaxFeeRate *float64 `json:"maxfeerate,omitempty"`
}

// TestMempoolAcceptResp holds the response to the TestMempoolAccept call.
//  [                               (json array) The result of the mempool acceptance test for each raw transaction in the input array.
//                                  Returns results for each transaction in the same order they were passed in.
//                                  It is possible for transactions to not be fully validated ('allowed' unset) if another transaction failed.
//    {                             (json object)
//      "txid" : "hex",             (string) The transaction hash in hex
//      "wtxid" : "hex",            (string) The transaction witness hash in hex
//      "package-error" : "str",    (string) Package validation error, if any (only possible if rawtxs had more than 1 transaction).
//      "allowed" : true|false,     (boolean) Whether this tx would be accepted to the mempool and pass client-specified maxfeerate.If not present, the tx was not fully validated due to a failure in another tx in the list.
//      "vsize" : n,                (numeric) Virtual transaction size as defined in BIP 141. This is different from actual serialized size for witness transactions as witness data is discounted (only present when 'allowed' is true)
//      "fees" : {                  (json object) Transaction fees (only present if 'allowed' is true)
//        "base" : n                (numeric) transaction fee in BTC
//      },
//      "reject-reason" : "str"     (string) Rejection string (only present when 'allowed' is false)
//    },
//    ...
//  ]
type TestMempoolAcceptResp struct {
	// The result of the mempool acceptance test for each raw transaction in the input array.
	// Returns results for each transaction in the same order they were passed in.
	// It is possible for transactions to not be fully validated ('allowed' unset) if another transaction failed.
	Array []TestMempoolAcceptRespElement
}

func (alts TestMempoolAcceptResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.Array)
}

func (alts *TestMempoolAcceptResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Array) == nil {
		return nil
	}
	alts.Array = reset.Array
	return &UnmarshalError{B: b, structName: "TestMempoolAcceptResp"}
}

// The result of the mempool acceptance test for each raw transaction in the input array.
// Returns results for each transaction in the same order they were passed in.
// It is possible for transactions to not be fully validated ('allowed' unset) if another transaction failed.
type TestMempoolAcceptRespElement struct {
	// The transaction hash in hex
	TxID string `json:"txid"`

	// The transaction witness hash in hex
	WTxID string `json:"wtxid"`

	// Package validation error, if any (only possible if rawtxs had more than 1 transaction).
	PackageError string `json:"package-error"`

	// Whether this tx would be accepted to the mempool and pass client-specified maxfeerate.If not present, the tx was not fully validated due to a failure in another tx in the list.
	Allowed bool `json:"allowed"`

	// Virtual transaction size as defined in BIP 141. This is different from actual serialized size for witness transactions as witness data is discounted (only present when 'allowed' is true)
	VSize float64 `json:"vsize"`

	// Transaction fees (only present if 'allowed' is true)
	Fees struct {
		// transaction fee in BTC
		Base float64 `json:"base"`
	} `json:"fees"`

	// Rejection string (only present when 'allowed' is false)
	RejectReason string `json:"reject-reason"`
}

// TestMempoolAccept RPC method.
// Returns result of mempool acceptance tests indicating if raw transaction(s) (serialized, hex-encoded) would be accepted by mempool.
// If multiple transactions are passed in, parents must come before children and package policies apply: the transactions cannot conflict with any mempool transactions or each other.
// If one transaction fails, other transactions may not be fully validated (the 'allowed' key will be blank).
// The maximum number of transactions allowed is 25.
// This checks if transactions violate the consensus or policy rules.
// See sendrawtransaction call.
func (bc *BitcoindClient) TestMempoolAccept(ctx context.Context, args TestMempoolAcceptReq) (result TestMempoolAcceptResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "testmempoolaccept", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// UtxoUpdatePsbtReq holds the arguments for the UtxoUpdatePsbt call.
//  1. psbt                          (string, required) A base64 string of a PSBT
//  2. descriptors                   (json array, optional) An array of either strings or objects
//       [
//         "",                       (string) An output descriptor
//         {                         (json object) An object with an output descriptor and extra information
//           "desc": "str",          (string, required) An output descriptor
//           "range": n or [n,n],    (numeric or array, optional, default=1000) Up to what index HD chains should be explored (either end or [begin,end])
//         },
//         ...
//       ]
type UtxoUpdatePsbtReq struct {
	// A base64 string of a PSBT
	Psbt string `json:"psbt"`

	// An array of either strings or objects
	Descriptors []UtxoUpdatePsbtReqDescriptors `json:"descriptors,omitempty"`
}

// Holder of alternative parameter formats, only one will be used, the first that is non-zero.
type UtxoUpdatePsbtReqDescriptors struct {
	// An output descriptor
	A string

	// An object with an output descriptor and extra information
	B struct {
		// An output descriptor
		Desc string `json:"desc"`

		// Up to what index HD chains should be explored (either end or [begin,end])
		// Default: 1000
		Range *[2]int64 `json:"range,omitempty"`
	}
}

func (alts UtxoUpdatePsbtReqDescriptors) MarshalJSON() ([]byte, error) {
	if !reflect.ValueOf(alts.A).IsZero() {
		return json.Marshal(alts.A)
	}
	return json.Marshal(alts.B)
}

func (alts *UtxoUpdatePsbtReqDescriptors) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.A) == nil {
		return nil
	}
	alts.A = reset.A
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.B) == nil {
		return nil
	}
	alts.B = reset.B
	return &UnmarshalError{B: b, structName: "UtxoUpdatePsbtReqDescriptors"}
}

// UtxoUpdatePsbtResp holds the response to the UtxoUpdatePsbt call.
//  "str"    (string) The base64-encoded partially signed transaction with inputs updated
type UtxoUpdatePsbtResp struct {
	// The base64-encoded partially signed transaction with inputs updated
	Str string
}

func (alts UtxoUpdatePsbtResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.Str)
}

func (alts *UtxoUpdatePsbtResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Str) == nil {
		return nil
	}
	alts.Str = reset.Str
	return &UnmarshalError{B: b, structName: "UtxoUpdatePsbtResp"}
}

// UtxoUpdatePsbt RPC method.
// Updates all segwit inputs and outputs in a PSBT with data from output descriptors, the UTXO set or the mempool.
func (bc *BitcoindClient) UtxoUpdatePsbt(ctx context.Context, args UtxoUpdatePsbtReq) (result UtxoUpdatePsbtResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "utxoupdatepsbt", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}
