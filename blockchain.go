// THIS FILE IS GENERATED CODE, MANUAL EDITS WILL BE OVERWRITTEN

package bitcoindclient

import (
	"bytes"
	"context"
	"encoding/json"
	"reflect"
)

// GetBestBlockhashResp holds the response to the GetBestBlockhash call.
//  "hex"    (string) the block hash, hex-encoded
type GetBestBlockhashResp struct {
	// the block hash, hex-encoded
	Hex string
}

func (alts GetBestBlockhashResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.Hex)
}

func (alts *GetBestBlockhashResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Hex) == nil {
		return nil
	}
	alts.Hex = reset.Hex
	return &UnmarshalError{B: b, structName: "GetBestBlockhashResp"}
}

// GetBestBlockhash RPC method.
// Returns the hash of the best (tip) block in the most-work fully-validated chain.
func (bc *BitcoindClient) GetBestBlockhash(ctx context.Context) (result GetBestBlockhashResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "getbestblockhash", nil); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// GetBlockReq holds the arguments for the GetBlock call.
//  1. blockhash    (string, required) The block hash
//  2. verbosity    (numeric, optional, default=1) 0 for hex-encoded data, 1 for a JSON object, 2 for JSON object with transaction data, and 3 for JSON object with transaction data including prevout information for inputs
type GetBlockReq struct {
	// The block hash
	Blockhash string `json:"blockhash"`

	// 0 for hex-encoded data, 1 for a JSON object, 2 for JSON object with transaction data, and 3 for JSON object with transaction data including prevout information for inputs
	// Default: 1
	Verbosity *float64 `json:"verbosity,omitempty"`
}

// GetBlockResp holds the response to the GetBlock call.
//
// ALTERNATIVE (for verbosity = 0)
//  "hex"    (string) A string that is serialized, hex-encoded data for block 'hash'
//
// ALTERNATIVE (for verbosity = 1)
//  {                                 (json object)
//    "hash" : "hex",                 (string) the block hash (same as provided)
//    "confirmations" : n,            (numeric) The number of confirmations, or -1 if the block is not on the main chain
//    "size" : n,                     (numeric) The block size
//    "strippedsize" : n,             (numeric) The block size excluding witness data
//    "weight" : n,                   (numeric) The block weight as defined in BIP 141
//    "height" : n,                   (numeric) The block height or index
//    "version" : n,                  (numeric) The block version
//    "versionHex" : "hex",           (string) The block version formatted in hexadecimal
//    "merkleroot" : "hex",           (string) The merkle root
//    "tx" : [                        (json array) The transaction ids
//      "hex",                        (string) The transaction id
//      ...
//    ],
//    "time" : xxx,                   (numeric) The block time expressed in UNIX epoch time
//    "mediantime" : xxx,             (numeric) The median block time expressed in UNIX epoch time
//    "nonce" : n,                    (numeric) The nonce
//    "bits" : "hex",                 (string) The bits
//    "difficulty" : n,               (numeric) The difficulty
//    "chainwork" : "hex",            (string) Expected number of hashes required to produce the chain up to this block (in hex)
//    "nTx" : n,                      (numeric) The number of transactions in the block
//    "previousblockhash" : "hex",    (string, optional) The hash of the previous block (if available)
//    "nextblockhash" : "hex"         (string, optional) The hash of the next block (if available)
//  }
//
// ALTERNATIVE (for verbosity = 2)
//  {                   (json object)
//    ...,              Same output as verbosity = 1
//    "tx" : [          (json array)
//      {               (json object)
//        ...,          The transactions in the format of the getrawtransaction RPC. Different from verbosity = 1 "tx" result
//        "fee" : n     (numeric) The transaction fee in BTC, omitted if block undo data is not available
//      },
//      ...
//    ]
//  }
//
// ALTERNATIVE (for verbosity = 3)
//  {                                        (json object)
//    ...,                                   Same output as verbosity = 2
//    "tx" : [                               (json array)
//      {                                    (json object)
//        "vin" : [                          (json array)
//          {                                (json object)
//            ...,                           The same output as verbosity = 2
//            "prevout" : {                  (json object) (Only if undo information is available)
//              "generated" : true|false,    (boolean) Coinbase or not
//              "height" : n,                (numeric) The height of the prevout
//              "value" : n,                 (numeric) The value in BTC
//              "scriptPubKey" : {           (json object)
//                "asm" : "str",             (string) The asm
//                "hex" : "str",             (string) The hex
//                "address" : "str",         (string, optional) The Bitcoin address (only if a well-defined address exists)
//                "type" : "str"             (string) The type (one of: nonstandard, pubkey, pubkeyhash, scripthash, multisig, nulldata, witness_v0_scripthash, witness_v0_keyhash, witness_v1_taproot, witness_unknown)
//              }
//            }
//          },
//          ...
//        ]
//      },
//      ...
//    ]
//  }
type GetBlockResp struct {
	// A string that is serialized, hex-encoded data for block 'hash'
	Hex string

	ForVerbosityEquals1 GetBlockRespForVerbosityEquals1

	ForVerbosityEquals2 GetBlockRespForVerbosityEquals2

	ForVerbosityEquals3 GetBlockRespForVerbosityEquals3
}

func (alts GetBlockResp) MarshalJSON() ([]byte, error) {
	if !reflect.ValueOf(alts.Hex).IsZero() {
		return json.Marshal(alts.Hex)
	}
	if !reflect.ValueOf(alts.ForVerbosityEquals1).IsZero() {
		return json.Marshal(alts.ForVerbosityEquals1)
	}
	if !reflect.ValueOf(alts.ForVerbosityEquals2).IsZero() {
		return json.Marshal(alts.ForVerbosityEquals2)
	}
	return json.Marshal(alts.ForVerbosityEquals3)
}

func (alts *GetBlockResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Hex) == nil {
		return nil
	}
	alts.Hex = reset.Hex
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.ForVerbosityEquals1) == nil {
		return nil
	}
	alts.ForVerbosityEquals1 = reset.ForVerbosityEquals1
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.ForVerbosityEquals2) == nil {
		return nil
	}
	alts.ForVerbosityEquals2 = reset.ForVerbosityEquals2
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.ForVerbosityEquals3) == nil {
		return nil
	}
	alts.ForVerbosityEquals3 = reset.ForVerbosityEquals3
	return &UnmarshalError{B: b, structName: "GetBlockResp"}
}

type GetBlockRespForVerbosityEquals1 struct {
	// the block hash (same as provided)
	Hash string `json:"hash"`

	// The number of confirmations, or -1 if the block is not on the main chain
	Confirmations float64 `json:"confirmations"`

	// The block size
	Size float64 `json:"size"`

	// The block size excluding witness data
	StrippedSize float64 `json:"strippedsize"`

	// The block weight as defined in BIP 141
	Weight float64 `json:"weight"`

	// The block height or index
	Height float64 `json:"height"`

	// The block version
	Version float64 `json:"version"`

	// The block version formatted in hexadecimal
	VersionHex string `json:"versionHex"`

	// The merkle root
	MerkleRoot string `json:"merkleroot"`

	// The transaction ids
	// Element: Hex    The transaction id
	Tx []string `json:"tx"`

	// The block time expressed in UNIX epoch time
	Time float64 `json:"time"`

	// The median block time expressed in UNIX epoch time
	MedianTime float64 `json:"mediantime"`

	// The nonce
	Nonce float64 `json:"nonce"`

	// The bits
	Bits string `json:"bits"`

	// The difficulty
	Difficulty float64 `json:"difficulty"`

	// Expected number of hashes required to produce the chain up to this block (in hex)
	ChainWork string `json:"chainwork"`

	// The number of transactions in the block
	NTx float64 `json:"nTx"`

	// The hash of the previous block (if available)
	PreviousBlockhash string `json:"previousblockhash,omitempty"`

	// The hash of the next block (if available)
	NextBlockhash string `json:"nextblockhash,omitempty"`
}

type GetBlockRespForVerbosityEquals2 struct {
	// Same output as verbosity = 1
	GetBlockRespForVerbosityEquals1

	Tx []GetBlockRespForVerbosityEquals2Tx `json:"tx"`
}

type GetBlockRespForVerbosityEquals2Tx struct {
	// The transactions in the format of the getrawtransaction RPC. Different from verbosity = 1 "tx" result
	GetRawTransactionRespIfVerboseIsSetToTrue

	// The transaction fee in BTC, omitted if block undo data is not available
	Fee float64 `json:"fee"`
}

type GetBlockRespForVerbosityEquals3 struct {
	// Same output as verbosity = 2
	GetBlockRespForVerbosityEquals2

	Tx []GetBlockRespForVerbosityEquals3Tx `json:"tx"`
}

type GetBlockRespForVerbosityEquals3Tx struct {
	Vin []GetBlockRespForVerbosityEquals3TxVin `json:"vin"`
}

type GetBlockRespForVerbosityEquals3TxVin struct {
	// The same output as verbosity = 2
	GetBlockRespForVerbosityEquals2Tx

	// (Only if undo information is available)
	PrevOut struct {
		// Coinbase or not
		Generated bool `json:"generated"`

		// The height of the prevout
		Height float64 `json:"height"`

		// The value in BTC
		Value float64 `json:"value"`

		ScriptPubkey struct {
			// The asm
			Asm string `json:"asm"`

			// The hex
			Hex string `json:"hex"`

			// The Bitcoin address (only if a well-defined address exists)
			Address string `json:"address,omitempty"`

			// The type (one of: nonstandard, pubkey, pubkeyhash, scripthash, multisig, nulldata, witness_v0_scripthash, witness_v0_keyhash, witness_v1_taproot, witness_unknown)
			Type string `json:"type"`
		} `json:"scriptPubKey"`
	} `json:"prevout"`
}

