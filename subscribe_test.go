package bitcoindclient

import (
	"context"
	"encoding/hex"
	"math/rand"
	"testing"
	"time"
)

// TestSubscribe tests the Subscribe... functions, that ZMQ messages from bitcoind are received
// and pushed as events onto the subscription channels.
func TestSubscribe(t *testing.T) {
	for name, tc := range map[string]struct {
		rounds    int
		hashtx    int
		hashblock int
		rawtx     int
		rawblock  int
		sequence  int
	}{
		"hashtx":      {rounds: 3, hashtx: 10},
		"hashblock":   {rounds: 3, hashblock: 10},
		"rawtx":       {rounds: 3, rawtx: 10},
		"rawblock":    {rounds: 3, rawblock: 10},
		"sequence":    {rounds: 3, sequence: 10},
		"mixed":       {rounds: 40, hashtx: 1, hashblock: 1, rawtx: 1, rawblock: 1, sequence: 1},
		"mixed-multi": {rounds: 20, hashtx: 10, hashblock: 10, rawtx: 10, rawblock: 10, sequence: 10},
	} {
		t.Run(name, func(t *testing.T) {
			seed := time.Now().UnixNano()
			t.Logf("seed = %v", seed)
			rand.Seed(seed)

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

			var hashtxCh [](chan HashMsg)
			var hashblockCh [](chan HashMsg)
			var rawtxCh [](chan RawMsg)
			var rawblockCh [](chan RawMsg)
			var sequenceCh [](chan SequenceMsg)
			var activeHashtxCh []int
			var activeHashblockCh []int
			var activeRawtxCh []int
			var activeRawblockCh []int
			var activeSequenceCh []int
			var hashtxCancel []func()
			var hashblockCancel []func()
			var rawtxCancel []func()
			var rawblockCancel []func()
			var sequenceCancel []func()
			for i := 0; i < tc.hashtx; i++ {
				ch, cancel, err := bc.SubscribeHashTx()
				if err != nil {
					t.Fatal(err)
				}
				hashtxCh = append(hashtxCh, ch)
				activeHashtxCh = append(activeHashtxCh, i)
				hashtxCancel = append(hashtxCancel, cancel)
			}
			for i := 0; i < tc.hashblock; i++ {
				ch, cancel, err := bc.SubscribeHashBlock()
				if err != nil {
					t.Fatal(err)
				}
				hashblockCh = append(hashblockCh, ch)
				activeHashblockCh = append(activeHashblockCh, i)
				hashblockCancel = append(hashblockCancel, cancel)
			}
			for i := 0; i < tc.rawtx; i++ {
				ch, cancel, err := bc.SubscribeRawTx()
				if err != nil {
					t.Fatal(err)
				}
				rawtxCh = append(rawtxCh, ch)
				activeRawtxCh = append(activeRawtxCh, i)
				rawtxCancel = append(rawtxCancel, cancel)
			}
			for i := 0; i < tc.rawblock; i++ {
				ch, cancel, err := bc.SubscribeRawBlock()
				if err != nil {
					t.Fatal(err)
				}
				rawblockCh = append(rawblockCh, ch)
				activeRawblockCh = append(activeRawblockCh, i)
				rawblockCancel = append(rawblockCancel, cancel)
			}
			for i := 0; i < tc.sequence; i++ {
				ch, cancel, err := bc.SubscribeSequence()
				if err != nil {
					t.Fatal(err)
				}
				sequenceCh = append(sequenceCh, ch)
				activeSequenceCh = append(activeSequenceCh, i)
				sequenceCancel = append(sequenceCancel, cancel)
			}

			// Create a wallet and get an address.
			ctx, ctxCancel := context.WithTimeout(context.Background(), 20*time.Second)
			defer ctxCancel()
			ctx = UseConnectionRetries(ctx, 2)
			_, err = bc.CreateWallet(ctx, CreateWalletReq{})
			if err != nil {
				t.Fatal(err)
			}
			address, err := bc.GetNewAddress(ctx, GetNewAddressReq{})
			if err != nil {
				t.Fatal(err)
			}

			prevRound := time.Now()
			for n := 0; n < tc.rounds; n++ {
				if n != 0 {
					t.Logf("[%d] %v", n-1, time.Since(prevRound))
					prevRound = time.Now()
				}
				ctx, ctxCancel = context.WithTimeout(context.Background(), 20*time.Second)
				defer ctxCancel()
				// New blocks, and one transaction per block. Run concurrent with reading from the channels.
				blockHashesCh := make(chan ([]string), 1)
				const nBlocks = 101
				go func() {
					resp, err := bc.GenerateToAddress(ctx, GenerateToAddressReq{Address: address.Str, NBlocks: nBlocks})
					if err != nil {
						t.Error(err)
					}
					blockHashesCh <- resp.Hex
				}()
				var rpcBlockhashes []string

				// Catch the block notifications.
				for _, i := range rand.Perm(len(activeHashblockCh)) {
					ch := hashblockCh[activeHashblockCh[i]]
				RECV:
					for {
						select {
						case hashMsg := <-ch:
							if int(hashMsg.Seq) < len(rpcBlockhashes) {
								if hex.EncodeToString(hashMsg.Hash[:]) != rpcBlockhashes[hashMsg.Seq] {
									t.Errorf("[%d] expected %v, got %v", hashMsg.Seq, rpcBlockhashes[hashMsg.Seq], hex.EncodeToString(hashMsg.Hash[:]))
								}
							}
							if (hashMsg.Seq+1)%nBlocks == 0 {
								break RECV
							}
						case rpcBlockhashes = <-blockHashesCh:
						case <-ctx.Done():
							t.Fatal()
						}
					}
				}
				for _, i := range rand.Perm(len(activeRawblockCh)) {
					ch := rawblockCh[activeRawblockCh[i]]
				RECV2:
					for {
						select {
						case rawMsg := <-ch:
							if len(rawMsg.Serialized) < 1 {
								t.Error("empty")
							}
							if (rawMsg.Seq+1)%nBlocks == 0 {
								break RECV2
							}
						case <-ctx.Done():
							t.Fatal()
						}
					}
				}
				// Catch the tx notifications.
				for _, i := range rand.Perm(len(activeHashtxCh)) {
					ch := hashtxCh[activeHashtxCh[i]]
				RECV3:
					for {
						select {
						case hashMsg := <-ch:
							if hashMsg.Hash == [32]byte{} {
								t.Error("empty")
							}
							if hashMsg.Seq%nBlocks == 0 && hashMsg.Seq != 0 {
								break RECV3
							}
						case <-ctx.Done():
							t.Fatal()
						}
					}
				}
				for _, i := range rand.Perm(len(activeRawtxCh)) {
					ch := rawtxCh[activeRawtxCh[i]]
				RECV4:
					for {
						select {
						case rawMsg := <-ch:
							if len(rawMsg.Serialized) < 1 {
								t.Error("empty")
							}
							if rawMsg.Seq%nBlocks == 0 && rawMsg.Seq != 0 {
								break RECV4
							}
						case <-ctx.Done():
							t.Fatal()
						}
					}
				}
				// Catch the sequence notifications.
				for _, i := range rand.Perm(len(activeSequenceCh)) {
					ch := sequenceCh[activeSequenceCh[i]]
					var lastHash [32]byte
				RECV5:
					for {
						select {
						case sequenceMsg := <-ch:
							if sequenceMsg.Event != BlockConnected {
								t.Error(sequenceMsg.Event)
							}
							lastHash = sequenceMsg.Hash
							if len(rpcBlockhashes) > 0 && hex.EncodeToString(lastHash[:]) == rpcBlockhashes[len(rpcBlockhashes)-1] {
								break RECV5
							}
						case rpcBlockhashes = <-blockHashesCh:
							if hex.EncodeToString(lastHash[:]) == rpcBlockhashes[len(rpcBlockhashes)-1] {
								break RECV5
							}
						case <-ctx.Done():
							t.Fatal()
						}
					}
				}

				// Done! Make sure we wait for the RPC to finish before manipulating subscriptions for the next round.
				if len(rpcBlockhashes) == 0 {
					rpcBlockhashes = <-blockHashesCh
				}

				// Cancel some subscriptions at random.
				const div = 2
				var toCancel []int
				activeHashtxCh, toCancel = divideRange(activeHashtxCh, div)
				for _, i := range toCancel {
					hashtxCancel[i]()
					if msg, open := <-hashtxCh[i]; open {
						t.Fatal("leaky message", msg)
					}
				}
				activeHashblockCh, toCancel = divideRange(activeHashblockCh, div)
				for _, i := range toCancel {
					hashblockCancel[i]()
					if msg, open := <-hashblockCh[i]; open {
						t.Fatal("leaky message", msg)
					}
				}
				activeRawtxCh, toCancel = divideRange(activeRawtxCh, div)
				for _, i := range toCancel {
					rawtxCancel[i]()
					if msg, open := <-rawtxCh[i]; open {
						t.Fatal("leaky message", msg)
					}
				}
				activeRawblockCh, toCancel = divideRange(activeRawblockCh, div)
				for _, i := range toCancel {
					rawblockCancel[i]()
					if msg, open := <-rawblockCh[i]; open {
						t.Fatal("leaky message", msg)
					}
				}
				activeSequenceCh, toCancel = divideRange(activeSequenceCh, div)
				for _, i := range toCancel {
					sequenceCancel[i]()
					if msg, open := <-sequenceCh[i]; open {
						t.Fatal("leaky message", msg)
					}
				}
				// Add some subscriptions at random.
				for i := 0; i < tc.hashtx; i++ {
					if len(activeHashtxCh) == tc.hashtx {
						break
					}
					if rand.Intn(div) != 0 {
						continue
					}
					ch, cancel, err := bc.SubscribeHashTx()
					if err != nil {
						t.Fatal(err)
					}
					hashtxCh = append(hashtxCh, ch)
					activeHashtxCh = append(activeHashtxCh, len(hashtxCh)-1)
					hashtxCancel = append(hashtxCancel, cancel)
				}
				for i := 0; i < tc.hashblock; i++ {
					if len(activeHashblockCh) == tc.hashblock {
						break
					}
					if rand.Intn(div) != 0 {
						continue
					}
					ch, cancel, err := bc.SubscribeHashBlock()
					if err != nil {
						t.Fatal(err)
					}
					hashblockCh = append(hashblockCh, ch)
					activeHashblockCh = append(activeHashblockCh, len(hashblockCh)-1)
					hashblockCancel = append(hashblockCancel, cancel)
				}
				for i := 0; i < tc.rawtx; i++ {
					if len(activeRawtxCh) == tc.rawtx {
						break
					}
					if rand.Intn(div) != 0 {
						continue
					}
					ch, cancel, err := bc.SubscribeRawTx()
					if err != nil {
						t.Fatal(err)
					}
					rawtxCh = append(rawtxCh, ch)
					activeRawtxCh = append(activeRawtxCh, len(rawtxCh)-1)
					rawtxCancel = append(rawtxCancel, cancel)
				}
				for i := 0; i < tc.rawblock; i++ {
					if len(activeRawblockCh) == tc.rawblock {
						break
					}
					if rand.Intn(div) != 0 {
						continue
					}
					ch, cancel, err := bc.SubscribeRawBlock()
					if err != nil {
						t.Fatal(err)
					}
					rawblockCh = append(rawblockCh, ch)
					activeRawblockCh = append(activeRawblockCh, len(rawblockCh)-1)
					rawblockCancel = append(rawblockCancel, cancel)
				}
				for i := 0; i < tc.sequence; i++ {
					if len(activeSequenceCh) == tc.sequence {
						break
					}
					if rand.Intn(div) != 0 {
						continue
					}
					ch, cancel, err := bc.SubscribeSequence()
					if err != nil {
						t.Fatal(err)
					}
					sequenceCh = append(sequenceCh, ch)
					activeSequenceCh = append(activeSequenceCh, len(sequenceCh)-1)
					sequenceCancel = append(sequenceCancel, cancel)
				}
			}

			err = bc.Close()
			if err != nil {
				t.Fatal(err)
			}
			// See that all channels have been drained (by test) and closed (by bc.Close()).
			for _, ch := range hashtxCh {
				if msg, open := <-ch; open {
					t.Error(msg)
				}
			}
			for _, ch := range hashblockCh {
				if msg, open := <-ch; open {
					t.Error(msg)
				}
			}
			for _, ch := range rawtxCh {
				if msg, open := <-ch; open {
					t.Error(msg)
				}
			}
			for _, ch := range rawblockCh {
				if msg, open := <-ch; open {
					t.Error(msg)
				}
			}
			for _, ch := range sequenceCh {
				if msg, open := <-ch; open {
					t.Error(msg)
				}
			}
		})
	}
}

func divideRange(s []int, div int) ([]int, []int) {
	var remaining []int
	var taken []int
	for _, i := range s {
		if rand.Intn(div) == 0 {
			taken = append(taken, i)
		} else {
			remaining = append(remaining, i)
		}
	}
	return remaining, taken
}
