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
		"AbandonTransaction",
		"AbortRescan",
		"AddMultisigAddress",
		"AddNode",
		"AnalyzePsbt",
		"BackupWallet",
		"BumpFee",
		"ClearBanned",
		"CombinePsbt",
		"CombineRawTransaction",
		"ConvertToPsbt",
		"CreateMultisig",
		"CreatePsbt",
		"CreateRawTransaction",
		"DecodePsbt",
		"DecodeRawTransaction",
		"DecodeScript",
		"DeriveAddresses",
		"DisconnectNode",
		"DumpPrivkey",
		"DumpWallet",
		"EncryptWallet",
		"EnumerateSigners",
		"EstimateSmartFee",
		"FinalizePsbt",
		"FundRawTransaction",
		"GenerateBlock",
		"GenerateToAddress",
		"GenerateToDescriptor",
		"GetAddedNodeInfo",
		"GetAddressInfo",
		"GetAddressesByLabel",
		"GetBalance",
		"GetBalances",
		"GetBestBlockhash",
		"GetBlock",
		"GetBlockCount",
		"GetBlockFilter",
		"GetBlockFromPeer",
		"GetBlockHeader",
		"GetBlockStats",
		"GetBlockTemplate",
		"GetBlockchainInfo",
		"GetBlockhash",
		"GetChainTips",
		"GetChainTxStats",
		"GetConnectionCount",
		"GetDeploymentInfo",
		"GetDescriptorInfo",
		"GetDifficulty",
		"GetIndexInfo",
		"GetMemoryInfo",
		"GetMempoolAncestors",
		"GetMempoolDescendants",
		"GetMempoolEntry",
		"GetMempoolInfo",
		"GetMiningInfo",
		"GetNetTotals",
		"GetNetworkHashPs",
		"GetNetworkInfo",
		"GetNewAddress",
		"GetNodeAddresses",
		"GetPeerInfo",
		"GetRawChangeaddress",
		"GetRawMempool",
		"GetRawTransaction",
		"GetReceivedByAddress",
		"GetReceivedByLabel",
		"GetRpcInfo",
		"GetTransaction",
		"GetTxOut",
		"GetTxOutProof",
		"GetTxOutSetInfo",
		"GetUnconfirmedBalance",
		"GetWalletInfo",
		"GetZmqNotifications",
		"Help",
		"ImportAddress",
		"ImportDescriptors",
		"ImportMulti",
		"ImportPrivkey",
		"ImportPrunedFunds",
		"ImportPubkey",
		"ImportWallet",
		"JoinPsbts",
		"KeypoolRefill",
		"ListAddressGroupings",
		"ListBanned",
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
		"Logging",
		"NewKeypool",
		"Ping",
		"PreciousBlock",
		"PrioritiseTransaction",
		"PruneBlockchain",
		"PsbtBumpFee",
		"RemovePrunedFunds",
		"RescanBlockchain",
		"RestoreWallet",
		"SaveMempool",
		"ScanTxOutSet",
		"Send",
		"SendMany",
		"SendRawTransaction",
		"SendToAddress",
		"SetBan",
		"SetHDSeed",
		"SetLabel",
		"SetNetworkActive",
		"SetTxFee",
		"SetWalletFlag",
		"SignMessage",
		"SignMessageWithPrivkey",
		"SignRawTransactionWithKey",
		"SignRawTransactionWithWallet",
		"SubmitBlock",
		"SubmitHeader",
		"TestMempoolAccept",
		"UnloadWallet",
		"UpgradeWallet",
		"Uptime",
		"UtxoUpdatePsbt",
		"ValidateAddress",
		"VerifyChain",
		"VerifyMessage",
		"VerifyTxOutProof",
		"WalletCreateFundedPsbt",
		"WalletDisplayAddress",
		"WalletLock",
		"WalletPassphrase",
		"WalletPassphraseChange",
		"WalletProcessPsbt",
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
			case "AbandonTransaction":
				err = bc.AbandonTransaction(ctx, AbandonTransactionReq{})
			case "AbortRescan":
				_, err = bc.AbortRescan(ctx)
			case "AddMultisigAddress":
				_, err = bc.AddMultisigAddress(ctx, AddMultisigAddressReq{})
			case "AddNode":
				err = bc.AddNode(ctx, AddNodeReq{Command: "add"})
			case "AnalyzePsbt":
				_, err = bc.AnalyzePsbt(ctx, AnalyzePsbtReq{})
			case "BackupWallet":
				err = bc.BackupWallet(ctx, BackupWalletReq{})
			case "BumpFee":
				_, err = bc.BumpFee(ctx, BumpFeeReq{})
			case "ClearBanned":
				err = bc.ClearBanned(ctx)
			case "CombinePsbt":
				_, err = bc.CombinePsbt(ctx, CombinePsbtReq{})
			case "CombineRawTransaction":
				_, err = bc.CombineRawTransaction(ctx, CombineRawTransactionReq{})
			case "ConvertToPsbt":
				_, err = bc.ConvertToPsbt(ctx, ConvertToPsbtReq{})
			case "CreateMultisig":
				_, err = bc.CreateMultisig(ctx, CreateMultisigReq{})
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
			case "DeriveAddresses":
				_, err = bc.DeriveAddresses(ctx, DeriveAddressesReq{})
			case "DisconnectNode":
				err = bc.DisconnectNode(ctx, DisconnectNodeReq{Address: "abc"})
			case "DumpPrivkey":
				_, err = bc.DumpPrivkey(ctx, DumpPrivkeyReq{})
			case "DumpWallet":
				_, err = bc.DumpWallet(ctx, DumpWalletReq{})
			case "EncryptWallet":
				_, err = bc.EncryptWallet(ctx, EncryptWalletReq{})
			case "EnumerateSigners":
				_, err = bc.EnumerateSigners(ctx)
			case "EstimateSmartFee":
				_, err = bc.EstimateSmartFee(ctx, EstimateSmartFeeReq{})
			case "FinalizePsbt":
				_, err = bc.FinalizePsbt(ctx, FinalizePsbtReq{})
			case "FundRawTransaction":
				_, err = bc.FundRawTransaction(ctx, FundRawTransactionReq{})
			case "GenerateBlock":
				_, err = bc.GenerateBlock(ctx, GenerateBlockReq{})
			case "GenerateToAddress":
				_, err = bc.GenerateToAddress(ctx, GenerateToAddressReq{})
			case "GenerateToDescriptor":
				_, err = bc.GenerateToDescriptor(ctx, GenerateToDescriptorReq{})
			case "GetAddedNodeInfo":
				_, err = bc.GetAddedNodeInfo(ctx, GetAddedNodeInfoReq{})
			case "GetAddressInfo":
				_, err = bc.GetAddressInfo(ctx, GetAddressInfoReq{})
			case "GetAddressesByLabel":
				_, err = bc.GetAddressesByLabel(ctx, GetAddressesByLabelReq{})
			case "GetBalance":
				_, err = bc.GetBalance(ctx, GetBalanceReq{})
			case "GetBalances":
				_, err = bc.GetBalances(ctx)
			case "GetBestBlockhash":
				_, err = bc.GetBestBlockhash(ctx)
			case "GetBlock":
				_, err = bc.GetBlock(ctx, GetBlockReq{})
			case "GetBlockCount":
				_, err = bc.GetBlockCount(ctx)
			case "GetBlockFilter":
				_, err = bc.GetBlockFilter(ctx, GetBlockFilterReq{})
			case "GetBlockFromPeer":
				err = bc.GetBlockFromPeer(ctx, GetBlockFromPeerReq{})
			case "GetBlockHeader":
				_, err = bc.GetBlockHeader(ctx, GetBlockHeaderReq{})
			case "GetBlockStats":
				_, err = bc.GetBlockStats(ctx, GetBlockStatsReq{})
			case "GetBlockTemplate":
				_, err = bc.GetBlockTemplate(ctx, GetBlockTemplateReq{})
			case "GetBlockchainInfo":
				_, err = bc.GetBlockchainInfo(ctx)
			case "GetBlockhash":
				_, err = bc.GetBlockhash(ctx, GetBlockhashReq{})
			case "GetChainTips":
				_, err = bc.GetChainTips(ctx)
			case "GetChainTxStats":
				_, err = bc.GetChainTxStats(ctx, GetChainTxStatsReq{})
			case "GetConnectionCount":
				_, err = bc.GetConnectionCount(ctx)
			case "GetDeploymentInfo":
				_, err = bc.GetDeploymentInfo(ctx, GetDeploymentInfoReq{})
			case "GetDescriptorInfo":
				_, err = bc.GetDescriptorInfo(ctx, GetDescriptorInfoReq{})
			case "GetDifficulty":
				_, err = bc.GetDifficulty(ctx)
			case "GetIndexInfo":
				_, err = bc.GetIndexInfo(ctx, GetIndexInfoReq{})
			case "GetMemoryInfo":
				_, err = bc.GetMemoryInfo(ctx, GetMemoryInfoReq{})
			case "GetMempoolAncestors":
				_, err = bc.GetMempoolAncestors(ctx, GetMempoolAncestorsReq{})
			case "GetMempoolDescendants":
				_, err = bc.GetMempoolDescendants(ctx, GetMempoolDescendantsReq{})
			case "GetMempoolEntry":
				_, err = bc.GetMempoolEntry(ctx, GetMempoolEntryReq{})
			case "GetMempoolInfo":
				_, err = bc.GetMempoolInfo(ctx)
			case "GetMiningInfo":
				_, err = bc.GetMiningInfo(ctx)
			case "GetNetTotals":
				_, err = bc.GetNetTotals(ctx)
			case "GetNetworkHashPs":
				_, err = bc.GetNetworkHashPs(ctx, GetNetworkHashPsReq{})
			case "GetNetworkInfo":
				_, err = bc.GetNetworkInfo(ctx)
			case "GetNewAddress":
				_, err = bc.GetNewAddress(ctx, GetNewAddressReq{})
			case "GetNodeAddresses":
				_, err = bc.GetNodeAddresses(ctx, GetNodeAddressesReq{})
			case "GetPeerInfo":
				_, err = bc.GetPeerInfo(ctx)
			case "GetRawChangeaddress":
				_, err = bc.GetRawChangeaddress(ctx, GetRawChangeaddressReq{})
			case "GetRawMempool":
				_, err = bc.GetRawMempool(ctx, GetRawMempoolReq{})
			case "GetRawTransaction":
				_, err = bc.GetRawTransaction(ctx, GetRawTransactionReq{})
			case "GetReceivedByAddress":
				_, err = bc.GetReceivedByAddress(ctx, GetReceivedByAddressReq{})
			case "GetReceivedByLabel":
				_, err = bc.GetReceivedByLabel(ctx, GetReceivedByLabelReq{})
			case "GetRpcInfo":
				_, err = bc.GetRpcInfo(ctx)
			case "GetTransaction":
				_, err = bc.GetTransaction(ctx, GetTransactionReq{})
			case "GetTxOut":
				_, err = bc.GetTxOut(ctx, GetTxOutReq{})
			case "GetTxOutProof":
				_, err = bc.GetTxOutProof(ctx, GetTxOutProofReq{})
			case "GetTxOutSetInfo":
				_, err = bc.GetTxOutSetInfo(ctx, GetTxOutSetInfoReq{})
			case "GetUnconfirmedBalance":
				_, err = bc.GetUnconfirmedBalance(ctx)
			case "GetWalletInfo":
				_, err = bc.GetWalletInfo(ctx)
			case "GetZmqNotifications":
				_, err = bc.GetZmqNotifications(ctx)
			case "Help":
				_, err = bc.Help(ctx, HelpReq{})
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
			case "JoinPsbts":
				_, err = bc.JoinPsbts(ctx, JoinPsbtsReq{})
			case "KeypoolRefill":
				err = bc.KeypoolRefill(ctx, KeypoolRefillReq{})
			case "ListAddressGroupings":
				_, err = bc.ListAddressGroupings(ctx)
			case "ListBanned":
				_, err = bc.ListBanned(ctx)
			case "ListDescriptors":
				_, err = bc.ListDescriptors(ctx, ListDescriptorsReq{})
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
			case "Logging":
				_, err = bc.Logging(ctx, LoggingReq{})
			case "NewKeypool":
				err = bc.NewKeypool(ctx)
			case "Ping":
				err = bc.Ping(ctx)
			case "PreciousBlock":
				err = bc.PreciousBlock(ctx, PreciousBlockReq{})
			case "PrioritiseTransaction":
				_, err = bc.PrioritiseTransaction(ctx, PrioritiseTransactionReq{})
			case "PruneBlockchain":
				_, err = bc.PruneBlockchain(ctx, PruneBlockchainReq{})
			case "PsbtBumpFee":
				_, err = bc.PsbtBumpFee(ctx, PsbtBumpFeeReq{})
			case "RemovePrunedFunds":
				err = bc.RemovePrunedFunds(ctx, RemovePrunedFundsReq{})
			case "RescanBlockchain":
				_, err = bc.RescanBlockchain(ctx, RescanBlockchainReq{})
			case "RestoreWallet":
				_, err = bc.RestoreWallet(ctx, RestoreWalletReq{})
			case "SaveMempool":
				_, err = bc.SaveMempool(ctx)
			case "ScanTxOutSet":
				_, err = bc.ScanTxOutSet(ctx, ScanTxOutSetReq{})
			case "Send":
				_, err = bc.Send(ctx, SendReq{})
			case "SendMany":
				_, err = bc.SendMany(ctx, SendManyReq{})
			case "SendRawTransaction":
				_, err = bc.SendRawTransaction(ctx, SendRawTransactionReq{})
			case "SendToAddress":
				_, err = bc.SendToAddress(ctx, SendToAddressReq{})
			case "SetBan":
				err = bc.SetBan(ctx, SetBanReq{Command: "add"})
			case "SetHDSeed":
				err = bc.SetHDSeed(ctx, SetHDSeedReq{})
			case "SetLabel":
				err = bc.SetLabel(ctx, SetLabelReq{})
			case "SetNetworkActive":
				_, err = bc.SetNetworkActive(ctx, SetNetworkActiveReq{})
			case "SetTxFee":
				_, err = bc.SetTxFee(ctx, SetTxFeeReq{})
			case "SetWalletFlag":
				_, err = bc.SetWalletFlag(ctx, SetWalletFlagReq{})
			case "SignMessage":
				_, err = bc.SignMessage(ctx, SignMessageReq{})
			case "SignMessageWithPrivkey":
				_, err = bc.SignMessageWithPrivkey(ctx, SignMessageWithPrivkeyReq{})
			case "SignRawTransactionWithKey":
				_, err = bc.SignRawTransactionWithKey(ctx, SignRawTransactionWithKeyReq{})
			case "SignRawTransactionWithWallet":
				_, err = bc.SignRawTransactionWithWallet(ctx, SignRawTransactionWithWalletReq{})
			case "SubmitBlock":
				_, err = bc.SubmitBlock(ctx, SubmitBlockReq{})
			case "SubmitHeader":
				err = bc.SubmitHeader(ctx, SubmitHeaderReq{})
			case "TestMempoolAccept":
				_, err = bc.TestMempoolAccept(ctx, TestMempoolAcceptReq{})
			case "UnloadWallet":
				_, err = bc.UnloadWallet(ctx, UnloadWalletReq{WalletName: "abc"})
			case "UpgradeWallet":
				_, err = bc.UpgradeWallet(ctx, UpgradeWalletReq{})
			case "Uptime":
				_, err = bc.Uptime(ctx)
			case "UtxoUpdatePsbt":
				_, err = bc.UtxoUpdatePsbt(ctx, UtxoUpdatePsbtReq{})
			case "ValidateAddress":
				_, err = bc.ValidateAddress(ctx, ValidateAddressReq{})
			case "VerifyChain":
				_, err = bc.VerifyChain(ctx, VerifyChainReq{})
			case "VerifyMessage":
				_, err = bc.VerifyMessage(ctx, VerifyMessageReq{})
			case "VerifyTxOutProof":
				_, err = bc.VerifyTxOutProof(ctx, VerifyTxOutProofReq{})
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