// GetBlock RPC method.
// If verbosity is 0, returns a string that is serialized, hex-encoded data for block 'hash'.
// If verbosity is 1, returns an Object with information about block <hash>.
// If verbosity is 2, returns an Object with information about block <hash> and information about each transaction.
// If verbosity is 3, returns an Object with information about block <hash> and information about each transaction, including prevout information for inputs (only for unpruned blocks in the current best chain).
func (bc *BitcoindClient) GetBlock(ctx context.Context, args GetBlockReq) (result GetBlockResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "getblock", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// GetBlockchainInfoResp holds the response to the GetBlockchainInfo call.
//  {                                         (json object)
//    "chain" : "str",                        (string) current network name (main, test, signet, regtest)
//    "blocks" : n,                           (numeric) the height of the most-work fully-validated chain. The genesis block has height 0
//    "headers" : n,                          (numeric) the current number of headers we have validated
//    "bestblockhash" : "str",                (string) the hash of the currently best block
//    "difficulty" : n,                       (numeric) the current difficulty
//    "time" : xxx,                           (numeric) The block time expressed in UNIX epoch time
//    "mediantime" : xxx,                     (numeric) The median block time expressed in UNIX epoch time
//    "verificationprogress" : n,             (numeric) estimate of verification progress [0..1]
//    "initialblockdownload" : true|false,    (boolean) (debug information) estimate of whether this node is in Initial Block Download mode
//    "chainwork" : "hex",                    (string) total amount of work in active chain, in hexadecimal
//    "size_on_disk" : n,                     (numeric) the estimated size of the block and undo files on disk
//    "pruned" : true|false,                  (boolean) if the blocks are subject to pruning
//    "pruneheight" : n,                      (numeric, optional) lowest-height complete block stored (only present if pruning is enabled)
//    "automatic_pruning" : true|false,       (boolean, optional) whether automatic pruning is enabled (only present if pruning is enabled)
//    "prune_target_size" : n,                (numeric, optional) the target size used by pruning (only present if automatic pruning is enabled)
//    "softforks" : {                         (json object) (DEPRECATED, returned only if config option -deprecatedrpc=softforks is passed) status of softforks
//      "xxxx" : {                            (json object) name of the softfork
//        "type" : "str",                     (string) one of "buried", "bip9"
//        "height" : n,                       (numeric, optional) height of the first block which the rules are or will be enforced (only for "buried" type, or "bip9" type with "active" status)
//        "active" : true|false,              (boolean) true if the rules are enforced for the mempool and the next block
//        "bip9" : {                          (json object, optional) status of bip9 softforks (only for "bip9" type)
//          "bit" : n,                        (numeric, optional) the bit (0-28) in the block version field used to signal this softfork (only for "started" and "locked_in" status)
//          "start_time" : xxx,               (numeric) the minimum median time past of a block at which the bit gains its meaning
//          "timeout" : xxx,                  (numeric) the median time past of a block at which the deployment is considered failed if not yet locked in
//          "min_activation_height" : n,      (numeric) minimum height of blocks for which the rules may be enforced
//          "status" : "str",                 (string) status of deployment at specified block (one of "defined", "started", "locked_in", "active", "failed")
//          "since" : n,                      (numeric) height of the first block to which the status applies
//          "status_next" : "str",            (string) status of deployment at the next block
//          "statistics" : {                  (json object, optional) numeric statistics about signalling for a softfork (only for "started" and "locked_in" status)
//            "period" : n,                   (numeric) the length in blocks of the signalling period
//            "threshold" : n,                (numeric, optional) the number of blocks with the version bit set required to activate the feature (only for "started" status)
//            "elapsed" : n,                  (numeric) the number of blocks elapsed since the beginning of the current period
//            "count" : n,                    (numeric) the number of blocks with the version bit set in the current period
//            "possible" : true|false         (boolean, optional) returns false if there are not enough blocks left in this period to pass activation threshold (only for "started" status)
//          },
//          "signalling" : "str"              (string) indicates blocks that signalled with a # and blocks that did not with a -
//        }
//      },
//      ...
//    },
//    "warnings" : "str"                      (string) any network and blockchain warnings
//  }
type GetBlockchainInfoResp struct {
	// current network name (main, test, signet, regtest)
	Chain string `json:"chain"`

	// the height of the most-work fully-validated chain. The genesis block has height 0
	Blocks float64 `json:"blocks"`

	// the current number of headers we have validated
	Headers float64 `json:"headers"`

	// the hash of the currently best block
	BestBlockhash string `json:"bestblockhash"`

	// the current difficulty
	Difficulty float64 `json:"difficulty"`

	// The block time expressed in UNIX epoch time
	Time float64 `json:"time"`

	// The median block time expressed in UNIX epoch time
	MedianTime float64 `json:"mediantime"`

	// estimate of verification progress [0..1]
	VerificationProgress float64 `json:"verificationprogress"`

	// (debug information) estimate of whether this node is in Initial Block Download mode
	InitialBlockDownload bool `json:"initialblockdownload"`

	// total amount of work in active chain, in hexadecimal
	ChainWork string `json:"chainwork"`

	// the estimated size of the block and undo files on disk
	SizeOnDisk float64 `json:"size_on_disk"`

	// if the blocks are subject to pruning
	Pruned bool `json:"pruned"`

	// lowest-height complete block stored (only present if pruning is enabled)
	PruneHeight *float64 `json:"pruneheight,omitempty"`

	// whether automatic pruning is enabled (only present if pruning is enabled)
	AutomaticPruning *bool `json:"automatic_pruning,omitempty"`

	// the target size used by pruning (only present if automatic pruning is enabled)
	PruneTargetSize *float64 `json:"prune_target_size,omitempty"`

	// (DEPRECATED, returned only if config option -deprecatedrpc=softforks is passed) status of softforks
	// name of the softfork
	// Key: xxxx, Value: struct
	SoftForks map[string]GetBlockchainInfoRespSoftForks `json:"softforks"`

	// any network and blockchain warnings
	Warnings string `json:"warnings"`
}

type GetBlockchainInfoRespSoftForks struct {
	// one of "buried", "bip9"
	Type string `json:"type"`

	// height of the first block which the rules are or will be enforced (only for "buried" type, or "bip9" type with "active" status)
	Height *float64 `json:"height,omitempty"`

	// true if the rules are enforced for the mempool and the next block
	Active bool `json:"active"`

	// status of bip9 softforks (only for "bip9" type)
	BIP9 *GetBlockchainInfoRespSoftForksBIP9 `json:"bip9,omitempty"`
}

type GetBlockchainInfoRespSoftForksBIP9 struct {
	// the bit (0-28) in the block version field used to signal this softfork (only for "started" and "locked_in" status)
	Bit *float64 `json:"bit,omitempty"`

	// the minimum median time past of a block at which the bit gains its meaning
	StartTime float64 `json:"start_time"`

	// the median time past of a block at which the deployment is considered failed if not yet locked in
	TimeOut float64 `json:"timeout"`

	// minimum height of blocks for which the rules may be enforced
	MinActivationHeight float64 `json:"min_activation_height"`

	// status of deployment at specified block (one of "defined", "started", "locked_in", "active", "failed")
	Status string `json:"status"`

	// height of the first block to which the status applies
	Since float64 `json:"since"`

	// status of deployment at the next block
	StatusNext string `json:"status_next"`

	// numeric statistics about signalling for a softfork (only for "started" and "locked_in" status)
	Statistics *GetBlockchainInfoRespSoftForksBIP9Statistics `json:"statistics,omitempty"`

	// indicates blocks that signalled with a # and blocks that did not with a -
	Signalling string `json:"signalling"`
}

type GetBlockchainInfoRespSoftForksBIP9Statistics struct {
	// the length in blocks of the signalling period
	Period float64 `json:"period"`

	// the number of blocks with the version bit set required to activate the feature (only for "started" status)
	Threshold *float64 `json:"threshold,omitempty"`

	// the number of blocks elapsed since the beginning of the current period
	Elapsed float64 `json:"elapsed"`

	// the number of blocks with the version bit set in the current period
	Count float64 `json:"count"`

	// returns false if there are not enough blocks left in this period to pass activation threshold (only for "started" status)
	Possible *bool `json:"possible,omitempty"`
}

// GetBlockchainInfo RPC method.
// Returns an object containing various state info regarding blockchain processing.
func (bc *BitcoindClient) GetBlockchainInfo(ctx context.Context) (result GetBlockchainInfoResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "getblockchaininfo", nil); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// GetBlockCountResp holds the response to the GetBlockCount call.
//  n    (numeric) The current block count
type GetBlockCountResp struct {
	// The current block count
	N float64
}

func (alts GetBlockCountResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.N)
}

func (alts *GetBlockCountResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.N) == nil {
		return nil
	}
	alts.N = reset.N
	return &UnmarshalError{B: b, structName: "GetBlockCountResp"}
}

// GetBlockCount RPC method.
// Returns the height of the most-work fully-validated chain.
// The genesis block has height 0.
func (bc *BitcoindClient) GetBlockCount(ctx context.Context) (result GetBlockCountResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "getblockcount", nil); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// GetBlockFilterReq holds the arguments for the GetBlockFilter call.
//  1. blockhash     (string, required) The hash of the block
//  2. filtertype    (string, optional, default="basic") The type name of the filter
type GetBlockFilterReq struct {
	// The hash of the block
	Blockhash string `json:"blockhash"`

	// The type name of the filter
	// Default: "basic"
	FilterType string `json:"filtertype,omitempty"`
}

// GetBlockFilterResp holds the response to the GetBlockFilter call.
//  {                      (json object)
//    "filter" : "hex",    (string) the hex-encoded filter data
//    "header" : "hex"     (string) the hex-encoded filter header
//  }
type GetBlockFilterResp struct {
	// the hex-encoded filter data
	Filter string `json:"filter"`

	// the hex-encoded filter header
	Header string `json:"header"`
}

// GetBlockFilter RPC method.
// Retrieve a BIP 157 content filter for a particular block.
func (bc *BitcoindClient) GetBlockFilter(ctx context.Context, args GetBlockFilterReq) (result GetBlockFilterResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "getblockfilter", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// GetBlockFromPeerReq holds the arguments for the GetBlockFromPeer call.
//  1. blockhash    (string, required) The block hash to try to fetch
//  2. peer_id      (numeric, required) The peer to fetch it from (see getpeerinfo for peer IDs)
type GetBlockFromPeerReq struct {
	// The block hash to try to fetch
	Blockhash string `json:"blockhash"`

	// The peer to fetch it from (see getpeerinfo for peer IDs)
	PeerID float64 `json:"peer_id"`
}

// GetBlockFromPeer RPC method.
// Attempt to fetch block from a given peer.
// We must have the header for this block, e.g. using submitheader.
// Subsequent calls for the same block and a new peer will cause the response from the previous peer to be ignored.
// Returns an empty JSON object if the request was successfully scheduled.
func (bc *BitcoindClient) GetBlockFromPeer(ctx context.Context, args GetBlockFromPeerReq) (err error) {
	_, err = bc.sendRequest(ctx, "getblockfrompeer", args)
	return
}

// GetBlockhashReq holds the arguments for the GetBlockhash call.
//  1. height    (numeric, required) The height index
type GetBlockhashReq struct {
	// The height index
	Height float64 `json:"height"`
}

// GetBlockhashResp holds the response to the GetBlockhash call.
//  "hex"    (string) The block hash
type GetBlockhashResp struct {
	// The block hash
	Hex string
}

func (alts GetBlockhashResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.Hex)
}

func (alts *GetBlockhashResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Hex) == nil {
		return nil
	}
	alts.Hex = reset.Hex
	return &UnmarshalError{B: b, structName: "GetBlockhashResp"}
}

// GetBlockhash RPC method.
// Returns hash of block in best-block-chain at height provided.
func (bc *BitcoindClient) GetBlockhash(ctx context.Context, args GetBlockhashReq) (result GetBlockhashResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "getblockhash", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// GetBlockHeaderReq holds the arguments for the GetBlockHeader call.
//  1. blockhash    (string, required) The block hash
//  2. verbose      (boolean, optional, default=true) true for a json object, false for the hex-encoded data
type GetBlockHeaderReq struct {
	// The block hash
	Blockhash string `json:"blockhash"`

	// true for a json object, false for the hex-encoded data
	// Default: true
	Verbose *bool `json:"verbose,omitempty"`
}

// GetBlockHeaderResp holds the response to the GetBlockHeader call.
//
// ALTERNATIVE (for verbose = true)
//  {                                 (json object)
//    "hash" : "hex",                 (string) the block hash (same as provided)
//    "confirmations" : n,            (numeric) The number of confirmations, or -1 if the block is not on the main chain
//    "height" : n,                   (numeric) The block height or index
//    "version" : n,                  (numeric) The block version
//    "versionHex" : "hex",           (string) The block version formatted in hexadecimal
//    "merkleroot" : "hex",           (string) The merkle root
//    "time" : xxx,                   (numeric) The block time expressed in UNIX epoch time
//    "mediantime" : xxx,             (numeric) The median block time expressed in UNIX epoch time
//    "nonce" : n,                    (numeric) The nonce
//    "bits" : "hex",                 (string) The bits
//    "difficulty" : n,               (numeric) The difficulty
//    "chainwork" : "hex",            (string) Expected number of hashes required to produce the current chain
//    "nTx" : n,                      (numeric) The number of transactions in the block
//    "previousblockhash" : "hex",    (string, optional) The hash of the previous block (if available)
//    "nextblockhash" : "hex"         (string, optional) The hash of the next block (if available)
//  }
//
// ALTERNATIVE (for verbose=false)
//  "hex"    (string) A string that is serialized, hex-encoded data for block 'hash'
type GetBlockHeaderResp struct {
	ForVerboseEqualsTrue GetBlockHeaderRespForVerboseEqualsTrue

	// A string that is serialized, hex-encoded data for block 'hash'
	Hex string
}

func (alts GetBlockHeaderResp) MarshalJSON() ([]byte, error) {
	if !reflect.ValueOf(alts.ForVerboseEqualsTrue).IsZero() {
		return json.Marshal(alts.ForVerboseEqualsTrue)
	}
	return json.Marshal(alts.Hex)
}

func (alts *GetBlockHeaderResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.ForVerboseEqualsTrue) == nil {
		return nil
	}
	alts.ForVerboseEqualsTrue = reset.ForVerboseEqualsTrue
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Hex) == nil {
		return nil
	}
	alts.Hex = reset.Hex
	return &UnmarshalError{B: b, structName: "GetBlockHeaderResp"}
}

type GetBlockHeaderRespForVerboseEqualsTrue struct {
	// the block hash (same as provided)
	Hash string `json:"hash"`

	// The number of confirmations, or -1 if the block is not on the main chain
	Confirmations float64 `json:"confirmations"`

	// The block height or index
	Height float64 `json:"height"`

	// The block version
	Version float64 `json:"version"`

	// The block version formatted in hexadecimal
	VersionHex string `json:"versionHex"`

	// The merkle root
	MerkleRoot string `json:"merkleroot"`

	// The block time expressed in UNIX epoch time
	Time float64 `json:"time"`

	// The median block time expressed in UNIX epoch time
	MedianTime float64 `json:"mediantime"`

	// The nonce
	Nonce float64 `json:"nonce"`

	// The bits
	Bits string `json:"bits"`

	// The difficulty
	Difficulty float64 `json:"difficulty"`

	// Expected number of hashes required to produce the current chain
	ChainWork string `json:"chainwork"`

	// The number of transactions in the block
	NTx float64 `json:"nTx"`

	// The hash of the previous block (if available)
	PreviousBlockhash string `json:"previousblockhash,omitempty"`

	// The hash of the next block (if available)
	NextBlockhash string `json:"nextblockhash,omitempty"`
}

