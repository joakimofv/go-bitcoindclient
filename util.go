// THIS FILE IS GENERATED CODE, MANUAL EDITS WILL BE OVERWRITTEN

package bitcoindclient

import (
	"bytes"
	"context"
	"encoding/json"
)

// CreateMultisigReq holds the arguments for the CreateMultisig call.
//  1. nrequired       (numeric, required) The number of required signatures out of the n keys.
//  2. keys            (json array, required) The hex-encoded public keys.
//       [
//         "key",      (string) The hex-encoded public key
//         ...
//       ]
//  3. address_type    (string, optional, default="legacy") The address type to use. Options are "legacy", "p2sh-segwit", and "bech32".
type CreateMultisigReq struct {
	// The number of required signatures out of the n keys.
	NRequired float64 `json:"nrequired"`

	// The hex-encoded public keys.
	// Element: Key    The hex-encoded public key
	Keys []string `json:"keys"`

	// The address type to use. Options are "legacy", "p2sh-segwit", and "bech32".
	// Default: "legacy"
	AddressType string `json:"address_type,omitempty"`
}

// CreateMultisigResp holds the response to the CreateMultisig call.
//  {                            (json object)
//    "address" : "str",         (string) The value of the new multisig address.
//    "redeemScript" : "hex",    (string) The string value of the hex-encoded redemption script.
//    "descriptor" : "str"       (string) The descriptor for this multisig
//  }
type CreateMultisigResp struct {
	// The value of the new multisig address.
	Address string `json:"address"`

	// The string value of the hex-encoded redemption script.
	RedeemScript string `json:"redeemScript"`

	// The descriptor for this multisig
	Descriptor string `json:"descriptor"`
}

// CreateMultisig RPC method.
// Creates a multi-signature address with n signature of m keys required.
// It returns a json object with the address and redeemScript.
func (bc *BitcoindClient) CreateMultisig(ctx context.Context, args CreateMultisigReq) (result CreateMultisigResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "createmultisig", args, false); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// DeriveAddressesReq holds the arguments for the DeriveAddresses call.
//  1. descriptor    (string, required) The descriptor.
//  2. range         (numeric or array, optional) If a ranged descriptor is used, this specifies the end or the range (in [begin,end] notation) to derive.
type DeriveAddressesReq struct {
	// The descriptor.
	Descriptor string `json:"descriptor"`

	// If a ranged descriptor is used, this specifies the end or the range (in [begin,end] notation) to derive.
	Range *[2]int64 `json:"range,omitempty"`
}

// DeriveAddressesResp holds the response to the DeriveAddresses call.
//  [           (json array)
//    "str",    (string) the derived addresses
//    ...
//  ]
type DeriveAddressesResp struct {
	// Element: Str    the derived addresses
	Str []string
}

func (alts DeriveAddressesResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.Str)
}

func (alts *DeriveAddressesResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Str) == nil {
		return nil
	}
	alts.Str = reset.Str
	return &UnmarshalError{B: b, structName: "DeriveAddressesResp"}
}

// DeriveAddresses RPC method.
// Derives one or more addresses corresponding to an output descriptor.
// Examples of output descriptors are:
//     pkh(<pubkey>)                        P2PKH outputs for the given pubkey
//     wpkh(<pubkey>)                       Native segwit P2PKH outputs for the given pubkey
//     sh(multi(<n>,<pubkey>,<pubkey>,...)) P2SH-multisig outputs for the given threshold and pubkeys
//     raw(<hex script>)                    Outputs whose scriptPubKey equals the specified hex scripts
// In the above, <pubkey> either refers to a fixed public key in hexadecimal notation, or to an xpub/xprv optionally followed by one
// or more path elements separated by "/", where "h" represents a hardened child key.
// For more information on output descriptors, see the documentation in the doc/descriptors.md file.
func (bc *BitcoindClient) DeriveAddresses(ctx context.Context, args DeriveAddressesReq) (result DeriveAddressesResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "deriveaddresses", args, false); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// EstimateSmartFeeReq holds the arguments for the EstimateSmartFee call.
//  1. conf_target      (numeric, required) Confirmation target in blocks (1 - 1008)
//  2. estimate_mode    (string, optional, default="conservative") The fee estimate mode.
//                      Whether to return a more conservative estimate which also satisfies
//                      a longer history. A conservative estimate potentially returns a
//                      higher feerate and is more likely to be sufficient for the desired
//                      target, but is not as responsive to short term drops in the
//                      prevailing fee market. Must be one of (case insensitive):
//                      "unset"
//                      "economical"
//                      "conservative"
type EstimateSmartFeeReq struct {
	// Confirmation target in blocks (1 - 1008)
	ConfTarget float64 `json:"conf_target"`

	// The fee estimate mode.
	// Whether to return a more conservative estimate which also satisfies
	// a longer history. A conservative estimate potentially returns a
	// higher feerate and is more likely to be sufficient for the desired
	// target, but is not as responsive to short term drops in the
	// prevailing fee market. Must be one of (case insensitive):
	// "unset"
	// "economical"
	// "conservative"
	// Default: "conservative"
	EstimateMode string `json:"estimate_mode,omitempty"`
}

