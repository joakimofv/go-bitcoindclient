// THIS FILE IS GENERATED CODE, MANUAL EDITS WILL BE OVERWRITTEN

package bitcoindclient

import (
	"bytes"
	"context"
	"encoding/json"
	"reflect"
)

// AbandonTransactionReq holds the arguments for the AbandonTransaction call.
//  1. txid    (string, required) The transaction id
type AbandonTransactionReq struct {
	// The transaction id
	TxID string `json:"txid"`
}

// AbandonTransaction RPC method.
// Mark in-wallet transaction <txid> as abandoned
// This will mark this transaction and all its in-wallet descendants as abandoned which will allow
// for their inputs to be respent.  It can be used to replace "stuck" or evicted transactions.
// It only works on transactions which are not included in a block and are not currently in the mempool.
// It has no effect on transactions which are already abandoned.
func (bc *BitcoindClient) AbandonTransaction(ctx context.Context, args AbandonTransactionReq) (err error) {
	_, err = bc.sendRequest(ctx, "abandontransaction", args)
	return
}

// AbortRescanResp holds the response to the AbortRescan call.
//  true|false    (boolean) Whether the abort was successful
type AbortRescanResp struct {
	// Whether the abort was successful
	TrueOrFalse bool
}

func (alts AbortRescanResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.TrueOrFalse)
}

func (alts *AbortRescanResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.TrueOrFalse) == nil {
		return nil
	}
	alts.TrueOrFalse = reset.TrueOrFalse
	return &UnmarshalError{B: b, structName: "AbortRescanResp"}
}

// AbortRescan RPC method.
// Stops current wallet rescan triggered by an RPC call, e.g. by an importprivkey call.
// Note: Use "getwalletinfo" to query the scanning progress.
func (bc *BitcoindClient) AbortRescan(ctx context.Context) (result AbortRescanResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "abortrescan", nil); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// AddMultisigAddressReq holds the arguments for the AddMultisigAddress call.
//  1. nrequired       (numeric, required) The number of required signatures out of the n keys or addresses.
//  2. keys            (json array, required) The bitcoin addresses or hex-encoded public keys
//       [
//         "key",      (string) bitcoin address or hex-encoded public key
//         ...
//       ]
//  3. label           (string, optional) A label to assign the addresses to.
//  4. address_type    (string, optional, default=set by -addresstype) The address type to use. Options are "legacy", "p2sh-segwit", and "bech32".
type AddMultisigAddressReq struct {
	// The number of required signatures out of the n keys or addresses.
	NRequired float64 `json:"nrequired"`

	// The bitcoin addresses or hex-encoded public keys
	// Element: Key    bitcoin address or hex-encoded public key
	Keys []string `json:"keys"`

	// A label to assign the addresses to.
	Label string `json:"label,omitempty"`

	// The address type to use. Options are "legacy", "p2sh-segwit", and "bech32".
	// Default: set by -addresstype
	AddressType string `json:"address_type,omitempty"`
}

// AddMultisigAddressResp holds the response to the AddMultisigAddress call.
//  {                            (json object)
//    "address" : "str",         (string) The value of the new multisig address
//    "redeemScript" : "hex",    (string) The string value of the hex-encoded redemption script
//    "descriptor" : "str"       (string) The descriptor for this multisig
//  }
type AddMultisigAddressResp struct {
	// The value of the new multisig address
	Address string `json:"address"`

	// The string value of the hex-encoded redemption script
	RedeemScript string `json:"redeemScript"`

	// The descriptor for this multisig
	Descriptor string `json:"descriptor"`
}

// AddMultisigAddress RPC method.
// Add an nrequired-to-sign multisignature address to the wallet. Requires a new wallet backup.
// Each key is a Bitcoin address or hex-encoded public key.
// This functionality is only intended for use with non-watchonly addresses.
// See `importaddress` for watchonly p2sh address support.
// If 'label' is specified, assign address to that label.
func (bc *BitcoindClient) AddMultisigAddress(ctx context.Context, args AddMultisigAddressReq) (result AddMultisigAddressResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "addmultisigaddress", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// BackupWalletReq holds the arguments for the BackupWallet call.
//  1. destination    (string, required) The destination directory or file
type BackupWalletReq struct {
	// The destination directory or file
	Destination string `json:"destination"`
}

// BackupWallet RPC method.
// Safely copies current wallet file to destination, which can be a directory or a path with filename.
func (bc *BitcoindClient) BackupWallet(ctx context.Context, args BackupWalletReq) (err error) {
	_, err = bc.sendRequest(ctx, "backupwallet", args)
	return
}

// BumpFeeReq holds the arguments for the BumpFee call.
//  1. txid                           (string, required) The txid to be bumped
//  2. options                        (json object, optional)
//       {
//         "conf_target": n,          (numeric, optional, default=wallet -txconfirmtarget) Confirmation target in blocks
type BumpFeeReq struct {
	// The txid to be bumped
	TxID string `json:"txid"`

	Options *BumpFeeReqOptions `json:"options,omitempty"`
}

type BumpFeeReqOptions struct {
	// Confirmation target in blocks
	// Default: wallet -txconfirmtarget
	ConfTarget *float64 `json:"conf_target,omitempty"`
}

// BumpFeeResp holds the response to the BumpFee call.
//  {                    (json object)
//    "txid" : "hex",    (string) The id of the new transaction.
//    "origfee" : n,     (numeric) The fee of the replaced transaction.
//    "fee" : n,         (numeric) The fee of the new transaction.
//    "errors" : [       (json array) Errors encountered during processing (may be empty).
//      "str",           (string)
//      ...
//    ]
//  }
type BumpFeeResp struct {
	// The id of the new transaction.
	TxID string `json:"txid"`

	// The fee of the replaced transaction.
	OrigFee float64 `json:"origfee"`

	// The fee of the new transaction.
	Fee float64 `json:"fee"`

	// Errors encountered during processing (may be empty).
	// Element: Str
	Errors []string `json:"errors"`
}

// BumpFee RPC method.
// Bumps the fee of an opt-in-RBF transaction T, replacing it with a new transaction B.
// An opt-in RBF transaction with the given txid must be in the wallet.
// The command will pay the additional fee by reducing change outputs or adding inputs when necessary.
// It may add a new change output if one does not already exist.
// All inputs in the original transaction will be included in the replacement transaction.
// The command will fail if the wallet or mempool contains a transaction that spends one of T's outputs.
// By default, the new fee will be calculated automatically using the estimatesmartfee RPC.
// The user can specify a confirmation target for estimatesmartfee.
// Alternatively, the user can specify a fee rate in sat/vB for the new transaction.
// At a minimum, the new fee rate must be high enough to pay an additional new relay fee (incrementalfee
// returned by getnetworkinfo) to enter the node's mempool.
// * WARNING: before version 0.21, fee_rate was in BTC/kvB. As of 0.21, fee_rate is in sat/vB. *
func (bc *BitcoindClient) BumpFee(ctx context.Context, args BumpFeeReq) (result BumpFeeResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "bumpfee", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// CreateWalletReq holds the arguments for the CreateWallet call.
//  1. wallet_name             (string, required) The name for the new wallet. If this is a path, the wallet will be created at the path location.
//  2. disable_private_keys    (boolean, optional, default=false) Disable the possibility of private keys (only watchonlys are possible in this mode).
//  3. blank                   (boolean, optional, default=false) Create a blank wallet. A blank wallet has no keys or HD seed. One can be set using sethdseed.
//  4. passphrase              (string, optional) Encrypt the wallet with this passphrase.
//  5. avoid_reuse             (boolean, optional, default=false) Keep track of coin reuse, and treat dirty and clean coins differently with privacy considerations in mind.
//  6. descriptors             (boolean, optional, default=false) Create a native descriptor wallet. The wallet will use descriptors internally to handle address creation
//  7. load_on_startup         (boolean, optional) Save wallet name to persistent settings and load on startup. True to add wallet to startup list, false to remove, null to leave unchanged.
//  8. external_signer         (boolean, optional, default=false) Use an external signer such as a hardware wallet. Requires -signer to be configured. Wallet creation will fail if keys cannot be fetched. Requires disable_private_keys and descriptors set to true.
type CreateWalletReq struct {
	// The name for the new wallet. If this is a path, the wallet will be created at the path location.
	WalletName string `json:"wallet_name"`

	// Disable the possibility of private keys (only watchonlys are possible in this mode).
	// Default: false
	DisablePrivateKeys bool `json:"disable_private_keys,omitempty"`

	// Create a blank wallet. A blank wallet has no keys or HD seed. One can be set using sethdseed.
	// Default: false
	Blank bool `json:"blank,omitempty"`

	// Encrypt the wallet with this passphrase.
	Passphrase string `json:"passphrase,omitempty"`

	// Keep track of coin reuse, and treat dirty and clean coins differently with privacy considerations in mind.
	// Default: false
	AvoidReuse bool `json:"avoid_reuse,omitempty"`

	// Create a native descriptor wallet. The wallet will use descriptors internally to handle address creation
	// Default: false
	Descriptors bool `json:"descriptors,omitempty"`

	// Save wallet name to persistent settings and load on startup. True to add wallet to startup list, false to remove, null to leave unchanged.
	LoadOnStartup *bool `json:"load_on_startup,omitempty"`

	// Use an external signer such as a hardware wallet. Requires -signer to be configured. Wallet creation will fail if keys cannot be fetched. Requires disable_private_keys and descriptors set to true.
	// Default: false
	ExternalSigner bool `json:"external_signer,omitempty"`
}

// CreateWalletResp holds the response to the CreateWallet call.
//  {                       (json object)
//    "name" : "str",       (string) The wallet name if created successfully. If the wallet was created using a full path, the wallet_name will be the full path.
//    "warning" : "str"     (string) Warning message if wallet was not loaded cleanly.
//  }
type CreateWalletResp struct {
	// The wallet name if created successfully. If the wallet was created using a full path, the wallet_name will be the full path.
	Name string `json:"name"`

	// Warning message if wallet was not loaded cleanly.
	Warning string `json:"warning"`
}

// CreateWallet RPC method.
// Creates and loads a new wallet.
func (bc *BitcoindClient) CreateWallet(ctx context.Context, args CreateWalletReq) (result CreateWalletResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "createwallet", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// DumpPrivkeyReq holds the arguments for the DumpPrivkey call.
//  1. address    (string, required) The bitcoin address for the private key
type DumpPrivkeyReq struct {
	// The bitcoin address for the private key
	Address string `json:"address"`
}

// DumpPrivkeyResp holds the response to the DumpPrivkey call.
//  "str"    (string) The private key
type DumpPrivkeyResp struct {
	// The private key
	Str string
}

func (alts DumpPrivkeyResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.Str)
}

func (alts *DumpPrivkeyResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Str) == nil {
		return nil
	}
	alts.Str = reset.Str
	return &UnmarshalError{B: b, structName: "DumpPrivkeyResp"}
}

// DumpPrivkey RPC method.
// Reveals the private key corresponding to 'address'.
// Then the importprivkey can be used with this output
func (bc *BitcoindClient) DumpPrivkey(ctx context.Context, args DumpPrivkeyReq) (result DumpPrivkeyResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "dumpprivkey", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// DumpWalletReq holds the arguments for the DumpWallet call.
//  1. filename    (string, required) The filename with path (absolute path recommended)
type DumpWalletReq struct {
	// The filename with path (absolute path recommended)
	FileName string `json:"filename"`
}

// DumpWalletResp holds the response to the DumpWallet call.
//  {                        (json object)
//    "filename" : "str"     (string) The filename with full absolute path
//  }
type DumpWalletResp struct {
	// The filename with full absolute path
	FileName string `json:"filename"`
}

// DumpWallet RPC method.
// Dumps all wallet keys in a human-readable format to a server-side file. This does not allow overwriting existing files.
// Imported scripts are included in the dumpfile, but corresponding BIP173 addresses, etc. may not be added automatically by importwallet.
// Note that if your wallet contains keys which are not derived from your HD seed (e.g. imported keys), these are not covered by
// only backing up the seed itself, and must be backed up too (e.g. ensure you back up the whole dumpfile).
func (bc *BitcoindClient) DumpWallet(ctx context.Context, args DumpWalletReq) (result DumpWalletResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "dumpwallet", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// EncryptWalletReq holds the arguments for the EncryptWallet call.
//  1. passphrase    (string, required) The pass phrase to encrypt the wallet with. It must be at least 1 character, but should be long.
type EncryptWalletReq struct {
	// The pass phrase to encrypt the wallet with. It must be at least 1 character, but should be long.
	Passphrase string `json:"passphrase"`
}

// EncryptWalletResp holds the response to the EncryptWallet call.
//  "str"    (string) A string with further instructions
type EncryptWalletResp struct {
	// A string with further instructions
	Str string
}

func (alts EncryptWalletResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.Str)
}

func (alts *EncryptWalletResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Str) == nil {
		return nil
	}
	alts.Str = reset.Str
	return &UnmarshalError{B: b, structName: "EncryptWalletResp"}
}

// EncryptWallet RPC method.
// Encrypts the wallet with 'passphrase'. This is for first time encryption.
// After this, any calls that interact with private keys such as sending or signing
// will require the passphrase to be set prior the making these calls.
// Use the walletpassphrase call for this, and then walletlock call.
// If the wallet is already encrypted, use the walletpassphrasechange call.
func (bc *BitcoindClient) EncryptWallet(ctx context.Context, args EncryptWalletReq) (result EncryptWalletResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "encryptwallet", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// GetAddressesByLabelReq holds the arguments for the GetAddressesByLabel call.
//  1. label    (string, required) The label.
type GetAddressesByLabelReq struct {
	// The label.
	Label string `json:"label"`
}

// GetAddressesByLabel RPC method.
// Returns the list of addresses assigned the specified label.
func (bc *BitcoindClient) GetAddressesByLabel(ctx context.Context, args GetAddressesByLabelReq) (err error) {
	_, err = bc.sendRequest(ctx, "getaddressesbylabel", args)
	return
}

// GetAddressInfoReq holds the arguments for the GetAddressInfo call.
//  1. address    (string, required) The bitcoin address for which to get information.
type GetAddressInfoReq struct {
	// The bitcoin address for which to get information.
	Address string `json:"address"`
}

// GetAddressInfoResp holds the response to the GetAddressInfo call.
//  {                                   (json object)
//    "address" : "str",                (string) The bitcoin address validated.
//    "scriptPubKey" : "hex",           (string) The hex-encoded scriptPubKey generated by the address.
//    "ismine" : true|false,            (boolean) If the address is yours.
//    "iswatchonly" : true|false,       (boolean) If the address is watchonly.
//    "solvable" : true|false,          (boolean) If we know how to spend coins sent to this address, ignoring the possible lack of private keys.
//    "desc" : "str",                   (string, optional) A descriptor for spending coins sent to this address (only when solvable).
//    "parent_desc" : "str",            (string, optional) The descriptor used to derive this address if this is a descriptor wallet
//    "isscript" : true|false,          (boolean) If the key is a script.
//    "ischange" : true|false,          (boolean) If the address was used for change output.
//    "iswitness" : true|false,         (boolean) If the address is a witness address.
//    "witness_version" : n,            (numeric, optional) The version number of the witness program.
//    "witness_program" : "hex",        (string, optional) The hex value of the witness program.
//    "script" : "str",                 (string, optional) The output script type. Only if isscript is true and the redeemscript is known. Possible
//                                      types: nonstandard, pubkey, pubkeyhash, scripthash, multisig, nulldata, witness_v0_keyhash,
//                                      witness_v0_scripthash, witness_unknown.
//    "hex" : "hex",                    (string, optional) The redeemscript for the p2sh address.
//    "pubkeys" : [                     (json array, optional) Array of pubkeys associated with the known redeemscript (only if script is multisig).
//      "str",                          (string)
//      ...
//    ],
//    "sigsrequired" : n,               (numeric, optional) The number of signatures required to spend multisig output (only if script is multisig).
//    "pubkey" : "hex",                 (string, optional) The hex value of the raw public key for single-key addresses (possibly embedded in P2SH or P2WSH).
//    "embedded" : {                    (json object, optional) Information about the address embedded in P2SH or P2WSH, if relevant and known.
//      ...                             Includes all getaddressinfo output fields for the embedded address, excluding metadata (timestamp, hdkeypath, hdseedid)
//                                      and relation to the wallet (ismine, iswatchonly).
//    },
//    "iscompressed" : true|false,      (boolean, optional) If the pubkey is compressed.
//    "timestamp" : xxx,                (numeric, optional) The creation time of the key, if available, expressed in UNIX epoch time.
//    "hdkeypath" : "str",              (string, optional) The HD keypath, if the key is HD and available.
//    "hdseedid" : "hex",               (string, optional) The Hash160 of the HD seed.
//    "hdmasterfingerprint" : "hex",    (string, optional) The fingerprint of the master key.
//    "labels" : [                      (json array) Array of labels associated with the address. Currently limited to one label but returned
//                                      as an array to keep the API stable if multiple labels are enabled in the future.
//      "str",                          (string) Label name (defaults to "").
//      ...
//    ]
//  }
type GetAddressInfoResp struct {
	// The bitcoin address validated.
	Address string `json:"address"`

	// The hex-encoded scriptPubKey generated by the address.
	ScriptPubkey string `json:"scriptPubKey"`

	// If the address is yours.
	IsMine bool `json:"ismine"`

	// If the address is watchonly.
	IsWatchOnly bool `json:"iswatchonly"`

	// If we know how to spend coins sent to this address, ignoring the possible lack of private keys.
	Solvable bool `json:"solvable"`

	// A descriptor for spending coins sent to this address (only when solvable).
	Desc string `json:"desc,omitempty"`

	// The descriptor used to derive this address if this is a descriptor wallet
	ParentDesc string `json:"parent_desc,omitempty"`

	// If the key is a script.
	IsScript bool `json:"isscript"`

	// If the address was used for change output.
	IsChange bool `json:"ischange"`

	// If the address is a witness address.
	IsWitness bool `json:"iswitness"`

	// The version number of the witness program.
	WitnessVersion *float64 `json:"witness_version,omitempty"`

	// The hex value of the witness program.
	WitnessProgram string `json:"witness_program,omitempty"`

	// The output script type. Only if isscript is true and the redeemscript is known. Possible
	// types: nonstandard, pubkey, pubkeyhash, scripthash, multisig, nulldata, witness_v0_keyhash,
	// witness_v0_scripthash, witness_unknown.
	Script string `json:"script,omitempty"`

	// The redeemscript for the p2sh address.
	Hex string `json:"hex,omitempty"`

	// Array of pubkeys associated with the known redeemscript (only if script is multisig).
	// Element: Str
	Pubkeys []string `json:"pubkeys,omitempty"`

	// The number of signatures required to spend multisig output (only if script is multisig).
	SigsRequired *float64 `json:"sigsrequired,omitempty"`

	// The hex value of the raw public key for single-key addresses (possibly embedded in P2SH or P2WSH).
	Pubkey string `json:"pubkey,omitempty"`

	// Information about the address embedded in P2SH or P2WSH, if relevant and known.
	// Includes all getaddressinfo output fields for the embedded address, excluding metadata (timestamp, hdkeypath, hdseedid)
	// and relation to the wallet (ismine, iswatchonly).
	Embedded *GetAddressInfoResp `json:"embedded,omitempty"`

	// If the pubkey is compressed.
	IsCompressed *bool `json:"iscompressed,omitempty"`

	// The creation time of the key, if available, expressed in UNIX epoch time.
	Timestamp *float64 `json:"timestamp,omitempty"`

	// The HD keypath, if the key is HD and available.
	HDKeyPath string `json:"hdkeypath,omitempty"`

	// The Hash160 of the HD seed.
	HDSeedID string `json:"hdseedid,omitempty"`

	// The fingerprint of the master key.
	HDMasterFingerprint string `json:"hdmasterfingerprint,omitempty"`

	// Array of labels associated with the address. Currently limited to one label but returned
	// as an array to keep the API stable if multiple labels are enabled in the future.
	// Element: Str    Label name (defaults to "").
	Labels []string `json:"labels"`
}

// GetAddressInfo RPC method.
// Return information about the given bitcoin address.
// Some of the information will only be present if the address is in the active wallet.
func (bc *BitcoindClient) GetAddressInfo(ctx context.Context, args GetAddressInfoReq) (result GetAddressInfoResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "getaddressinfo", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// GetBalanceReq holds the arguments for the GetBalance call.
//  1. dummy                (string, optional) Remains for backward compatibility. Must be excluded or set to "*".
//  2. minconf              (numeric, optional, default=0) Only include transactions confirmed at least this many times.
//  3. include_watchonly    (boolean, optional, default=true for watch-only wallets, otherwise false) Also include balance in watch-only addresses (see 'importaddress')
//  4. avoid_reuse          (boolean, optional, default=true) (only available if avoid_reuse wallet flag is set) Do not include balance in dirty outputs; addresses are considered dirty if they have previously been used in a transaction.
type GetBalanceReq struct {
	// Only include transactions confirmed at least this many times.
	// Default: 0
	MinConf float64 `json:"minconf,omitempty"`

	// Also include balance in watch-only addresses (see 'importaddress')
	// Default: true for watch-only wallets, otherwise false
	IncludeWatchOnly *bool `json:"include_watchonly,omitempty"`

	// (only available if avoid_reuse wallet flag is set) Do not include balance in dirty outputs; addresses are considered dirty if they have previously been used in a transaction.
	// Default: true
	AvoidReuse *bool `json:"avoid_reuse,omitempty"`
}

// GetBalanceResp holds the response to the GetBalance call.
//  n    (numeric) The total amount in BTC received for this wallet.
type GetBalanceResp struct {
	// The total amount in BTC received for this wallet.
	N float64
}

func (alts GetBalanceResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.N)
}

func (alts *GetBalanceResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.N) == nil {
		return nil
	}
	alts.N = reset.N
	return &UnmarshalError{B: b, structName: "GetBalanceResp"}
}

// GetBalance RPC method.
// Returns the total available balance.
// The available balance is what the wallet considers currently spendable, and is
// thus affected by options which limit spendability such as -spendzeroconfchange.
func (bc *BitcoindClient) GetBalance(ctx context.Context, args GetBalanceReq) (result GetBalanceResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "getbalance", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// GetBalancesResp holds the response to the GetBalances call.
//  {                               (json object)
//    "mine" : {                    (json object) balances from outputs that the wallet can sign
//      "trusted" : n,              (numeric) trusted balance (outputs created by the wallet or confirmed outputs)
//      "untrusted_pending" : n,    (numeric) untrusted pending balance (outputs created by others that are in the mempool)
//      "immature" : n,             (numeric) balance from immature coinbase outputs
//      "used" : n                  (numeric) (only present if avoid_reuse is set) balance from coins sent to addresses that were previously spent from (potentially privacy violating)
//    },
//    "watchonly" : {               (json object) watchonly balances (not present if wallet does not watch anything)
//      "trusted" : n,              (numeric) trusted balance (outputs created by the wallet or confirmed outputs)
//      "untrusted_pending" : n,    (numeric) untrusted pending balance (outputs created by others that are in the mempool)
//      "immature" : n              (numeric) balance from immature coinbase outputs
//    }
//  }
type GetBalancesResp struct {
	// balances from outputs that the wallet can sign
	Mine struct {
		// trusted balance (outputs created by the wallet or confirmed outputs)
		Trusted float64 `json:"trusted"`

		// untrusted pending balance (outputs created by others that are in the mempool)
		UntrustedPending float64 `json:"untrusted_pending"`

		// balance from immature coinbase outputs
		Immature float64 `json:"immature"`

		// (only present if avoid_reuse is set) balance from coins sent to addresses that were previously spent from (potentially privacy violating)
		Used float64 `json:"used"`
	} `json:"mine"`

	// watchonly balances (not present if wallet does not watch anything)
	WatchOnly struct {
		// trusted balance (outputs created by the wallet or confirmed outputs)
		Trusted float64 `json:"trusted"`

		// untrusted pending balance (outputs created by others that are in the mempool)
		UntrustedPending float64 `json:"untrusted_pending"`

		// balance from immature coinbase outputs
		Immature float64 `json:"immature"`
	} `json:"watchonly"`
}

// GetBalances RPC method.
// Returns an object with all balances in BTC.
func (bc *BitcoindClient) GetBalances(ctx context.Context) (result GetBalancesResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "getbalances", nil); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// GetNewAddressReq holds the arguments for the GetNewAddress call.
//  1. label           (string, optional, default="") The label name for the address to be linked to. It can also be set to the empty string "" to represent the default label. The label does not need to exist, it will be created if there is no label by the given name.
//  2. address_type    (string, optional, default=set by -addresstype) The address type to use. Options are "legacy", "p2sh-segwit", and "bech32".
type GetNewAddressReq struct {
	// The label name for the address to be linked to. It can also be set to the empty string "" to represent the default label. The label does not need to exist, it will be created if there is no label by the given name.
	// Default: ""
	Label string `json:"label,omitempty"`

	// The address type to use. Options are "legacy", "p2sh-segwit", and "bech32".
	// Default: set by -addresstype
	AddressType string `json:"address_type,omitempty"`
}

// GetNewAddressResp holds the response to the GetNewAddress call.
//  "str"    (string) The new bitcoin address
type GetNewAddressResp struct {
	// The new bitcoin address
	Str string
}

func (alts GetNewAddressResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.Str)
}

func (alts *GetNewAddressResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Str) == nil {
		return nil
	}
	alts.Str = reset.Str
	return &UnmarshalError{B: b, structName: "GetNewAddressResp"}
}