// GetBlockHeader RPC method.
// If verbose is false, returns a string that is serialized, hex-encoded data for blockheader 'hash'.
// If verbose is true, returns an Object with information about blockheader <hash>.
func (bc *BitcoindClient) GetBlockHeader(ctx context.Context, args GetBlockHeaderReq) (result GetBlockHeaderResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "getblockheader", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// GetBlockStatsReq holds the arguments for the GetBlockStats call.
//  1. hash_or_height    (string or numeric, required) The block hash or height of the target block
//  2. stats             (json array, optional, default=all values) Values to plot (see result below)
//       [
//         "height",     (string) Selected statistic
//         "time",       (string) Selected statistic
//         ...
//       ]
type GetBlockStatsReq struct {
	// The block hash or height of the target block
	HashOrHeight string `json:"hash_or_height"`

	// Values to plot (see result below)
	// Default: all values
	// Element: height    Selected statistic
	// Element: time      Selected statistic
	Stats []string `json:"stats,omitempty"`
}

// GetBlockStatsResp holds the response to the GetBlockStats call.
//  {                              (json object)
//    "avgfee" : n,                (numeric, optional) Average fee in the block
//    "avgfeerate" : n,            (numeric, optional) Average feerate (in satoshis per virtual byte)
//    "avgtxsize" : n,             (numeric, optional) Average transaction size
//    "blockhash" : "hex",         (string, optional) The block hash (to check for potential reorgs)
//    "feerate_percentiles" : [    (json array, optional) Feerates at the 10th, 25th, 50th, 75th, and 90th percentile weight unit (in satoshis per virtual byte)
//      n,                         (numeric) The 10th percentile feerate
//      n,                         (numeric) The 25th percentile feerate
//      n,                         (numeric) The 50th percentile feerate
//      n,                         (numeric) The 75th percentile feerate
//      n                          (numeric) The 90th percentile feerate
//    ],
//    "height" : n,                (numeric, optional) The height of the block
//    "ins" : n,                   (numeric, optional) The number of inputs (excluding coinbase)
//    "maxfee" : n,                (numeric, optional) Maximum fee in the block
//    "maxfeerate" : n,            (numeric, optional) Maximum feerate (in satoshis per virtual byte)
//    "maxtxsize" : n,             (numeric, optional) Maximum transaction size
//    "medianfee" : n,             (numeric, optional) Truncated median fee in the block
//    "mediantime" : n,            (numeric, optional) The block median time past
//    "mediantxsize" : n,          (numeric, optional) Truncated median transaction size
//    "minfee" : n,                (numeric, optional) Minimum fee in the block
//    "minfeerate" : n,            (numeric, optional) Minimum feerate (in satoshis per virtual byte)
//    "mintxsize" : n,             (numeric, optional) Minimum transaction size
//    "outs" : n,                  (numeric, optional) The number of outputs
//    "subsidy" : n,               (numeric, optional) The block subsidy
//    "swtotal_size" : n,          (numeric, optional) Total size of all segwit transactions
//    "swtotal_weight" : n,        (numeric, optional) Total weight of all segwit transactions
//    "swtxs" : n,                 (numeric, optional) The number of segwit transactions
//    "time" : n,                  (numeric, optional) The block time
//    "total_out" : n,             (numeric, optional) Total amount in all outputs (excluding coinbase and thus reward [ie subsidy + totalfee])
//    "total_size" : n,            (numeric, optional) Total size of all non-coinbase transactions
//    "total_weight" : n,          (numeric, optional) Total weight of all non-coinbase transactions
//    "totalfee" : n,              (numeric, optional) The fee total
//    "txs" : n,                   (numeric, optional) The number of transactions (including coinbase)
//    "utxo_increase" : n,         (numeric, optional) The increase/decrease in the number of unspent outputs
//    "utxo_size_inc" : n          (numeric, optional) The increase/decrease in size for the utxo index (not discounting op_return and similar)
//  }
type GetBlockStatsResp struct {
	// Average fee in the block
	AvgFee *float64 `json:"avgfee,omitempty"`

	// Average feerate (in satoshis per virtual byte)
	AvgFeeRate *float64 `json:"avgfeerate,omitempty"`

	// Average transaction size
	AvgTxSize *float64 `json:"avgtxsize,omitempty"`

	// The block hash (to check for potential reorgs)
	Blockhash string `json:"blockhash,omitempty"`

	// Feerates at the 10th, 25th, 50th, 75th, and 90th percentile weight unit (in satoshis per virtual byte)
	// Element: n    The 10th percentile feerate
	// Element: n    The 25th percentile feerate
	// Element: n    The 50th percentile feerate
	// Element: n    The 75th percentile feerate
	// n                          (numeric) The 90th percentile feerate
	FeeRatePercentiles []float64 `json:"feerate_percentiles,omitempty"`

	// The height of the block
	Height *float64 `json:"height,omitempty"`

	// The number of inputs (excluding coinbase)
	Ins *float64 `json:"ins,omitempty"`

	// Maximum fee in the block
	MaxFee *float64 `json:"maxfee,omitempty"`

	// Maximum feerate (in satoshis per virtual byte)
	MaxFeeRate *float64 `json:"maxfeerate,omitempty"`

	// Maximum transaction size
	MaxTxSize *float64 `json:"maxtxsize,omitempty"`

	// Truncated median fee in the block
	MedianFee *float64 `json:"medianfee,omitempty"`

	// The block median time past
	MedianTime *float64 `json:"mediantime,omitempty"`

	// Truncated median transaction size
	MedianTxSize *float64 `json:"mediantxsize,omitempty"`

	// Minimum fee in the block
	MinFee *float64 `json:"minfee,omitempty"`

	// Minimum feerate (in satoshis per virtual byte)
	MinFeeRate *float64 `json:"minfeerate,omitempty"`

	// Minimum transaction size
	MinTxSize *float64 `json:"mintxsize,omitempty"`

	// The number of outputs
	Outs *float64 `json:"outs,omitempty"`

	// The block subsidy
	Subsidy *float64 `json:"subsidy,omitempty"`

	// Total size of all segwit transactions
	SwTotalSize *float64 `json:"swtotal_size,omitempty"`

	// Total weight of all segwit transactions
	SwTotalWeight *float64 `json:"swtotal_weight,omitempty"`

	// The number of segwit transactions
	SwTxs *float64 `json:"swtxs,omitempty"`

	// The block time
	Time *float64 `json:"time,omitempty"`

	// Total amount in all outputs (excluding coinbase and thus reward [ie subsidy + totalfee])
	TotalOut *float64 `json:"total_out,omitempty"`

	// Total size of all non-coinbase transactions
	TotalSize *float64 `json:"total_size,omitempty"`

	// Total weight of all non-coinbase transactions
	TotalWeight *float64 `json:"total_weight,omitempty"`

	// The fee total
	TotalFee *float64 `json:"totalfee,omitempty"`

	// The number of transactions (including coinbase)
	Txs *float64 `json:"txs,omitempty"`

	// The increase/decrease in the number of unspent outputs
	UtxoIncrease *float64 `json:"utxo_increase,omitempty"`

	// The increase/decrease in size for the utxo index (not discounting op_return and similar)
	UtxoSizeInc *float64 `json:"utxo_size_inc,omitempty"`
}

// GetBlockStats RPC method.
// Compute per block statistics for a given window. All amounts are in satoshis.
// It won't work for some heights with pruning.
func (bc *BitcoindClient) GetBlockStats(ctx context.Context, args GetBlockStatsReq) (result GetBlockStatsResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "getblockstats", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// GetChainTipsResp holds the response to the GetChainTips call.
//  [                        (json array)
//    {                      (json object)
//      "height" : n,        (numeric) height of the chain tip
//      "hash" : "hex",      (string) block hash of the tip
//      "branchlen" : n,     (numeric) zero for main chain, otherwise length of branch connecting the tip to the main chain
//      "status" : "str"     (string) status of the chain, "active" for the main chain
//                           Possible values for status:
//                           1.  "invalid"               This branch contains at least one invalid block
//                           2.  "headers-only"          Not all blocks for this branch are available, but the headers are valid
//                           3.  "valid-headers"         All blocks are available for this branch, but they were never fully validated
//                           4.  "valid-fork"            This branch is not part of the active chain, but is fully validated
//                           5.  "active"                This is the tip of the active main chain, which is certainly valid
//    },
//    ...
//  ]
type GetChainTipsResp struct {
	Array []GetChainTipsRespElement
}

func (alts GetChainTipsResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.Array)
}

func (alts *GetChainTipsResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Array) == nil {
		return nil
	}
	alts.Array = reset.Array
	return &UnmarshalError{B: b, structName: "GetChainTipsResp"}
}

type GetChainTipsRespElement struct {
	// height of the chain tip
	Height float64 `json:"height"`

	// block hash of the tip
	Hash string `json:"hash"`

	// zero for main chain, otherwise length of branch connecting the tip to the main chain
	BranchLen float64 `json:"branchlen"`

	// status of the chain, "active" for the main chain
	// Possible values for status:
	// 1.  "invalid"               This branch contains at least one invalid block
	// 2.  "headers-only"          Not all blocks for this branch are available, but the headers are valid
	// 3.  "valid-headers"         All blocks are available for this branch, but they were never fully validated
	// 4.  "valid-fork"            This branch is not part of the active chain, but is fully validated
	// 5.  "active"                This is the tip of the active main chain, which is certainly valid
	Status string `json:"status"`
}

// GetChainTips RPC method.
// Return information about all known tips in the block tree, including the main chain as well as orphaned branches.
func (bc *BitcoindClient) GetChainTips(ctx context.Context) (result GetChainTipsResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "getchaintips", nil); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// GetChainTxStatsReq holds the arguments for the GetChainTxStats call.
//  1. nblocks      (numeric, optional, default=one month) Size of the window in number of blocks
//  2. blockhash    (string, optional, default=chain tip) The hash of the block that ends the window.
type GetChainTxStatsReq struct {
	// Size of the window in number of blocks
	// Default: one month
	NBlocks *float64 `json:"nblocks,omitempty"`

	// The hash of the block that ends the window.
	// Default: chain tip
	Blockhash string `json:"blockhash,omitempty"`
}

// GetChainTxStatsResp holds the response to the GetChainTxStats call.
//  {                                       (json object)
//    "time" : xxx,                         (numeric) The timestamp for the final block in the window, expressed in UNIX epoch time
//    "txcount" : n,                        (numeric) The total number of transactions in the chain up to that point
//    "window_final_block_hash" : "hex",    (string) The hash of the final block in the window
//    "window_final_block_height" : n,      (numeric) The height of the final block in the window.
//    "window_block_count" : n,             (numeric) Size of the window in number of blocks
//    "window_tx_count" : n,                (numeric, optional) The number of transactions in the window. Only returned if "window_block_count" is > 0
//    "window_interval" : n,                (numeric, optional) The elapsed time in the window in seconds. Only returned if "window_block_count" is > 0
//    "txrate" : n                          (numeric, optional) The average rate of transactions per second in the window. Only returned if "window_interval" is > 0
//  }
type GetChainTxStatsResp struct {
	// The timestamp for the final block in the window, expressed in UNIX epoch time
	Time float64 `json:"time"`

	// The total number of transactions in the chain up to that point
	TxCount float64 `json:"txcount"`

	// The hash of the final block in the window
	WindowFinalBlockhash string `json:"window_final_block_hash"`

	// The height of the final block in the window.
	WindowFinalBlockHeight float64 `json:"window_final_block_height"`

	// Size of the window in number of blocks
	WindowBlockCount float64 `json:"window_block_count"`

	// The number of transactions in the window. Only returned if "window_block_count" is > 0
	WindowTxCount *float64 `json:"window_tx_count,omitempty"`

	// The elapsed time in the window in seconds. Only returned if "window_block_count" is > 0
	WindowInterval *float64 `json:"window_interval,omitempty"`

	// The average rate of transactions per second in the window. Only returned if "window_interval" is > 0
	TxRate *float64 `json:"txrate,omitempty"`
}

