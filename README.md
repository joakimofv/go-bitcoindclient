[![Go Reference](https://pkg.go.dev/badge/github.com/joakimofv/go-bitcoindclient/v22.svg)](https://pkg.go.dev/github.com/joakimofv/go-bitcoindclient/v22)

go-bitcoindclient
=================

A client for connecting to bitcoind from Go code.
All RPC methods are code-generated from the bitcoind help output.

# Import

```go
	bitcoindclient "github.com/joakimofv/go-bitcoindclient/v22"
```

# Dependencies

## libzmq

`libzmq` must be installed in order to compile, because it is needed by [this import](github.com/pebbe/zmq4) that handles ZMQ.
See the instructions for your OS on the [ZMQ website](https://zeromq.org/download/).

## bitcoind

The major version of this package tracks the major version of bitcoind (starting with v22).
Running the client against another version might or might not work, depending on specific changes to the RPC methods.
There is no explicit version check by default.

You don't need to have bitcoind installed in order to compile. To run the tests or re-generate the code you would need it.

# Basic Usage

### New

```go
bc, err := bitcoindclient.New(bitcoindclient.Config{
	RpcAddress:    "localhost:8332",
	RpcUser:       "name",
	RpcPassword:   "password",
	ZmqPubAddress: "tcp://localhost:12345",
})
if err != nil {
	// Handle err
}
```

See [Config](https://pkg.go.dev/github.com/joakimofv/go-bitcoindclient/v22#Config) for options.

### Ready: Check Connection

```go
for {
	if err := bc.Ready(); err != nil {
		// Inspect error, maybe give up
	} else {
		// Success!
		break
	}
}
```

There are also options that can be given to `Ready` to tell it to check the version or ZMQ messages received, see [ReadyOption](https://pkg.go.dev/github.com/joakimofv/go-bitcoindclient/v22#ReadyOption).

### Call RPC Methods

```go
resp, err := bc.GetBlock(ctx, GetBlockReq{
	Blockhash: "000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f",
})
if err != nil {
	// Handle err
}
fmt.Println(resp.Hex)
```

See [pkg.go.dev](https://pkg.go.dev/github.com/joakimofv/go-bitcoindclient/v22#pkg-index) for the full list of methods. Or see [BitcoinCore RPC Docs](https://bitcoincore.org/en/doc/), the methods are the same (in camelcase). Input is always incapsulated in a struct called `<Method>Req` (or empty) and output is always in a struct called `<Method>Resp` (or empty).

### Close

```go
if err := bc.Close(); err != nil {
	// Handle err
}
```

# ZMQ Messages

### Subscribe

```go
ch, cancel, err := bc.SubscribeHashBlock()
if err != nil {
	// Handle err
}
```

There are [SubscribeHashTx](https://pkg.go.dev/github.com/joakimofv/go-bitcoindclient/v22#SubscribeHashTx), [SubscribeHashBlock](https://pkg.go.dev/github.com/joakimofv/go-bitcoindclient/v22#SubscribeHashBlock), [SubscribeRawTx](https://pkg.go.dev/github.com/joakimofv/go-bitcoindclient/v22#SubscribeRawTx), [SubscribeRawBlock](https://pkg.go.dev/github.com/joakimofv/go-bitcoindclient/v22#SubscribeRawBlock), and [SubscribeSequence](https://pkg.go.dev/github.com/joakimofv/go-bitcoindclient/v22#SubscribeSequence).

### Receive

```go
select {
case msg, open := <-ch:
	if !open {
		// return
	}
	fmt.Println(hex.EncodeToString(msg.Hash[:]))
}
```

The message types are [HashMsg](https://pkg.go.dev/github.com/joakimofv/go-bitcoindclient/v22#HashMsg), [RawMsg](https://pkg.go.dev/github.com/joakimofv/go-bitcoindclient/v22#RawMsg) or [SequenceMsg](https://pkg.go.dev/github.com/joakimofv/go-bitcoindclient/v22#SequenceMsg).

### Cancel

```go
cancel()
```

This closes the channel and releases resources. Closing the client will do that too.

# Error Handling

The RPC methods may return errors that are of the type [\*BitcoindError](https://pkg.go.dev/github.com/joakimofv/go-bitcoindclient/v22#BitcoindError), that means the connection was successful but bitcoind had something to complain about related to the user input.

```go
var bErr *bitcoindclient.BitcoindError
if errors.As(err, &bErr) {
	fmt.Println(bErr.Code)
	fmt.Println(bErr.Message)
}
```

Other expected errors are `*url.Error` if the connection failed, or `context.Canceled`/`context.DeadlineExceeded` if the context was cancelled/expired.

If running against a different version of bitcoind you might get an error of type [\*UnmarshalError](https://pkg.go.dev/github.com/joakimofv/go-bitcoindclient/v22#UnmarshalError) because the response format was not as expected. If you get this when running against the matching version then please report an issue.

# Call Options

### Connection Retries

```go
ctx = bitcoindclient.UseConnectionRetries(ctx, 2)
```

Enable retry on connection error by modifying the context with [UseConnectionRetries](https://pkg.go.dev/github.com/joakimofv/go-bitcoindclient/v22#UseConnectionRetries).

### URI Path

```go
ctx = bitcoindclient.UseUriPath(ctx, "/wallet/mywallet")
```

Change the URI path for a call by modifying the context with [UseUriPath](https://pkg.go.dev/github.com/joakimofv/go-bitcoindclient/v22#UseUriPath).