// GetNewAddress RPC method.
// Returns a new Bitcoin address for receiving payments.
// If 'label' is specified, it is added to the address book
// so payments received with the address will be associated with 'label'.
func (bc *BitcoindClient) GetNewAddress(ctx context.Context, args GetNewAddressReq) (result GetNewAddressResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "getnewaddress", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// GetRawChangeaddressReq holds the arguments for the GetRawChangeaddress call.
//  1. address_type    (string, optional, default=set by -changetype) The address type to use. Options are "legacy", "p2sh-segwit", and "bech32".
type GetRawChangeaddressReq struct {
	// The address type to use. Options are "legacy", "p2sh-segwit", and "bech32".
	// Default: set by -changetype
	AddressType string `json:"address_type,omitempty"`
}

// GetRawChangeaddressResp holds the response to the GetRawChangeaddress call.
//  "str"    (string) The address
type GetRawChangeaddressResp struct {
	// The address
	Str string
}

func (alts GetRawChangeaddressResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.Str)
}

func (alts *GetRawChangeaddressResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Str) == nil {
		return nil
	}
	alts.Str = reset.Str
	return &UnmarshalError{B: b, structName: "GetRawChangeaddressResp"}
}

// GetRawChangeaddress RPC method.
// Returns a new Bitcoin address, for receiving change.
// This is for use with raw transactions, NOT normal use.
func (bc *BitcoindClient) GetRawChangeaddress(ctx context.Context, args GetRawChangeaddressReq) (result GetRawChangeaddressResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "getrawchangeaddress", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// GetReceivedByAddressReq holds the arguments for the GetReceivedByAddress call.
//  1. address    (string, required) The bitcoin address for transactions.
//  2. minconf    (numeric, optional, default=1) Only include transactions confirmed at least this many times.
type GetReceivedByAddressReq struct {
	// The bitcoin address for transactions.
	Address string `json:"address"`

	// Only include transactions confirmed at least this many times.
	// Default: 1
	MinConf *float64 `json:"minconf,omitempty"`
}

// GetReceivedByAddressResp holds the response to the GetReceivedByAddress call.
//  n    (numeric) The total amount in BTC received at this address.
type GetReceivedByAddressResp struct {
	// The total amount in BTC received at this address.
	N float64
}

func (alts GetReceivedByAddressResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.N)
}

func (alts *GetReceivedByAddressResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.N) == nil {
		return nil
	}
	alts.N = reset.N
	return &UnmarshalError{B: b, structName: "GetReceivedByAddressResp"}
}

// GetReceivedByAddress RPC method.
// Returns the total amount received by the given address in transactions with at least minconf confirmations.
func (bc *BitcoindClient) GetReceivedByAddress(ctx context.Context, args GetReceivedByAddressReq) (result GetReceivedByAddressResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "getreceivedbyaddress", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// GetReceivedByLabelReq holds the arguments for the GetReceivedByLabel call.
//  1. label      (string, required) The selected label, may be the default label using "".
//  2. minconf    (numeric, optional, default=1) Only include transactions confirmed at least this many times.
type GetReceivedByLabelReq struct {
	// The selected label, may be the default label using "".
	Label string `json:"label"`

	// Only include transactions confirmed at least this many times.
	// Default: 1
	MinConf *float64 `json:"minconf,omitempty"`
}

// GetReceivedByLabelResp holds the response to the GetReceivedByLabel call.
//  n    (numeric) The total amount in BTC received for this label.
type GetReceivedByLabelResp struct {
	// The total amount in BTC received for this label.
	N float64
}

func (alts GetReceivedByLabelResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.N)
}

func (alts *GetReceivedByLabelResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.N) == nil {
		return nil
	}
	alts.N = reset.N
	return &UnmarshalError{B: b, structName: "GetReceivedByLabelResp"}
}

// GetReceivedByLabel RPC method.
// Returns the total amount received by addresses with <label> in transactions with at least [minconf] confirmations.
func (bc *BitcoindClient) GetReceivedByLabel(ctx context.Context, args GetReceivedByLabelReq) (result GetReceivedByLabelResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "getreceivedbylabel", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// GetTransactionReq holds the arguments for the GetTransaction call.
//  1. txid                 (string, required) The transaction id
//  2. include_watchonly    (boolean, optional, default=true for watch-only wallets, otherwise false) Whether to include watch-only addresses in balance calculation and details[]
//  3. verbose              (boolean, optional, default=false) Whether to include a `decoded` field containing the decoded transaction (equivalent to RPC decoderawtransaction)
type GetTransactionReq struct {
	// The transaction id
	TxID string `json:"txid"`

	// Whether to include watch-only addresses in balance calculation and details[]
	// Default: true for watch-only wallets, otherwise false
	IncludeWatchOnly *bool `json:"include_watchonly,omitempty"`

	// Whether to include a `decoded` field containing the decoded transaction (equivalent to RPC decoderawtransaction)
	// Default: false
	Verbose bool `json:"verbose,omitempty"`
}

// GetTransactionResp holds the response to the GetTransaction call.
//  {                                          (json object)
//    "amount" : n,                            (numeric) The amount in BTC
//    "fee" : n,                               (numeric) The amount of the fee in BTC. This is negative and only available for the
//                                             'send' category of transactions.
//    "confirmations" : n,                     (numeric) The number of confirmations for the transaction. Negative confirmations means the
//                                             transaction conflicted that many blocks ago.
//    "generated" : true|false,                (boolean) Only present if transaction only input is a coinbase one.
//    "trusted" : true|false,                  (boolean) Only present if we consider transaction to be trusted and so safe to spend from.
//    "blockhash" : "hex",                     (string) The block hash containing the transaction.
//    "blockheight" : n,                       (numeric) The block height containing the transaction.
//    "blockindex" : n,                        (numeric) The index of the transaction in the block that includes it.
//    "blocktime" : xxx,                       (numeric) The block time expressed in UNIX epoch time.
//    "txid" : "hex",                          (string) The transaction id.
//    "walletconflicts" : [                    (json array) Conflicting transaction ids.
//      "hex",                                 (string) The transaction id.
//      ...
//    ],
//    "time" : xxx,                            (numeric) The transaction time expressed in UNIX epoch time.
//    "timereceived" : xxx,                    (numeric) The time received expressed in UNIX epoch time.
//    "comment" : "str",                       (string) If a comment is associated with the transaction, only present if not empty.
//    "bip125-replaceable" : "str",            (string) ("yes|no|unknown") Whether this transaction could be replaced due to BIP125 (replace-by-fee);
//                                             may be unknown for unconfirmed transactions not in the mempool
//    "details" : [                            (json array)
//      {                                      (json object)
//        "involvesWatchonly" : true|false,    (boolean) Only returns true if imported addresses were involved in transaction.
//        "address" : "str",                   (string) The bitcoin address involved in the transaction.
//        "category" : "str",                  (string) The transaction category.
//                                             "send"                  Transactions sent.
//                                             "receive"               Non-coinbase transactions received.
//                                             "generate"              Coinbase transactions received with more than 100 confirmations.
//                                             "immature"              Coinbase transactions received with 100 or fewer confirmations.
//                                             "orphan"                Orphaned coinbase transactions received.
//        "amount" : n,                        (numeric) The amount in BTC
//        "label" : "str",                     (string) A comment for the address/transaction, if any
//        "vout" : n,                          (numeric) the vout value
//        "fee" : n,                           (numeric) The amount of the fee in BTC. This is negative and only available for the
//                                             'send' category of transactions.
//        "abandoned" : true|false             (boolean) 'true' if the transaction has been abandoned (inputs are respendable). Only available for the
//                                             'send' category of transactions.
//      },
//      ...
//    ],
//    "hex" : "hex",                           (string) Raw data for transaction
//    "decoded" : {                            (json object) Optional, the decoded transaction (only present when `verbose` is passed)
//      ...                                    Equivalent to the RPC decoderawtransaction method, or the RPC getrawtransaction method when `verbose` is passed.
//    }
//  }
type GetTransactionResp struct {
	// The amount in BTC
	Amount float64 `json:"amount"`

	// The amount of the fee in BTC. This is negative and only available for the
	// 'send' category of transactions.
	Fee float64 `json:"fee"`

	// The number of confirmations for the transaction. Negative confirmations means the
	// transaction conflicted that many blocks ago.
	Confirmations float64 `json:"confirmations"`

	// Only present if transaction only input is a coinbase one.
	Generated bool `json:"generated"`

	// Only present if we consider transaction to be trusted and so safe to spend from.
	Trusted bool `json:"trusted"`

	// The block hash containing the transaction.
	Blockhash string `json:"blockhash"`

	// The block height containing the transaction.
	BlockHeight float64 `json:"blockheight"`

	// The index of the transaction in the block that includes it.
	BlockIndex float64 `json:"blockindex"`

	// The block time expressed in UNIX epoch time.
	BlockTime float64 `json:"blocktime"`

	// The transaction id.
	TxID string `json:"txid"`

	// Conflicting transaction ids.
	// Element: Hex    The transaction id.
	WalletConflicts []string `json:"walletconflicts"`

	// The transaction time expressed in UNIX epoch time.
	Time float64 `json:"time"`

	// The time received expressed in UNIX epoch time.
	TimeReceived float64 `json:"timereceived"`

	// If a comment is associated with the transaction, only present if not empty.
	Comment string `json:"comment"`

	// ("yes|no|unknown") Whether this transaction could be replaced due to BIP125 (replace-by-fee);
	// may be unknown for unconfirmed transactions not in the mempool
	BIP125Replaceable string `json:"bip125-replaceable"`

	Details []GetTransactionRespDetails `json:"details"`

	// Raw data for transaction
	Hex string `json:"hex"`

	// Optional, the decoded transaction (only present when `verbose` is passed)
	// Equivalent to the RPC decoderawtransaction method, or the RPC getrawtransaction method when `verbose` is passed.
	Decoded DecodeRawTransactionResp `json:"decoded"`
}

type GetTransactionRespDetails struct {
	// Only returns true if imported addresses were involved in transaction.
	InvolvesWatchOnly bool `json:"involvesWatchonly"`

	// The bitcoin address involved in the transaction.
	Address string `json:"address"`

	// The transaction category.
	// "send"                  Transactions sent.
	// "receive"               Non-coinbase transactions received.
	// "generate"              Coinbase transactions received with more than 100 confirmations.
	// "immature"              Coinbase transactions received with 100 or fewer confirmations.
	// "orphan"                Orphaned coinbase transactions received.
	Category string `json:"category"`

	// The amount in BTC
	Amount float64 `json:"amount"`

	// A comment for the address/transaction, if any
	Label string `json:"label"`

	// the vout value
	Vout float64 `json:"vout"`

	// The amount of the fee in BTC. This is negative and only available for the
	// 'send' category of transactions.
	Fee float64 `json:"fee"`

	// 'true' if the transaction has been abandoned (inputs are respendable). Only available for the
	// 'send' category of transactions.
	Abandoned bool `json:"abandoned"`
}

// GetTransaction RPC method.
// Get detailed information about in-wallet transaction <txid>
func (bc *BitcoindClient) GetTransaction(ctx context.Context, args GetTransactionReq) (result GetTransactionResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "gettransaction", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// GetUnconfirmedBalanceResp holds the response to the GetUnconfirmedBalance call.
//  n    (numeric) The balance
type GetUnconfirmedBalanceResp struct {
	// The balance
	N float64
}

func (alts GetUnconfirmedBalanceResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.N)
}

func (alts *GetUnconfirmedBalanceResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.N) == nil {
		return nil
	}
	alts.N = reset.N
	return &UnmarshalError{B: b, structName: "GetUnconfirmedBalanceResp"}
}