// GetChainTxStats RPC method.
// Compute statistics about the total number and rate of transactions in the chain.
func (bc *BitcoindClient) GetChainTxStats(ctx context.Context, args GetChainTxStatsReq) (result GetChainTxStatsResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "getchaintxstats", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// GetDeploymentInfoReq holds the arguments for the GetDeploymentInfo call.
//  1. blockhash    (string, optional, default="hash of current chain tip") The block hash at which to query deployment state
type GetDeploymentInfoReq struct {
	// The block hash at which to query deployment state
	// Default: "hash of current chain tip"
	Blockhash string `json:"blockhash,omitempty"`
}

// GetDeploymentInfoResp holds the response to the GetDeploymentInfo call.
//  {                                       (json object)
//    "hash" : "str",                       (string) requested block hash (or tip)
//    "height" : n,                         (numeric) requested block height (or tip)
//    "deployments" : {                     (json object)
//      "xxxx" : {                          (json object) name of the deployment
//        "type" : "str",                   (string) one of "buried", "bip9"
//        "height" : n,                     (numeric, optional) height of the first block which the rules are or will be enforced (only for "buried" type, or "bip9" type with "active" status)
//        "active" : true|false,            (boolean) true if the rules are enforced for the mempool and the next block
//        "bip9" : {                        (json object, optional) status of bip9 softforks (only for "bip9" type)
//          "bit" : n,                      (numeric, optional) the bit (0-28) in the block version field used to signal this softfork (only for "started" and "locked_in" status)
//          "start_time" : xxx,             (numeric) the minimum median time past of a block at which the bit gains its meaning
//          "timeout" : xxx,                (numeric) the median time past of a block at which the deployment is considered failed if not yet locked in
//          "min_activation_height" : n,    (numeric) minimum height of blocks for which the rules may be enforced
//          "status" : "str",               (string) status of deployment at specified block (one of "defined", "started", "locked_in", "active", "failed")
//          "since" : n,                    (numeric) height of the first block to which the status applies
//          "status_next" : "str",          (string) status of deployment at the next block
//          "statistics" : {                (json object, optional) numeric statistics about signalling for a softfork (only for "started" and "locked_in" status)
//            "period" : n,                 (numeric) the length in blocks of the signalling period
//            "threshold" : n,              (numeric, optional) the number of blocks with the version bit set required to activate the feature (only for "started" status)
//            "elapsed" : n,                (numeric) the number of blocks elapsed since the beginning of the current period
//            "count" : n,                  (numeric) the number of blocks with the version bit set in the current period
//            "possible" : true|false       (boolean, optional) returns false if there are not enough blocks left in this period to pass activation threshold (only for "started" status)
//          },
//          "signalling" : "str"            (string) indicates blocks that signalled with a # and blocks that did not with a -
//        }
//      }
//    }
//  }
type GetDeploymentInfoResp struct {
	// requested block hash (or tip)
	Hash string `json:"hash"`

	// requested block height (or tip)
	Height float64 `json:"height"`

	// name of the deployment
	// Key: xxxx, Value: struct
	Deployments map[string]GetDeploymentInfoRespDeployments `json:"deployments"`
}

type GetDeploymentInfoRespDeployments struct {
	// one of "buried", "bip9"
	Type string `json:"type"`

	// height of the first block which the rules are or will be enforced (only for "buried" type, or "bip9" type with "active" status)
	Height *float64 `json:"height,omitempty"`

	// true if the rules are enforced for the mempool and the next block
	Active bool `json:"active"`

	// status of bip9 softforks (only for "bip9" type)
	BIP9 *GetDeploymentInfoRespDeploymentsBIP9 `json:"bip9,omitempty"`
}

type GetDeploymentInfoRespDeploymentsBIP9 struct {
	// the bit (0-28) in the block version field used to signal this softfork (only for "started" and "locked_in" status)
	Bit *float64 `json:"bit,omitempty"`

	// the minimum median time past of a block at which the bit gains its meaning
	StartTime float64 `json:"start_time"`

	// the median time past of a block at which the deployment is considered failed if not yet locked in
	TimeOut float64 `json:"timeout"`

	// minimum height of blocks for which the rules may be enforced
	MinActivationHeight float64 `json:"min_activation_height"`

	// status of deployment at specified block (one of "defined", "started", "locked_in", "active", "failed")
	Status string `json:"status"`

	// height of the first block to which the status applies
	Since float64 `json:"since"`

	// status of deployment at the next block
	StatusNext string `json:"status_next"`

	// numeric statistics about signalling for a softfork (only for "started" and "locked_in" status)
	Statistics *GetDeploymentInfoRespDeploymentsBIP9Statistics `json:"statistics,omitempty"`

	// indicates blocks that signalled with a # and blocks that did not with a -
	Signalling string `json:"signalling"`
}

type GetDeploymentInfoRespDeploymentsBIP9Statistics struct {
	// the length in blocks of the signalling period
	Period float64 `json:"period"`

	// the number of blocks with the version bit set required to activate the feature (only for "started" status)
	Threshold *float64 `json:"threshold,omitempty"`

	// the number of blocks elapsed since the beginning of the current period
	Elapsed float64 `json:"elapsed"`

	// the number of blocks with the version bit set in the current period
	Count float64 `json:"count"`

	// returns false if there are not enough blocks left in this period to pass activation threshold (only for "started" status)
	Possible *bool `json:"possible,omitempty"`
}

// GetDeploymentInfo RPC method.
// Returns an object containing various state info regarding deployments of consensus changes.
func (bc *BitcoindClient) GetDeploymentInfo(ctx context.Context, args GetDeploymentInfoReq) (result GetDeploymentInfoResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "getdeploymentinfo", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// GetDifficultyResp holds the response to the GetDifficulty call.
//  n    (numeric) the proof-of-work difficulty as a multiple of the minimum difficulty.
type GetDifficultyResp struct {
	// the proof-of-work difficulty as a multiple of the minimum difficulty.
	N float64
}

func (alts GetDifficultyResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.N)
}

func (alts *GetDifficultyResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.N) == nil {
		return nil
	}
	alts.N = reset.N
	return &UnmarshalError{B: b, structName: "GetDifficultyResp"}
}

// GetDifficulty RPC method.
// Returns the proof-of-work difficulty as a multiple of the minimum difficulty.
func (bc *BitcoindClient) GetDifficulty(ctx context.Context) (result GetDifficultyResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "getdifficulty", nil); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// GetMempoolAncestorsReq holds the arguments for the GetMempoolAncestors call.
//  1. txid       (string, required) The transaction id (must be in mempool)
//  2. verbose    (boolean, optional, default=false) True for a json object, false for array of transaction ids
type GetMempoolAncestorsReq struct {
	// The transaction id (must be in mempool)
	TxID string `json:"txid"`

	// True for a json object, false for array of transaction ids
	// Default: false
	Verbose bool `json:"verbose,omitempty"`
}

// GetMempoolAncestorsResp holds the response to the GetMempoolAncestors call.
//
// ALTERNATIVE (for verbose = false)
//  [           (json array)
//    "hex",    (string) The transaction id of an in-mempool ancestor transaction
//    ...
//  ]
//
// ALTERNATIVE (for verbose = true)
//  {                                         (json object)
//    "transactionid" : {                     (json object)
//      "vsize" : n,                          (numeric) virtual transaction size as defined in BIP 141. This is different from actual serialized size for witness transactions as witness data is discounted.
//      "weight" : n,                         (numeric) transaction weight as defined in BIP 141.
//      "fee" : n,                            (numeric, optional) transaction fee, denominated in BTC (DEPRECATED, returned only if config option -deprecatedrpc=fees is passed)
//      "modifiedfee" : n,                    (numeric, optional) transaction fee with fee deltas used for mining priority, denominated in BTC (DEPRECATED, returned only if config option -deprecatedrpc=fees is passed)
//      "time" : xxx,                         (numeric) local time transaction entered pool in seconds since 1 Jan 1970 GMT
//      "height" : n,                         (numeric) block height when transaction entered pool
//      "descendantcount" : n,                (numeric) number of in-mempool descendant transactions (including this one)
//      "descendantsize" : n,                 (numeric) virtual transaction size of in-mempool descendants (including this one)
//      "descendantfees" : n,                 (numeric, optional) transaction fees of in-mempool descendants (including this one) with fee deltas used for mining priority, denominated in sats (DEPRECATED, returned only if config option -deprecatedrpc=fees is passed)
//      "ancestorcount" : n,                  (numeric) number of in-mempool ancestor transactions (including this one)
//      "ancestorsize" : n,                   (numeric) virtual transaction size of in-mempool ancestors (including this one)
//      "ancestorfees" : n,                   (numeric, optional) transaction fees of in-mempool ancestors (including this one) with fee deltas used for mining priority, denominated in sats (DEPRECATED, returned only if config option -deprecatedrpc=fees is passed)
//      "wtxid" : "hex",                      (string) hash of serialized transaction, including witness data
//      "fees" : {                            (json object)
//        "base" : n,                         (numeric) transaction fee, denominated in BTC
//        "modified" : n,                     (numeric) transaction fee with fee deltas used for mining priority, denominated in BTC
//        "ancestor" : n,                     (numeric) transaction fees of in-mempool ancestors (including this one) with fee deltas used for mining priority, denominated in BTC
//        "descendant" : n                    (numeric) transaction fees of in-mempool descendants (including this one) with fee deltas used for mining priority, denominated in BTC
//      },
//      "depends" : [                         (json array) unconfirmed transactions used as inputs for this transaction
//        "hex",                              (string) parent transaction id
//        ...
//      ],
//      "spentby" : [                         (json array) unconfirmed transactions spending outputs from this transaction
//        "hex",                              (string) child transaction id
//        ...
//      ],
//      "bip125-replaceable" : true|false,    (boolean) Whether this transaction could be replaced due to BIP125 (replace-by-fee)
//      "unbroadcast" : true|false            (boolean) Whether this transaction is currently unbroadcast (initial broadcast not yet acknowledged by any peers)
//    },
//    ...
//  }
type GetMempoolAncestorsResp struct {
	// Element: Hex    The transaction id of an in-mempool ancestor transaction
	Hex []string

	// Key: transactionid, Value: struct
	ForVerboseEqualsTrue map[string]GetMempoolAncestorsRespForVerboseEqualsTrue
}

func (alts GetMempoolAncestorsResp) MarshalJSON() ([]byte, error) {
	if !reflect.ValueOf(alts.Hex).IsZero() {
		return json.Marshal(alts.Hex)
	}
	return json.Marshal(alts.ForVerboseEqualsTrue)
}

func (alts *GetMempoolAncestorsResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Hex) == nil {
		return nil
	}
	alts.Hex = reset.Hex
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.ForVerboseEqualsTrue) == nil {
		return nil
	}
	alts.ForVerboseEqualsTrue = reset.ForVerboseEqualsTrue
	return &UnmarshalError{B: b, structName: "GetMempoolAncestorsResp"}
}