// EstimateSmartFeeResp holds the response to the EstimateSmartFee call.
//  {                   (json object)
//    "feerate" : n,    (numeric, optional) estimate fee rate in BTC/kvB (only present if no errors were encountered)
//    "errors" : [      (json array, optional) Errors encountered during processing (if there are any)
//      "str",          (string) error
//      ...
//    ],
//    "blocks" : n      (numeric) block number where estimate was found
//                      The request target will be clamped between 2 and the highest target
//                      fee estimation is able to return based on how long it has been running.
//                      An error is returned if not enough transactions and blocks
//                      have been observed to make an estimate for any number of blocks.
//  }
type EstimateSmartFeeResp struct {
	// estimate fee rate in BTC/kvB (only present if no errors were encountered)
	FeeRate *float64 `json:"feerate,omitempty"`

	// Errors encountered during processing (if there are any)
	// Element: Str    error
	Errors []string `json:"errors,omitempty"`

	// block number where estimate was found
	// The request target will be clamped between 2 and the highest target
	// fee estimation is able to return based on how long it has been running.
	// An error is returned if not enough transactions and blocks
	// have been observed to make an estimate for any number of blocks.
	Blocks float64 `json:"blocks"`
}

// EstimateSmartFee RPC method.
// Estimates the approximate fee per kilobyte needed for a transaction to begin
// confirmation within conf_target blocks if possible and return the number of blocks
// for which the estimate is valid. Uses virtual transaction size as defined
// in BIP 141 (witness data is discounted).
func (bc *BitcoindClient) EstimateSmartFee(ctx context.Context, args EstimateSmartFeeReq) (result EstimateSmartFeeResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "estimatesmartfee", args, false); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// GetDescriptorInfoReq holds the arguments for the GetDescriptorInfo call.
//  1. descriptor    (string, required) The descriptor.
type GetDescriptorInfoReq struct {
	// The descriptor.
	Descriptor string `json:"descriptor"`
}

// GetDescriptorInfoResp holds the response to the GetDescriptorInfo call.
//  {                                   (json object)
//    "descriptor" : "str",             (string) The descriptor in canonical form, without private keys
//    "checksum" : "str",               (string) The checksum for the input descriptor
//    "isrange" : true|false,           (boolean) Whether the descriptor is ranged
//    "issolvable" : true|false,        (boolean) Whether the descriptor is solvable
//    "hasprivatekeys" : true|false     (boolean) Whether the input descriptor contained at least one private key
//  }
type GetDescriptorInfoResp struct {
	// The descriptor in canonical form, without private keys
	Descriptor string `json:"descriptor"`

	// The checksum for the input descriptor
	CheckSum string `json:"checksum"`

	// Whether the descriptor is ranged
	IsRange bool `json:"isrange"`

	// Whether the descriptor is solvable
	IsSolvable bool `json:"issolvable"`

	// Whether the input descriptor contained at least one private key
	HasPrivateKeys bool `json:"hasprivatekeys"`
}

// GetDescriptorInfo RPC method.
// Analyses a descriptor.
func (bc *BitcoindClient) GetDescriptorInfo(ctx context.Context, args GetDescriptorInfoReq) (result GetDescriptorInfoResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "getdescriptorinfo", args, false); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// GetIndexInfoReq holds the arguments for the GetIndexInfo call.
//  1. index_name    (string, optional) Filter results for an index with a specific name.
type GetIndexInfoReq struct {
	// Filter results for an index with a specific name.
	IndexName string `json:"index_name,omitempty"`
}

// GetIndexInfoResp holds the response to the GetIndexInfo call.
//  {                               (json object)
//    "name" : {                    (json object) The name of the index
//      "synced" : true|false,      (boolean) Whether the index is synced or not
//      "best_block_height" : n     (numeric) The block height to which the index is synced
//    }
//  }
type GetIndexInfoResp struct {
	// The name of the index
	Name struct {
		// Whether the index is synced or not
		Synced bool `json:"synced"`

		// The block height to which the index is synced
		BestBlockHeight float64 `json:"best_block_height"`
	} `json:"name"`
}