// GetUnconfirmedBalance RPC method.
// DEPRECATED
// Identical to getbalances().mine.untrusted_pending
func (bc *BitcoindClient) GetUnconfirmedBalance(ctx context.Context) (result GetUnconfirmedBalanceResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "getunconfirmedbalance", nil); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// GetWalletInfoResp holds the response to the GetWalletInfo call.
//  {                                         (json object)
//    "walletname" : "str",                   (string) the wallet name
//    "walletversion" : n,                    (numeric) the wallet version
//    "format" : "str",                       (string) the database format (bdb or sqlite)
//    "balance" : n,                          (numeric) DEPRECATED. Identical to getbalances().mine.trusted
//    "unconfirmed_balance" : n,              (numeric) DEPRECATED. Identical to getbalances().mine.untrusted_pending
//    "immature_balance" : n,                 (numeric) DEPRECATED. Identical to getbalances().mine.immature
//    "txcount" : n,                          (numeric) the total number of transactions in the wallet
//    "keypoololdest" : xxx,                  (numeric) the UNIX epoch time of the oldest pre-generated key in the key pool. Legacy wallets only.
//    "keypoolsize" : n,                      (numeric) how many new keys are pre-generated (only counts external keys)
//    "keypoolsize_hd_internal" : n,          (numeric) how many new keys are pre-generated for internal use (used for change outputs, only appears if the wallet is using this feature, otherwise external keys are used)
//    "unlocked_until" : xxx,                 (numeric, optional) the UNIX epoch time until which the wallet is unlocked for transfers, or 0 if the wallet is locked (only present for passphrase-encrypted wallets)
//    "paytxfee" : n,                         (numeric) the transaction fee configuration, set in BTC/kvB
//    "hdseedid" : "hex",                     (string, optional) the Hash160 of the HD seed (only present when HD is enabled)
//    "private_keys_enabled" : true|false,    (boolean) false if privatekeys are disabled for this wallet (enforced watch-only wallet)
//    "avoid_reuse" : true|false,             (boolean) whether this wallet tracks clean/dirty coins in terms of reuse
//    "scanning" : {                          (json object) current scanning details, or false if no scan is in progress
//      "duration" : n,                       (numeric) elapsed seconds since scan start
//      "progress" : n                        (numeric) scanning progress percentage [0.0, 1.0]
//    },
//    "descriptors" : true|false              (boolean) whether this wallet uses descriptors for scriptPubKey management
//  }
type GetWalletInfoResp struct {
	// the wallet name
	WalletName string `json:"walletname"`

	// the wallet version
	WalletVersion float64 `json:"walletversion"`

	// the database format (bdb or sqlite)
	Format string `json:"format"`

	// DEPRECATED. Identical to getbalances().mine.trusted
	Balance float64 `json:"balance"`

	// DEPRECATED. Identical to getbalances().mine.untrusted_pending
	UnconfirmedBalance float64 `json:"unconfirmed_balance"`

	// DEPRECATED. Identical to getbalances().mine.immature
	ImmatureBalance float64 `json:"immature_balance"`

	// the total number of transactions in the wallet
	TxCount float64 `json:"txcount"`

	// the UNIX epoch time of the oldest pre-generated key in the key pool. Legacy wallets only.
	KeypoolOldest float64 `json:"keypoololdest"`

	// how many new keys are pre-generated (only counts external keys)
	KeypoolSize float64 `json:"keypoolsize"`

	// how many new keys are pre-generated for internal use (used for change outputs, only appears if the wallet is using this feature, otherwise external keys are used)
	KeypoolSizeHDInternal float64 `json:"keypoolsize_hd_internal"`

	// the UNIX epoch time until which the wallet is unlocked for transfers, or 0 if the wallet is locked (only present for passphrase-encrypted wallets)
	UnlockedUntil *float64 `json:"unlocked_until,omitempty"`

	// the transaction fee configuration, set in BTC/kvB
	PayTxFee float64 `json:"paytxfee"`

	// the Hash160 of the HD seed (only present when HD is enabled)
	HDSeedID string `json:"hdseedid,omitempty"`

	// false if privatekeys are disabled for this wallet (enforced watch-only wallet)
	PrivateKeysEnabled bool `json:"private_keys_enabled"`

	// whether this wallet tracks clean/dirty coins in terms of reuse
	AvoidReuse bool `json:"avoid_reuse"`

	// current scanning details, or false if no scan is in progress
	Scanning *GetWalletInfoRespScanning `json:"scanning,omitempty"`

	// whether this wallet uses descriptors for scriptPubKey management
	Descriptors bool `json:"descriptors"`
}

type GetWalletInfoRespScanning struct {
	GetWalletInfoRespScanningContents
}

func (mayBeFalse *GetWalletInfoRespScanning) UnmarshalJSON(b []byte) error {
	if bytes.HasPrefix(b, []byte("false")) {
		return nil
	}
	return json.Unmarshal(b, &mayBeFalse.GetWalletInfoRespScanningContents)
}

type GetWalletInfoRespScanningContents struct {
	// elapsed seconds since scan start
	Duration float64 `json:"duration"`

	// scanning progress percentage [0.0, 1.0]
	Progress float64 `json:"progress"`
}

// GetWalletInfo RPC method.
// Returns an object containing various wallet state info.
func (bc *BitcoindClient) GetWalletInfo(ctx context.Context) (result GetWalletInfoResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "getwalletinfo", nil); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// ImportAddressReq holds the arguments for the ImportAddress call.
//  1. address    (string, required) The Bitcoin address (or hex-encoded script)
//  2. label      (string, optional, default="") An optional label
//  3. rescan     (boolean, optional, default=true) Rescan the wallet for transactions
//  4. p2sh       (boolean, optional, default=false) Add the P2SH version of the script as well
type ImportAddressReq struct {
	// The Bitcoin address (or hex-encoded script)
	Address string `json:"address"`

	// An optional label
	// Default: ""
	Label string `json:"label,omitempty"`

	// Rescan the wallet for transactions
	// Default: true
	Rescan *bool `json:"rescan,omitempty"`

	// Add the P2SH version of the script as well
	// Default: false
	P2SH bool `json:"p2sh,omitempty"`
}

// ImportAddress RPC method.
// Adds an address or script (in hex) that can be watched as if it were in your wallet but cannot be used to spend. Requires a new wallet backup.
// Note: This call can take over an hour to complete if rescan is true, during that time, other rpc calls
// may report that the imported address exists but related transactions are still missing, leading to temporarily incorrect/bogus balances and unspent outputs until rescan completes.
// If you have the full public key, you should call importpubkey instead of this.
// Hint: use importmulti to import more than one address.
// Note: If you import a non-standard raw script in hex form, outputs sending to it will be treated
// as change, and not show up in many RPCs.
// Note: Use "getwalletinfo" to query the scanning progress.
func (bc *BitcoindClient) ImportAddress(ctx context.Context, args ImportAddressReq) (err error) {
	_, err = bc.sendRequest(ctx, "importaddress", args)
	return
}

// ImportDescriptorsReq holds the arguments for the ImportDescriptors call.
//  1. requests                                 (json array, required) Data to be imported
//       [
//         {                                    (json object)
//           "desc": "str",                     (string, required) Descriptor to import.
//           "active": bool,                    (boolean, optional, default=false) Set this descriptor to be the active descriptor for the corresponding output type/externality
//           "range": n or [n,n],               (numeric or array) If a ranged descriptor is used, this specifies the end or the range (in the form [begin,end]) to import
//           "next_index": n,                   (numeric) If a ranged descriptor is set to active, this specifies the next index to generate addresses from
//           "timestamp": timestamp | "now",    (integer / string, required) Time from which to start rescanning the blockchain for this descriptor, in UNIX epoch time
//                                              Use the string "now" to substitute the current synced blockchain time.
//                                              "now" can be specified to bypass scanning, for outputs which are known to never have been used, and
//                                              0 can be specified to scan the entire blockchain. Blocks up to 2 hours before the earliest timestamp
//                                              of all descriptors being imported will be scanned.
//           "internal": bool,                  (boolean, optional, default=false) Whether matching outputs should be treated as not incoming payments (e.g. change)
//           "label": "str",                    (string, optional, default="") Label to assign to the address, only allowed with internal=false
//         },
//         ...
//       ]
type ImportDescriptorsReq struct {
	// Data to be imported
	Requests []ImportDescriptorsReqRequests `json:"requests"`
}

type ImportDescriptorsReqRequests struct {
	// Descriptor to import.
	Desc string `json:"desc"`

	// Set this descriptor to be the active descriptor for the corresponding output type/externality
	// Default: false
	Active bool `json:"active,omitempty"`

	// If a ranged descriptor is used, this specifies the end or the range (in the form [begin,end]) to import
	Range [2]int64 `json:"range"`

	// If a ranged descriptor is set to active, this specifies the next index to generate addresses from
	NextIndex float64 `json:"next_index"`

	// Time from which to start rescanning the blockchain for this descriptor, in UNIX epoch time
	// Use the string "now" to substitute the current synced blockchain time.
	// "now" can be specified to bypass scanning, for outputs which are known to never have been used, and
	// 0 can be specified to scan the entire blockchain. Blocks up to 2 hours before the earliest timestamp
	// of all descriptors being imported will be scanned.
	Timestamp int64 `json:"timestamp"`

	// Whether matching outputs should be treated as not incoming payments (e.g. change)
	// Default: false
	Internal bool `json:"internal,omitempty"`

	// Label to assign to the address, only allowed with internal=false
	// Default: ""
	Label string `json:"label,omitempty"`
}

// ImportDescriptorsResp holds the response to the ImportDescriptors call.
//  [                              (json array) Response is an array with the same size as the input that has the execution result
//    {                            (json object)
//      "success" : true|false,    (boolean)
//      "warnings" : [             (json array, optional)
//        "str",                   (string)
//        ...
//      ],
//      "error" : {                (json object, optional)
//        ...                      JSONRPC error
//      }
//    },
//    ...
//  ]
type ImportDescriptorsResp struct {
	// Response is an array with the same size as the input that has the execution result
	Array []ImportDescriptorsRespElement
}

func (alts ImportDescriptorsResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.Array)
}

func (alts *ImportDescriptorsResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Array) == nil {
		return nil
	}
	alts.Array = reset.Array
	return &UnmarshalError{B: b, structName: "ImportDescriptorsResp"}
}

// Response is an array with the same size as the input that has the execution result
type ImportDescriptorsRespElement struct {
	Success bool `json:"success"`

	// Element: Str
	Warnings []string `json:"warnings,omitempty"`

	Error *JsonRPCError `json:"error,omitempty"`
}

// ImportDescriptors RPC method.
// Import descriptors. This will trigger a rescan of the blockchain based on the earliest timestamp of all descriptors being imported. Requires a new wallet backup.
// Note: This call can take over an hour to complete if using an early timestamp; during that time, other rpc calls
// may report that the imported keys, addresses or scripts exist but related transactions are still missing.
func (bc *BitcoindClient) ImportDescriptors(ctx context.Context, args ImportDescriptorsReq) (result ImportDescriptorsResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "importdescriptors", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// ImportMultiReq holds the arguments for the ImportMulti call.
//  1. requests                                                         (json array, required) Data to be imported
//       [
//         {                                                            (json object)
//           "desc": "str",                                             (string) Descriptor to import. If using descriptor, do not also provide address/scriptPubKey, scripts, or pubkeys
//           "scriptPubKey": "<script>" | { "address":"<address>" },    (string / json, required) Type of scriptPubKey (string for script, json for address). Should not be provided if using a descriptor
//           "timestamp": timestamp | "now",                            (integer / string, required) Creation time of the key expressed in UNIX epoch time,
//                                                                      or the string "now" to substitute the current synced blockchain time. The timestamp of the oldest
//                                                                      key will determine how far back blockchain rescans need to begin for missing wallet transactions.
//                                                                      "now" can be specified to bypass scanning, for keys which are known to never have been used, and
//                                                                      0 can be specified to scan the entire blockchain. Blocks up to 2 hours before the earliest key
//                                                                      creation time of all keys being imported by the importmulti call will be scanned.
//           "redeemscript": "str",                                     (string) Allowed only if the scriptPubKey is a P2SH or P2SH-P2WSH address/scriptPubKey
//           "witnessscript": "str",                                    (string) Allowed only if the scriptPubKey is a P2SH-P2WSH or P2WSH address/scriptPubKey
//           "pubkeys": [                                               (json array, optional, default=[]) Array of strings giving pubkeys to import. They must occur in P2PKH or P2WPKH scripts. They are not required when the private key is also provided (see the "keys" argument).
//             "pubKey",                                                (string)
//             ...
//           ],
//           "keys": [                                                  (json array, optional, default=[]) Array of strings giving private keys to import. The corresponding public keys must occur in the output or redeemscript.
//             "key",                                                   (string)
//             ...
//           ],
//           "range": n or [n,n],                                       (numeric or array) If a ranged descriptor is used, this specifies the end or the range (in the form [begin,end]) to import
//           "internal": bool,                                          (boolean, optional, default=false) Stating whether matching outputs should be treated as not incoming payments (also known as change)
//           "watchonly": bool,                                         (boolean, optional, default=false) Stating whether matching outputs should be considered watchonly.
//           "label": "str",                                            (string, optional, default="") Label to assign to the address, only allowed with internal=false
//           "keypool": bool,                                           (boolean, optional, default=false) Stating whether imported public keys should be added to the keypool for when users request new addresses. Only allowed when wallet private keys are disabled
//         },
//         ...
//       ]
//  2. options                                                          (json object, optional)
//       {
//         "rescan": bool,                                              (boolean, optional, default=true) Stating if should rescan the blockchain after all imports
//       }
type ImportMultiReq struct {
	// Data to be imported
	Requests []ImportMultiReqRequests `json:"requests"`

	Options *ImportMultiReqOptions `json:"options,omitempty"`
}

type ImportMultiReqRequests struct {
	// Descriptor to import. If using descriptor, do not also provide address/scriptPubKey, scripts, or pubkeys
	Desc string `json:"desc"`

	// Type of scriptPubKey (string for script, json for address). Should not be provided if using a descriptor
	ScriptPubkey string `json:"scriptPubKey"`

	// Creation time of the key expressed in UNIX epoch time,
	// or the string "now" to substitute the current synced blockchain time. The timestamp of the oldest
	// key will determine how far back blockchain rescans need to begin for missing wallet transactions.
	// "now" can be specified to bypass scanning, for keys which are known to never have been used, and
	// 0 can be specified to scan the entire blockchain. Blocks up to 2 hours before the earliest key
	// creation time of all keys being imported by the importmulti call will be scanned.
	Timestamp int64 `json:"timestamp"`

	// Allowed only if the scriptPubKey is a P2SH or P2SH-P2WSH address/scriptPubKey
	RedeemScript string `json:"redeemscript"`

	// Allowed only if the scriptPubKey is a P2SH-P2WSH or P2WSH address/scriptPubKey
	WitnessScript string `json:"witnessscript"`

	// Array of strings giving pubkeys to import. They must occur in P2PKH or P2WPKH scripts. They are not required when the private key is also provided (see the "keys" argument).
	// Element: Pubkey
	Pubkeys []string `json:"pubkeys,omitempty"`

	// Array of strings giving private keys to import. The corresponding public keys must occur in the output or redeemscript.
	// Element: Key
	Keys []string `json:"keys,omitempty"`

	// If a ranged descriptor is used, this specifies the end or the range (in the form [begin,end]) to import
	Range [2]int64 `json:"range"`

	// Stating whether matching outputs should be treated as not incoming payments (also known as change)
	// Default: false
	Internal bool `json:"internal,omitempty"`

	// Stating whether matching outputs should be considered watchonly.
	// Default: false
	WatchOnly bool `json:"watchonly,omitempty"`

	// Label to assign to the address, only allowed with internal=false
	// Default: ""
	Label string `json:"label,omitempty"`

	// Stating whether imported public keys should be added to the keypool for when users request new addresses. Only allowed when wallet private keys are disabled
	// Default: false
	Keypool bool `json:"keypool,omitempty"`
}

type ImportMultiReqOptions struct {
	// Stating if should rescan the blockchain after all imports
	// Default: true
	Rescan *bool `json:"rescan,omitempty"`
}

// ImportMultiResp holds the response to the ImportMulti call.
//  [                              (json array) Response is an array with the same size as the input that has the execution result
//    {                            (json object)
//      "success" : true|false,    (boolean)
//      "warnings" : [             (json array, optional)
//        "str",                   (string)
//        ...
//      ],
//      "error" : {                (json object, optional)
//        ...                      JSONRPC error
//      }
//    },
//    ...
//  ]
type ImportMultiResp struct {
	// Response is an array with the same size as the input that has the execution result
	Array []ImportMultiRespElement
}

func (alts ImportMultiResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.Array)
}