type GetMempoolAncestorsRespForVerboseEqualsTrue struct {
	// virtual transaction size as defined in BIP 141. This is different from actual serialized size for witness transactions as witness data is discounted.
	VSize float64 `json:"vsize"`

	// transaction weight as defined in BIP 141.
	Weight float64 `json:"weight"`

	// transaction fee, denominated in BTC (DEPRECATED, returned only if config option -deprecatedrpc=fees is passed)
	Fee *float64 `json:"fee,omitempty"`

	// transaction fee with fee deltas used for mining priority, denominated in BTC (DEPRECATED, returned only if config option -deprecatedrpc=fees is passed)
	ModifiedFee *float64 `json:"modifiedfee,omitempty"`

	// local time transaction entered pool in seconds since 1 Jan 1970 GMT
	Time float64 `json:"time"`

	// block height when transaction entered pool
	Height float64 `json:"height"`

	// number of in-mempool descendant transactions (including this one)
	DescendantCount float64 `json:"descendantcount"`

	// virtual transaction size of in-mempool descendants (including this one)
	DescendantSize float64 `json:"descendantsize"`

	// transaction fees of in-mempool descendants (including this one) with fee deltas used for mining priority, denominated in sats (DEPRECATED, returned only if config option -deprecatedrpc=fees is passed)
	DescendantFees *float64 `json:"descendantfees,omitempty"`

	// number of in-mempool ancestor transactions (including this one)
	AncestorCount float64 `json:"ancestorcount"`

	// virtual transaction size of in-mempool ancestors (including this one)
	AncestorSize float64 `json:"ancestorsize"`

	// transaction fees of in-mempool ancestors (including this one) with fee deltas used for mining priority, denominated in sats (DEPRECATED, returned only if config option -deprecatedrpc=fees is passed)
	AncestorFees *float64 `json:"ancestorfees,omitempty"`

	// hash of serialized transaction, including witness data
	WTxID string `json:"wtxid"`

	Fees struct {
		// transaction fee, denominated in BTC
		Base float64 `json:"base"`

		// transaction fee with fee deltas used for mining priority, denominated in BTC
		Modified float64 `json:"modified"`

		// transaction fees of in-mempool ancestors (including this one) with fee deltas used for mining priority, denominated in BTC
		Ancestor float64 `json:"ancestor"`

		// transaction fees of in-mempool descendants (including this one) with fee deltas used for mining priority, denominated in BTC
		Descendant float64 `json:"descendant"`
	} `json:"fees"`

	// unconfirmed transactions used as inputs for this transaction
	// Element: Hex    parent transaction id
	Depends []string `json:"depends"`

	// unconfirmed transactions spending outputs from this transaction
	// Element: Hex    child transaction id
	SpentBy []string `json:"spentby"`

	// Whether this transaction could be replaced due to BIP125 (replace-by-fee)
	BIP125Replaceable bool `json:"bip125-replaceable"`

	// Whether this transaction is currently unbroadcast (initial broadcast not yet acknowledged by any peers)
	Unbroadcast bool `json:"unbroadcast"`
}

// GetMempoolAncestors RPC method.
// If txid is in the mempool, returns all in-mempool ancestors.
func (bc *BitcoindClient) GetMempoolAncestors(ctx context.Context, args GetMempoolAncestorsReq) (result GetMempoolAncestorsResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "getmempoolancestors", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// GetMempoolDescendantsReq holds the arguments for the GetMempoolDescendants call.
//  1. txid       (string, required) The transaction id (must be in mempool)
//  2. verbose    (boolean, optional, default=false) True for a json object, false for array of transaction ids
type GetMempoolDescendantsReq struct {
	// The transaction id (must be in mempool)
	TxID string `json:"txid"`

	// True for a json object, false for array of transaction ids
	// Default: false
	Verbose bool `json:"verbose,omitempty"`
}

// GetMempoolDescendantsResp holds the response to the GetMempoolDescendants call.
//
// ALTERNATIVE (for verbose = false)
//  [           (json array)
//    "hex",    (string) The transaction id of an in-mempool descendant transaction
//    ...
//  ]
//
// ALTERNATIVE (for verbose = true)
//  {                                         (json object)
//    "transactionid" : {                     (json object)
//      "vsize" : n,                          (numeric) virtual transaction size as defined in BIP 141. This is different from actual serialized size for witness transactions as witness data is discounted.
//      "weight" : n,                         (numeric) transaction weight as defined in BIP 141.
//      "fee" : n,                            (numeric, optional) transaction fee, denominated in BTC (DEPRECATED, returned only if config option -deprecatedrpc=fees is passed)
//      "modifiedfee" : n,                    (numeric, optional) transaction fee with fee deltas used for mining priority, denominated in BTC (DEPRECATED, returned only if config option -deprecatedrpc=fees is passed)
//      "time" : xxx,                         (numeric) local time transaction entered pool in seconds since 1 Jan 1970 GMT
//      "height" : n,                         (numeric) block height when transaction entered pool
//      "descendantcount" : n,                (numeric) number of in-mempool descendant transactions (including this one)
//      "descendantsize" : n,                 (numeric) virtual transaction size of in-mempool descendants (including this one)
//      "descendantfees" : n,                 (numeric, optional) transaction fees of in-mempool descendants (including this one) with fee deltas used for mining priority, denominated in sats (DEPRECATED, returned only if config option -deprecatedrpc=fees is passed)
//      "ancestorcount" : n,                  (numeric) number of in-mempool ancestor transactions (including this one)
//      "ancestorsize" : n,                   (numeric) virtual transaction size of in-mempool ancestors (including this one)
//      "ancestorfees" : n,                   (numeric, optional) transaction fees of in-mempool ancestors (including this one) with fee deltas used for mining priority, denominated in sats (DEPRECATED, returned only if config option -deprecatedrpc=fees is passed)
//      "wtxid" : "hex",                      (string) hash of serialized transaction, including witness data
//      "fees" : {                            (json object)
//        "base" : n,                         (numeric) transaction fee, denominated in BTC
//        "modified" : n,                     (numeric) transaction fee with fee deltas used for mining priority, denominated in BTC
//        "ancestor" : n,                     (numeric) transaction fees of in-mempool ancestors (including this one) with fee deltas used for mining priority, denominated in BTC
//        "descendant" : n                    (numeric) transaction fees of in-mempool descendants (including this one) with fee deltas used for mining priority, denominated in BTC
//      },
//      "depends" : [                         (json array) unconfirmed transactions used as inputs for this transaction
//        "hex",                              (string) parent transaction id
//        ...
//      ],
//      "spentby" : [                         (json array) unconfirmed transactions spending outputs from this transaction
//        "hex",                              (string) child transaction id
//        ...
//      ],
//      "bip125-replaceable" : true|false,    (boolean) Whether this transaction could be replaced due to BIP125 (replace-by-fee)
//      "unbroadcast" : true|false            (boolean) Whether this transaction is currently unbroadcast (initial broadcast not yet acknowledged by any peers)
//    },
//    ...
//  }
type GetMempoolDescendantsResp struct {
	// Element: Hex    The transaction id of an in-mempool descendant transaction
	Hex []string

	// Key: transactionid, Value: struct
	ForVerboseEqualsTrue map[string]GetMempoolDescendantsRespForVerboseEqualsTrue
}

func (alts GetMempoolDescendantsResp) MarshalJSON() ([]byte, error) {
	if !reflect.ValueOf(alts.Hex).IsZero() {
		return json.Marshal(alts.Hex)
	}
	return json.Marshal(alts.ForVerboseEqualsTrue)
}

func (alts *GetMempoolDescendantsResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Hex) == nil {
		return nil
	}
	alts.Hex = reset.Hex
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.ForVerboseEqualsTrue) == nil {
		return nil
	}
	alts.ForVerboseEqualsTrue = reset.ForVerboseEqualsTrue
	return &UnmarshalError{B: b, structName: "GetMempoolDescendantsResp"}
}

type GetMempoolDescendantsRespForVerboseEqualsTrue struct {
	// virtual transaction size as defined in BIP 141. This is different from actual serialized size for witness transactions as witness data is discounted.
	VSize float64 `json:"vsize"`

	// transaction weight as defined in BIP 141.
	Weight float64 `json:"weight"`

	// transaction fee, denominated in BTC (DEPRECATED, returned only if config option -deprecatedrpc=fees is passed)
	Fee *float64 `json:"fee,omitempty"`

	// transaction fee with fee deltas used for mining priority, denominated in BTC (DEPRECATED, returned only if config option -deprecatedrpc=fees is passed)
	ModifiedFee *float64 `json:"modifiedfee,omitempty"`

	// local time transaction entered pool in seconds since 1 Jan 1970 GMT
	Time float64 `json:"time"`

	// block height when transaction entered pool
	Height float64 `json:"height"`

	// number of in-mempool descendant transactions (including this one)
	DescendantCount float64 `json:"descendantcount"`

	// virtual transaction size of in-mempool descendants (including this one)
	DescendantSize float64 `json:"descendantsize"`

	// transaction fees of in-mempool descendants (including this one) with fee deltas used for mining priority, denominated in sats (DEPRECATED, returned only if config option -deprecatedrpc=fees is passed)
	DescendantFees *float64 `json:"descendantfees,omitempty"`

	// number of in-mempool ancestor transactions (including this one)
	AncestorCount float64 `json:"ancestorcount"`

	// virtual transaction size of in-mempool ancestors (including this one)
	AncestorSize float64 `json:"ancestorsize"`

	// transaction fees of in-mempool ancestors (including this one) with fee deltas used for mining priority, denominated in sats (DEPRECATED, returned only if config option -deprecatedrpc=fees is passed)
	AncestorFees *float64 `json:"ancestorfees,omitempty"`

	// hash of serialized transaction, including witness data
	WTxID string `json:"wtxid"`

	Fees struct {
		// transaction fee, denominated in BTC
		Base float64 `json:"base"`

		// transaction fee with fee deltas used for mining priority, denominated in BTC
		Modified float64 `json:"modified"`

		// transaction fees of in-mempool ancestors (including this one) with fee deltas used for mining priority, denominated in BTC
		Ancestor float64 `json:"ancestor"`

		// transaction fees of in-mempool descendants (including this one) with fee deltas used for mining priority, denominated in BTC
		Descendant float64 `json:"descendant"`
	} `json:"fees"`

	// unconfirmed transactions used as inputs for this transaction
	// Element: Hex    parent transaction id
	Depends []string `json:"depends"`

	// unconfirmed transactions spending outputs from this transaction
	// Element: Hex    child transaction id
	SpentBy []string `json:"spentby"`

	// Whether this transaction could be replaced due to BIP125 (replace-by-fee)
	BIP125Replaceable bool `json:"bip125-replaceable"`

	// Whether this transaction is currently unbroadcast (initial broadcast not yet acknowledged by any peers)
	Unbroadcast bool `json:"unbroadcast"`
}

