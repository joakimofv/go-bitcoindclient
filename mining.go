// THIS FILE IS GENERATED CODE, MANUAL EDITS WILL BE OVERWRITTEN

package bitcoindclient

import (
	"bytes"
	"context"
	"encoding/json"
	"reflect"
)

// GetBlockTemplateReq holds the arguments for the GetBlockTemplate call.
//  1. template_request         (json object, optional, default={}) Format of the template
//       {
//         "mode": "str",       (string, optional) This must be set to "template", "proposal" (see BIP 23), or omitted
//         "capabilities": [    (json array, optional) A list of strings
//           "str",             (string) client side supported feature, 'longpoll', 'coinbasevalue', 'proposal', 'serverlist', 'workid'
//           ...
//         ],
//         "rules": [           (json array, required) A list of strings
//           "segwit",          (string, required) (literal) indicates client side segwit support
//           "str",             (string) other client side supported softfork deployment
//           ...
//         ],
//       }
type GetBlockTemplateReq struct {
	// Format of the template
	// Default: {}
	TemplateRequest *GetBlockTemplateReqTemplateRequest `json:"template_request,omitempty"`
}

type GetBlockTemplateReqTemplateRequest struct {
	// This must be set to "template", "proposal" (see BIP 23), or omitted
	Mode string `json:"mode,omitempty"`

	// A list of strings
	// Element: Str    client side supported feature, 'longpoll', 'coinbasevalue', 'proposal', 'serverlist', 'workid'
	Capabilities []string `json:"capabilities,omitempty"`

	// A list of strings
	// Element: segwit    (literal) indicates client side segwit support
	// Element: str       other client side supported softfork deployment
	Rules []string `json:"rules"`
}

// GetBlockTemplateResp holds the response to the GetBlockTemplate call.
//
// ALTERNATIVE (If the proposal was accepted with mode=='proposal')
//  null    (json null)
//
// ALTERNATIVE (If the proposal was not accepted with mode=='proposal')
//  "str"    (string) According to BIP22
//
// ALTERNATIVE (Otherwise)
//  {                                          (json object)
//    "version" : n,                           (numeric) The preferred block version
//    "rules" : [                              (json array) specific block rules that are to be enforced
//      "str",                                 (string) name of a rule the client must understand to some extent; see BIP 9 for format
//      ...
//    ],
//    "vbavailable" : {                        (json object) set of pending, supported versionbit (BIP 9) softfork deployments
//      "rulename" : n,                        (numeric) identifies the bit number as indicating acceptance and readiness for the named softfork rule
//      ...
//    },
//    "vbrequired" : n,                        (numeric) bit mask of versionbits the server requires set in submissions
//    "previousblockhash" : "str",             (string) The hash of current highest block
//    "transactions" : [                       (json array) contents of non-coinbase transactions that should be included in the next block
//      {                                      (json object)
//        "data" : "hex",                      (string) transaction data encoded in hexadecimal (byte-for-byte)
//        "txid" : "hex",                      (string) transaction id encoded in little-endian hexadecimal
//        "hash" : "hex",                      (string) hash encoded in little-endian hexadecimal (including witness data)
//        "depends" : [                        (json array) array of numbers
//          n,                                 (numeric) transactions before this one (by 1-based index in 'transactions' list) that must be present in the final block if this one is
//          ...
//        ],
//        "fee" : n,                           (numeric) difference in value between transaction inputs and outputs (in satoshis); for coinbase transactions, this is a negative Number of the total collected block fees (ie, not including the block subsidy); if key is not present, fee is unknown and clients MUST NOT assume there isn't one
//        "sigops" : n,                        (numeric) total SigOps cost, as counted for purposes of block limits; if key is not present, sigop cost is unknown and clients MUST NOT assume it is zero
//        "weight" : n                         (numeric) total transaction weight, as counted for purposes of block limits
//      },
//      ...
//    ],
//    "coinbaseaux" : {                        (json object) data that should be included in the coinbase's scriptSig content
//      "key" : "hex",                         (string) values must be in the coinbase (keys may be ignored)
//      ...
//    },
//    "coinbasevalue" : n,                     (numeric) maximum allowable input to coinbase transaction, including the generation award and transaction fees (in satoshis)
//    "longpollid" : "str",                    (string) an id to include with a request to longpoll on an update to this template
//    "target" : "str",                        (string) The hash target
//    "mintime" : xxx,                         (numeric) The minimum timestamp appropriate for the next block time, expressed in UNIX epoch time
//    "mutable" : [                            (json array) list of ways the block template may be changed
//      "str",                                 (string) A way the block template may be changed, e.g. 'time', 'transactions', 'prevblock'
//      ...
//    ],
//    "noncerange" : "hex",                    (string) A range of valid nonces
//    "sigoplimit" : n,                        (numeric) limit of sigops in blocks
//    "sizelimit" : n,                         (numeric) limit of block size
//    "weightlimit" : n,                       (numeric) limit of block weight
//    "curtime" : xxx,                         (numeric) current timestamp in UNIX epoch time
//    "bits" : "str",                          (string) compressed target of next block
//    "height" : n,                            (numeric) The height of the next block
//    "default_witness_commitment" : "str"     (string, optional) a valid witness commitment for the unmodified block template
//  }
type GetBlockTemplateResp struct {
	// According to BIP22
	Str string

	Otherwise GetBlockTemplateRespOtherwise
}