func (alts *ImportMultiResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Array) == nil {
		return nil
	}
	alts.Array = reset.Array
	return &UnmarshalError{B: b, structName: "ImportMultiResp"}
}

// Response is an array with the same size as the input that has the execution result
type ImportMultiRespElement struct {
	Success bool `json:"success"`

	// Element: Str
	Warnings []string `json:"warnings,omitempty"`

	Error *JsonRPCError `json:"error,omitempty"`
}

// ImportMulti RPC method.
// Import addresses/scripts (with private or public keys, redeem script (P2SH)), optionally rescanning the blockchain from the earliest creation time of the imported scripts. Requires a new wallet backup.
// If an address/script is imported without all of the private keys required to spend from that address, it will be watchonly. The 'watchonly' option must be set to true in this case or a warning will be returned.
// Conversely, if all the private keys are provided and the address/script is spendable, the watchonly option must be set to false, or a warning will be returned.
// Note: This call can take over an hour to complete if rescan is true, during that time, other rpc calls
// may report that the imported keys, addresses or scripts exist but related transactions are still missing.
// Note: Use "getwalletinfo" to query the scanning progress.
func (bc *BitcoindClient) ImportMulti(ctx context.Context, args ImportMultiReq) (result ImportMultiResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "importmulti", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// ImportPrivkeyReq holds the arguments for the ImportPrivkey call.
//  1. privkey    (string, required) The private key (see dumpprivkey)
//  2. label      (string, optional, default=current label if address exists, otherwise "") An optional label
//  3. rescan     (boolean, optional, default=true) Rescan the wallet for transactions
type ImportPrivkeyReq struct {
	// The private key (see dumpprivkey)
	Privkey string `json:"privkey"`

	// An optional label
	// Default: current label if address exists, otherwise ""
	Label string `json:"label,omitempty"`

	// Rescan the wallet for transactions
	// Default: true
	Rescan *bool `json:"rescan,omitempty"`
}

// ImportPrivkey RPC method.
// Adds a private key (as returned by dumpprivkey) to your wallet. Requires a new wallet backup.
// Hint: use importmulti to import more than one private key.
// Note: This call can take over an hour to complete if rescan is true, during that time, other rpc calls
// may report that the imported key exists but related transactions are still missing, leading to temporarily incorrect/bogus balances and unspent outputs until rescan completes.
// Note: Use "getwalletinfo" to query the scanning progress.
func (bc *BitcoindClient) ImportPrivkey(ctx context.Context, args ImportPrivkeyReq) (err error) {
	_, err = bc.sendRequest(ctx, "importprivkey", args)
	return
}

// ImportPrunedFundsReq holds the arguments for the ImportPrunedFunds call.
//  1. rawtransaction    (string, required) A raw transaction in hex funding an already-existing address in wallet
//  2. txoutproof        (string, required) The hex output from gettxoutproof that contains the transaction
type ImportPrunedFundsReq struct {
	// A raw transaction in hex funding an already-existing address in wallet
	RawTransaction string `json:"rawtransaction"`

	// The hex output from gettxoutproof that contains the transaction
	TxOutProof string `json:"txoutproof"`
}

// ImportPrunedFunds RPC method.
// Imports funds without rescan. Corresponding address or script must previously be included in wallet. Aimed towards pruned wallets. The end-user is responsible to import additional transactions that subsequently spend the imported outputs or rescan after the point in the blockchain the transaction is included.
func (bc *BitcoindClient) ImportPrunedFunds(ctx context.Context, args ImportPrunedFundsReq) (err error) {
	_, err = bc.sendRequest(ctx, "importprunedfunds", args)
	return
}

// ImportPubkeyReq holds the arguments for the ImportPubkey call.
//  1. pubkey    (string, required) The hex-encoded public key
//  2. label     (string, optional, default="") An optional label
//  3. rescan    (boolean, optional, default=true) Rescan the wallet for transactions
type ImportPubkeyReq struct {
	// The hex-encoded public key
	Pubkey string `json:"pubkey"`

	// An optional label
	// Default: ""
	Label string `json:"label,omitempty"`

	// Rescan the wallet for transactions
	// Default: true
	Rescan *bool `json:"rescan,omitempty"`
}

// ImportPubkey RPC method.
// Adds a public key (in hex) that can be watched as if it were in your wallet but cannot be used to spend. Requires a new wallet backup.
// Hint: use importmulti to import more than one public key.
// Note: This call can take over an hour to complete if rescan is true, during that time, other rpc calls
// may report that the imported pubkey exists but related transactions are still missing, leading to temporarily incorrect/bogus balances and unspent outputs until rescan completes.
// Note: Use "getwalletinfo" to query the scanning progress.
func (bc *BitcoindClient) ImportPubkey(ctx context.Context, args ImportPubkeyReq) (err error) {
	_, err = bc.sendRequest(ctx, "importpubkey", args)
	return
}

// ImportWalletReq holds the arguments for the ImportWallet call.
//  1. filename    (string, required) The wallet file
type ImportWalletReq struct {
	// The wallet file
	FileName string `json:"filename"`
}

// ImportWallet RPC method.
// Imports keys from a wallet dump file (see dumpwallet). Requires a new wallet backup to include imported keys.
// Note: Use "getwalletinfo" to query the scanning progress.
func (bc *BitcoindClient) ImportWallet(ctx context.Context, args ImportWalletReq) (err error) {
	_, err = bc.sendRequest(ctx, "importwallet", args)
	return
}

// KeypoolRefillReq holds the arguments for the KeypoolRefill call.
//  1. newsize    (numeric, optional, default=100) The new keypool size
type KeypoolRefillReq struct {
	// The new keypool size
	// Default: 100
	NewSize *float64 `json:"newsize,omitempty"`
}

// KeypoolRefill RPC method.
// Fills the keypool.
// Requires wallet passphrase to be set with walletpassphrase call if wallet is encrypted.
func (bc *BitcoindClient) KeypoolRefill(ctx context.Context, args KeypoolRefillReq) (err error) {
	_, err = bc.sendRequest(ctx, "keypoolrefill", args)
	return
}

// ListAddressGroupingsResp holds the response to the ListAddressGroupings call.
//  [               (json array)
//    [             (json array)
//      [           (json array)
//        "str",    (string) The bitcoin address
//        n,        (numeric) The amount in BTC
//        "str"     (string, optional) The label
//      ],
//      ...
//    ],
//    ...
//  ]
type ListAddressGroupingsResp struct {
	// Holder of alternative parameter formats, only one will be used, the first that is non-zero.
	Array [][][]ListAddressGroupingsRespElement
}

func (alts ListAddressGroupingsResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.Array)
}

func (alts *ListAddressGroupingsResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Array) == nil {
		return nil
	}
	alts.Array = reset.Array
	return &UnmarshalError{B: b, structName: "ListAddressGroupingsResp"}
}

// Holder of alternative parameter formats, only one will be used, the first that is non-zero.
type ListAddressGroupingsRespElement struct {
	// The bitcoin address
	Str string

	// The amount in BTC
	// "str"     (string, optional) The label
	N float64
}

func (alts ListAddressGroupingsRespElement) MarshalJSON() ([]byte, error) {
	if !reflect.ValueOf(alts.Str).IsZero() {
		return json.Marshal(alts.Str)
	}
	return json.Marshal(alts.N)
}

func (alts *ListAddressGroupingsRespElement) UnmarshalJSON(b []byte) error {
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
	if decoder.Decode(&alts.N) == nil {
		return nil
	}
	alts.N = reset.N
	return &UnmarshalError{B: b, structName: "ListAddressGroupingsRespElement"}
}

// ListAddressGroupings RPC method.
// Lists groups of addresses which have had their common ownership
// made public by common use as inputs or as the resulting change
// in past transactions
func (bc *BitcoindClient) ListAddressGroupings(ctx context.Context) (result ListAddressGroupingsResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "listaddressgroupings", nil); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// ListDescriptorsResp holds the response to the ListDescriptors call.
//  {                                 (json object)
//    "wallet_name" : "str",          (string) Name of wallet this operation was performed on
//    "descriptors" : [               (json array) Array of descriptor objects
//      {                             (json object)
//        "desc" : "str",             (string) Descriptor string representation
//        "timestamp" : n,            (numeric) The creation time of the descriptor
//        "active" : true|false,      (boolean) Activeness flag
//        "internal" : true|false,    (boolean, optional) Whether this is an internal or external descriptor; defined only for active descriptors
//        "range" : [                 (json array, optional) Defined only for ranged descriptors
//          n,                        (numeric) Range start inclusive
//          n                         (numeric) Range end inclusive
//        ],
//        "next" : n                  (numeric, optional) The next index to generate addresses from; defined only for ranged descriptors
//      },
//      ...
//    ]
//  }
type ListDescriptorsResp struct {
	// Name of wallet this operation was performed on
	WalletName string `json:"wallet_name"`

	// Array of descriptor objects
	Descriptors []ListDescriptorsRespDescriptors `json:"descriptors"`
}

type ListDescriptorsRespDescriptors struct {
	// Descriptor string representation
	Desc string `json:"desc"`

	// The creation time of the descriptor
	Timestamp float64 `json:"timestamp"`

	// Activeness flag
	Active bool `json:"active"`

	// Whether this is an internal or external descriptor; defined only for active descriptors
	Internal *bool `json:"internal,omitempty"`

	// Defined only for ranged descriptors
	// Element: N    Range start inclusive
	// n                         (numeric) Range end inclusive
	Range []float64 `json:"range,omitempty"`

	// The next index to generate addresses from; defined only for ranged descriptors
	Next *float64 `json:"next,omitempty"`
}

// ListDescriptors RPC method.
// List descriptors imported into a descriptor-enabled wallet.
func (bc *BitcoindClient) ListDescriptors(ctx context.Context) (result ListDescriptorsResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "listdescriptors", nil); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// ListLabelsReq holds the arguments for the ListLabels call.
//  1. purpose    (string, optional) Address purpose to list labels for ('send','receive'). An empty string is the same as not providing this argument.
type ListLabelsReq struct {
	// Address purpose to list labels for ('send','receive'). An empty string is the same as not providing this argument.
	Purpose string `json:"purpose,omitempty"`
}

// ListLabelsResp holds the response to the ListLabels call.
//  [           (json array)
//    "str",    (string) Label name
//    ...
//  ]
type ListLabelsResp struct {
	// Element: Str    Label name
	Str []string
}

func (alts ListLabelsResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.Str)
}

func (alts *ListLabelsResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Str) == nil {
		return nil
	}
	alts.Str = reset.Str
	return &UnmarshalError{B: b, structName: "ListLabelsResp"}
}

// ListLabels RPC method.
// Returns the list of all labels, or labels that are assigned to addresses with a specific purpose.
func (bc *BitcoindClient) ListLabels(ctx context.Context, args ListLabelsReq) (result ListLabelsResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "listlabels", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// ListLockUnspentResp holds the response to the ListLockUnspent call.
//  [                      (json array)
//    {                    (json object)
//      "txid" : "hex",    (string) The transaction id locked
//      "vout" : n         (numeric) The vout value
//    },
//    ...
//  ]
type ListLockUnspentResp struct {
	Array []ListLockUnspentRespElement
}

func (alts ListLockUnspentResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.Array)
}

func (alts *ListLockUnspentResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Array) == nil {
		return nil
	}
	alts.Array = reset.Array
	return &UnmarshalError{B: b, structName: "ListLockUnspentResp"}
}

type ListLockUnspentRespElement struct {
	// The transaction id locked
	TxID string `json:"txid"`

	// The vout value
	Vout float64 `json:"vout"`
}

// ListLockUnspent RPC method.
// Returns list of temporarily unspendable outputs.
// See the lockunspent call to lock and unlock transactions for spending.
func (bc *BitcoindClient) ListLockUnspent(ctx context.Context) (result ListLockUnspentResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "listlockunspent", nil); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// ListReceivedByAddressReq holds the arguments for the ListReceivedByAddress call.
//  1. minconf              (numeric, optional, default=1) The minimum number of confirmations before payments are included.
//  2. include_empty        (boolean, optional, default=false) Whether to include addresses that haven't received any payments.
//  3. include_watchonly    (boolean, optional, default=true for watch-only wallets, otherwise false) Whether to include watch-only addresses (see 'importaddress')
//  4. address_filter       (string, optional) If present, only return information on this address.
type ListReceivedByAddressReq struct {
	// The minimum number of confirmations before payments are included.
	// Default: 1
	MinConf *float64 `json:"minconf,omitempty"`

	// Whether to include addresses that haven't received any payments.
	// Default: false
	IncludeEmpty bool `json:"include_empty,omitempty"`

	// Whether to include watch-only addresses (see 'importaddress')
	// Default: true for watch-only wallets, otherwise false
	IncludeWatchOnly *bool `json:"include_watchonly,omitempty"`

	// If present, only return information on this address.
	AddressFilter string `json:"address_filter,omitempty"`
}

// ListReceivedByAddressResp holds the response to the ListReceivedByAddress call.
//  [                                        (json array)
//    {                                      (json object)
//      "involvesWatchonly" : true|false,    (boolean) Only returns true if imported addresses were involved in transaction
//      "address" : "str",                   (string) The receiving address
//      "amount" : n,                        (numeric) The total amount in BTC received by the address
//      "confirmations" : n,                 (numeric) The number of confirmations of the most recent transaction included
//      "label" : "str",                     (string) The label of the receiving address. The default label is ""
//      "txids" : [                          (json array)
//        "hex",                             (string) The ids of transactions received with the address
//        ...
//      ]
//    },
//    ...
//  ]
type ListReceivedByAddressResp struct {
	Array []ListReceivedByAddressRespElement
}

func (alts ListReceivedByAddressResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.Array)
}

func (alts *ListReceivedByAddressResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Array) == nil {
		return nil
	}
	alts.Array = reset.Array
	return &UnmarshalError{B: b, structName: "ListReceivedByAddressResp"}
}

type ListReceivedByAddressRespElement struct {
	// Only returns true if imported addresses were involved in transaction
	InvolvesWatchOnly bool `json:"involvesWatchonly"`

	// The receiving address
	Address string `json:"address"`

	// The total amount in BTC received by the address
	Amount float64 `json:"amount"`

	// The number of confirmations of the most recent transaction included
	Confirmations float64 `json:"confirmations"`

	// The label of the receiving address. The default label is ""
	Label string `json:"label"`

	// Element: Hex    The ids of transactions received with the address
	TxIDs []string `json:"txids"`
}

// ListReceivedByAddress RPC method.
// List balances by receiving address.
func (bc *BitcoindClient) ListReceivedByAddress(ctx context.Context, args ListReceivedByAddressReq) (result ListReceivedByAddressResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "listreceivedbyaddress", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// ListReceivedByLabelReq holds the arguments for the ListReceivedByLabel call.
//  1. minconf              (numeric, optional, default=1) The minimum number of confirmations before payments are included.
//  2. include_empty        (boolean, optional, default=false) Whether to include labels that haven't received any payments.
//  3. include_watchonly    (boolean, optional, default=true for watch-only wallets, otherwise false) Whether to include watch-only addresses (see 'importaddress')
type ListReceivedByLabelReq struct {
	// The minimum number of confirmations before payments are included.
	// Default: 1
	MinConf *float64 `json:"minconf,omitempty"`

	// Whether to include labels that haven't received any payments.
	// Default: false
	IncludeEmpty bool `json:"include_empty,omitempty"`

	// Whether to include watch-only addresses (see 'importaddress')
	// Default: true for watch-only wallets, otherwise false
	IncludeWatchOnly *bool `json:"include_watchonly,omitempty"`
}

// ListReceivedByLabelResp holds the response to the ListReceivedByLabel call.
//  [                                        (json array)
//    {                                      (json object)
//      "involvesWatchonly" : true|false,    (boolean) Only returns true if imported addresses were involved in transaction
//      "amount" : n,                        (numeric) The total amount received by addresses with this label
//      "confirmations" : n,                 (numeric) The number of confirmations of the most recent transaction included
//      "label" : "str"                      (string) The label of the receiving address. The default label is ""
//    },
//    ...
//  ]
type ListReceivedByLabelResp struct {
	Array []ListReceivedByLabelRespElement
}

func (alts ListReceivedByLabelResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.Array)
}

func (alts *ListReceivedByLabelResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Array) == nil {
		return nil
	}
	alts.Array = reset.Array
	return &UnmarshalError{B: b, structName: "ListReceivedByLabelResp"}
}

type ListReceivedByLabelRespElement struct {
	// Only returns true if imported addresses were involved in transaction
	InvolvesWatchOnly bool `json:"involvesWatchonly"`

	// The total amount received by addresses with this label
	Amount float64 `json:"amount"`

	// The number of confirmations of the most recent transaction included
	Confirmations float64 `json:"confirmations"`

	// The label of the receiving address. The default label is ""
	Label string `json:"label"`
}