// GetMempoolDescendants RPC method.
// If txid is in the mempool, returns all in-mempool descendants.
func (bc *BitcoindClient) GetMempoolDescendants(ctx context.Context, args GetMempoolDescendantsReq) (result GetMempoolDescendantsResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "getmempooldescendants", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// GetMempoolEntryReq holds the arguments for the GetMempoolEntry call.
//  1. txid    (string, required) The transaction id (must be in mempool)
type GetMempoolEntryReq struct {
	// The transaction id (must be in mempool)
	TxID string `json:"txid"`
}

// GetMempoolEntryResp holds the response to the GetMempoolEntry call.
//  {                                       (json object)
//    "vsize" : n,                          (numeric) virtual transaction size as defined in BIP 141. This is different from actual serialized size for witness transactions as witness data is discounted.
//    "weight" : n,                         (numeric) transaction weight as defined in BIP 141.
//    "fee" : n,                            (numeric, optional) transaction fee, denominated in BTC (DEPRECATED, returned only if config option -deprecatedrpc=fees is passed)
//    "modifiedfee" : n,                    (numeric, optional) transaction fee with fee deltas used for mining priority, denominated in BTC (DEPRECATED, returned only if config option -deprecatedrpc=fees is passed)
//    "time" : xxx,                         (numeric) local time transaction entered pool in seconds since 1 Jan 1970 GMT
//    "height" : n,                         (numeric) block height when transaction entered pool
//    "descendantcount" : n,                (numeric) number of in-mempool descendant transactions (including this one)
//    "descendantsize" : n,                 (numeric) virtual transaction size of in-mempool descendants (including this one)
//    "descendantfees" : n,                 (numeric, optional) transaction fees of in-mempool descendants (including this one) with fee deltas used for mining priority, denominated in sats (DEPRECATED, returned only if config option -deprecatedrpc=fees is passed)
//    "ancestorcount" : n,                  (numeric) number of in-mempool ancestor transactions (including this one)
//    "ancestorsize" : n,                   (numeric) virtual transaction size of in-mempool ancestors (including this one)
//    "ancestorfees" : n,                   (numeric, optional) transaction fees of in-mempool ancestors (including this one) with fee deltas used for mining priority, denominated in sats (DEPRECATED, returned only if config option -deprecatedrpc=fees is passed)
//    "wtxid" : "hex",                      (string) hash of serialized transaction, including witness data
//    "fees" : {                            (json object)
//      "base" : n,                         (numeric) transaction fee, denominated in BTC
//      "modified" : n,                     (numeric) transaction fee with fee deltas used for mining priority, denominated in BTC
//      "ancestor" : n,                     (numeric) transaction fees of in-mempool ancestors (including this one) with fee deltas used for mining priority, denominated in BTC
//      "descendant" : n                    (numeric) transaction fees of in-mempool descendants (including this one) with fee deltas used for mining priority, denominated in BTC
//    },
//    "depends" : [                         (json array) unconfirmed transactions used as inputs for this transaction
//      "hex",                              (string) parent transaction id
//      ...
//    ],
//    "spentby" : [                         (json array) unconfirmed transactions spending outputs from this transaction
//      "hex",                              (string) child transaction id
//      ...
//    ],
//    "bip125-replaceable" : true|false,    (boolean) Whether this transaction could be replaced due to BIP125 (replace-by-fee)
//    "unbroadcast" : true|false            (boolean) Whether this transaction is currently unbroadcast (initial broadcast not yet acknowledged by any peers)
//  }
type GetMempoolEntryResp struct {
	// virtual transaction size as defined in BIP 141. This is different from actual serialized size for witness transactions as witness data is discounted.
	VSize float64 `json:"vsize"`

	// transaction weight as defined in BIP 141.
	Weight float64 `json:"weight"`

	// transaction fee, denominated in BTC (DEPRECATED, returned only if config option -deprecatedrpc=fees is passed)
	Fee *float64 `json:"fee,omitempty"`

	// transaction fee with fee deltas used for mining priority, denominated in BTC (DEPRECATED, returned only if config option -deprecatedrpc=fees is passed)
	ModifiedFee *float64 `json:"modifiedfee,omitempty"`

	// local time transaction entered pool in seconds since 1 Jan 1970 GMT
	Time float64 `json:"time"`

	// block height when transaction entered pool
	Height float64 `json:"height"`

	// number of in-mempool descendant transactions (including this one)
	DescendantCount float64 `json:"descendantcount"`

	// virtual transaction size of in-mempool descendants (including this one)
	DescendantSize float64 `json:"descendantsize"`

	// transaction fees of in-mempool descendants (including this one) with fee deltas used for mining priority, denominated in sats (DEPRECATED, returned only if config option -deprecatedrpc=fees is passed)
	DescendantFees *float64 `json:"descendantfees,omitempty"`

	// number of in-mempool ancestor transactions (including this one)
	AncestorCount float64 `json:"ancestorcount"`

	// virtual transaction size of in-mempool ancestors (including this one)
	AncestorSize float64 `json:"ancestorsize"`

	// transaction fees of in-mempool ancestors (including this one) with fee deltas used for mining priority, denominated in sats (DEPRECATED, returned only if config option -deprecatedrpc=fees is passed)
	AncestorFees *float64 `json:"ancestorfees,omitempty"`

	// hash of serialized transaction, including witness data
	WTxID string `json:"wtxid"`

	Fees struct {
		// transaction fee, denominated in BTC
		Base float64 `json:"base"`

		// transaction fee with fee deltas used for mining priority, denominated in BTC
		Modified float64 `json:"modified"`

		// transaction fees of in-mempool ancestors (including this one) with fee deltas used for mining priority, denominated in BTC
		Ancestor float64 `json:"ancestor"`

		// transaction fees of in-mempool descendants (including this one) with fee deltas used for mining priority, denominated in BTC
		Descendant float64 `json:"descendant"`
	} `json:"fees"`

	// unconfirmed transactions used as inputs for this transaction
	// Element: Hex    parent transaction id
	Depends []string `json:"depends"`

	// unconfirmed transactions spending outputs from this transaction
	// Element: Hex    child transaction id
	SpentBy []string `json:"spentby"`

	// Whether this transaction could be replaced due to BIP125 (replace-by-fee)
	BIP125Replaceable bool `json:"bip125-replaceable"`

	// Whether this transaction is currently unbroadcast (initial broadcast not yet acknowledged by any peers)
	Unbroadcast bool `json:"unbroadcast"`
}

// GetMempoolEntry RPC method.
// Returns mempool data for given transaction
func (bc *BitcoindClient) GetMempoolEntry(ctx context.Context, args GetMempoolEntryReq) (result GetMempoolEntryResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "getmempoolentry", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// GetMempoolInfoResp holds the response to the GetMempoolInfo call.
//  {                            (json object)
//    "loaded" : true|false,     (boolean) True if the mempool is fully loaded
//    "size" : n,                (numeric) Current tx count
//    "bytes" : n,               (numeric) Sum of all virtual transaction sizes as defined in BIP 141. Differs from actual serialized size because witness data is discounted
//    "usage" : n,               (numeric) Total memory usage for the mempool
//    "total_fee" : n,           (numeric) Total fees for the mempool in BTC, ignoring modified fees through prioritisetransaction
//    "maxmempool" : n,          (numeric) Maximum memory usage for the mempool
//    "mempoolminfee" : n,       (numeric) Minimum fee rate in BTC/kvB for tx to be accepted. Is the maximum of minrelaytxfee and minimum mempool fee
//    "minrelaytxfee" : n,       (numeric) Current minimum relay fee for transactions
//    "unbroadcastcount" : n     (numeric) Current number of transactions that haven't passed initial broadcast yet
//  }
type GetMempoolInfoResp struct {
	// True if the mempool is fully loaded
	Loaded bool `json:"loaded"`

	// Current tx count
	Size float64 `json:"size"`

	// Sum of all virtual transaction sizes as defined in BIP 141. Differs from actual serialized size because witness data is discounted
	Bytes float64 `json:"bytes"`

	// Total memory usage for the mempool
	Usage float64 `json:"usage"`

	// Total fees for the mempool in BTC, ignoring modified fees through prioritisetransaction
	TotalFee float64 `json:"total_fee"`

	// Maximum memory usage for the mempool
	MaxMempool float64 `json:"maxmempool"`

	// Minimum fee rate in BTC/kvB for tx to be accepted. Is the maximum of minrelaytxfee and minimum mempool fee
	MempoolMinFee float64 `json:"mempoolminfee"`

	// Current minimum relay fee for transactions
	MinRelayTxFee float64 `json:"minrelaytxfee"`

	// Current number of transactions that haven't passed initial broadcast yet
	UnbroadcastCount float64 `json:"unbroadcastcount"`
}

// GetMempoolInfo RPC method.
// Returns details on the active state of the TX memory pool.
func (bc *BitcoindClient) GetMempoolInfo(ctx context.Context) (result GetMempoolInfoResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "getmempoolinfo", nil); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// GetRawMempoolReq holds the arguments for the GetRawMempool call.
//  1. verbose             (boolean, optional, default=false) True for a json object, false for array of transaction ids
//  2. mempool_sequence    (boolean, optional, default=false) If verbose=false, returns a json object with transaction list and mempool sequence number attached.
type GetRawMempoolReq struct {
	// True for a json object, false for array of transaction ids
	// Default: false
	Verbose bool `json:"verbose,omitempty"`

	// If verbose=false, returns a json object with transaction list and mempool sequence number attached.
	// Default: false
	MempoolSequence bool `json:"mempool_sequence,omitempty"`
}

// GetRawMempoolResp holds the response to the GetRawMempool call.
//
// ALTERNATIVE (for verbose = false)
//  [           (json array)
//    "hex",    (string) The transaction id
//    ...
//  ]
//
// ALTERNATIVE (for verbose = true)
//  {                                         (json object)
//    "transactionid" : {                     (json object)
//      "vsize" : n,                          (numeric) virtual transaction size as defined in BIP 141. This is different from actual serialized size for witness transactions as witness data is discounted.
//      "weight" : n,                         (numeric) transaction weight as defined in BIP 141.
//      "fee" : n,                            (numeric, optional) transaction fee, denominated in BTC (DEPRECATED, returned only if config option -deprecatedrpc=fees is passed)
//      "modifiedfee" : n,                    (numeric, optional) transaction fee with fee deltas used for mining priority, denominated in BTC (DEPRECATED, returned only if config option -deprecatedrpc=fees is passed)
//      "time" : xxx,                         (numeric) local time transaction entered pool in seconds since 1 Jan 1970 GMT
//      "height" : n,                         (numeric) block height when transaction entered pool
//      "descendantcount" : n,                (numeric) number of in-mempool descendant transactions (including this one)
//      "descendantsize" : n,                 (numeric) virtual transaction size of in-mempool descendants (including this one)
//      "descendantfees" : n,                 (numeric, optional) transaction fees of in-mempool descendants (including this one) with fee deltas used for mining priority, denominated in sats (DEPRECATED, returned only if config option -deprecatedrpc=fees is passed)
//      "ancestorcount" : n,                  (numeric) number of in-mempool ancestor transactions (including this one)
//      "ancestorsize" : n,                   (numeric) virtual transaction size of in-mempool ancestors (including this one)
//      "ancestorfees" : n,                   (numeric, optional) transaction fees of in-mempool ancestors (including this one) with fee deltas used for mining priority, denominated in sats (DEPRECATED, returned only if config option -deprecatedrpc=fees is passed)
//      "wtxid" : "hex",                      (string) hash of serialized transaction, including witness data
//      "fees" : {                            (json object)
//        "base" : n,                         (numeric) transaction fee, denominated in BTC
//        "modified" : n,                     (numeric) transaction fee with fee deltas used for mining priority, denominated in BTC
//        "ancestor" : n,                     (numeric) transaction fees of in-mempool ancestors (including this one) with fee deltas used for mining priority, denominated in BTC
//        "descendant" : n                    (numeric) transaction fees of in-mempool descendants (including this one) with fee deltas used for mining priority, denominated in BTC
//      },
//      "depends" : [                         (json array) unconfirmed transactions used as inputs for this transaction
//        "hex",                              (string) parent transaction id
//        ...
//      ],
//      "spentby" : [                         (json array) unconfirmed transactions spending outputs from this transaction
//        "hex",                              (string) child transaction id
//        ...
//      ],
//      "bip125-replaceable" : true|false,    (boolean) Whether this transaction could be replaced due to BIP125 (replace-by-fee)
//      "unbroadcast" : true|false            (boolean) Whether this transaction is currently unbroadcast (initial broadcast not yet acknowledged by any peers)
//    },
//    ...
//  }
//
// ALTERNATIVE (for verbose = false and mempool_sequence = true)
//  {                            (json object)
//    "txids" : [                (json array)
//      "hex",                   (string) The transaction id
//      ...
//    ],
//    "mempool_sequence" : n     (numeric) The mempool sequence value.
//  }
type GetRawMempoolResp struct {
	// Element: Hex    The transaction id
	Hex []string

	// Key: transactionid, Value: struct
	ForVerboseEqualsTrue map[string]GetRawMempoolRespForVerboseEqualsTrue

	ForVerboseEqualsFalseAndMempoolSequenceEqualsTrue GetRawMempoolRespForVerboseEqualsFalseAndMempoolSequenceEqualsTrue
}

func (alts GetRawMempoolResp) MarshalJSON() ([]byte, error) {
	if !reflect.ValueOf(alts.Hex).IsZero() {
		return json.Marshal(alts.Hex)
	}
	if !reflect.ValueOf(alts.ForVerboseEqualsTrue).IsZero() {
		return json.Marshal(alts.ForVerboseEqualsTrue)
	}
	return json.Marshal(alts.ForVerboseEqualsFalseAndMempoolSequenceEqualsTrue)
}

func (alts *GetRawMempoolResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Hex) == nil {
		return nil
	}
	alts.Hex = reset.Hex
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.ForVerboseEqualsTrue) == nil {
		return nil
	}
	alts.ForVerboseEqualsTrue = reset.ForVerboseEqualsTrue
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.ForVerboseEqualsFalseAndMempoolSequenceEqualsTrue) == nil {
		return nil
	}
	alts.ForVerboseEqualsFalseAndMempoolSequenceEqualsTrue = reset.ForVerboseEqualsFalseAndMempoolSequenceEqualsTrue
	return &UnmarshalError{B: b, structName: "GetRawMempoolResp"}
}

