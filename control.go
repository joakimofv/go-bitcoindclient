// THIS FILE IS GENERATED CODE, MANUAL EDITS WILL BE OVERWRITTEN

package bitcoindclient

import (
	"bytes"
	"context"
	"encoding/json"
	"reflect"
)

// GetMemoryInfoReq holds the arguments for the GetMemoryInfo call.
//  1. mode    (string, optional, default="stats") determines what kind of information is returned.
//             - "stats" returns general statistics about memory usage in the daemon.
//             - "mallocinfo" returns an XML string describing low-level heap state (only available if compiled with glibc 2.10+).
type GetMemoryInfoReq struct {
	// determines what kind of information is returned.
	// - "stats" returns general statistics about memory usage in the daemon.
	// - "mallocinfo" returns an XML string describing low-level heap state (only available if compiled with glibc 2.10+).
	// Default: "stats"
	Mode string `json:"mode,omitempty"`
}

// GetMemoryInfoResp holds the response to the GetMemoryInfo call.
//
// ALTERNATIVE (mode "stats")
//  {                         (json object)
//    "locked" : {            (json object) Information about locked memory manager
//      "used" : n,           (numeric) Number of bytes used
//      "free" : n,           (numeric) Number of bytes available in current arenas
//      "total" : n,          (numeric) Total number of bytes managed
//      "locked" : n,         (numeric) Amount of bytes that succeeded locking. If this number is smaller than total, locking pages failed at some point and key data could be swapped to disk.
//      "chunks_used" : n,    (numeric) Number allocated chunks
//      "chunks_free" : n     (numeric) Number unused chunks
//    }
//  }
//
// ALTERNATIVE (mode "mallocinfo")
//  "str"    (string) "<malloc version="1">..."
type GetMemoryInfoResp struct {
	ModeStats GetMemoryInfoRespModeStats

	// "<malloc version="1">..."
	Str string
}

func (alts GetMemoryInfoResp) MarshalJSON() ([]byte, error) {
	if !reflect.ValueOf(alts.ModeStats).IsZero() {
		return json.Marshal(alts.ModeStats)
	}
	return json.Marshal(alts.Str)
}

func (alts *GetMemoryInfoResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.ModeStats) == nil {
		return nil
	}
	alts.ModeStats = reset.ModeStats
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Str) == nil {
		return nil
	}
	alts.Str = reset.Str
	return &UnmarshalError{B: b, structName: "GetMemoryInfoResp"}
}

type GetMemoryInfoRespModeStats struct {
	// Information about locked memory manager
	Locked struct {
		// Number of bytes used
		Used float64 `json:"used"`

		// Number of bytes available in current arenas
		Free float64 `json:"free"`

		// Total number of bytes managed
		Total float64 `json:"total"`

		// Amount of bytes that succeeded locking. If this number is smaller than total, locking pages failed at some point and key data could be swapped to disk.
		Locked float64 `json:"locked"`

		// Number allocated chunks
		ChunksUsed float64 `json:"chunks_used"`

		// Number unused chunks
		ChunksFree float64 `json:"chunks_free"`
	} `json:"locked"`
}

// GetMemoryInfo RPC method.
// Returns an object containing information about memory usage.
func (bc *BitcoindClient) GetMemoryInfo(ctx context.Context, args GetMemoryInfoReq) (result GetMemoryInfoResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "getmemoryinfo", args, false); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// GetRpcInfoResp holds the response to the GetRpcInfo call.
//  {                          (json object)
//    "active_commands" : [    (json array) All active commands
//      {                      (json object) Information about an active command
//        "method" : "str",    (string) The name of the RPC command
//        "duration" : n       (numeric) The running time in microseconds
//      },
//      ...
//    ],
//    "logpath" : "str"        (string) The complete file path to the debug log
//  }
type GetRpcInfoResp struct {
	// All active commands
	ActiveCommands []GetRpcInfoRespActiveCommands `json:"active_commands"`

	// The complete file path to the debug log
	LogPath string `json:"logpath"`
}