// ListReceivedByLabel RPC method.
// List received transactions by label.
func (bc *BitcoindClient) ListReceivedByLabel(ctx context.Context, args ListReceivedByLabelReq) (result ListReceivedByLabelResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "listreceivedbylabel", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// ListSinceBlockReq holds the arguments for the ListSinceBlock call.
//  1. blockhash               (string, optional) If set, the block hash to list transactions since, otherwise list all transactions.
//  2. target_confirmations    (numeric, optional, default=1) Return the nth block hash from the main chain. e.g. 1 would mean the best block hash. Note: this is not used as a filter, but only affects [lastblock] in the return value
//  3. include_watchonly       (boolean, optional, default=true for watch-only wallets, otherwise false) Include transactions to watch-only addresses (see 'importaddress')
//  4. include_removed         (boolean, optional, default=true) Show transactions that were removed due to a reorg in the "removed" array
//                             (not guaranteed to work on pruned nodes)
type ListSinceBlockReq struct {
	// If set, the block hash to list transactions since, otherwise list all transactions.
	Blockhash string `json:"blockhash,omitempty"`

	// Return the nth block hash from the main chain. e.g. 1 would mean the best block hash. Note: this is not used as a filter, but only affects [lastblock] in the return value
	// Default: 1
	TargetConfirmations *float64 `json:"target_confirmations,omitempty"`

	// Include transactions to watch-only addresses (see 'importaddress')
	// Default: true for watch-only wallets, otherwise false
	IncludeWatchOnly *bool `json:"include_watchonly,omitempty"`

	// Show transactions that were removed due to a reorg in the "removed" array
	// (not guaranteed to work on pruned nodes)
	// Default: true
	IncludeRemoved *bool `json:"include_removed,omitempty"`
}

// ListSinceBlockResp holds the response to the ListSinceBlock call.
//  {                                          (json object)
//    "transactions" : [                       (json array)
//      {                                      (json object)
//        "involvesWatchonly" : true|false,    (boolean) Only returns true if imported addresses were involved in transaction.
//        "address" : "str",                   (string) The bitcoin address of the transaction.
//        "category" : "str",                  (string) The transaction category.
//                                             "send"                  Transactions sent.
//                                             "receive"               Non-coinbase transactions received.
//                                             "generate"              Coinbase transactions received with more than 100 confirmations.
//                                             "immature"              Coinbase transactions received with 100 or fewer confirmations.
//                                             "orphan"                Orphaned coinbase transactions received.
//        "amount" : n,                        (numeric) The amount in BTC. This is negative for the 'send' category, and is positive
//                                             for all other categories
//        "vout" : n,                          (numeric) the vout value
//        "fee" : n,                           (numeric) The amount of the fee in BTC. This is negative and only available for the
//                                             'send' category of transactions.
//        "confirmations" : n,                 (numeric) The number of confirmations for the transaction. Negative confirmations means the
//                                             transaction conflicted that many blocks ago.
//        "generated" : true|false,            (boolean) Only present if transaction only input is a coinbase one.
//        "trusted" : true|false,              (boolean) Only present if we consider transaction to be trusted and so safe to spend from.
//        "blockhash" : "hex",                 (string) The block hash containing the transaction.
//        "blockheight" : n,                   (numeric) The block height containing the transaction.
//        "blockindex" : n,                    (numeric) The index of the transaction in the block that includes it.
//        "blocktime" : xxx,                   (numeric) The block time expressed in UNIX epoch time.
//        "txid" : "hex",                      (string) The transaction id.
//        "walletconflicts" : [                (json array) Conflicting transaction ids.
//          "hex",                             (string) The transaction id.
//          ...
//        ],
//        "time" : xxx,                        (numeric) The transaction time expressed in UNIX epoch time.
//        "timereceived" : xxx,                (numeric) The time received expressed in UNIX epoch time.
//        "comment" : "str",                   (string) If a comment is associated with the transaction, only present if not empty.
//        "bip125-replaceable" : "str",        (string) ("yes|no|unknown") Whether this transaction could be replaced due to BIP125 (replace-by-fee);
//                                             may be unknown for unconfirmed transactions not in the mempool
//        "abandoned" : true|false,            (boolean) 'true' if the transaction has been abandoned (inputs are respendable). Only available for the
//                                             'send' category of transactions.
//        "label" : "str",                     (string) A comment for the address/transaction, if any
//        "to" : "str"                         (string) If a comment to is associated with the transaction.
//      },
//      ...
//    ],
//    "removed" : [                            (json array) <structure is the same as "transactions" above, only present if include_removed=true>
//                                             Note: transactions that were re-added in the active chain will appear as-is in this array, and may thus have a positive confirmation count.
//      ...
//    ],
//    "lastblock" : "hex"                      (string) The hash of the block (target_confirmations-1) from the best block on the main chain, or the genesis hash if the referenced block does not exist yet. This is typically used to feed back into listsinceblock the next time you call it. So you would generally use a target_confirmations of say 6, so you will be continually re-notified of transactions until they've reached 6 confirmations plus any new ones
//  }
type ListSinceBlockResp struct {
	Transactions []ListSinceBlockRespTransactions `json:"transactions"`

	// <structure is the same as "transactions" above, only present if include_removed=true>
	// Note: transactions that were re-added in the active chain will appear as-is in this array, and may thus have a positive confirmation count.
	Removed []ListSinceBlockRespTransactions `json:"removed"`

	// The hash of the block (target_confirmations-1) from the best block on the main chain, or the genesis hash if the referenced block does not exist yet. This is typically used to feed back into listsinceblock the next time you call it. So you would generally use a target_confirmations of say 6, so you will be continually re-notified of transactions until they've reached 6 confirmations plus any new ones
	LastBlock string `json:"lastblock"`
}

type ListSinceBlockRespTransactions struct {
	// Only returns true if imported addresses were involved in transaction.
	InvolvesWatchOnly bool `json:"involvesWatchonly"`

	// The bitcoin address of the transaction.
	Address string `json:"address"`

	// The transaction category.
	// "send"                  Transactions sent.
	// "receive"               Non-coinbase transactions received.
	// "generate"              Coinbase transactions received with more than 100 confirmations.
	// "immature"              Coinbase transactions received with 100 or fewer confirmations.
	// "orphan"                Orphaned coinbase transactions received.
	Category string `json:"category"`

	// The amount in BTC. This is negative for the 'send' category, and is positive
	// for all other categories
	Amount float64 `json:"amount"`

	// the vout value
	Vout float64 `json:"vout"`

	// The amount of the fee in BTC. This is negative and only available for the
	// 'send' category of transactions.
	Fee float64 `json:"fee"`

	// The number of confirmations for the transaction. Negative confirmations means the
	// transaction conflicted that many blocks ago.
	Confirmations float64 `json:"confirmations"`

	// Only present if transaction only input is a coinbase one.
	Generated bool `json:"generated"`

	// Only present if we consider transaction to be trusted and so safe to spend from.
	Trusted bool `json:"trusted"`

	// The block hash containing the transaction.
	Blockhash string `json:"blockhash"`

	// The block height containing the transaction.
	BlockHeight float64 `json:"blockheight"`

	// The index of the transaction in the block that includes it.
	BlockIndex float64 `json:"blockindex"`

	// The block time expressed in UNIX epoch time.
	BlockTime float64 `json:"blocktime"`

	// The transaction id.
	TxID string `json:"txid"`

	// Conflicting transaction ids.
	// Element: Hex    The transaction id.
	WalletConflicts []string `json:"walletconflicts"`

	// The transaction time expressed in UNIX epoch time.
	Time float64 `json:"time"`

	// The time received expressed in UNIX epoch time.
	TimeReceived float64 `json:"timereceived"`

	// If a comment is associated with the transaction, only present if not empty.
	Comment string `json:"comment"`

	// ("yes|no|unknown") Whether this transaction could be replaced due to BIP125 (replace-by-fee);
	// may be unknown for unconfirmed transactions not in the mempool
	BIP125Replaceable string `json:"bip125-replaceable"`

	// 'true' if the transaction has been abandoned (inputs are respendable). Only available for the
	// 'send' category of transactions.
	Abandoned bool `json:"abandoned"`

	// A comment for the address/transaction, if any
	Label string `json:"label"`

	// If a comment to is associated with the transaction.
	To string `json:"to"`
}

// ListSinceBlock RPC method.
// Get all transactions in blocks since block [blockhash], or all transactions if omitted.
// If "blockhash" is no longer a part of the main chain, transactions from the fork point onward are included.
// Additionally, if include_removed is set, transactions affecting the wallet which were removed are returned in the "removed" array.
func (bc *BitcoindClient) ListSinceBlock(ctx context.Context, args ListSinceBlockReq) (result ListSinceBlockResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "listsinceblock", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// ListTransactionsReq holds the arguments for the ListTransactions call.
//  1. label                (string, optional) If set, should be a valid label name to return only incoming transactions
//                          with the specified label, or "*" to disable filtering and return all transactions.
//  2. count                (numeric, optional, default=10) The number of transactions to return
//  3. skip                 (numeric, optional, default=0) The number of transactions to skip
//  4. include_watchonly    (boolean, optional, default=true for watch-only wallets, otherwise false) Include transactions to watch-only addresses (see 'importaddress')
type ListTransactionsReq struct {
	// If set, should be a valid label name to return only incoming transactions
	// with the specified label, or "*" to disable filtering and return all transactions.
	Label string `json:"label,omitempty"`

	// The number of transactions to return
	// Default: 10
	Count *float64 `json:"count,omitempty"`

	// The number of transactions to skip
	// Default: 0
	Skip float64 `json:"skip,omitempty"`

	// Include transactions to watch-only addresses (see 'importaddress')
	// Default: true for watch-only wallets, otherwise false
	IncludeWatchOnly *bool `json:"include_watchonly,omitempty"`
}

// ListTransactionsResp holds the response to the ListTransactions call.
//  [                                        (json array)
//    {                                      (json object)
//      "involvesWatchonly" : true|false,    (boolean) Only returns true if imported addresses were involved in transaction.
//      "address" : "str",                   (string) The bitcoin address of the transaction.
//      "category" : "str",                  (string) The transaction category.
//                                           "send"                  Transactions sent.
//                                           "receive"               Non-coinbase transactions received.
//                                           "generate"              Coinbase transactions received with more than 100 confirmations.
//                                           "immature"              Coinbase transactions received with 100 or fewer confirmations.
//                                           "orphan"                Orphaned coinbase transactions received.
//      "amount" : n,                        (numeric) The amount in BTC. This is negative for the 'send' category, and is positive
//                                           for all other categories
//      "label" : "str",                     (string) A comment for the address/transaction, if any
//      "vout" : n,                          (numeric) the vout value
//      "fee" : n,                           (numeric) The amount of the fee in BTC. This is negative and only available for the
//                                           'send' category of transactions.
//      "confirmations" : n,                 (numeric) The number of confirmations for the transaction. Negative confirmations means the
//                                           transaction conflicted that many blocks ago.
//      "generated" : true|false,            (boolean) Only present if transaction only input is a coinbase one.
//      "trusted" : true|false,              (boolean) Only present if we consider transaction to be trusted and so safe to spend from.
//      "blockhash" : "hex",                 (string) The block hash containing the transaction.
//      "blockheight" : n,                   (numeric) The block height containing the transaction.
//      "blockindex" : n,                    (numeric) The index of the transaction in the block that includes it.
//      "blocktime" : xxx,                   (numeric) The block time expressed in UNIX epoch time.
//      "txid" : "hex",                      (string) The transaction id.
//      "walletconflicts" : [                (json array) Conflicting transaction ids.
//        "hex",                             (string) The transaction id.
//        ...
//      ],
//      "time" : xxx,                        (numeric) The transaction time expressed in UNIX epoch time.
//      "timereceived" : xxx,                (numeric) The time received expressed in UNIX epoch time.
//      "comment" : "str",                   (string) If a comment is associated with the transaction, only present if not empty.
//      "bip125-replaceable" : "str",        (string) ("yes|no|unknown") Whether this transaction could be replaced due to BIP125 (replace-by-fee);
//                                           may be unknown for unconfirmed transactions not in the mempool
//      "abandoned" : true|false             (boolean) 'true' if the transaction has been abandoned (inputs are respendable). Only available for the
//                                           'send' category of transactions.
//    },
//    ...
//  ]
type ListTransactionsResp struct {
	Array []ListTransactionsRespElement
}

func (alts ListTransactionsResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.Array)
}

func (alts *ListTransactionsResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Array) == nil {
		return nil
	}
	alts.Array = reset.Array
	return &UnmarshalError{B: b, structName: "ListTransactionsResp"}
}

type ListTransactionsRespElement struct {
	// Only returns true if imported addresses were involved in transaction.
	InvolvesWatchOnly bool `json:"involvesWatchonly"`

	// The bitcoin address of the transaction.
	Address string `json:"address"`

	// The transaction category.
	// "send"                  Transactions sent.
	// "receive"               Non-coinbase transactions received.
	// "generate"              Coinbase transactions received with more than 100 confirmations.
	// "immature"              Coinbase transactions received with 100 or fewer confirmations.
	// "orphan"                Orphaned coinbase transactions received.
	Category string `json:"category"`

	// The amount in BTC. This is negative for the 'send' category, and is positive
	// for all other categories
	Amount float64 `json:"amount"`

	// A comment for the address/transaction, if any
	Label string `json:"label"`

	// the vout value
	Vout float64 `json:"vout"`

	// The amount of the fee in BTC. This is negative and only available for the
	// 'send' category of transactions.
	Fee float64 `json:"fee"`

	// The number of confirmations for the transaction. Negative confirmations means the
	// transaction conflicted that many blocks ago.
	Confirmations float64 `json:"confirmations"`

	// Only present if transaction only input is a coinbase one.
	Generated bool `json:"generated"`

	// Only present if we consider transaction to be trusted and so safe to spend from.
	Trusted bool `json:"trusted"`

	// The block hash containing the transaction.
	Blockhash string `json:"blockhash"`

	// The block height containing the transaction.
	BlockHeight float64 `json:"blockheight"`

	// The index of the transaction in the block that includes it.
	BlockIndex float64 `json:"blockindex"`

	// The block time expressed in UNIX epoch time.
	BlockTime float64 `json:"blocktime"`

	// The transaction id.
	TxID string `json:"txid"`

	// Conflicting transaction ids.
	// Element: Hex    The transaction id.
	WalletConflicts []string `json:"walletconflicts"`

	// The transaction time expressed in UNIX epoch time.
	Time float64 `json:"time"`

	// The time received expressed in UNIX epoch time.
	TimeReceived float64 `json:"timereceived"`

	// If a comment is associated with the transaction, only present if not empty.
	Comment string `json:"comment"`

	// ("yes|no|unknown") Whether this transaction could be replaced due to BIP125 (replace-by-fee);
	// may be unknown for unconfirmed transactions not in the mempool
	BIP125Replaceable string `json:"bip125-replaceable"`

	// 'true' if the transaction has been abandoned (inputs are respendable). Only available for the
	// 'send' category of transactions.
	Abandoned bool `json:"abandoned"`
}

// ListTransactions RPC method.
// If a label name is provided, this will return only incoming transactions paying to addresses with the specified label.
// Returns up to 'count' most recent transactions skipping the first 'from' transactions.
func (bc *BitcoindClient) ListTransactions(ctx context.Context, args ListTransactionsReq) (result ListTransactionsResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "listtransactions", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// ListUnspentReq holds the arguments for the ListUnspent call.
//  1. minconf                            (numeric, optional, default=1) The minimum confirmations to filter
//  2. maxconf                            (numeric, optional, default=9999999) The maximum confirmations to filter
//  3. addresses                          (json array, optional, default=[]) The bitcoin addresses to filter
//       [
//         "address",                     (string) bitcoin address
//         ...
//       ]
//  4. include_unsafe                     (boolean, optional, default=true) Include outputs that are not safe to spend
//                                        See description of "safe" attribute below.
//  5. query_options                      (json object, optional) JSON with query options
//       {
//         "minimumAmount": amount,       (numeric or string, optional, default="0.00") Minimum value of each UTXO in BTC
//         "maximumAmount": amount,       (numeric or string, optional, default=unlimited) Maximum value of each UTXO in BTC
//         "maximumCount": n,             (numeric, optional, default=unlimited) Maximum number of UTXOs
//         "minimumSumAmount": amount,    (numeric or string, optional, default=unlimited) Minimum sum value of all UTXOs in BTC
//       }
type ListUnspentReq struct {
	// The minimum confirmations to filter
	// Default: 1
	MinConf *float64 `json:"minconf,omitempty"`

	// The maximum confirmations to filter
	// Default: 9999999
	MaxConf *float64 `json:"maxconf,omitempty"`

	// The bitcoin addresses to filter
	// Default: []
	// Element: Address    bitcoin address
	Addresses []string `json:"addresses,omitempty"`

	// Include outputs that are not safe to spend
	// See description of "safe" attribute below.
	// Default: true
	IncludeUnsafe *bool `json:"include_unsafe,omitempty"`

	// JSON with query options
	QueryOptions *ListUnspentReqQueryOptions `json:"query_options,omitempty"`
}

type ListUnspentReqQueryOptions struct {
	// Minimum value of each UTXO in BTC
	// Default: "0.00"
	MinimumAmount *float64 `json:"minimumAmount,omitempty"`

	// Maximum value of each UTXO in BTC
	// Default: unlimited
	MaximumAmount *float64 `json:"maximumAmount,omitempty"`

	// Maximum number of UTXOs
	// Default: unlimited
	MaximumCount *float64 `json:"maximumCount,omitempty"`

	// Minimum sum value of all UTXOs in BTC
	// Default: unlimited
	MinimumSumAmount *float64 `json:"minimumSumAmount,omitempty"`
}

// ListUnspentResp holds the response to the ListUnspent call.
//  [                                (json array)
//    {                              (json object)
//      "txid" : "hex",              (string) the transaction id
//      "vout" : n,                  (numeric) the vout value
//      "address" : "str",           (string) the bitcoin address
//      "label" : "str",             (string) The associated label, or "" for the default label
//      "scriptPubKey" : "str",      (string) the script key
//      "amount" : n,                (numeric) the transaction output amount in BTC
//      "confirmations" : n,         (numeric) The number of confirmations
//      "redeemScript" : "hex",      (string) The redeemScript if scriptPubKey is P2SH
//      "witnessScript" : "str",     (string) witnessScript if the scriptPubKey is P2WSH or P2SH-P2WSH
//      "spendable" : true|false,    (boolean) Whether we have the private keys to spend this output
//      "solvable" : true|false,     (boolean) Whether we know how to spend this output, ignoring the lack of keys
//      "reused" : true|false,       (boolean) (only present if avoid_reuse is set) Whether this output is reused/dirty (sent to an address that was previously spent from)
//      "desc" : "str",              (string) (only when solvable) A descriptor for spending this output
//      "safe" : true|false          (boolean) Whether this output is considered safe to spend. Unconfirmed transactions
//                                   from outside keys and unconfirmed replacement transactions are considered unsafe
//                                   and are not eligible for spending by fundrawtransaction and sendtoaddress.
//    },
//    ...
//  ]
type ListUnspentResp struct {
	Array []ListUnspentRespElement
}

func (alts ListUnspentResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.Array)
}