type GetRawMempoolRespForVerboseEqualsTrue struct {
	// virtual transaction size as defined in BIP 141. This is different from actual serialized size for witness transactions as witness data is discounted.
	VSize float64 `json:"vsize"`

	// transaction weight as defined in BIP 141.
	Weight float64 `json:"weight"`

	// transaction fee, denominated in BTC (DEPRECATED, returned only if config option -deprecatedrpc=fees is passed)
	Fee *float64 `json:"fee,omitempty"`

	// transaction fee with fee deltas used for mining priority, denominated in BTC (DEPRECATED, returned only if config option -deprecatedrpc=fees is passed)
	ModifiedFee *float64 `json:"modifiedfee,omitempty"`

	// local time transaction entered pool in seconds since 1 Jan 1970 GMT
	Time float64 `json:"time"`

	// block height when transaction entered pool
	Height float64 `json:"height"`

	// number of in-mempool descendant transactions (including this one)
	DescendantCount float64 `json:"descendantcount"`

	// virtual transaction size of in-mempool descendants (including this one)
	DescendantSize float64 `json:"descendantsize"`

	// transaction fees of in-mempool descendants (including this one) with fee deltas used for mining priority, denominated in sats (DEPRECATED, returned only if config option -deprecatedrpc=fees is passed)
	DescendantFees *float64 `json:"descendantfees,omitempty"`

	// number of in-mempool ancestor transactions (including this one)
	AncestorCount float64 `json:"ancestorcount"`

	// virtual transaction size of in-mempool ancestors (including this one)
	AncestorSize float64 `json:"ancestorsize"`

	// transaction fees of in-mempool ancestors (including this one) with fee deltas used for mining priority, denominated in sats (DEPRECATED, returned only if config option -deprecatedrpc=fees is passed)
	AncestorFees *float64 `json:"ancestorfees,omitempty"`

	// hash of serialized transaction, including witness data
	WTxID string `json:"wtxid"`

	Fees struct {
		// transaction fee, denominated in BTC
		Base float64 `json:"base"`

		// transaction fee with fee deltas used for mining priority, denominated in BTC
		Modified float64 `json:"modified"`

		// transaction fees of in-mempool ancestors (including this one) with fee deltas used for mining priority, denominated in BTC
		Ancestor float64 `json:"ancestor"`

		// transaction fees of in-mempool descendants (including this one) with fee deltas used for mining priority, denominated in BTC
		Descendant float64 `json:"descendant"`
	} `json:"fees"`

	// unconfirmed transactions used as inputs for this transaction
	// Element: Hex    parent transaction id
	Depends []string `json:"depends"`

	// unconfirmed transactions spending outputs from this transaction
	// Element: Hex    child transaction id
	SpentBy []string `json:"spentby"`

	// Whether this transaction could be replaced due to BIP125 (replace-by-fee)
	BIP125Replaceable bool `json:"bip125-replaceable"`

	// Whether this transaction is currently unbroadcast (initial broadcast not yet acknowledged by any peers)
	Unbroadcast bool `json:"unbroadcast"`
}

type GetRawMempoolRespForVerboseEqualsFalseAndMempoolSequenceEqualsTrue struct {
	// Element: Hex    The transaction id
	TxIDs []string `json:"txids"`

	// The mempool sequence value.
	MempoolSequence float64 `json:"mempool_sequence"`
}

// GetRawMempool RPC method.
// Returns all transaction ids in memory pool as a json array of string transaction ids.
// Hint: use getmempoolentry to fetch a specific transaction from the mempool.
func (bc *BitcoindClient) GetRawMempool(ctx context.Context, args GetRawMempoolReq) (result GetRawMempoolResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "getrawmempool", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// GetTxOutReq holds the arguments for the GetTxOut call.
//  1. txid               (string, required) The transaction id
//  2. n                  (numeric, required) vout number
//  3. include_mempool    (boolean, optional, default=true) Whether to include the mempool. Note that an unspent output that is spent in the mempool won't appear.
type GetTxOutReq struct {
	// The transaction id
	TxID string `json:"txid"`

	// vout number
	N float64 `json:"n"`

	// Whether to include the mempool. Note that an unspent output that is spent in the mempool won't appear.
	// Default: true
	IncludeMempool *bool `json:"include_mempool,omitempty"`
}

// GetTxOutResp holds the response to the GetTxOut call.
//
// ALTERNATIVE (If the UTXO was not found)
//  null    (json null)
//
// ALTERNATIVE (Otherwise)
//  {                             (json object)
//    "bestblock" : "hex",        (string) The hash of the block at the tip of the chain
//    "confirmations" : n,        (numeric) The number of confirmations
//    "value" : n,                (numeric) The transaction value in BTC
//    "scriptPubKey" : {          (json object)
//      "asm" : "str",            (string)
//      "desc" : "str",           (string) Inferred descriptor for the output
//      "hex" : "hex",            (string)
//      "type" : "str",           (string) The type, eg pubkeyhash
//      "address" : "str"         (string, optional) The Bitcoin address (only if a well-defined address exists)
//    },
//    "coinbase" : true|false     (boolean) Coinbase or not
//  }
type GetTxOutResp struct {
	// The hash of the block at the tip of the chain
	BestBlock string `json:"bestblock"`

	// The number of confirmations
	Confirmations float64 `json:"confirmations"`

	// The transaction value in BTC
	Value float64 `json:"value"`

	ScriptPubkey struct {
		Asm string `json:"asm"`

		// Inferred descriptor for the output
		Desc string `json:"desc"`

		Hex string `json:"hex"`

		// The type, eg pubkeyhash
		Type string `json:"type"`

		// The Bitcoin address (only if a well-defined address exists)
		Address string `json:"address,omitempty"`
	} `json:"scriptPubKey"`

	// Coinbase or not
	Coinbase bool `json:"coinbase"`
}

// GetTxOut RPC method.
// Returns details about an unspent transaction output.
func (bc *BitcoindClient) GetTxOut(ctx context.Context, args GetTxOutReq) (result GetTxOutResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "gettxout", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// GetTxOutProofReq holds the arguments for the GetTxOutProof call.
//  1. txids          (json array, required) The txids to filter
//       [
//         "txid",    (string) A transaction hash
//         ...
//       ]
//  2. blockhash      (string, optional) If specified, looks for txid in the block with this hash
type GetTxOutProofReq struct {
	// The txids to filter
	// Element: TxID    A transaction hash
	TxIDs []string `json:"txids"`

	// If specified, looks for txid in the block with this hash
	Blockhash string `json:"blockhash,omitempty"`
}

// GetTxOutProofResp holds the response to the GetTxOutProof call.
//  "str"    (string) A string that is a serialized, hex-encoded data for the proof.
type GetTxOutProofResp struct {
	// A string that is a serialized, hex-encoded data for the proof.
	Str string
}

func (alts GetTxOutProofResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.Str)
}

func (alts *GetTxOutProofResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Str) == nil {
		return nil
	}
	alts.Str = reset.Str
	return &UnmarshalError{B: b, structName: "GetTxOutProofResp"}
}

// GetTxOutProof RPC method.
// Returns a hex-encoded proof that "txid" was included in a block.
// NOTE: By default this function only works sometimes. This is when there is an
// unspent output in the utxo for this transaction. To make it always work,
// you need to maintain a transaction index, using the -txindex command line option or
// specify the block in which the transaction is included manually (by blockhash).
func (bc *BitcoindClient) GetTxOutProof(ctx context.Context, args GetTxOutProofReq) (result GetTxOutProofResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "gettxoutproof", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// GetTxOutSetInfoReq holds the arguments for the GetTxOutSetInfo call.
//  1. hash_type         (string, optional, default="hash_serialized_2") Which UTXO set hash should be calculated. Options: 'hash_serialized_2' (the legacy algorithm), 'muhash', 'none'.
//  2. hash_or_height    (string or numeric, optional) The block hash or height of the target height (only available with coinstatsindex).
//  3. use_index         (boolean, optional, default=true) Use coinstatsindex, if available.
type GetTxOutSetInfoReq struct {
	// Which UTXO set hash should be calculated. Options: 'hash_serialized_2' (the legacy algorithm), 'muhash', 'none'.
	// Default: "hash_serialized_2"
	HashType string `json:"hash_type,omitempty"`

	// The block hash or height of the target height (only available with coinstatsindex).
	HashOrHeight string `json:"hash_or_height,omitempty"`

	// Use coinstatsindex, if available.
	// Default: true
	UseIndex *bool `json:"use_index,omitempty"`
}

// GetTxOutSetInfoResp holds the response to the GetTxOutSetInfo call.
//  {                                     (json object)
//    "height" : n,                       (numeric) The block height (index) of the returned statistics
//    "bestblock" : "hex",                (string) The hash of the block at which these statistics are calculated
//    "txouts" : n,                       (numeric) The number of unspent transaction outputs
//    "bogosize" : n,                     (numeric) Database-independent, meaningless metric indicating the UTXO set size
//    "hash_serialized_2" : "hex",        (string, optional) The serialized hash (only present if 'hash_serialized_2' hash_type is chosen)
//    "muhash" : "hex",                   (string, optional) The serialized hash (only present if 'muhash' hash_type is chosen)
//    "transactions" : n,                 (numeric, optional) The number of transactions with unspent outputs (not available when coinstatsindex is used)
//    "disk_size" : n,                    (numeric, optional) The estimated size of the chainstate on disk (not available when coinstatsindex is used)
//    "total_amount" : n,                 (numeric) The total amount of coins in the UTXO set
//    "total_unspendable_amount" : n,     (numeric, optional) The total amount of coins permanently excluded from the UTXO set (only available if coinstatsindex is used)
//    "block_info" : {                    (json object, optional) Info on amounts in the block at this block height (only available if coinstatsindex is used)
//      "prevout_spent" : n,              (numeric) Total amount of all prevouts spent in this block
//      "coinbase" : n,                   (numeric) Coinbase subsidy amount of this block
//      "new_outputs_ex_coinbase" : n,    (numeric) Total amount of new outputs created by this block
//      "unspendable" : n,                (numeric) Total amount of unspendable outputs created in this block
//      "unspendables" : {                (json object) Detailed view of the unspendable categories
//        "genesis_block" : n,            (numeric) The unspendable amount of the Genesis block subsidy
//        "bip30" : n,                    (numeric) Transactions overridden by duplicates (no longer possible with BIP30)
//        "scripts" : n,                  (numeric) Amounts sent to scripts that are unspendable (for example OP_RETURN outputs)
//        "unclaimed_rewards" : n         (numeric) Fee rewards that miners did not claim in their coinbase transaction
//      }
//    }
//  }
type GetTxOutSetInfoResp struct {
	// The block height (index) of the returned statistics
	Height float64 `json:"height"`

	// The hash of the block at which these statistics are calculated
	BestBlock string `json:"bestblock"`

	// The number of unspent transaction outputs
	TxOuts float64 `json:"txouts"`

	// Database-independent, meaningless metric indicating the UTXO set size
	BogoSize float64 `json:"bogosize"`

	// The serialized hash (only present if 'hash_serialized_2' hash_type is chosen)
	HashSerialized2 string `json:"hash_serialized_2,omitempty"`

	// The serialized hash (only present if 'muhash' hash_type is chosen)
	MuHash string `json:"muhash,omitempty"`

	// The number of transactions with unspent outputs (not available when coinstatsindex is used)
	Transactions *float64 `json:"transactions,omitempty"`

	// The estimated size of the chainstate on disk (not available when coinstatsindex is used)
	DiskSize *float64 `json:"disk_size,omitempty"`

	// The total amount of coins in the UTXO set
	TotalAmount float64 `json:"total_amount"`

	// The total amount of coins permanently excluded from the UTXO set (only available if coinstatsindex is used)
	TotalUnspendableAmount *float64 `json:"total_unspendable_amount,omitempty"`

	// Info on amounts in the block at this block height (only available if coinstatsindex is used)
	BlockInfo *GetTxOutSetInfoRespBlockInfo `json:"block_info,omitempty"`
}

type GetTxOutSetInfoRespBlockInfo struct {
	// Total amount of all prevouts spent in this block
	PrevOutSpent float64 `json:"prevout_spent"`

	// Coinbase subsidy amount of this block
	Coinbase float64 `json:"coinbase"`

	// Total amount of new outputs created by this block
	NewOutputsExCoinbase float64 `json:"new_outputs_ex_coinbase"`

	// Total amount of unspendable outputs created in this block
	Unspendable float64 `json:"unspendable"`

	// Detailed view of the unspendable categories
	Unspendables struct {
		// The unspendable amount of the Genesis block subsidy
		GenesisBlock float64 `json:"genesis_block"`

		// Transactions overridden by duplicates (no longer possible with BIP30)
		BIP30 float64 `json:"bip30"`

		// Amounts sent to scripts that are unspendable (for example OP_RETURN outputs)
		Scripts float64 `json:"scripts"`

		// Fee rewards that miners did not claim in their coinbase transaction
		UnclaimedRewards float64 `json:"unclaimed_rewards"`
	} `json:"unspendables"`
}

// GetTxOutSetInfo RPC method.
// Returns statistics about the unspent transaction output set.
// Note this call may take some time if you are not using coinstatsindex.
func (bc *BitcoindClient) GetTxOutSetInfo(ctx context.Context, args GetTxOutSetInfoReq) (result GetTxOutSetInfoResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "gettxoutsetinfo", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// PreciousBlockReq holds the arguments for the PreciousBlock call.
//  1. blockhash    (string, required) the hash of the block to mark as precious
type PreciousBlockReq struct {
	// the hash of the block to mark as precious
	Blockhash string `json:"blockhash"`
}

// PreciousBlock RPC method.
// Treats a block as if it were received before others with the same work.
// A later preciousblock call can override the effect of an earlier one.
// The effects of preciousblock are not retained across restarts.
func (bc *BitcoindClient) PreciousBlock(ctx context.Context, args PreciousBlockReq) (err error) {
	_, err = bc.sendRequest(ctx, "preciousblock", args)
	return
}

// PruneBlockchainReq holds the arguments for the PruneBlockchain call.
//  1. height    (numeric, required) The block height to prune up to. May be set to a discrete height, or to a UNIX epoch time
//               to prune blocks whose block time is at least 2 hours older than the provided timestamp.
type PruneBlockchainReq struct {
	// The block height to prune up to. May be set to a discrete height, or to a UNIX epoch time
	// to prune blocks whose block time is at least 2 hours older than the provided timestamp.
	Height float64 `json:"height"`
}

// PruneBlockchainResp holds the response to the PruneBlockchain call.
//  n    (numeric) Height of the last block pruned
type PruneBlockchainResp struct {
	// Height of the last block pruned
	N float64
}

func (alts PruneBlockchainResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.N)
}

func (alts *PruneBlockchainResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.N) == nil {
		return nil
	}
	alts.N = reset.N
	return &UnmarshalError{B: b, structName: "PruneBlockchainResp"}
}

