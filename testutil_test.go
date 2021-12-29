package bitcoindclient

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"
)

const (
	datadir = "testtmp"
)

var (
	bitcoindProcess *exec.Cmd
	stderr          *bytes.Buffer
)

// startBitcoind starts a bitcoind process in regtest mode. To be used inside a test and stopped
// (with stopBitcoind) when the test ends.
//
// After a test crashes then the bitcoind might have to be stopped manually. Do this in an UNIX shell:
// > ps -a | grep bitcoind
// > kill <number>
func startBitcoind() (rpcAddress, rpcUser, rpcPassword, zmqPubAddress string, err error) {
	// Check that bitcoind is installed and has the matching version.
	var bitcoindVersion []byte
	bitcoindVersion, err = exec.Command("bitcoind", "-version").CombinedOutput()
	if err != nil {
		return
	}
	thisPackageVersion := fmt.Sprintf("v%d.", MAJOR_VERSION)
	if !strings.Contains(string(bitcoindVersion), thisPackageVersion) {
		err = fmt.Errorf("bitcoind -version %q doesn't match MAJOR_VERSION in this package %q.", bitcoindVersion, thisPackageVersion)
		return
	}

	rpcPort := "28332"
	rpcAddress = "localhost:" + rpcPort
	rpcUser = "testuser"
	rpcPassword = "test1234"
	zmqPubAddress = "tcp://127.0.0.1:28331"
	if err = os.RemoveAll(datadir); err != nil {
		return
	}
	if err = os.Mkdir(datadir, 0777); err != nil {
		return
	}
	bitcoindProcess = exec.Command("bitcoind", "-regtest", "-datadir="+datadir,
		"-port=28444", // P2P port irrelevant in regtest, but needs to be set to non-colliding value.
		"-rpcport="+fmt.Sprint(rpcPort),
		"-rpcuser="+rpcUser,
		"-rpcpassword="+rpcPassword,
		"-zmqpubhashblock="+zmqPubAddress,
		"-zmqpubhashtx="+zmqPubAddress,
		"-zmqpubrawblock="+zmqPubAddress,
		"-zmqpubrawtx="+zmqPubAddress,
		"-zmqpubsequence="+zmqPubAddress,
	)
	stderr = new(bytes.Buffer)
	bitcoindProcess.Stderr = stderr
	if err = bitcoindProcess.Start(); err != nil {
		return
	}
	return
}

func stopBitcoind(t *testing.T) {
	errCh := make(chan error)
	go func() {
		errCh <- nil
		errCh <- bitcoindProcess.Wait()
	}()
	<-errCh
	select {
	case err := <-errCh:
		// Process already exited (failed).
		if err != nil {
			t.Error(fmt.Errorf("bitcoind failure: %s", stderr.String()))
		}
	case <-time.After(200 * time.Millisecond):
		// Go on and interrupt the process.
		if err := bitcoindProcess.Process.Signal(os.Interrupt); err != nil {
			t.Error(err)
		}
		<-errCh // Expecting error: "signal: interrupt". No conventient way to check on it, so just discard.
	}
	bitcoindProcess = nil
	if err := os.RemoveAll(datadir); err != nil {
		t.Error(err)
	}
}