func (alts GetBlockTemplateResp) MarshalJSON() ([]byte, error) {
	if !reflect.ValueOf(alts.Str).IsZero() {
		return json.Marshal(alts.Str)
	}
	return json.Marshal(alts.Otherwise)
}

func (alts *GetBlockTemplateResp) UnmarshalJSON(b []byte) error {
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
	if decoder.Decode(&alts.Otherwise) == nil {
		return nil
	}
	alts.Otherwise = reset.Otherwise
	return &UnmarshalError{B: b, structName: "GetBlockTemplateResp"}
}

type GetBlockTemplateRespOtherwise struct {
	// The preferred block version
	Version float64 `json:"version"`

	// specific block rules that are to be enforced
	// Element: Str    name of a rule the client must understand to some extent; see BIP 9 for format
	Rules []string `json:"rules"`

	// set of pending, supported versionbit (BIP 9) softfork deployments
	VbAvailable map[string]float64 `json:"vbavailable"`

	// bit mask of versionbits the server requires set in submissions
	VbRequired float64 `json:"vbrequired"`

	// The hash of current highest block
	PreviousBlockhash string `json:"previousblockhash"`

	// contents of non-coinbase transactions that should be included in the next block
	Transactions []GetBlockTemplateRespOtherwiseTransactions `json:"transactions"`

	// data that should be included in the coinbase's scriptSig content
	CoinbaseAux map[string]string `json:"coinbaseaux"`

	// maximum allowable input to coinbase transaction, including the generation award and transaction fees (in satoshis)
	CoinbaseValue float64 `json:"coinbasevalue"`

	// an id to include with a request to longpoll on an update to this template
	LongPollID string `json:"longpollid"`

	// The hash target
	Target string `json:"target"`

	// The minimum timestamp appropriate for the next block time, expressed in UNIX epoch time
	MinTime float64 `json:"mintime"`

	// list of ways the block template may be changed
	// Element: Str    A way the block template may be changed, e.g. 'time', 'transactions', 'prevblock'
	MuTable []string `json:"mutable"`

	// A range of valid nonces
	NonceRange string `json:"noncerange"`

	// limit of sigops in blocks
	SigOpLimit float64 `json:"sigoplimit"`

	// limit of block size
	SizeLimit float64 `json:"sizelimit"`

	// limit of block weight
	WeightLimit float64 `json:"weightlimit"`

	// current timestamp in UNIX epoch time
	CurTime float64 `json:"curtime"`

	// compressed target of next block
	Bits string `json:"bits"`

	// The height of the next block
	Height float64 `json:"height"`

	// a valid witness commitment for the unmodified block template
	DefaultWitnessCommitment string `json:"default_witness_commitment,omitempty"`
}