func (alts *ListUnspentResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Array) == nil {
		return nil
	}
	alts.Array = reset.Array
	return &UnmarshalError{B: b, structName: "ListUnspentResp"}
}

type ListUnspentRespElement struct {
	// the transaction id
	TxID string `json:"txid"`

	// the vout value
	Vout float64 `json:"vout"`

	// the bitcoin address
	Address string `json:"address"`

	// The associated label, or "" for the default label
	Label string `json:"label"`

	// the script key
	ScriptPubkey string `json:"scriptPubKey"`

	// the transaction output amount in BTC
	Amount float64 `json:"amount"`

	// The number of confirmations
	Confirmations float64 `json:"confirmations"`

	// The redeemScript if scriptPubKey is P2SH
	RedeemScript string `json:"redeemScript"`

	// witnessScript if the scriptPubKey is P2WSH or P2SH-P2WSH
	WitnessScript string `json:"witnessScript"`

	// Whether we have the private keys to spend this output
	Spendable bool `json:"spendable"`

	// Whether we know how to spend this output, ignoring the lack of keys
	Solvable bool `json:"solvable"`

	// (only present if avoid_reuse is set) Whether this output is reused/dirty (sent to an address that was previously spent from)
	Reused bool `json:"reused"`

	// (only when solvable) A descriptor for spending this output
	Desc string `json:"desc"`

	// Whether this output is considered safe to spend. Unconfirmed transactions
	// from outside keys and unconfirmed replacement transactions are considered unsafe
	// and are not eligible for spending by fundrawtransaction and sendtoaddress.
	Safe bool `json:"safe"`
}

// ListUnspent RPC method.
// Returns array of unspent transaction outputs
// with between minconf and maxconf (inclusive) confirmations.
// Optionally filter to only include txouts paid to specified addresses.
func (bc *BitcoindClient) ListUnspent(ctx context.Context, args ListUnspentReq) (result ListUnspentResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "listunspent", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// ListWalletDirResp holds the response to the ListWalletDir call.
//  {                        (json object)
//    "wallets" : [          (json array)
//      {                    (json object)
//        "name" : "str"     (string) The wallet name
//      },
//      ...
//    ]
//  }
type ListWalletDirResp struct {
	Wallets []ListWalletDirRespWallets `json:"wallets"`
}

type ListWalletDirRespWallets struct {
	// The wallet name
	Name string `json:"name"`
}

// ListWalletDir RPC method.
// Returns a list of wallets in the wallet directory.
func (bc *BitcoindClient) ListWalletDir(ctx context.Context) (result ListWalletDirResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "listwalletdir", nil); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// ListWalletsResp holds the response to the ListWallets call.
//  [           (json array)
//    "str",    (string) the wallet name
//    ...
//  ]
type ListWalletsResp struct {
	// Element: Str    the wallet name
	Str []string
}

func (alts ListWalletsResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.Str)
}

func (alts *ListWalletsResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Str) == nil {
		return nil
	}
	alts.Str = reset.Str
	return &UnmarshalError{B: b, structName: "ListWalletsResp"}
}

// ListWallets RPC method.
// Returns a list of currently loaded wallets.
// For full information on the wallet, use "getwalletinfo"
func (bc *BitcoindClient) ListWallets(ctx context.Context) (result ListWalletsResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "listwallets", nil); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// LoadWalletReq holds the arguments for the LoadWallet call.
//  1. filename           (string, required) The wallet directory or .dat file.
//  2. load_on_startup    (boolean, optional) Save wallet name to persistent settings and load on startup. True to add wallet to startup list, false to remove, null to leave unchanged.
type LoadWalletReq struct {
	// The wallet directory or .dat file.
	FileName string `json:"filename"`

	// Save wallet name to persistent settings and load on startup. True to add wallet to startup list, false to remove, null to leave unchanged.
	LoadOnStartup *bool `json:"load_on_startup,omitempty"`
}

// LoadWalletResp holds the response to the LoadWallet call.
//  {                       (json object)
//    "name" : "str",       (string) The wallet name if loaded successfully.
//    "warning" : "str"     (string) Warning message if wallet was not loaded cleanly.
//  }
type LoadWalletResp struct {
	// The wallet name if loaded successfully.
	Name string `json:"name"`

	// Warning message if wallet was not loaded cleanly.
	Warning string `json:"warning"`
}

// LoadWallet RPC method.
// Loads a wallet from a wallet file or directory.
// Note that all wallet command-line options used when starting bitcoind will be
// applied to the new wallet (eg -rescan, etc).
func (bc *BitcoindClient) LoadWallet(ctx context.Context, args LoadWalletReq) (result LoadWalletResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "loadwallet", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// LockUnspentReq holds the arguments for the LockUnspent call.
//  1. unlock                  (boolean, required) Whether to unlock (true) or lock (false) the specified transactions
//  2. transactions            (json array, optional, default=[]) The transaction outputs and within each, the txid (string) vout (numeric).
//       [
//         {                   (json object)
//           "txid": "hex",    (string, required) The transaction id
//           "vout": n,        (numeric, required) The output number
//         },
//         ...
//       ]
type LockUnspentReq struct {
	// Whether to unlock (true) or lock (false) the specified transactions
	Unlock bool `json:"unlock"`

	// The transaction outputs and within each, the txid (string) vout (numeric).
	// Default: []
	Transactions []LockUnspentReqTransactions `json:"transactions,omitempty"`
}

type LockUnspentReqTransactions struct {
	// The transaction id
	TxID string `json:"txid"`

	// The output number
	Vout float64 `json:"vout"`
}

// LockUnspentResp holds the response to the LockUnspent call.
//  true|false    (boolean) Whether the command was successful or not
type LockUnspentResp struct {
	// Whether the command was successful or not
	TrueOrFalse bool
}

func (alts LockUnspentResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.TrueOrFalse)
}

func (alts *LockUnspentResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.TrueOrFalse) == nil {
		return nil
	}
	alts.TrueOrFalse = reset.TrueOrFalse
	return &UnmarshalError{B: b, structName: "LockUnspentResp"}
}

// LockUnspent RPC method.
// Updates list of temporarily unspendable outputs.
// Temporarily lock (unlock=false) or unlock (unlock=true) specified transaction outputs.
// If no transaction outputs are specified when unlocking then all current locked transaction outputs are unlocked.
// A locked transaction output will not be chosen by automatic coin selection, when spending bitcoins.
// Manually selected coins are automatically unlocked.
// Locks are stored in memory only. Nodes start with zero locked outputs, and the locked output list
// is always cleared (by virtue of process exit) when a node stops or fails.
// Also see the listunspent call
func (bc *BitcoindClient) LockUnspent(ctx context.Context, args LockUnspentReq) (result LockUnspentResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "lockunspent", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// PsbtBumpFeeReq holds the arguments for the PsbtBumpFee call.
//  1. txid                           (string, required) The txid to be bumped
//  2. options                        (json object, optional)
//       {
//         "conf_target": n,          (numeric, optional, default=wallet -txconfirmtarget) Confirmation target in blocks
type PsbtBumpFeeReq struct {
	// The txid to be bumped
	TxID string `json:"txid"`

	Options *PsbtBumpFeeReqOptions `json:"options,omitempty"`
}

type PsbtBumpFeeReqOptions struct {
	// Confirmation target in blocks
	// Default: wallet -txconfirmtarget
	ConfTarget *float64 `json:"conf_target,omitempty"`
}

// PsbtBumpFeeResp holds the response to the PsbtBumpFee call.
//  {                    (json object)
//    "psbt" : "str",    (string) The base64-encoded unsigned PSBT of the new transaction.
//    "origfee" : n,     (numeric) The fee of the replaced transaction.
//    "fee" : n,         (numeric) The fee of the new transaction.
//    "errors" : [       (json array) Errors encountered during processing (may be empty).
//      "str",           (string)
//      ...
//    ]
//  }
type PsbtBumpFeeResp struct {
	// The base64-encoded unsigned PSBT of the new transaction.
	Psbt string `json:"psbt"`

	// The fee of the replaced transaction.
	OrigFee float64 `json:"origfee"`

	// The fee of the new transaction.
	Fee float64 `json:"fee"`

	// Errors encountered during processing (may be empty).
	// Element: Str
	Errors []string `json:"errors"`
}

// PsbtBumpFee RPC method.
// Bumps the fee of an opt-in-RBF transaction T, replacing it with a new transaction B.
// Returns a PSBT instead of creating and signing a new transaction.
// An opt-in RBF transaction with the given txid must be in the wallet.
// The command will pay the additional fee by reducing change outputs or adding inputs when necessary.
// It may add a new change output if one does not already exist.
// All inputs in the original transaction will be included in the replacement transaction.
// The command will fail if the wallet or mempool contains a transaction that spends one of T's outputs.
// By default, the new fee will be calculated automatically using the estimatesmartfee RPC.
// The user can specify a confirmation target for estimatesmartfee.
// Alternatively, the user can specify a fee rate in sat/vB for the new transaction.
// At a minimum, the new fee rate must be high enough to pay an additional new relay fee (incrementalfee
// returned by getnetworkinfo) to enter the node's mempool.
// * WARNING: before version 0.21, fee_rate was in BTC/kvB. As of 0.21, fee_rate is in sat/vB. *
func (bc *BitcoindClient) PsbtBumpFee(ctx context.Context, args PsbtBumpFeeReq) (result PsbtBumpFeeResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "psbtbumpfee", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// RemovePrunedFundsReq holds the arguments for the RemovePrunedFunds call.
//  1. txid    (string, required) The hex-encoded id of the transaction you are deleting
type RemovePrunedFundsReq struct {
	// The hex-encoded id of the transaction you are deleting
	TxID string `json:"txid"`
}

// RemovePrunedFunds RPC method.
// Deletes the specified transaction from the wallet. Meant for use with pruned wallets and as a companion to importprunedfunds. This will affect wallet balances.
func (bc *BitcoindClient) RemovePrunedFunds(ctx context.Context, args RemovePrunedFundsReq) (err error) {
	_, err = bc.sendRequest(ctx, "removeprunedfunds", args)
	return
}

// RescanBlockchainReq holds the arguments for the RescanBlockchain call.
//  1. start_height    (numeric, optional, default=0) block height where the rescan should start
//  2. stop_height     (numeric, optional) the last block height that should be scanned. If none is provided it will rescan up to the tip at return time of this call.
type RescanBlockchainReq struct {
	// block height where the rescan should start
	// Default: 0
	StartHeight float64 `json:"start_height,omitempty"`

	// the last block height that should be scanned. If none is provided it will rescan up to the tip at return time of this call.
	StopHeight *float64 `json:"stop_height,omitempty"`
}

// RescanBlockchainResp holds the response to the RescanBlockchain call.
//  {                        (json object)
//    "start_height" : n,    (numeric) The block height where the rescan started (the requested height or 0)
//    "stop_height" : n      (numeric) The height of the last rescanned block. May be null in rare cases if there was a reorg and the call didn't scan any blocks because they were already scanned in the background.
//  }
type RescanBlockchainResp struct {
	// The block height where the rescan started (the requested height or 0)
	StartHeight float64 `json:"start_height"`

	// The height of the last rescanned block. May be null in rare cases if there was a reorg and the call didn't scan any blocks because they were already scanned in the background.
	StopHeight float64 `json:"stop_height"`
}

// RescanBlockchain RPC method.
// Rescan the local blockchain for wallet related transactions.
// Note: Use "getwalletinfo" to query the scanning progress.
func (bc *BitcoindClient) RescanBlockchain(ctx context.Context, args RescanBlockchainReq) (result RescanBlockchainResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "rescanblockchain", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// SendReq holds the arguments for the Send call.
//  1. outputs                               (json array, required) The outputs (key-value pairs), where none of the keys are duplicated.
//                                           That is, each address can only appear once and there can only be one 'data' object.
//                                           For convenience, a dictionary, which holds the key-value pairs directly, is also accepted.
//       [
//         {                                 (json object)
//           "address": amount,              (numeric or string, required) A key-value pair. The key (string) is the bitcoin address, the value (float or string) is the amount in BTC
//           ...
//         },
//         {                                 (json object)
//           "data": "hex",                  (string, required) A key-value pair. The key must be "data", the value is hex-encoded data
//         },
//         ...
//       ]
//  2. conf_target                           (numeric, optional, default=wallet -txconfirmtarget) Confirmation target in blocks
//  3. estimate_mode                         (string, optional, default="unset") The fee estimate mode, must be one of (case insensitive):
//                                           "unset"
//                                           "economical"
//                                           "conservative"
//  4. fee_rate                              (numeric or string, optional, default=not set, fall back to wallet fee estimation) Specify a fee rate in sat/vB.
//  5. options                               (json object, optional)
//       {
//         "add_inputs": bool,               (boolean, optional, default=false) If inputs are specified, automatically include more if they are not enough.
//         "include_unsafe": bool,           (boolean, optional, default=false) Include inputs that are not safe to spend (unconfirmed transactions from outside keys and unconfirmed replacement transactions).
//                                           Warning: the resulting transaction may become invalid if one of the unsafe inputs disappears.
//                                           If that happens, you will need to fund the transaction with different inputs and republish it.
//         "add_to_wallet": bool,            (boolean, optional, default=true) When false, returns a serialized transaction which will not be added to the wallet or broadcast
//         "change_address": "hex",          (string, optional, default=pool address) The bitcoin address to receive the change
//         "change_position": n,             (numeric, optional, default=random) The index of the change output
//         "change_type": "str",             (string, optional, default=set by -changetype) The output type to use. Only valid if change_address is not specified. Options are "legacy", "p2sh-segwit", and "bech32".
//         "conf_target": n,                 (numeric, optional, default=wallet -txconfirmtarget) Confirmation target in blocks
//         "estimate_mode": "str",           (string, optional, default="unset") The fee estimate mode, must be one of (case insensitive):
//                                           "unset"
//                                           "economical"
//                                           "conservative"
//         "fee_rate": amount,               (numeric or string, optional, default=not set, fall back to wallet fee estimation) Specify a fee rate in sat/vB.
//         "include_watching": bool,         (boolean, optional, default=true for watch-only wallets, otherwise false) Also select inputs which are watch only.
//                                           Only solvable inputs can be used. Watch-only destinations are solvable if the public key and/or output script was imported,
//                                           e.g. with 'importpubkey' or 'importmulti' with the 'pubkeys' or 'desc' field.
//         "inputs": [                       (json array, optional, default=[]) Specify inputs instead of adding them automatically. A JSON array of JSON objects
//           "txid",                         (string, required) The transaction id
//           vout,                           (numeric, required) The output number
//           sequence,                       (numeric, required) The sequence number
//           ...
//         ],
//         "locktime": n,                    (numeric, optional, default=0) Raw locktime. Non-0 value also locktime-activates inputs
//         "lock_unspents": bool,            (boolean, optional, default=false) Lock selected unspent outputs
//         "psbt": bool,                     (boolean, optional, default=automatic) Always return a PSBT, implies add_to_wallet=false.
//         "subtract_fee_from_outputs": [    (json array, optional, default=[]) Outputs to subtract the fee from, specified as integer indices.
//                                           The fee will be equally deducted from the amount of each specified output.
//                                           Those recipients will receive less bitcoins than you enter in their corresponding amount field.
//                                           If no outputs are specified here, the sender pays the fee.
//           vout_index,                     (numeric) The zero-based output index, before a change output is added.
//           ...
//         ],
//         "replaceable": bool,              (boolean, optional, default=wallet default) Marks this transaction as BIP125 replaceable.
//                                           Allows this transaction to be replaced by a transaction with higher fees
//       }
type SendReq struct {
	// The outputs (key-value pairs), where none of the keys are duplicated.
	// That is, each address can only appear once and there can only be one 'data' object.
	// For convenience, a dictionary, which holds the key-value pairs directly, is also accepted.
	Outputs []SendReqOutputs `json:"outputs"`

	// Confirmation target in blocks
	// Default: wallet -txconfirmtarget
	ConfTarget *float64 `json:"conf_target,omitempty"`

	// The fee estimate mode, must be one of (case insensitive):
	// "unset"
	// "economical"
	// "conservative"
	// Default: "unset"
	EstimateMode string `json:"estimate_mode,omitempty"`

	// Specify a fee rate in sat/vB.
	// Default: not set, fall back to wallet fee estimation
	FeeRate *float64 `json:"fee_rate,omitempty"`

	Options *SendReqOptions `json:"options,omitempty"`
}

// Holder of alternative parameter formats, only one will be used, the first that is non-zero.
type SendReqOutputs struct {
	// A key-value pair. The key (string) is the bitcoin address, the value (float or string) is the amount in BTC
	A map[string]float64

	B struct {
		// A key-value pair. The key must be "data", the value is hex-encoded data
		Data string `json:"data"`
	}
}

func (alts SendReqOutputs) MarshalJSON() ([]byte, error) {
	if !reflect.ValueOf(alts.A).IsZero() {
		return json.Marshal(alts.A)
	}
	return json.Marshal(alts.B)
}

func (alts *SendReqOutputs) UnmarshalJSON(b []byte) error {
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
	return &UnmarshalError{B: b, structName: "SendReqOutputs"}
}

type SendReqOptions struct {
	// If inputs are specified, automatically include more if they are not enough.
	// Default: false
	AddInputs bool `json:"add_inputs,omitempty"`

	// Include inputs that are not safe to spend (unconfirmed transactions from outside keys and unconfirmed replacement transactions).
	// Warning: the resulting transaction may become invalid if one of the unsafe inputs disappears.
	// If that happens, you will need to fund the transaction with different inputs and republish it.
	// Default: false
	IncludeUnsafe bool `json:"include_unsafe,omitempty"`

	// When false, returns a serialized transaction which will not be added to the wallet or broadcast
	// Default: true
	AddToWallet *bool `json:"add_to_wallet,omitempty"`

	// The bitcoin address to receive the change
	// Default: pool address
	Changeaddress string `json:"change_address,omitempty"`

	// The index of the change output
	// Default: random
	ChangePosition *float64 `json:"change_position,omitempty"`

	// The output type to use. Only valid if change_address is not specified. Options are "legacy", "p2sh-segwit", and "bech32".
	// Default: set by -changetype
	ChangeType string `json:"change_type,omitempty"`

	// Confirmation target in blocks
	// Default: wallet -txconfirmtarget
	ConfTarget *float64 `json:"conf_target,omitempty"`

	// The fee estimate mode, must be one of (case insensitive):
	// "unset"
	// "economical"
	// "conservative"
	// Default: "unset"
	EstimateMode string `json:"estimate_mode,omitempty"`

	// Specify a fee rate in sat/vB.
	// Default: not set, fall back to wallet fee estimation
	FeeRate *float64 `json:"fee_rate,omitempty"`

	// Also select inputs which are watch only.
	// Only solvable inputs can be used. Watch-only destinations are solvable if the public key and/or output script was imported,
	// e.g. with 'importpubkey' or 'importmulti' with the 'pubkeys' or 'desc' field.
	// Default: true for watch-only wallets, otherwise false
	IncludeWatching *bool `json:"include_watching,omitempty"`

	// Specify inputs instead of adding them automatically. A JSON array of JSON objects
	Inputs []SendReqOptionsInputs `json:"inputs,omitempty"`

	// Raw locktime. Non-0 value also locktime-activates inputs
	// Default: 0
	LockTime float64 `json:"locktime,omitempty"`

	// Lock selected unspent outputs
	// Default: false
	LockUnspents bool `json:"lock_unspents,omitempty"`

	// Always return a PSBT, implies add_to_wallet=false.
	// Default: automatic
	Psbt *bool `json:"psbt,omitempty"`

	// Outputs to subtract the fee from, specified as integer indices.
	// The fee will be equally deducted from the amount of each specified output.
	// Those recipients will receive less bitcoins than you enter in their corresponding amount field.
	// If no outputs are specified here, the sender pays the fee.
	// Element: VoutIndex    The zero-based output index, before a change output is added.
	SubtractFeeFromOutputs []float64 `json:"subtract_fee_from_outputs,omitempty"`

	// Marks this transaction as BIP125 replaceable.
	// Allows this transaction to be replaced by a transaction with higher fees
	// Default: wallet default
	Replaceable *bool `json:"replaceable,omitempty"`
}

// Holder of alternative parameter formats, only one will be used, the first that is non-zero.
type SendReqOptionsInputs struct {
	// The transaction id
	TxID string

	// The output number
	Vout float64

	// The sequence number
	Sequence float64
}

func (alts SendReqOptionsInputs) MarshalJSON() ([]byte, error) {
	if !reflect.ValueOf(alts.TxID).IsZero() {
		return json.Marshal(alts.TxID)
	}
	if !reflect.ValueOf(alts.Vout).IsZero() {
		return json.Marshal(alts.Vout)
	}
	return json.Marshal(alts.Sequence)
}

func (alts *SendReqOptionsInputs) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.TxID) == nil {
		return nil
	}
	alts.TxID = reset.TxID
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Vout) == nil {
		return nil
	}
	alts.Vout = reset.Vout
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Sequence) == nil {
		return nil
	}
	alts.Sequence = reset.Sequence
	return &UnmarshalError{B: b, structName: "SendReqOptionsInputs"}
}