// PruneBlockchain RPC method.
func (bc *BitcoindClient) PruneBlockchain(ctx context.Context, args PruneBlockchainReq) (result PruneBlockchainResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "pruneblockchain", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// SaveMempoolResp holds the response to the SaveMempool call.
//  {                        (json object)
//    "filename" : "str"     (string) the directory and file where the mempool was saved
//  }
type SaveMempoolResp struct {
	// the directory and file where the mempool was saved
	FileName string `json:"filename"`
}

// SaveMempool RPC method.
// Dumps the mempool to disk. It will fail until the previous dump is fully loaded.
func (bc *BitcoindClient) SaveMempool(ctx context.Context) (result SaveMempoolResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "savemempool", nil); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// ScanTxOutSetReq holds the arguments for the ScanTxOutSet call.
//  1. action                        (string, required) The action to execute
//                                   "start" for starting a scan
//                                   "abort" for aborting the current scan (returns true when abort was successful)
//                                   "status" for progress report (in %) of the current scan
//  2. scanobjects                   (json array) Array of scan objects. Required for "start" action
//                                   Every scan object is either a string descriptor or an object:
//       [
//         "descriptor",             (string) An output descriptor
//         {                         (json object) An object with output descriptor and metadata
//           "desc": "str",          (string, required) An output descriptor
//           "range": n or [n,n],    (numeric or array, optional, default=1000) The range of HD chain indexes to explore (either end or [begin,end])
//         },
//         ...
//       ]
type ScanTxOutSetReq struct {
	// The action to execute
	// "start" for starting a scan
	// "abort" for aborting the current scan (returns true when abort was successful)
	// "status" for progress report (in %) of the current scan
	Action string `json:"action"`

	// Array of scan objects. Required for "start" action
	// Every scan object is either a string descriptor or an object:
	ScanObjects []ScanTxOutSetReqScanObjects `json:"scanobjects"`
}

// Holder of alternative parameter formats, only one will be used, the first that is non-zero.
type ScanTxOutSetReqScanObjects struct {
	// An output descriptor
	Descriptor string

	// An object with output descriptor and metadata
	A struct {
		// An output descriptor
		Desc string `json:"desc"`

		// The range of HD chain indexes to explore (either end or [begin,end])
		// Default: 1000
		Range *[2]int64 `json:"range,omitempty"`
	}
}

func (alts ScanTxOutSetReqScanObjects) MarshalJSON() ([]byte, error) {
	if !reflect.ValueOf(alts.Descriptor).IsZero() {
		return json.Marshal(alts.Descriptor)
	}
	return json.Marshal(alts.A)
}

func (alts *ScanTxOutSetReqScanObjects) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Descriptor) == nil {
		return nil
	}
	alts.Descriptor = reset.Descriptor
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.A) == nil {
		return nil
	}
	alts.A = reset.A
	return &UnmarshalError{B: b, structName: "ScanTxOutSetReqScanObjects"}
}

// ScanTxOutSetResp holds the response to the ScanTxOutSet call.
//
// ALTERNATIVE (When action=='abort')
//  true|false    (boolean)
//
// ALTERNATIVE (When action=='status' and no scan is in progress)
//  null    (json null)
//
// ALTERNATIVE (When action=='status' and scan is in progress)
//  {                    (json object)
//    "progress" : n     (numeric) The scan progress
//  }
//
// ALTERNATIVE (When action=='start')
//  {                                (json object)
//    "success" : true|false,        (boolean) Whether the scan was completed
//    "txouts" : n,                  (numeric) The number of unspent transaction outputs scanned
//    "height" : n,                  (numeric) The current block height (index)
//    "bestblock" : "hex",           (string) The hash of the block at the tip of the chain
//    "unspents" : [                 (json array)
//      {                            (json object)
//        "txid" : "hex",            (string) The transaction id
//        "vout" : n,                (numeric) The vout value
//        "scriptPubKey" : "hex",    (string) The script key
//        "desc" : "str",            (string) A specialized descriptor for the matched scriptPubKey
//        "amount" : n,              (numeric) The total amount in BTC of the unspent output
//        "height" : n               (numeric) Height of the unspent transaction output
//      },
//      ...
//    ],
//    "total_amount" : n             (numeric) The total amount of all found unspent outputs in BTC
//  }
type ScanTxOutSetResp struct {
	TrueOrFalse bool

	WhenActionEqualsStatusAndScanIsInProgress ScanTxOutSetRespWhenActionEqualsStatusAndScanIsInProgress

	WhenActionEqualsStart ScanTxOutSetRespWhenActionEqualsStart
}

func (alts ScanTxOutSetResp) MarshalJSON() ([]byte, error) {
	if !reflect.ValueOf(alts.TrueOrFalse).IsZero() {
		return json.Marshal(alts.TrueOrFalse)
	}
	if !reflect.ValueOf(alts.WhenActionEqualsStatusAndScanIsInProgress).IsZero() {
		return json.Marshal(alts.WhenActionEqualsStatusAndScanIsInProgress)
	}
	return json.Marshal(alts.WhenActionEqualsStart)
}

func (alts *ScanTxOutSetResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.TrueOrFalse) == nil {
		return nil
	}
	alts.TrueOrFalse = reset.TrueOrFalse
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.WhenActionEqualsStatusAndScanIsInProgress) == nil {
		return nil
	}
	alts.WhenActionEqualsStatusAndScanIsInProgress = reset.WhenActionEqualsStatusAndScanIsInProgress
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.WhenActionEqualsStart) == nil {
		return nil
	}
	alts.WhenActionEqualsStart = reset.WhenActionEqualsStart
	return &UnmarshalError{B: b, structName: "ScanTxOutSetResp"}
}

type ScanTxOutSetRespWhenActionEqualsStatusAndScanIsInProgress struct {
	// The scan progress
	Progress float64 `json:"progress"`
}

type ScanTxOutSetRespWhenActionEqualsStart struct {
	// Whether the scan was completed
	Success bool `json:"success"`

	// The number of unspent transaction outputs scanned
	TxOuts float64 `json:"txouts"`

	// The current block height (index)
	Height float64 `json:"height"`

	// The hash of the block at the tip of the chain
	BestBlock string `json:"bestblock"`

	Unspents []ScanTxOutSetRespWhenActionEqualsStartUnspents `json:"unspents"`

	// The total amount of all found unspent outputs in BTC
	TotalAmount float64 `json:"total_amount"`
}

type ScanTxOutSetRespWhenActionEqualsStartUnspents struct {
	// The transaction id
	TxID string `json:"txid"`

	// The vout value
	Vout float64 `json:"vout"`

	// The script key
	ScriptPubkey string `json:"scriptPubKey"`

	// A specialized descriptor for the matched scriptPubKey
	Desc string `json:"desc"`

	// The total amount in BTC of the unspent output
	Amount float64 `json:"amount"`

	// Height of the unspent transaction output
	Height float64 `json:"height"`
}

// ScanTxOutSet RPC method.
// Scans the unspent transaction output set for entries that match certain output descriptors.
// Examples of output descriptors are:
//     addr(<address>)                      Outputs whose scriptPubKey corresponds to the specified address (does not include P2PK)
//     raw(<hex script>)                    Outputs whose scriptPubKey equals the specified hex scripts
//     combo(<pubkey>)                      P2PK, P2PKH, P2WPKH, and P2SH-P2WPKH outputs for the given pubkey
//     pkh(<pubkey>)                        P2PKH outputs for the given pubkey
//     sh(multi(<n>,<pubkey>,<pubkey>,...)) P2SH-multisig outputs for the given threshold and pubkeys
// In the above, <pubkey> either refers to a fixed public key in hexadecimal notation, or to an xpub/xprv optionally followed by one
// or more path elements separated by "/", and optionally ending in "/*" (unhardened), or "/*'" or "/*h" (hardened) to specify all
// unhardened or hardened child keys.
// In the latter case, a range needs to be specified by below if different from 1000.
// For more information on output descriptors, see the documentation in the doc/descriptors.md file.
func (bc *BitcoindClient) ScanTxOutSet(ctx context.Context, args ScanTxOutSetReq) (result ScanTxOutSetResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "scantxoutset", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// VerifyChainReq holds the arguments for the VerifyChain call.
//  1. checklevel    (numeric, optional, default=3, range=0-4) How thorough the block verification is:
//                   - level 0 reads the blocks from disk
//                   - level 1 verifies block validity
//                   - level 2 verifies undo data
//                   - level 3 checks disconnection of tip blocks
//                   - level 4 tries to reconnect the blocks
//                   - each level includes the checks of the previous levels
//  2. nblocks       (numeric, optional, default=6, 0=all) The number of blocks to check.
type VerifyChainReq struct {
	// How thorough the block verification is:
	// - level 0 reads the blocks from disk
	// - level 1 verifies block validity
	// - level 2 verifies undo data
	// - level 3 checks disconnection of tip blocks
	// - level 4 tries to reconnect the blocks
	// - each level includes the checks of the previous levels
	// Default: 3, range=0-4
	CheckLevel *float64 `json:"checklevel,omitempty"`

	// The number of blocks to check.
	// Default: 6, 0=all
	NBlocks *float64 `json:"nblocks,omitempty"`
}

// VerifyChainResp holds the response to the VerifyChain call.
//  true|false    (boolean) Verified or not
type VerifyChainResp struct {
	// Verified or not
	TrueOrFalse bool
}

func (alts VerifyChainResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.TrueOrFalse)
}

func (alts *VerifyChainResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.TrueOrFalse) == nil {
		return nil
	}
	alts.TrueOrFalse = reset.TrueOrFalse
	return &UnmarshalError{B: b, structName: "VerifyChainResp"}
}

// VerifyChain RPC method.
// Verifies blockchain database.
func (bc *BitcoindClient) VerifyChain(ctx context.Context, args VerifyChainReq) (result VerifyChainResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "verifychain", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// VerifyTxOutProofReq holds the arguments for the VerifyTxOutProof call.
//  1. proof    (string, required) The hex-encoded proof generated by gettxoutproof
type VerifyTxOutProofReq struct {
	// The hex-encoded proof generated by gettxoutproof
	Proof string `json:"proof"`
}

// VerifyTxOutProofResp holds the response to the VerifyTxOutProof call.
//  [           (json array)
//    "hex",    (string) The txid(s) which the proof commits to, or empty array if the proof cannot be validated.
//    ...
//  ]
type VerifyTxOutProofResp struct {
	// Element: Hex    The txid(s) which the proof commits to, or empty array if the proof cannot be validated.
	Hex []string
}

func (alts VerifyTxOutProofResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.Hex)
}

func (alts *VerifyTxOutProofResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Hex) == nil {
		return nil
	}
	alts.Hex = reset.Hex
	return &UnmarshalError{B: b, structName: "VerifyTxOutProofResp"}
}

// VerifyTxOutProof RPC method.
// Verifies that a proof points to a transaction in a block, returning the transaction it commits to
// and throwing an RPC error if the block is not in our best chain
func (bc *BitcoindClient) VerifyTxOutProof(ctx context.Context, args VerifyTxOutProofReq) (result VerifyTxOutProofResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "verifytxoutproof", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}