type GetBlockTemplateRespOtherwiseTransactions struct {
	// transaction data encoded in hexadecimal (byte-for-byte)
	Data string `json:"data"`

	// transaction id encoded in little-endian hexadecimal
	TxID string `json:"txid"`

	// hash encoded in little-endian hexadecimal (including witness data)
	Hash string `json:"hash"`

	// array of numbers
	// Element: N    transactions before this one (by 1-based index in 'transactions' list) that must be present in the final block if this one is
	Depends []float64 `json:"depends"`

	// difference in value between transaction inputs and outputs (in satoshis); for coinbase transactions, this is a negative Number of the total collected block fees (ie, not including the block subsidy); if key is not present, fee is unknown and clients MUST NOT assume there isn't one
	Fee float64 `json:"fee"`

	// total SigOps cost, as counted for purposes of block limits; if key is not present, sigop cost is unknown and clients MUST NOT assume it is zero
	SigOps float64 `json:"sigops"`

	// total transaction weight, as counted for purposes of block limits
	Weight float64 `json:"weight"`
}

// GetBlockTemplate RPC method.
// If the request parameters include a 'mode' key, that is used to explicitly select between the default 'template' request or a 'proposal'.
// It returns data needed to construct a block to work on.
// For full specification, see BIPs 22, 23, 9, and 145:
//     https://github.com/bitcoin/bips/blob/master/bip-0022.mediawiki
//     https://github.com/bitcoin/bips/blob/master/bip-0023.mediawiki
//     https://github.com/bitcoin/bips/blob/master/bip-0009.mediawiki#getblocktemplate_changes
//     https://github.com/bitcoin/bips/blob/master/bip-0145.mediawiki
func (bc *BitcoindClient) GetBlockTemplate(ctx context.Context, args GetBlockTemplateReq) (result GetBlockTemplateResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "getblocktemplate", args, false); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// GetMiningInfoResp holds the response to the GetMiningInfo call.
//  {                              (json object)
//    "blocks" : n,                (numeric) The current block
//    "currentblockweight" : n,    (numeric, optional) The block weight of the last assembled block (only present if a block was ever assembled)
//    "currentblocktx" : n,        (numeric, optional) The number of block transactions of the last assembled block (only present if a block was ever assembled)
//    "difficulty" : n,            (numeric) The current difficulty
//    "networkhashps" : n,         (numeric) The network hashes per second
//    "pooledtx" : n,              (numeric) The size of the mempool
//    "chain" : "str",             (string) current network name (main, test, signet, regtest)
//    "warnings" : "str"           (string) any network and blockchain warnings
//  }
type GetMiningInfoResp struct {
	// The current block
	Blocks float64 `json:"blocks"`

	// The block weight of the last assembled block (only present if a block was ever assembled)
	CurrentBlockWeight *float64 `json:"currentblockweight,omitempty"`

	// The number of block transactions of the last assembled block (only present if a block was ever assembled)
	CurrentBlockTx *float64 `json:"currentblocktx,omitempty"`

	// The current difficulty
	Difficulty float64 `json:"difficulty"`

	// The network hashes per second
	NetworkHashPs float64 `json:"networkhashps"`

	// The size of the mempool
	PooledTx float64 `json:"pooledtx"`

	// current network name (main, test, signet, regtest)
	Chain string `json:"chain"`

	// any network and blockchain warnings
	Warnings string `json:"warnings"`
}

// GetMiningInfo RPC method.
// Returns a json object containing mining-related information.
func (bc *BitcoindClient) GetMiningInfo(ctx context.Context) (result GetMiningInfoResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "getmininginfo", nil, false); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// GetNetworkHashPsReq holds the arguments for the GetNetworkHashPs call.
//  1. nblocks    (numeric, optional, default=120) The number of blocks, or -1 for blocks since last difficulty change.
//  2. height     (numeric, optional, default=-1) To estimate at the time of the given height.
type GetNetworkHashPsReq struct {
	// The number of blocks, or -1 for blocks since last difficulty change.
	// Default: 120
	NBlocks *float64 `json:"nblocks,omitempty"`

	// To estimate at the time of the given height.
	// Default: -1
	Height *float64 `json:"height,omitempty"`
}

// GetNetworkHashPsResp holds the response to the GetNetworkHashPs call.
//  n    (numeric) Hashes per second estimated
type GetNetworkHashPsResp struct {
	// Hashes per second estimated
	N float64
}

func (alts GetNetworkHashPsResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.N)
}

func (alts *GetNetworkHashPsResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.N) == nil {
		return nil
	}
	alts.N = reset.N
	return &UnmarshalError{B: b, structName: "GetNetworkHashPsResp"}
}