// SendResp holds the response to the Send call.
//  {                             (json object)
//    "complete" : true|false,    (boolean) If the transaction has a complete set of signatures
//    "txid" : "hex",             (string) The transaction id for the send. Only 1 transaction is created regardless of the number of addresses.
//    "hex" : "hex",              (string) If add_to_wallet is false, the hex-encoded raw transaction with signature(s)
//    "psbt" : "str"              (string) If more signatures are needed, or if add_to_wallet is false, the base64-encoded (partially) signed transaction
//  }
type SendResp struct {
	// If the transaction has a complete set of signatures
	Complete bool `json:"complete"`

	// The transaction id for the send. Only 1 transaction is created regardless of the number of addresses.
	TxID string `json:"txid"`

	// If add_to_wallet is false, the hex-encoded raw transaction with signature(s)
	Hex string `json:"hex"`

	// If more signatures are needed, or if add_to_wallet is false, the base64-encoded (partially) signed transaction
	Psbt string `json:"psbt"`
}

// Send RPC method.
// EXPERIMENTAL warning: this call may be changed in future releases.
// Send a transaction.
func (bc *BitcoindClient) Send(ctx context.Context, args SendReq) (result SendResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "send", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// SendManyReq holds the arguments for the SendMany call.
//  1. dummy                     (string, required) Must be set to "" for backwards compatibility.
//  2. amounts                   (json object, required) The addresses and amounts
//       {
//         "address": amount,    (numeric or string, required) The bitcoin address is the key, the numeric amount (can be string) in BTC is the value
//         ...
//       }
//  3. minconf                   (numeric, optional) Ignored dummy value
//  4. comment                   (string, optional) A comment
//  5. subtractfeefrom           (json array, optional) The addresses.
//                               The fee will be equally deducted from the amount of each selected address.
//                               Those recipients will receive less bitcoins than you enter in their corresponding amount field.
//                               If no addresses are specified here, the sender pays the fee.
//       [
//         "address",            (string) Subtract fee from this address
//         ...
//       ]
//  6. replaceable               (boolean, optional, default=wallet default) Allow this transaction to be replaced by a transaction with higher fees via BIP 125
//  7. conf_target               (numeric, optional, default=wallet -txconfirmtarget) Confirmation target in blocks
//  8. estimate_mode             (string, optional, default="unset") The fee estimate mode, must be one of (case insensitive):
//                               "unset"
//                               "economical"
//                               "conservative"
//  9. fee_rate                  (numeric or string, optional, default=not set, fall back to wallet fee estimation) Specify a fee rate in sat/vB.
//  10. verbose                  (boolean, optional, default=false) If true, return extra infomration about the transaction.
type SendManyReq struct {
	// Must be set to "" for backwards compatibility.
	Dummy string `json:"dummy"`

	// The addresses and amounts
	Amounts map[string]float64 `json:"amounts"`

	// A comment
	Comment string `json:"comment,omitempty"`

	// The addresses.
	// The fee will be equally deducted from the amount of each selected address.
	// Those recipients will receive less bitcoins than you enter in their corresponding amount field.
	// If no addresses are specified here, the sender pays the fee.
	// Element: Address    Subtract fee from this address
	SubtractFeeFrom []string `json:"subtractfeefrom,omitempty"`

	// Allow this transaction to be replaced by a transaction with higher fees via BIP 125
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

	// Specify a fee rate in sat/vB.
	// Default: not set, fall back to wallet fee estimation
	FeeRate *float64 `json:"fee_rate,omitempty"`

	// If true, return extra infomration about the transaction.
	// Default: false
	Verbose bool `json:"verbose,omitempty"`
}

// SendManyResp holds the response to the SendMany call.
//
// ALTERNATIVE (if verbose is not set or set to false)
//  "hex"    (string) The transaction id for the send. Only 1 transaction is created regardless of
//           the number of addresses.
//
// ALTERNATIVE (if verbose is set to true)
//  {                          (json object)
//    "txid" : "hex",          (string) The transaction id for the send. Only 1 transaction is created regardless of
//                             the number of addresses.
//    "fee reason" : "str"     (string) The transaction fee reason.
//  }
type SendManyResp struct {
	// The transaction id for the send. Only 1 transaction is created regardless of
	// the number of addresses.
	Hex string

	IfVerboseIsSetToTrue SendManyRespIfVerboseIsSetToTrue
}

func (alts SendManyResp) MarshalJSON() ([]byte, error) {
	if !reflect.ValueOf(alts.Hex).IsZero() {
		return json.Marshal(alts.Hex)
	}
	return json.Marshal(alts.IfVerboseIsSetToTrue)
}

func (alts *SendManyResp) UnmarshalJSON(b []byte) error {
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
	if decoder.Decode(&alts.IfVerboseIsSetToTrue) == nil {
		return nil
	}
	alts.IfVerboseIsSetToTrue = reset.IfVerboseIsSetToTrue
	return &UnmarshalError{B: b, structName: "SendManyResp"}
}

type SendManyRespIfVerboseIsSetToTrue struct {
	// The transaction id for the send. Only 1 transaction is created regardless of
	// the number of addresses.
	TxID string `json:"txid"`

	// The transaction fee reason.
	FeeReason string `json:"fee reason"`
}

// SendMany RPC method.
// Send multiple times. Amounts are double-precision floating point numbers.
// Requires wallet passphrase to be set with walletpassphrase call if wallet is encrypted.
func (bc *BitcoindClient) SendMany(ctx context.Context, args SendManyReq) (result SendManyResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "sendmany", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// SendToAddressReq holds the arguments for the SendToAddress call.
//  1. address                  (string, required) The bitcoin address to send to.
//  2. amount                   (numeric or string, required) The amount in BTC to send. eg 0.1
//  3. comment                  (string, optional) A comment used to store what the transaction is for.
//                              This is not part of the transaction, just kept in your wallet.
//  4. comment_to               (string, optional) A comment to store the name of the person or organization
//                              to which you're sending the transaction. This is not part of the
//                              transaction, just kept in your wallet.
//  5. subtractfeefromamount    (boolean, optional, default=false) The fee will be deducted from the amount being sent.
//                              The recipient will receive less bitcoins than you enter in the amount field.
//  6. replaceable              (boolean, optional, default=wallet default) Allow this transaction to be replaced by a transaction with higher fees via BIP 125
//  7. conf_target              (numeric, optional, default=wallet -txconfirmtarget) Confirmation target in blocks
//  8. estimate_mode            (string, optional, default="unset") The fee estimate mode, must be one of (case insensitive):
//                              "unset"
//                              "economical"
//                              "conservative"
//  9. avoid_reuse              (boolean, optional, default=true) (only available if avoid_reuse wallet flag is set) Avoid spending from dirty addresses; addresses are considered
//                              dirty if they have previously been used in a transaction. If true, this also activates avoidpartialspends, grouping outputs by their addresses.
//  10. fee_rate                (numeric or string, optional, default=not set, fall back to wallet fee estimation) Specify a fee rate in sat/vB.
//  11. verbose                 (boolean, optional, default=false) If true, return extra information about the transaction.
type SendToAddressReq struct {
	// The bitcoin address to send to.
	Address string `json:"address"`

	// The amount in BTC to send. eg 0.1
	Amount float64 `json:"amount"`

	// A comment used to store what the transaction is for.
	// This is not part of the transaction, just kept in your wallet.
	Comment string `json:"comment,omitempty"`

	// A comment to store the name of the person or organization
	// to which you're sending the transaction. This is not part of the
	// transaction, just kept in your wallet.
	CommentTo string `json:"comment_to,omitempty"`

	// The fee will be deducted from the amount being sent.
	// The recipient will receive less bitcoins than you enter in the amount field.
	// Default: false
	SubtractFeeFromAmount bool `json:"subtractfeefromamount,omitempty"`

	// Allow this transaction to be replaced by a transaction with higher fees via BIP 125
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

	// (only available if avoid_reuse wallet flag is set) Avoid spending from dirty addresses; addresses are considered
	// dirty if they have previously been used in a transaction. If true, this also activates avoidpartialspends, grouping outputs by their addresses.
	// Default: true
	AvoidReuse *bool `json:"avoid_reuse,omitempty"`

	// Specify a fee rate in sat/vB.
	// Default: not set, fall back to wallet fee estimation
	FeeRate *float64 `json:"fee_rate,omitempty"`

	// If true, return extra information about the transaction.
	// Default: false
	Verbose bool `json:"verbose,omitempty"`
}

// SendToAddressResp holds the response to the SendToAddress call.
//
// ALTERNATIVE (if verbose is not set or set to false)
//  "hex"    (string) The transaction id.
//
// ALTERNATIVE (if verbose is set to true)
//  {                          (json object)
//    "txid" : "hex",          (string) The transaction id.
//    "fee reason" : "str"     (string) The transaction fee reason.
//  }
type SendToAddressResp struct {
	// The transaction id.
	Hex string

	IfVerboseIsSetToTrue SendToAddressRespIfVerboseIsSetToTrue
}

func (alts SendToAddressResp) MarshalJSON() ([]byte, error) {
	if !reflect.ValueOf(alts.Hex).IsZero() {
		return json.Marshal(alts.Hex)
	}
	return json.Marshal(alts.IfVerboseIsSetToTrue)
}

func (alts *SendToAddressResp) UnmarshalJSON(b []byte) error {
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
	if decoder.Decode(&alts.IfVerboseIsSetToTrue) == nil {
		return nil
	}
	alts.IfVerboseIsSetToTrue = reset.IfVerboseIsSetToTrue
	return &UnmarshalError{B: b, structName: "SendToAddressResp"}
}

type SendToAddressRespIfVerboseIsSetToTrue struct {
	// The transaction id.
	TxID string `json:"txid"`

	// The transaction fee reason.
	FeeReason string `json:"fee reason"`
}

// SendToAddress RPC method.
// Send an amount to a given address.
// Requires wallet passphrase to be set with walletpassphrase call if wallet is encrypted.
func (bc *BitcoindClient) SendToAddress(ctx context.Context, args SendToAddressReq) (result SendToAddressResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "sendtoaddress", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// SetHDSeedReq holds the arguments for the SetHDSeed call.
//  1. newkeypool    (boolean, optional, default=true) Whether to flush old unused addresses, including change addresses, from the keypool and regenerate it.
//                   If true, the next address from getnewaddress and change address from getrawchangeaddress will be from this new seed.
//                   If false, addresses (including change addresses if the wallet already had HD Chain Split enabled) from the existing
//                   keypool will be used until it has been depleted.
//  2. seed          (string, optional, default=random seed) The WIF private key to use as the new HD seed.
//                   The seed value can be retrieved using the dumpwallet command. It is the private key marked hdseed=1
type SetHDSeedReq struct {
	// Whether to flush old unused addresses, including change addresses, from the keypool and regenerate it.
	// If true, the next address from getnewaddress and change address from getrawchangeaddress will be from this new seed.
	// If false, addresses (including change addresses if the wallet already had HD Chain Split enabled) from the existing
	// keypool will be used until it has been depleted.
	// Default: true
	NewKeypool *bool `json:"newkeypool,omitempty"`

	// The WIF private key to use as the new HD seed.
	// The seed value can be retrieved using the dumpwallet command. It is the private key marked hdseed=1
	// Default: random seed
	Seed string `json:"seed,omitempty"`
}

// SetHDSeed RPC method.
// Set or generate a new HD wallet seed. Non-HD wallets will not be upgraded to being a HD wallet. Wallets that are already
// HD will have a new HD seed set so that new keys added to the keypool will be derived from this new seed.
// Note that you will need to MAKE A NEW BACKUP of your wallet after setting the HD wallet seed.
// Requires wallet passphrase to be set with walletpassphrase call if wallet is encrypted.
func (bc *BitcoindClient) SetHDSeed(ctx context.Context, args SetHDSeedReq) (err error) {
	_, err = bc.sendRequest(ctx, "sethdseed", args)
	return
}

// SetLabelReq holds the arguments for the SetLabel call.
//  1. address    (string, required) The bitcoin address to be associated with a label.
//  2. label      (string, required) The label to assign to the address.
type SetLabelReq struct {
	// The bitcoin address to be associated with a label.
	Address string `json:"address"`

	// The label to assign to the address.
	Label string `json:"label"`
}

// SetLabel RPC method.
// Sets the label associated with the given address.
func (bc *BitcoindClient) SetLabel(ctx context.Context, args SetLabelReq) (err error) {
	_, err = bc.sendRequest(ctx, "setlabel", args)
	return
}

// SetTxFeeReq holds the arguments for the SetTxFee call.
//  1. amount    (numeric or string, required) The transaction fee rate in BTC/kvB
type SetTxFeeReq struct {
	// The transaction fee rate in BTC/kvB
	Amount float64 `json:"amount"`
}

// SetTxFeeResp holds the response to the SetTxFee call.
//  true|false    (boolean) Returns true if successful
type SetTxFeeResp struct {
	// Returns true if successful
	TrueOrFalse bool
}

func (alts SetTxFeeResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.TrueOrFalse)
}

func (alts *SetTxFeeResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.TrueOrFalse) == nil {
		return nil
	}
	alts.TrueOrFalse = reset.TrueOrFalse
	return &UnmarshalError{B: b, structName: "SetTxFeeResp"}
}

// SetTxFee RPC method.
// Set the transaction fee rate in BTC/kvB for this wallet. Overrides the global -paytxfee command line parameter.
// Can be deactivated by passing 0 as the fee. In that case automatic fee selection will be used by default.
func (bc *BitcoindClient) SetTxFee(ctx context.Context, args SetTxFeeReq) (result SetTxFeeResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "settxfee", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// SetWalletFlagReq holds the arguments for the SetWalletFlag call.
//  1. flag     (string, required) The name of the flag to change. Current available flags: avoid_reuse
//  2. value    (boolean, optional, default=true) The new state.
type SetWalletFlagReq struct {
	// The name of the flag to change. Current available flags: avoid_reuse
	Flag string `json:"flag"`

	// The new state.
	// Default: true
	Value *bool `json:"value,omitempty"`
}

// SetWalletFlagResp holds the response to the SetWalletFlag call.
//  {                               (json object)
//    "flag_name" : "str",          (string) The name of the flag that was modified
//    "flag_state" : true|false,    (boolean) The new state of the flag
//    "warnings" : "str"            (string) Any warnings associated with the change
//  }
type SetWalletFlagResp struct {
	// The name of the flag that was modified
	FlagName string `json:"flag_name"`

	// The new state of the flag
	FlagState bool `json:"flag_state"`

	// Any warnings associated with the change
	Warnings string `json:"warnings"`
}

// SetWalletFlag RPC method.
// Change the state of the given wallet flag for a wallet.
func (bc *BitcoindClient) SetWalletFlag(ctx context.Context, args SetWalletFlagReq) (result SetWalletFlagResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "setwalletflag", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// SignMessageReq holds the arguments for the SignMessage call.
//  1. address    (string, required) The bitcoin address to use for the private key.
//  2. message    (string, required) The message to create a signature of.
type SignMessageReq struct {
	// The bitcoin address to use for the private key.
	Address string `json:"address"`

	// The message to create a signature of.
	Message string `json:"message"`
}

// SignMessageResp holds the response to the SignMessage call.
//  "str"    (string) The signature of the message encoded in base 64
type SignMessageResp struct {
	// The signature of the message encoded in base 64
	Str string
}

func (alts SignMessageResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.Str)
}

func (alts *SignMessageResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Str) == nil {
		return nil
	}
	alts.Str = reset.Str
	return &UnmarshalError{B: b, structName: "SignMessageResp"}
}

// SignMessage RPC method.
// Sign a message with the private key of an address
// Requires wallet passphrase to be set with walletpassphrase call if wallet is encrypted.
func (bc *BitcoindClient) SignMessage(ctx context.Context, args SignMessageReq) (result SignMessageResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "signmessage", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// SignRawTransactionWithWalletReq holds the arguments for the SignRawTransactionWithWallet call.
//  1. hexstring                        (string, required) The transaction hex string
//  2. prevtxs                          (json array, optional) The previous dependent transaction outputs
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
//  3. sighashtype                      (string, optional, default="DEFAULT") The signature hash type. Must be one of
//                                      "DEFAULT"
//                                      "ALL"
//                                      "NONE"
//                                      "SINGLE"
//                                      "ALL|ANYONECANPAY"
//                                      "NONE|ANYONECANPAY"
//                                      "SINGLE|ANYONECANPAY"
type SignRawTransactionWithWalletReq struct {
	// The transaction hex string
	HexString string `json:"hexstring"`

	// The previous dependent transaction outputs
	PrevTxs []SignRawTransactionWithWalletReqPrevTxs `json:"prevtxs,omitempty"`

	// The signature hash type. Must be one of
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

type SignRawTransactionWithWalletReqPrevTxs struct {
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

// SignRawTransactionWithWalletResp holds the response to the SignRawTransactionWithWallet call.
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
type SignRawTransactionWithWalletResp struct {
	// The hex-encoded raw transaction with signature(s)
	Hex string `json:"hex"`

	// If the transaction has a complete set of signatures
	Complete bool `json:"complete"`

	// Script verification errors (if there are any)
	Errors []SignRawTransactionWithWalletRespErrors `json:"errors,omitempty"`
}

type SignRawTransactionWithWalletRespErrors struct {
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

// SignRawTransactionWithWallet RPC method.
// Sign inputs for raw transaction (serialized, hex-encoded).
// The second optional argument (may be null) is an array of previous transaction outputs that
// this transaction depends on but may not yet be in the block chain.
// Requires wallet passphrase to be set with walletpassphrase call if wallet is encrypted.
func (bc *BitcoindClient) SignRawTransactionWithWallet(ctx context.Context, args SignRawTransactionWithWalletReq) (result SignRawTransactionWithWalletResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "signrawtransactionwithwallet", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// UnloadWalletReq holds the arguments for the UnloadWallet call.
//  1. wallet_name        (string, optional, default=the wallet name from the RPC endpoint) The name of the wallet to unload. If provided both here and in the RPC endpoint, the two must be identical.
//  2. load_on_startup    (boolean, optional) Save wallet name to persistent settings and load on startup. True to add wallet to startup list, false to remove, null to leave unchanged.
type UnloadWalletReq struct {
	// The name of the wallet to unload. If provided both here and in the RPC endpoint, the two must be identical.
	// Default: the wallet name from the RPC endpoint
	WalletName string `json:"wallet_name,omitempty"`

	// Save wallet name to persistent settings and load on startup. True to add wallet to startup list, false to remove, null to leave unchanged.
	LoadOnStartup *bool `json:"load_on_startup,omitempty"`
}

// UnloadWalletResp holds the response to the UnloadWallet call.
//  {                       (json object)
//    "warning" : "str"     (string) Warning message if wallet was not unloaded cleanly.
//  }
type UnloadWalletResp struct {
	// Warning message if wallet was not unloaded cleanly.
	Warning string `json:"warning"`
}

// UnloadWallet RPC method.
// Unloads the wallet referenced by the request endpoint otherwise unloads the wallet specified in the argument.
// Specifying the wallet name on a wallet endpoint is invalid.
func (bc *BitcoindClient) UnloadWallet(ctx context.Context, args UnloadWalletReq) (result UnloadWalletResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "unloadwallet", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// UpgradeWalletReq holds the arguments for the UpgradeWallet call.
//  1. version    (numeric, optional, default=169900) The version number to upgrade to. Default is the latest wallet version.
type UpgradeWalletReq struct {
	// The version number to upgrade to. Default is the latest wallet version.
	// Default: 169900
	Version *float64 `json:"version,omitempty"`
}

// UpgradeWalletResp holds the response to the UpgradeWallet call.
//  {                            (json object)
//    "wallet_name" : "str",     (string) Name of wallet this operation was performed on
//    "previous_version" : n,    (numeric) Version of wallet before this operation
//    "current_version" : n,     (numeric) Version of wallet after this operation
//    "result" : "str",          (string, optional) Description of result, if no error
//    "error" : "str"            (string, optional) Error message (if there is one)
//  }
type UpgradeWalletResp struct {
	// Name of wallet this operation was performed on
	WalletName string `json:"wallet_name"`

	// Version of wallet before this operation
	PreviousVersion float64 `json:"previous_version"`

	// Version of wallet after this operation
	CurrentVersion float64 `json:"current_version"`

	// Description of result, if no error
	Result string `json:"result,omitempty"`

	// Error message (if there is one)
	Error string `json:"error,omitempty"`
}

// UpgradeWallet RPC method.
// Upgrade the wallet. Upgrades to the latest version if no version number is specified.
// New keys may be generated and a new wallet backup will need to be made.
func (bc *BitcoindClient) UpgradeWallet(ctx context.Context, args UpgradeWalletReq) (result UpgradeWalletResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "upgradewallet", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// WalletCreateFundedPsbtReq holds the arguments for the WalletCreateFundedPsbt call.
//  1. inputs                             (json array, optional) Leave empty to add inputs automatically. See add_inputs option.
//       [
//         {                              (json object)
//           "txid": "hex",               (string, required) The transaction id
//           "vout": n,                   (numeric, required) The output number
//           "sequence": n,               (numeric, optional, default=depends on the value of the 'locktime' and 'options.replaceable' arguments) The sequence number
//         },
//         ...
//       ]
//  2. outputs                            (json array, required) The outputs (key-value pairs), where none of the keys are duplicated.
//                                        That is, each address can only appear once and there can only be one 'data' object.
//                                        For compatibility reasons, a dictionary, which holds the key-value pairs directly, is also
//                                        accepted as second parameter.
//       [
//         {                              (json object)
//           "address": amount,           (numeric or string, required) A key-value pair. The key (string) is the bitcoin address, the value (float or string) is the amount in BTC
//           ...
//         },
//         {                              (json object)
//           "data": "hex",               (string, required) A key-value pair. The key must be "data", the value is hex-encoded data
//         },
//         ...
//       ]
//  3. locktime                           (numeric, optional, default=0) Raw locktime. Non-0 value also locktime-activates inputs
//  4. options                            (json object, optional)
//       {
//         "add_inputs": bool,            (boolean, optional, default=false) If inputs are specified, automatically include more if they are not enough.
//         "include_unsafe": bool,        (boolean, optional, default=false) Include inputs that are not safe to spend (unconfirmed transactions from outside keys and unconfirmed replacement transactions).
//                                        Warning: the resulting transaction may become invalid if one of the unsafe inputs disappears.
//                                        If that happens, you will need to fund the transaction with different inputs and republish it.
//         "changeAddress": "hex",        (string, optional, default=pool address) The bitcoin address to receive the change
//         "changePosition": n,           (numeric, optional, default=random) The index of the change output
//         "change_type": "str",          (string, optional, default=set by -changetype) The output type to use. Only valid if changeAddress is not specified. Options are "legacy", "p2sh-segwit", and "bech32".
//         "includeWatching": bool,       (boolean, optional, default=true for watch-only wallets, otherwise false) Also select inputs which are watch only
//         "lockUnspents": bool,          (boolean, optional, default=false) Lock selected unspent outputs
//         "fee_rate": amount,            (numeric or string, optional, default=not set, fall back to wallet fee estimation) Specify a fee rate in sat/vB.
//         "feeRate": amount,             (numeric or string, optional, default=not set, fall back to wallet fee estimation) Specify a fee rate in BTC/kvB.
//         "subtractFeeFromOutputs": [    (json array, optional, default=[]) The outputs to subtract the fee from.
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
//  5. bip32derivs                        (boolean, optional, default=true) Include BIP 32 derivation paths for public keys if we know them
type WalletCreateFundedPsbtReq struct {
	// Leave empty to add inputs automatically. See add_inputs option.
	Inputs []WalletCreateFundedPsbtReqInputs `json:"inputs,omitempty"`

	// The outputs (key-value pairs), where none of the keys are duplicated.
	// That is, each address can only appear once and there can only be one 'data' object.
	// For compatibility reasons, a dictionary, which holds the key-value pairs directly, is also
	// accepted as second parameter.
	Outputs []WalletCreateFundedPsbtReqOutputs `json:"outputs"`

	// Raw locktime. Non-0 value also locktime-activates inputs
	// Default: 0
	LockTime float64 `json:"locktime,omitempty"`

	Options *WalletCreateFundedPsbtReqOptions `json:"options,omitempty"`

	// Include BIP 32 derivation paths for public keys if we know them
	// Default: true
	BIP32Derivs *bool `json:"bip32derivs,omitempty"`
}

type WalletCreateFundedPsbtReqInputs struct {
	// The transaction id
	TxID string `json:"txid"`

	// The output number
	Vout float64 `json:"vout"`

	// The sequence number
	// Default: depends on the value of the 'locktime' and 'options.replaceable' arguments
	Sequence *float64 `json:"sequence,omitempty"`
}

// Holder of alternative parameter formats, only one will be used, the first that is non-zero.
type WalletCreateFundedPsbtReqOutputs struct {
	// A key-value pair. The key (string) is the bitcoin address, the value (float or string) is the amount in BTC
	A map[string]float64

	B struct {
		// A key-value pair. The key must be "data", the value is hex-encoded data
		Data string `json:"data"`
	}
}

func (alts WalletCreateFundedPsbtReqOutputs) MarshalJSON() ([]byte, error) {
	if !reflect.ValueOf(alts.A).IsZero() {
		return json.Marshal(alts.A)
	}
	return json.Marshal(alts.B)
}

func (alts *WalletCreateFundedPsbtReqOutputs) UnmarshalJSON(b []byte) error {
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
	return &UnmarshalError{B: b, structName: "WalletCreateFundedPsbtReqOutputs"}
}

type WalletCreateFundedPsbtReqOptions struct {
	// If inputs are specified, automatically include more if they are not enough.
	// Default: false
	AddInputs bool `json:"add_inputs,omitempty"`

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

	// Also select inputs which are watch only
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

	// The outputs to subtract the fee from.
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

// WalletCreateFundedPsbtResp holds the response to the WalletCreateFundedPsbt call.
//  {                     (json object)
//    "psbt" : "str",     (string) The resulting raw transaction (base64-encoded string)
//    "fee" : n,          (numeric) Fee in BTC the resulting transaction pays
//    "changepos" : n     (numeric) The position of the added change output, or -1
//  }
type WalletCreateFundedPsbtResp struct {
	// The resulting raw transaction (base64-encoded string)
	Psbt string `json:"psbt"`

	// Fee in BTC the resulting transaction pays
	Fee float64 `json:"fee"`

	// The position of the added change output, or -1
	ChangePos float64 `json:"changepos"`
}

// WalletCreateFundedPsbt RPC method.
// Creates and funds a transaction in the Partially Signed Transaction format.
// Implements the Creator and Updater roles.
func (bc *BitcoindClient) WalletCreateFundedPsbt(ctx context.Context, args WalletCreateFundedPsbtReq) (result WalletCreateFundedPsbtResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "walletcreatefundedpsbt", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// WalletDisplayAddressReq holds the arguments for the WalletDisplayAddress call.
//  1. address    (string, required)
type WalletDisplayAddressReq struct {
	Address string `json:"address"`
}

// WalletDisplayAddressResp holds the response to the WalletDisplayAddress call.
//  {                       (json object)
//    "address" : "str"     (string) The address as confirmed by the signer
//  }
type WalletDisplayAddressResp struct {
	// The address as confirmed by the signer
	Address string `json:"address"`
}

// WalletDisplayAddress RPC method.
// Display address on an external signer for verification.
func (bc *BitcoindClient) WalletDisplayAddress(ctx context.Context, args WalletDisplayAddressReq) (result WalletDisplayAddressResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "walletdisplayaddress", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// WalletLock RPC method.
// Removes the wallet encryption key from memory, locking the wallet.
// After calling this method, you will need to call walletpassphrase again
// before being able to call any methods which require the wallet to be unlocked.
func (bc *BitcoindClient) WalletLock(ctx context.Context) (err error) {
	_, err = bc.sendRequest(ctx, "walletlock", nil)
	return
}

// WalletPassphraseReq holds the arguments for the WalletPassphrase call.
//  1. passphrase    (string, required) The wallet passphrase
//  2. timeout       (numeric, required) The time to keep the decryption key in seconds; capped at 100000000 (~3 years).
type WalletPassphraseReq struct {
	// The wallet passphrase
	Passphrase string `json:"passphrase"`

	// The time to keep the decryption key in seconds; capped at 100000000 (~3 years).
	TimeOut float64 `json:"timeout"`
}

// WalletPassphrase RPC method.
// Stores the wallet decryption key in memory for 'timeout' seconds.
// This is needed prior to performing transactions related to private keys such as sending bitcoins
// Note:
// Issuing the walletpassphrase command while the wallet is already unlocked will set a new unlock
// time that overrides the old one.
func (bc *BitcoindClient) WalletPassphrase(ctx context.Context, args WalletPassphraseReq) (err error) {
	_, err = bc.sendRequest(ctx, "walletpassphrase", args)
	return
}

// WalletPassphraseChangeReq holds the arguments for the WalletPassphraseChange call.
//  1. oldpassphrase    (string, required) The current passphrase
//  2. newpassphrase    (string, required) The new passphrase
type WalletPassphraseChangeReq struct {
	// The current passphrase
	OldPassphrase string `json:"oldpassphrase"`

	// The new passphrase
	NewPassphrase string `json:"newpassphrase"`
}

// WalletPassphraseChange RPC method.
// Changes the wallet passphrase from 'oldpassphrase' to 'newpassphrase'.
func (bc *BitcoindClient) WalletPassphraseChange(ctx context.Context, args WalletPassphraseChangeReq) (err error) {
	_, err = bc.sendRequest(ctx, "walletpassphrasechange", args)
	return
}

// WalletProcessPsbtReq holds the arguments for the WalletProcessPsbt call.
//  1. psbt           (string, required) The transaction base64 string
//  2. sign           (boolean, optional, default=true) Also sign the transaction when updating
//  3. sighashtype    (string, optional, default="DEFAULT") The signature hash type to sign with if not specified by the PSBT. Must be one of
//                    "DEFAULT"
//                    "ALL"
//                    "NONE"
//                    "SINGLE"
//                    "ALL|ANYONECANPAY"
//                    "NONE|ANYONECANPAY"
//                    "SINGLE|ANYONECANPAY"
//  4. bip32derivs    (boolean, optional, default=true) Include BIP 32 derivation paths for public keys if we know them
type WalletProcessPsbtReq struct {
	// The transaction base64 string
	Psbt string `json:"psbt"`

	// Also sign the transaction when updating
	// Default: true
	Sign *bool `json:"sign,omitempty"`

	// The signature hash type to sign with if not specified by the PSBT. Must be one of
	// "DEFAULT"
	// "ALL"
	// "NONE"
	// "SINGLE"
	// "ALL|ANYONECANPAY"
	// "NONE|ANYONECANPAY"
	// "SINGLE|ANYONECANPAY"
	// Default: "DEFAULT"
	SigHashType string `json:"sighashtype,omitempty"`

	// Include BIP 32 derivation paths for public keys if we know them
	// Default: true
	BIP32Derivs *bool `json:"bip32derivs,omitempty"`
}

// WalletProcessPsbtResp holds the response to the WalletProcessPsbt call.
//  {                             (json object)
//    "psbt" : "str",             (string) The base64-encoded partially signed transaction
//    "complete" : true|false     (boolean) If the transaction has a complete set of signatures
//  }
type WalletProcessPsbtResp struct {
	// The base64-encoded partially signed transaction
	Psbt string `json:"psbt"`

	// If the transaction has a complete set of signatures
	Complete bool `json:"complete"`
}

// WalletProcessPsbt RPC method.
// Update a PSBT with input information from our wallet and then sign inputs
// that we can sign for.
// Requires wallet passphrase to be set with walletpassphrase call if wallet is encrypted.
func (bc *BitcoindClient) WalletProcessPsbt(ctx context.Context, args WalletProcessPsbtReq) (result WalletProcessPsbtResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "walletprocesspsbt", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}