// Information about an active command
type GetRpcInfoRespActiveCommands struct {
	// The name of the RPC command
	Method string `json:"method"`

	// The running time in microseconds
	Duration float64 `json:"duration"`
}

// GetRpcInfo RPC method.
// Returns details of the RPC server.
func (bc *BitcoindClient) GetRpcInfo(ctx context.Context) (result GetRpcInfoResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "getrpcinfo", nil, false); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// HelpReq holds the arguments for the Help call.
//  1. command    (string, optional, default=all commands) The command to get help on
type HelpReq struct {
	// The command to get help on
	// Default: all commands
	Command string `json:"command,omitempty"`
}

// HelpResp holds the response to the Help call.
//  "str"    (string) The help text
type HelpResp struct {
	// The help text
	Str string
}

func (alts HelpResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.Str)
}

func (alts *HelpResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Str) == nil {
		return nil
	}
	alts.Str = reset.Str
	return &UnmarshalError{B: b, structName: "HelpResp"}
}

// Help RPC method.
// List all commands, or get help for a specified command.
func (bc *BitcoindClient) Help(ctx context.Context, args HelpReq) (result HelpResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "help", args, false); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// LoggingReq holds the arguments for the Logging call.
//  1. include                    (json array, optional) The categories to add to debug logging
//       [
//         "include_category",    (string) the valid logging category
//         ...
//       ]
//  2. exclude                    (json array, optional) The categories to remove from debug logging
//       [
//         "exclude_category",    (string) the valid logging category
//         ...
//       ]
type LoggingReq struct {
	// The categories to add to debug logging
	// Element: IncludeCategory    the valid logging category
	Include []string `json:"include,omitempty"`

	// The categories to remove from debug logging
	// Element: ExcludeCategory    the valid logging category
	Exclude []string `json:"exclude,omitempty"`
}

// Logging RPC method.
// Gets and sets the logging configuration.
// When called without an argument, returns the list of categories with status that are currently being debug logged or not.
// When called with arguments, adds or removes categories from debug logging and return the lists above.
// The arguments are evaluated in order "include", "exclude".
// If an item is both included and excluded, it will thus end up being excluded.
// The valid logging categories are: net, tor, mempool, http, bench, zmq, walletdb, rpc, estimatefee, addrman, selectcoins, reindex, cmpctblock, rand, prune, proxy, mempoolrej, libevent, coindb, qt, leveldb, validation, i2p, ipc
// In addition, the following are available as category names with special meanings:
//   - "all",  "1" : represent all logging categories.
//   - "none", "0" : even if other logging categories are specified, ignore all of them.
func (bc *BitcoindClient) Logging(ctx context.Context, args LoggingReq) (err error) {
	_, err = bc.sendRequest(ctx, "logging", args, false)
	return
}

// StopResp holds the response to the Stop call.
//  "str"    (string) A string with the content 'Bitcoin Core stopping'
type StopResp struct {
	// A string with the content 'Bitcoin Core stopping'
	Str string
}

func (alts StopResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.Str)
}

func (alts *StopResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Str) == nil {
		return nil
	}
	alts.Str = reset.Str
	return &UnmarshalError{B: b, structName: "StopResp"}
}

// Stop RPC method.
// Request a graceful shutdown of Bitcoin Core.
func (bc *BitcoindClient) Stop(ctx context.Context) (result StopResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "stop", nil, false); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// UptimeResp holds the response to the Uptime call.
//  n    (numeric) The number of seconds that the server has been running
type UptimeResp struct {
	// The number of seconds that the server has been running
	N float64
}

func (alts UptimeResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.N)
}

func (alts *UptimeResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.N) == nil {
		return nil
	}
	alts.N = reset.N
	return &UnmarshalError{B: b, structName: "UptimeResp"}
}

// Uptime RPC method.
// Returns the total uptime of the server.
func (bc *BitcoindClient) Uptime(ctx context.Context) (result UptimeResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "uptime", nil, false); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}