// GetIndexInfo RPC method.
// Returns the status of one or all available indices currently running in the node.
func (bc *BitcoindClient) GetIndexInfo(ctx context.Context, args GetIndexInfoReq) (result GetIndexInfoResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "getindexinfo", args, false); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// SignMessageWithPrivkeyReq holds the arguments for the SignMessageWithPrivkey call.
//  1. privkey    (string, required) The private key to sign the message with.
//  2. message    (string, required) The message to create a signature of.
type SignMessageWithPrivkeyReq struct {
	// The private key to sign the message with.
	Privkey string `json:"privkey"`

	// The message to create a signature of.
	Message string `json:"message"`
}

// SignMessageWithPrivkeyResp holds the response to the SignMessageWithPrivkey call.
//  "str"    (string) The signature of the message encoded in base 64
type SignMessageWithPrivkeyResp struct {
	// The signature of the message encoded in base 64
	Str string
}

func (alts SignMessageWithPrivkeyResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.Str)
}

func (alts *SignMessageWithPrivkeyResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Str) == nil {
		return nil
	}
	alts.Str = reset.Str
	return &UnmarshalError{B: b, structName: "SignMessageWithPrivkeyResp"}
}

// SignMessageWithPrivkey RPC method.
// Sign a message with the private key of an address
func (bc *BitcoindClient) SignMessageWithPrivkey(ctx context.Context, args SignMessageWithPrivkeyReq) (result SignMessageWithPrivkeyResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "signmessagewithprivkey", args, false); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// ValidateAddressReq holds the arguments for the ValidateAddress call.
//  1. address    (string, required) The bitcoin address to validate
type ValidateAddressReq struct {
	// The bitcoin address to validate
	Address string `json:"address"`
}

// ValidateAddressResp holds the response to the ValidateAddress call.
//  {                               (json object)
//    "isvalid" : true|false,       (boolean) If the address is valid or not
//    "address" : "str",            (string) The bitcoin address validated
//    "scriptPubKey" : "hex",       (string) The hex-encoded scriptPubKey generated by the address
//    "isscript" : true|false,      (boolean) If the key is a script
//    "iswitness" : true|false,     (boolean) If the address is a witness address
//    "witness_version" : n,        (numeric, optional) The version number of the witness program
//    "witness_program" : "hex",    (string, optional) The hex value of the witness program
//    "error" : "str"               (string, optional) Error message, if any
//  }
type ValidateAddressResp struct {
	// If the address is valid or not
	IsValid bool `json:"isvalid"`

	// The bitcoin address validated
	Address string `json:"address"`

	// The hex-encoded scriptPubKey generated by the address
	ScriptPubkey string `json:"scriptPubKey"`

	// If the key is a script
	IsScript bool `json:"isscript"`

	// If the address is a witness address
	IsWitness bool `json:"iswitness"`

	// The version number of the witness program
	WitnessVersion *float64 `json:"witness_version,omitempty"`

	// The hex value of the witness program
	WitnessProgram string `json:"witness_program,omitempty"`

	// Error message, if any
	Error string `json:"error,omitempty"`
}

// ValidateAddress RPC method.
// Return information about the given bitcoin address.
func (bc *BitcoindClient) ValidateAddress(ctx context.Context, args ValidateAddressReq) (result ValidateAddressResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "validateaddress", args, false); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// VerifyMessageReq holds the arguments for the VerifyMessage call.
//  1. address      (string, required) The bitcoin address to use for the signature.
//  2. signature    (string, required) The signature provided by the signer in base 64 encoding (see signmessage).
//  3. message      (string, required) The message that was signed.
type VerifyMessageReq struct {
	// The bitcoin address to use for the signature.
	Address string `json:"address"`

	// The signature provided by the signer in base 64 encoding (see signmessage).
	Signature string `json:"signature"`

	// The message that was signed.
	Message string `json:"message"`
}

// VerifyMessageResp holds the response to the VerifyMessage call.
//  true|false    (boolean) If the signature is verified or not.
type VerifyMessageResp struct {
	// If the signature is verified or not.
	TrueOrFalse bool
}

func (alts VerifyMessageResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.TrueOrFalse)
}

func (alts *VerifyMessageResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.TrueOrFalse) == nil {
		return nil
	}
	alts.TrueOrFalse = reset.TrueOrFalse
	return &UnmarshalError{B: b, structName: "VerifyMessageResp"}
}

// VerifyMessage RPC method.
// Verify a signed message
func (bc *BitcoindClient) VerifyMessage(ctx context.Context, args VerifyMessageReq) (result VerifyMessageResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "verifymessage", args, false); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}
