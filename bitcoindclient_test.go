package bitcoindclient

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

// TestReady tests the Ready function.
func TestReady(t *testing.T) {
	for name, tc := range map[string]struct {
		versionCheck         []int
		expectFailure        bool
		subscribeEventWithin time.Duration
		generateBlocks       bool
	}{
		"basic":                               {},
		"versionCheck":                        {versionCheck: []int{}},
		"versionCheck-explicit":               {versionCheck: []int{MAJOR_VERSION}},
		"versionCheck-bad":                    {versionCheck: []int{MAJOR_VERSION - 1}, expectFailure: true},
		"versionCheck-bad-and-good":           {versionCheck: []int{MAJOR_VERSION - 1, MAJOR_VERSION, MAJOR_VERSION + 1}},
		"subscribeEventWithin":                {subscribeEventWithin: time.Hour, expectFailure: true},
		"subscribeEventWithin-generateBlocks": {subscribeEventWithin: time.Second, generateBlocks: true},
	} {
		t.Run(name, func(t *testing.T) {
			rpcAddress, rpcUser, rpcPassword, zmqPubAddress, err := startBitcoind()
			if err != nil {
				t.Fatal(err)
			}
			defer stopBitcoind(t)

			bc, err := New(Config{
				RpcAddress:    rpcAddress,
				RpcUser:       rpcUser,
				RpcPassword:   rpcPassword,
				ZmqPubAddress: zmqPubAddress,
			})
			if err != nil {
				t.Fatal(err)
			}

			if tc.generateBlocks {
				_, cancel, err := bc.SubscribeHashBlock()
				if err != nil {
					t.Fatal(err)
				}
				defer cancel()
				_, err = bc.CreateWallet(UseConnectionRetries(context.Background(), 2), CreateWalletReq{})
				if err != nil {
					t.Fatal(err)
				}
				result, err := bc.GetNewAddress(context.Background(), GetNewAddressReq{})
				if err != nil {
					t.Fatal(err)
				}
				_, err = bc.GenerateToAddress(context.Background(), GenerateToAddressReq{Address: result.Str, NBlocks: 1})
				if err != nil {
					t.Fatal(err)
				}
			}

			var opts []ReadyOption
			if tc.versionCheck != nil {
				if len(tc.versionCheck) == 0 {
					opts = append(opts, WithVersionCheck())
				} else {
					checkFn := func(version int) bool {
						for _, v := range tc.versionCheck {
							if version/10000 == v {
								return true
							}
						}
						return false
					}
					opts = append(opts, WithVersionCheckFn(checkFn))
				}
			}
			if tc.subscribeEventWithin > 0 {
				opts = append(opts, WithZmqMessageWithin(tc.subscribeEventWithin))
			}
			if err := bc.Ready(opts...); err != nil {
				t.Log("First try Ready() got error:", err)
				t.Log("Sleeping 300 msec before trying again...")
				time.Sleep(300 * time.Millisecond)
				if err := bc.Ready(opts...); err != nil {
					if tc.expectFailure {
						t.Log("[expected] Second try Ready() got error:", err)
					} else {
						t.Error("Second try Ready() got error:", err)
					}
				} else if tc.expectFailure {
					t.Error("Expected error, got nil")
				}
			} else if tc.expectFailure {
				t.Error("Expected error, got nil")
			}

			err = bc.Close()
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

// TestSomeRpc tests some specific RPC methods. Making it all-covering on the methods would be too much work
// and hard to maintain, just do a few to see that things work in general.
func TestSomeRpc(t *testing.T) {
	rpcAddress, rpcUser, rpcPassword, _, err := startBitcoind()
	if err != nil {
		t.Fatal(err)
	}
	defer stopBitcoind(t)

	bc, err := New(Config{
		RpcAddress:  rpcAddress,
		RpcUser:     rpcUser,
		RpcPassword: rpcPassword,
	})
	if err != nil {
		t.Fatal(err)
	}

	for name, tc := range map[string]struct {
		method   []string
		args     []interface{}
		expected []interface{}
	}{
		"Uptime": {method: []string{"Uptime"}, args: []interface{}{nil}, expected: []interface{}{isNonZero}},
		"CreateWallet": {
			method:   []string{"CreateWallet"},
			args:     []interface{}{CreateWalletReq{WalletName: "abcd", Passphrase: "efgh"}},
			expected: []interface{}{CreateWalletResp{Name: "abcd"}},
		},
		"GenerateToAddress": {
			method: []string{"CreateWallet", "GetNewAddress", "GenerateToAddress", "ListAddressGroupings"},
			args: []interface{}{
				CreateWalletReq{WalletName: "efgh", Passphrase: "efgh"},
				GetNewAddressReq{},
				GenerateToAddressReq{NBlocks: 101},
				nil,
			},
			expected: []interface{}{
				CreateWalletResp{Name: "efgh"},
				isNonZero,
				isNonZero,
				isNonZeroAndPrint,
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			address := ""
			wallet := ""
			for i, method := range tc.method {
				t.Run(method, func(t *testing.T) {
					ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
					defer cancel()
					ctx = UseConnectionRetries(ctx, -1)

					var result interface{}
					var err error
					switch method {
					case "Uptime":
						time.Sleep(time.Second)
						result, err = bc.Uptime(ctx)
					case "CreateWallet":
						result, err = bc.CreateWallet(ctx, tc.args[i].(CreateWalletReq))
						if resp, ok := result.(CreateWalletResp); ok {
							wallet = resp.Name
						}
					case "GetNewAddress":
						result, err = bc.GetNewAddress(UseUriPath(ctx, "/wallet/"+wallet), tc.args[i].(GetNewAddressReq))
						if resp, ok := result.(GetNewAddressResp); ok {
							address = resp.Str
						}
					case "GenerateToAddress":
						req := tc.args[i].(GenerateToAddressReq)
						req.Address = address
						result, err = bc.GenerateToAddress(ctx, req)
					case "ListAddressGroupings":
						result, err = bc.ListAddressGroupings(UseUriPath(ctx, "/wallet/"+wallet))
					default:
						t.Fatal("missing case for " + method)
					}

					if err != nil {
						t.Error(err)
					} else if tc.expected[i] != nil {
						if doPrint, checkNonZero := tc.expected[i].(nonZero); checkNonZero {
							if reflect.ValueOf(result).IsZero() {
								t.Errorf("IsZero: %T", result)
							} else if doPrint.doPrint {
								t.Log(result)
							}
						} else {
							diff := cmp.Diff(tc.expected[i], result)
							if diff != "" {
								t.Error(diff)
							}
						}
					} else {
						// Will want to do a check on this eventually, print it for now.
						t.Log(result)
					}
				})
			}
		})
	}

	err = bc.Close()
	if err != nil {
		t.Fatal(err)
	}
}

type nonZero struct {
	doPrint bool
}

var isNonZero = nonZero{}
var isNonZeroAndPrint = nonZero{true}
