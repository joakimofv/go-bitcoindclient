package bitcoindclient

import (
	"errors"
	"net/http"
	"sync"
	"sync/atomic"

	"github.com/joakimofv/sanity"
	zmq "github.com/pebbe/zmq4"
)

const (
	MAJOR_VERSION               = 23
	DefaultSubChannelBufferSize = 2
	DefaultRpcUriPath           = "/"
)

var (
	ErrRpcDisabled       = errors.New("RPC disabled (RpcAddress was not set).")
	ErrSubscribeDisabled = errors.New("Subscribe disabled (ZmqPubAddress was not set).")
	ErrSubscribeExited   = errors.New("Subscription backend has exited.")
)

type Config struct {
	// RpcAddress is the address, formatted as "host:port", of the JSON-RPC interface of bitcoind.
	//
	// Example: ":8332" (localhost, mainnet)
	//          ":18332" (localhost, testnet/regtest)
	//          "2.2.2.2:8332" (remote, mainnet)
	RpcAddress string

	// RpcUser is the 'rpcuser' option that bitcoind was configured with, or the equivalent in 'rpcauth'.
	RpcUser string

	// RpcUser is the 'rpcpassword' option that bitcoind was configured with, or the equivalent in 'rpcauth'.
	RpcPassword string

	// RpcUriPath is the URI path of requests. Can be modified on a per request basis by using ctx = UseUriPath(ctx, newPath).
	// Default "/".
	//
	// Example: "/wallet/<WalletName>" (specify which wallet to use for a wallet command)
	RpcUriPath string

	// ZmqPubAddress is the public address that the bitcoind instance uses for zmqpub,
	// corresponding to what is set when starting bitcoind through one or multiple of:
	// {-zmqpubhashtx=address -zmqpubhashblock=address -zmqpubrawblock=address -zmqpubrawtx=address -zmqpubsequence=address}.
	// Only a single address is supported in this client. Either use the same address for
	// all desired topics when starting bitcoind, or create a seperate client for each address.
	//
	// Example: "tcp://8.8.8.8:1234"
	// More examples at: https://github.com/bitcoin/bitcoin/blob/master/doc/zmq.md (the host part in those examples
	// are local IPs and should be replaced with public IPs here on the client side)
	//
	// If ZmqPubAddress is not set then the Subscribe functions will return ErrSubscribeDisabled when called.
	ZmqPubAddress string

	// SubChannelBufferSize sets the number of entries that a subscription channel can hold
	// before dropping entries, if it is not drained fast enough.
	// If not set (or set to zero) then defaults to DefaultSubChannelBufferSize.
	SubChannelBufferSize int
}

// BitcoindClient is a client that provides methods for interacting with bitcoind.
// Must be created with New and destroyed with Close.
//
// Clients are safe for concurrent use by multiple goroutines.
type BitcoindClient struct {
	closed int32 // Set atomically.
	wg     sync.WaitGroup
	quit   chan struct{}

	Cfg Config

	// httpClient is a reusable and concurrency safe client for all HTTP requests.
	// JSON-RPC is over HTTP.
	httpClient http.Client

	// ZMQ subscription related things.
	zctx *zmq.Context
	zsub *zmq.Socket
	subs subscriptions
	// subs.zfront --> zback is used like a channel to send messages to the zmqHandler goroutine.
	// Have to use zmq sockets in place of native channels for communication from
	// other functions to the goroutine, since it is constantly waiting on the zsub socket,
	// it can't select on a channel at the same time but can poll on multiple sockets.
	zback *zmq.Socket
}

// New returns an initiated client, or an error.
// Missing RpcAddress in Config will disable the RPC methods, and missing ZmqPubAddress
// will disable the Subscribe methods.
// New does not try using the RPC connection and can't detect if the ZMQ connection works,
// you need to call Ready in order to check connection health.
func New(cfg Config) (*BitcoindClient, error) {
	bc := &BitcoindClient{
		Cfg:  cfg,
		quit: make(chan struct{}),
	}

	// JSON-RPC.
	if bc.Cfg.RpcAddress != "" {
		if bc.Cfg.RpcUriPath == "" {
			bc.Cfg.RpcUriPath = DefaultRpcUriPath
		}
		if err := sanity.SpecificFieldsInitiated(bc.Cfg, "RpcUser", "RpcPassword"); err != nil {
			return nil, err
		}
	}

	// ZMQ Subscribe.
	if bc.Cfg.ZmqPubAddress != "" {
		if bc.Cfg.SubChannelBufferSize == 0 {
			bc.Cfg.SubChannelBufferSize = DefaultSubChannelBufferSize
		}

		zctx, err := zmq.NewContext()
		if err != nil {
			return nil, err
		}
		zsub, err := zctx.NewSocket(zmq.SUB)
		if err != nil {
			return nil, err
		}
		if err := zsub.Connect(bc.Cfg.ZmqPubAddress); err != nil {
			return nil, err
		}
		zback, err := zctx.NewSocket(zmq.PAIR)
		if err != nil {
			return nil, err
		}
		if err := zback.Bind("inproc://channel"); err != nil {
			return nil, err
		}
		zfront, err := zctx.NewSocket(zmq.PAIR)
		if err != nil {
			return nil, err
		}
		if err := zfront.Connect("inproc://channel"); err != nil {
			return nil, err
		}

		bc.zctx = zctx
		bc.zsub = zsub
		bc.subs.exited = make(chan struct{})
		bc.subs.zfront = zfront
		bc.zback = zback

		bc.wg.Add(1)
		go bc.zmqHandler()
	}

	return bc, nil
}

// Close terminates the client and releases resources.
func (bc *BitcoindClient) Close() (err error) {
	if !atomic.CompareAndSwapInt32(&bc.closed, 0, 1) {
		return errors.New("BitcoindClient already closed")
	}
	if bc.zctx != nil {
		bc.zctx.SetRetryAfterEINTR(false)
		bc.subs.Lock()
		select {
		case <-bc.subs.exited:
		default:
			bc.subs.zfront.SendMessage("term")
		}
		bc.subs.Unlock()
		<-bc.subs.exited
		err = bc.zctx.Term()
	}
	close(bc.quit)
	bc.wg.Wait()
	return nil
}
