// THIS FILE IS GENERATED CODE, MANUAL EDITS WILL BE OVERWRITTEN

package bitcoindclient

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"
)

func TestRpc(t *testing.T) {
	rpcAddress, rpcUser, rpcPassword, _, err := startBitcoind()
	if err != nil {
		t.Fatal(err)
	}
	defer stopBitcoind(t)

	for _, name := range []string{
		"CreateWallet",
		"GetBestBlockhash",
		"GetBlock",
		"GetBlockchainInfo",
		"GetBlockCount",
		"GetBlockFilter",
		"GetBlockhash",
		"GetBlockHeader",
		"GetBlockStats",
		"GetChainTips",
		"GetChainTxStats",
		"GetDifficulty",
		"GetMempoolAncestors",
		"GetMempoolDescendants",
		"GetMempoolEntry",
		"GetMempoolInfo",
		"GetRawMempool",
		"GetTxOut",
		"GetTxOutProof",
		"GetTxOutSetInfo",
		"PreciousBlock",
		"PruneBlockchain",
		"SaveMempool",
		"ScanTxOutSet",
		"VerifyChain",
		"VerifyTxOutProof",
		"GetMemoryInfo",
		"GetRpcInfo",
		"Help",
		"Logging",
		"Uptime",
		"GenerateBlock",
		"GenerateToAddress",
		"GenerateToDescriptor",
		"GetBlockTemplate",
		"GetMiningInfo",
		"GetNetworkHashPs",
		"PrioritiseTransaction",
		"SubmitBlock",
		"SubmitHeader",
		"AddNode",
		"ClearBanned",
		"DisconnectNode",
		"GetAddedNodeInfo",
		"GetConnectionCount",
		"GetNetTotals",
		"GetNetworkInfo",
		"GetNodeAddresses",
		"GetPeerInfo",
		"ListBanned",
		"Ping",
		"SetBan",
		"SetNetworkActive",
		"AnalyzePsbt",
		"CombinePsbt",
		"CombineRawTransaction",
		"ConvertToPsbt",
		"CreatePsbt",
		"CreateRawTransaction",
		"DecodePsbt",
		"DecodeRawTransaction",
		"DecodeScript",
		"FinalizePsbt",
		"FundRawTransaction",
		"GetRawTransaction",
		"JoinPsbts",
		"SendRawTransaction",
		"SignRawTransactionWithKey",
		"TestMempoolAccept",
		"UtxoUpdatePsbt",
		"EnumerateSigners",
		"CreateMultisig",
		"DeriveAddresses",
		"EstimateSmartFee",
		"GetDescriptorInfo",
		"GetIndexInfo",
		"SignMessageWithPrivkey",
		"ValidateAddress",
		"VerifyMessage",
		"AbandonTransaction",
		"AbortRescan",
		"AddMultisigAddress",
		"BackupWallet",
		"BumpFee",
		"DumpPrivkey",
		"DumpWallet",
		"EncryptWallet",
		"GetAddressesByLabel",
		"GetAddressInfo",
		"GetBalance",
		"GetBalances",
		"GetNewAddress",
		"GetRawChangeaddress",
		"GetReceivedByAddress",
		"GetReceivedByLabel",
		"GetTransaction",
		"GetUnconfirmedBalance",
		"GetWalletInfo",
		"ImportAddress",
		"ImportDescriptors",
		"ImportMulti",
		"ImportPrivkey",
		"ImportPrunedFunds",
		"ImportPubkey",
		"ImportWallet",
		"KeypoolRefill",
		"ListAddressGroupings",
		"ListDescriptors",
		"ListLabels",
		"ListLockUnspent",
		"ListReceivedByAddress",
		"ListReceivedByLabel",
		"ListSinceBlock",
		"ListTransactions",
		"ListUnspent",
		"ListWalletDir",
		"ListWallets",
		"LoadWallet",
		"LockUnspent",
		"PsbtBumpFee",
		"RemovePrunedFunds",
		"RescanBlockchain",
		"Send",
		"SendMany",
		"SendToAddress",
		"SetHDSeed",
		"SetLabel",
		"SetTxFee",
		"SetWalletFlag",
		"SignMessage",
		"SignRawTransactionWithWallet",
		"UnloadWallet",
		"UpgradeWallet",
		"WalletCreateFundedPsbt",
		"WalletDisplayAddress",
		"WalletLock",
		"WalletPassphrase",
		"WalletPassphraseChange",
		"WalletProcessPsbt",
		"GetZmqNotifications",
		"Stop",
	} {
		t.Run(name, func(t *testing.T) {
			bc, err := New(Config{
				RpcAddress:  rpcAddress,
				RpcUser:     rpcUser,
				RpcPassword: rpcPassword,
			})
			if err != nil {
				t.Fatal(err)
			}
			defer bc.Close()
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			ctx = UseConnectionRetries(ctx, 2)

			switch name {
			case "CreateWallet":
				_, err = bc.CreateWallet(ctx, CreateWalletReq{})
			case "GetBestBlockhash":
				_, err = bc.GetBestBlockhash(ctx)
			case "GetBlock":
				_, err = bc.GetBlock(ctx, GetBlockReq{})
			case "GetBlockchainInfo":
				_, err = bc.GetBlockchainInfo(ctx)
			case "GetBlockCount":
				_, err = bc.GetBlockCount(ctx)
			case "GetBlockFilter":
				_, err = bc.GetBlockFilter(ctx, GetBlockFilterReq{})
			case "GetBlockhash":
				_, err = bc.GetBlockhash(ctx, GetBlockhashReq{})
			case "GetBlockHeader":
				_, err = bc.GetBlockHeader(ctx, GetBlockHeaderReq{})
			case "GetBlockStats":
				_, err = bc.GetBlockStats(ctx, GetBlockStatsReq{})
			case "GetChainTips":
				_, err = bc.GetChainTips(ctx)
			case "GetChainTxStats":
				_, err = bc.GetChainTxStats(ctx, GetChainTxStatsReq{})
			case "GetDifficulty":
				_, err = bc.GetDifficulty(ctx)
			case "GetMempoolAncestors":
				_, err = bc.GetMempoolAncestors(ctx, GetMempoolAncestorsReq{})
			case "GetMempoolDescendants":
				_, err = bc.GetMempoolDescendants(ctx, GetMempoolDescendantsReq{})
			case "GetMempoolEntry":
				_, err = bc.GetMempoolEntry(ctx, GetMempoolEntryReq{})
			case "GetMempoolInfo":
				_, err = bc.GetMempoolInfo(ctx)
			case "GetRawMempool":
				_, err = bc.GetRawMempool(ctx, GetRawMempoolReq{})
			case "GetTxOut":
				_, err = bc.GetTxOut(ctx, GetTxOutReq{})
			case "GetTxOutProof":
				_, err = bc.GetTxOutProof(ctx, GetTxOutProofReq{})
			case "GetTxOutSetInfo":
				_, err = bc.GetTxOutSetInfo(ctx, GetTxOutSetInfoReq{})
			case "PreciousBlock":
				err = bc.PreciousBlock(ctx, PreciousBlockReq{})
			case "PruneBlockchain":
				_, err = bc.PruneBlockchain(ctx, PruneBlockchainReq{})
			case "SaveMempool":
				err = bc.SaveMempool(ctx)
			case "ScanTxOutSet":
				_, err = bc.ScanTxOutSet(ctx, ScanTxOutSetReq{})
			case "VerifyChain":
				_, err = bc.VerifyChain(ctx, VerifyChainReq{})
			case "VerifyTxOutProof":
				_, err = bc.VerifyTxOutProof(ctx, VerifyTxOutProofReq{})
			case "GetMemoryInfo":
				_, err = bc.GetMemoryInfo(ctx, GetMemoryInfoReq{})
			case "GetRpcInfo":
				_, err = bc.GetRpcInfo(ctx)
			case "Help":
				_, err = bc.Help(ctx, HelpReq{})
			case "Logging":
				err = bc.Logging(ctx, LoggingReq{})
			case "Uptime":
				_, err = bc.Uptime(ctx)
			case "GenerateBlock":
				_, err = bc.GenerateBlock(ctx, GenerateBlockReq{})
			case "GenerateToAddress":
				_, err = bc.GenerateToAddress(ctx, GenerateToAddressReq{})
			case "GenerateToDescriptor":
				_, err = bc.GenerateToDescriptor(ctx, GenerateToDescriptorReq{})
			case "GetBlockTemplate":
				_, err = bc.GetBlockTemplate(ctx, GetBlockTemplateReq{})
			case "GetMiningInfo":
				_, err = bc.GetMiningInfo(ctx)
			case "GetNetworkHashPs":
				_, err = bc.GetNetworkHashPs(ctx, GetNetworkHashPsReq{})
			case "PrioritiseTransaction":
				_, err = bc.PrioritiseTransaction(ctx, PrioritiseTransactionReq{})
			case "SubmitBlock":
				_, err = bc.SubmitBlock(ctx, SubmitBlockReq{})
			case "SubmitHeader":
				err = bc.SubmitHeader(ctx, SubmitHeaderReq{})
			case "AddNode":
				err = bc.AddNode(ctx, AddNodeReq{Command: "add"})
			case "ClearBanned":
				err = bc.ClearBanned(ctx)
			case "DisconnectNode":
				err = bc.DisconnectNode(ctx, DisconnectNodeReq{Address: "abc"})
			case "GetAddedNodeInfo":
				_, err = bc.GetAddedNodeInfo(ctx, GetAddedNodeInfoReq{})
			case "GetConnectionCount":
				_, err = bc.GetConnectionCount(ctx)
			case "GetNetTotals":
				_, err = bc.GetNetTotals(ctx)
			case "GetNetworkInfo":
				_, err = bc.GetNetworkInfo(ctx)
			case "GetNodeAddresses":
				_, err = bc.GetNodeAddresses(ctx, GetNodeAddressesReq{})
			case "GetPeerInfo":
				_, err = bc.GetPeerInfo(ctx)
			case "ListBanned":
				_, err = bc.ListBanned(ctx)
			case "Ping":
				err = bc.Ping(ctx)
			case "SetBan":
				err = bc.SetBan(ctx, SetBanReq{Command: "add"})
			case "SetNetworkActive":
				_, err = bc.SetNetworkActive(ctx, SetNetworkActiveReq{})
			case "AnalyzePsbt":
				_, err = bc.AnalyzePsbt(ctx, AnalyzePsbtReq{})
			case "CombinePsbt":
				_, err = bc.CombinePsbt(ctx, CombinePsbtReq{})
			case "CombineRawTransaction":
				_, err = bc.CombineRawTransaction(ctx, CombineRawTransactionReq{})
			case "ConvertToPsbt":
				_, err = bc.ConvertToPsbt(ctx, ConvertToPsbtReq{})
			case "CreatePsbt":
				_, err = bc.CreatePsbt(ctx, CreatePsbtReq{})
			case "CreateRawTransaction":
				_, err = bc.CreateRawTransaction(ctx, CreateRawTransactionReq{})
			case "DecodePsbt":
				_, err = bc.DecodePsbt(ctx, DecodePsbtReq{})
			case "DecodeRawTransaction":
				_, err = bc.DecodeRawTransaction(ctx, DecodeRawTransactionReq{})
			case "DecodeScript":
				_, err = bc.DecodeScript(ctx, DecodeScriptReq{})
			case "FinalizePsbt":
				_, err = bc.FinalizePsbt(ctx, FinalizePsbtReq{})
			case "FundRawTransaction":
				_, err = bc.FundRawTransaction(ctx, FundRawTransactionReq{})
			case "GetRawTransaction":
				_, err = bc.GetRawTransaction(ctx, GetRawTransactionReq{})
			case "JoinPsbts":
				_, err = bc.JoinPsbts(ctx, JoinPsbtsReq{})
			case "SendRawTransaction":
				_, err = bc.SendRawTransaction(ctx, SendRawTransactionReq{})
			case "SignRawTransactionWithKey":
				_, err = bc.SignRawTransactionWithKey(ctx, SignRawTransactionWithKeyReq{})
			case "TestMempoolAccept":
				_, err = bc.TestMempoolAccept(ctx, TestMempoolAcceptReq{})
			case "UtxoUpdatePsbt":
				_, err = bc.UtxoUpdatePsbt(ctx, UtxoUpdatePsbtReq{})
			case "EnumerateSigners":
				_, err = bc.EnumerateSigners(ctx)
			case "CreateMultisig":
				_, err = bc.CreateMultisig(ctx, CreateMultisigReq{})
			case "DeriveAddresses":
				_, err = bc.DeriveAddresses(ctx, DeriveAddressesReq{})
			case "EstimateSmartFee":
				_, err = bc.EstimateSmartFee(ctx, EstimateSmartFeeReq{})
			case "GetDescriptorInfo":
				_, err = bc.GetDescriptorInfo(ctx, GetDescriptorInfoReq{})
			case "GetIndexInfo":
				_, err = bc.GetIndexInfo(ctx, GetIndexInfoReq{})
			case "SignMessageWithPrivkey":
				_, err = bc.SignMessageWithPrivkey(ctx, SignMessageWithPrivkeyReq{})
			case "ValidateAddress":
				_, err = bc.ValidateAddress(ctx, ValidateAddressReq{})
			case "VerifyMessage":
				_, err = bc.VerifyMessage(ctx, VerifyMessageReq{})
			case "AbandonTransaction":
				err = bc.AbandonTransaction(ctx, AbandonTransactionReq{})
			case "AbortRescan":
				_, err = bc.AbortRescan(ctx)
			case "AddMultisigAddress":
				_, err = bc.AddMultisigAddress(ctx, AddMultisigAddressReq{})
			case "BackupWallet":
				err = bc.BackupWallet(ctx, BackupWalletReq{})
			case "BumpFee":
				_, err = bc.BumpFee(ctx, BumpFeeReq{})
			case "DumpPrivkey":
				_, err = bc.DumpPrivkey(ctx, DumpPrivkeyReq{})
			case "DumpWallet":
				_, err = bc.DumpWallet(ctx, DumpWalletReq{})
			case "EncryptWallet":
				_, err = bc.EncryptWallet(ctx, EncryptWalletReq{})
			case "GetAddressesByLabel":
				err = bc.GetAddressesByLabel(ctx, GetAddressesByLabelReq{})
			case "GetAddressInfo":
				_, err = bc.GetAddressInfo(ctx, GetAddressInfoReq{})
			case "GetBalance":
				_, err = bc.GetBalance(ctx, GetBalanceReq{})
			case "GetBalances":
				_, err = bc.GetBalances(ctx)
			case "GetNewAddress":
				_, err = bc.GetNewAddress(ctx, GetNewAddressReq{})
			case "GetRawChangeaddress":
				_, err = bc.GetRawChangeaddress(ctx, GetRawChangeaddressReq{})
			case "GetReceivedByAddress":
				_, err = bc.GetReceivedByAddress(ctx, GetReceivedByAddressReq{})
			case "GetReceivedByLabel":
				_, err = bc.GetReceivedByLabel(ctx, GetReceivedByLabelReq{})
			case "GetTransaction":
				_, err = bc.GetTransaction(ctx, GetTransactionReq{})
			case "GetUnconfirmedBalance":
				_, err = bc.GetUnconfirmedBalance(ctx)
			case "GetWalletInfo":
				_, err = bc.GetWalletInfo(ctx)
			case "ImportAddress":
				err = bc.ImportAddress(ctx, ImportAddressReq{})
			case "ImportDescriptors":
				_, err = bc.ImportDescriptors(ctx, ImportDescriptorsReq{})
			case "ImportMulti":
				_, err = bc.ImportMulti(ctx, ImportMultiReq{})
			case "ImportPrivkey":
				err = bc.ImportPrivkey(ctx, ImportPrivkeyReq{})
			case "ImportPrunedFunds":
				err = bc.ImportPrunedFunds(ctx, ImportPrunedFundsReq{})
			case "ImportPubkey":
				err = bc.ImportPubkey(ctx, ImportPubkeyReq{})
			case "ImportWallet":
				err = bc.ImportWallet(ctx, ImportWalletReq{})
			case "KeypoolRefill":
				err = bc.KeypoolRefill(ctx, KeypoolRefillReq{})
			case "ListAddressGroupings":
				_, err = bc.ListAddressGroupings(ctx)
			case "ListDescriptors":
				_, err = bc.ListDescriptors(ctx)
			case "ListLabels":
				_, err = bc.ListLabels(ctx, ListLabelsReq{})
			case "ListLockUnspent":
				_, err = bc.ListLockUnspent(ctx)
			case "ListReceivedByAddress":
				_, err = bc.ListReceivedByAddress(ctx, ListReceivedByAddressReq{})
			case "ListReceivedByLabel":
				_, err = bc.ListReceivedByLabel(ctx, ListReceivedByLabelReq{})
			case "ListSinceBlock":
				_, err = bc.ListSinceBlock(ctx, ListSinceBlockReq{})
			case "ListTransactions":
				_, err = bc.ListTransactions(ctx, ListTransactionsReq{})
			case "ListUnspent":
				_, err = bc.ListUnspent(ctx, ListUnspentReq{})
			case "ListWalletDir":
				_, err = bc.ListWalletDir(ctx)
			case "ListWallets":
				_, err = bc.ListWallets(ctx)
			case "LoadWallet":
				_, err = bc.LoadWallet(ctx, LoadWalletReq{})
			case "LockUnspent":
				_, err = bc.LockUnspent(ctx, LockUnspentReq{})
			case "PsbtBumpFee":
				_, err = bc.PsbtBumpFee(ctx, PsbtBumpFeeReq{})
			case "RemovePrunedFunds":
				err = bc.RemovePrunedFunds(ctx, RemovePrunedFundsReq{})
			case "RescanBlockchain":
				_, err = bc.RescanBlockchain(ctx, RescanBlockchainReq{})
			case "Send":
				_, err = bc.Send(ctx, SendReq{})
			case "SendMany":
				_, err = bc.SendMany(ctx, SendManyReq{})
			case "SendToAddress":
				_, err = bc.SendToAddress(ctx, SendToAddressReq{})
			case "SetHDSeed":
				err = bc.SetHDSeed(ctx, SetHDSeedReq{})
			case "SetLabel":
				err = bc.SetLabel(ctx, SetLabelReq{})
			case "SetTxFee":
				_, err = bc.SetTxFee(ctx, SetTxFeeReq{})
			case "SetWalletFlag":
				_, err = bc.SetWalletFlag(ctx, SetWalletFlagReq{})
			case "SignMessage":
				_, err = bc.SignMessage(ctx, SignMessageReq{})
			case "SignRawTransactionWithWallet":
				_, err = bc.SignRawTransactionWithWallet(ctx, SignRawTransactionWithWalletReq{})
			case "UnloadWallet":
				_, err = bc.UnloadWallet(ctx, UnloadWalletReq{WalletName: "abc"})
			case "UpgradeWallet":
				_, err = bc.UpgradeWallet(ctx, UpgradeWalletReq{})
			case "WalletCreateFundedPsbt":
				_, err = bc.WalletCreateFundedPsbt(ctx, WalletCreateFundedPsbtReq{})
			case "WalletDisplayAddress":
				_, err = bc.WalletDisplayAddress(ctx, WalletDisplayAddressReq{})
			case "WalletLock":
				err = bc.WalletLock(ctx)
			case "WalletPassphrase":
				err = bc.WalletPassphrase(ctx, WalletPassphraseReq{})
			case "WalletPassphraseChange":
				err = bc.WalletPassphraseChange(ctx, WalletPassphraseChangeReq{})
			case "WalletProcessPsbt":
				_, err = bc.WalletProcessPsbt(ctx, WalletProcessPsbtReq{})
			case "GetZmqNotifications":
				_, err = bc.GetZmqNotifications(ctx)
			case "Stop":
				_, err = bc.Stop(ctx)
			}

			if err != nil {
				var bErr *BitcoindError
				if errors.As(err, &bErr) {
					switch bErr.Code {
					case -32600:
						// RPC_INVALID_REQUEST
						t.Error(err)
					case -32601:
						// RPC_METHOD_NOT_FOUND
						t.Error(err)
					case -32602:
						// RPC_INVALID_PARAMS
						t.Error(err)
					case -32:
						// RPC_METHOD_DEPRECATED    RPC method is deprecated
						t.Error(err)
					case -1:
						// RPC_MISC_ERROR    std::exception thrown in command handling
						// Suppress error for known cases without feasible workarounds.
						switch name {
						case "PruneBlockchain":
							// Cannot prune blocks because node is not in prune mode.
							t.Log(err)
						case "EnumerateSigners":
							// Error: restart bitcoind with -signer=<cmd>
							t.Log(err)
						default:
							t.Error(err)
						}
					case -3:
						// RPC_TYPE_ERROR    Unexpected type was passed as parameter
						switch name {
						default:
							t.Error(err)
						}
					case -8:
						// RPC_INVALID_PARAMETER    Invalid, missing or duplicate parameter
						if strings.HasPrefix(bErr.Message, "Unknown named parameter") {
							t.Error(err)
						} else {
							// Likely malformed argument, application specific.
							t.Log(err)
						}
					default:
						// Sending empty arguments are sure to be rejected for various reasons,
						// but we don't fail on anything but JSON structure mismatch.
						t.Log(err)
					}
				} else {
					t.Errorf("Internal error: %v", err)
				}
			}
		})
	}
}