// GetNetworkHashPs RPC method.
// Returns the estimated network hashes per second based on the last n blocks.
// Pass in [blocks] to override # of blocks, -1 specifies since last difficulty change.
// Pass in [height] to estimate the network speed at the time when a certain block was found.
func (bc *BitcoindClient) GetNetworkHashPs(ctx context.Context, args GetNetworkHashPsReq) (result GetNetworkHashPsResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "getnetworkhashps", args, false); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// PrioritiseTransactionReq holds the arguments for the PrioritiseTransaction call.
//  1. txid         (string, required) The transaction id.
//  2. dummy        (numeric, optional) API-Compatibility for previous API. Must be zero or null.
//                  DEPRECATED. For forward compatibility use named arguments and omit this parameter.
//  3. fee_delta    (numeric, required) The fee value (in satoshis) to add (or subtract, if negative).
//                  Note, that this value is not a fee rate. It is a value to modify absolute fee of the TX.
//                  The fee is not actually paid, only the algorithm for selecting transactions into a block
//                  considers the transaction as it would have paid a higher (or lower) fee.
type PrioritiseTransactionReq struct {
	// The transaction id.
	TxID string `json:"txid"`

	// The fee value (in satoshis) to add (or subtract, if negative).
	// Note, that this value is not a fee rate. It is a value to modify absolute fee of the TX.
	// The fee is not actually paid, only the algorithm for selecting transactions into a block
	// considers the transaction as it would have paid a higher (or lower) fee.
	FeeDelta float64 `json:"fee_delta"`
}

// PrioritiseTransactionResp holds the response to the PrioritiseTransaction call.
//  true|false    (boolean) Returns true
type PrioritiseTransactionResp struct {
	// Returns true
	TrueOrFalse bool
}

func (alts PrioritiseTransactionResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.TrueOrFalse)
}

func (alts *PrioritiseTransactionResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.TrueOrFalse) == nil {
		return nil
	}
	alts.TrueOrFalse = reset.TrueOrFalse
	return &UnmarshalError{B: b, structName: "PrioritiseTransactionResp"}
}

// PrioritiseTransaction RPC method.
// Accepts the transaction into mined blocks at a higher (or lower) priority
func (bc *BitcoindClient) PrioritiseTransaction(ctx context.Context, args PrioritiseTransactionReq) (result PrioritiseTransactionResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "prioritisetransaction", args, false); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// SubmitBlockReq holds the arguments for the SubmitBlock call.
//  1. hexdata    (string, required) the hex-encoded block data to submit
//  2. dummy      (string, optional, default=ignored) dummy value, for compatibility with BIP22. This value is ignored.
type SubmitBlockReq struct {
	// the hex-encoded block data to submit
	HexData string `json:"hexdata"`
}

// SubmitBlockResp holds the response to the SubmitBlock call.
//
// ALTERNATIVE (If the block was accepted)
//  null    (json null)
//
// ALTERNATIVE (Otherwise)
//  "str"    (string) According to BIP22
type SubmitBlockResp struct {
	// According to BIP22
	Str string
}

func (alts SubmitBlockResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.Str)
}

func (alts *SubmitBlockResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Str) == nil {
		return nil
	}
	alts.Str = reset.Str
	return &UnmarshalError{B: b, structName: "SubmitBlockResp"}
}

// SubmitBlock RPC method.
// Attempts to submit new block to network.
// See https://en.bitcoin.it/wiki/BIP_0022 for full specification.
func (bc *BitcoindClient) SubmitBlock(ctx context.Context, args SubmitBlockReq) (result SubmitBlockResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "submitblock", args, false); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// SubmitHeaderReq holds the arguments for the SubmitHeader call.
//  1. hexdata    (string, required) the hex-encoded block header data
type SubmitHeaderReq struct {
	// the hex-encoded block header data
	HexData string `json:"hexdata"`
}

// SubmitHeader RPC method.
// Decode the given hexdata as a header and submit it as a candidate chain tip if valid.
// Throws when the header is invalid.
func (bc *BitcoindClient) SubmitHeader(ctx context.Context, args SubmitHeaderReq) (err error) {
	_, err = bc.sendRequest(ctx, "submitheader", args, false)
	return
}
