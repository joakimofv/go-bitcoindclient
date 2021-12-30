// THIS FILE IS GENERATED CODE, MANUAL EDITS WILL BE OVERWRITTEN

package bitcoindclient

import (
	"bytes"
	"context"
	"encoding/json"
)

// AddNodeReq holds the arguments for the AddNode call.
//  1. node       (string, required) The node (see getpeerinfo for nodes)
//  2. command    (string, required) 'add' to add a node to the list, 'remove' to remove a node from the list, 'onetry' to try a connection to the node once
type AddNodeReq struct {
	// The node (see getpeerinfo for nodes)
	Node string `json:"node"`

	// 'add' to add a node to the list, 'remove' to remove a node from the list, 'onetry' to try a connection to the node once
	Command string `json:"command"`
}

// AddNode RPC method.
// Attempts to add or remove a node from the addnode list.
// Or try a connection to a node once.
// Nodes added using addnode (or -connect) are protected from DoS disconnection and are not required to be
// full nodes/support SegWit as other outbound peers are (though such peers will not be synced from).
// Addnode connections are limited to 8 at a time and are counted separately from the -maxconnections limit.
func (bc *BitcoindClient) AddNode(ctx context.Context, args AddNodeReq) (err error) {
	_, err = bc.sendRequest(ctx, "addnode", args)
	return
}

// ClearBanned RPC method.
// Clear all banned IPs.
func (bc *BitcoindClient) ClearBanned(ctx context.Context) (err error) {
	_, err = bc.sendRequest(ctx, "clearbanned", nil)
	return
}

// DisconnectNodeReq holds the arguments for the DisconnectNode call.
//  1. address    (string, optional, default=fallback to nodeid) The IP address/port of the node
//  2. nodeid     (numeric, optional, default=fallback to address) The node ID (see getpeerinfo for node IDs)
type DisconnectNodeReq struct {
	// The IP address/port of the node
	// Default: fallback to nodeid
	Address string `json:"address,omitempty"`

	// The node ID (see getpeerinfo for node IDs)
	// Default: fallback to address
	NodeID *float64 `json:"nodeid,omitempty"`
}

// DisconnectNode RPC method.
// Immediately disconnects from the specified peer node.
// Strictly one out of 'address' and 'nodeid' can be provided to identify the node.
// To disconnect by nodeid, either set 'address' to the empty string, or call using the named 'nodeid' argument only.
func (bc *BitcoindClient) DisconnectNode(ctx context.Context, args DisconnectNodeReq) (err error) {
	_, err = bc.sendRequest(ctx, "disconnectnode", args)
	return
}

// GetAddedNodeInfoReq holds the arguments for the GetAddedNodeInfo call.
//  1. node    (string, optional, default=all nodes) If provided, return information about this specific node, otherwise all nodes are returned.
type GetAddedNodeInfoReq struct {
	// If provided, return information about this specific node, otherwise all nodes are returned.
	// Default: all nodes
	Node string `json:"node,omitempty"`
}

// GetAddedNodeInfoResp holds the response to the GetAddedNodeInfo call.
//  [                                (json array)
//    {                              (json object)
//      "addednode" : "str",         (string) The node IP address or name (as provided to addnode)
//      "connected" : true|false,    (boolean) If connected
//      "addresses" : [              (json array) Only when connected = true
//        {                          (json object)
//          "address" : "str",       (string) The bitcoin server IP and port we're connected to
//          "connected" : "str"      (string) connection, inbound or outbound
//        },
//        ...
//      ]
//    },
//    ...
//  ]
type GetAddedNodeInfoResp struct {
	Array []GetAddedNodeInfoRespElement
}

func (alts GetAddedNodeInfoResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.Array)
}

func (alts *GetAddedNodeInfoResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Array) == nil {
		return nil
	}
	alts.Array = reset.Array
	return &UnmarshalError{B: b, structName: "GetAddedNodeInfoResp"}
}

type GetAddedNodeInfoRespElement struct {
	// The node IP address or name (as provided to addnode)
	AddedNode string `json:"addednode"`

	// If connected
	Connected bool `json:"connected"`

	// Only when connected = true
	Addresses []GetAddedNodeInfoRespElementAddresses `json:"addresses"`
}

type GetAddedNodeInfoRespElementAddresses struct {
	// The bitcoin server IP and port we're connected to
	Address string `json:"address"`

	// connection, inbound or outbound
	Connected string `json:"connected"`
}

// GetAddedNodeInfo RPC method.
// Returns information about the given added node, or all added nodes
// (note that onetry addnodes are not listed here)
func (bc *BitcoindClient) GetAddedNodeInfo(ctx context.Context, args GetAddedNodeInfoReq) (result GetAddedNodeInfoResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "getaddednodeinfo", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// GetConnectionCountResp holds the response to the GetConnectionCount call.
//  n    (numeric) The connection count
type GetConnectionCountResp struct {
	// The connection count
	N float64
}

func (alts GetConnectionCountResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.N)
}

func (alts *GetConnectionCountResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.N) == nil {
		return nil
	}
	alts.N = reset.N
	return &UnmarshalError{B: b, structName: "GetConnectionCountResp"}
}

// GetConnectionCount RPC method.
// Returns the number of connections to other nodes.
func (bc *BitcoindClient) GetConnectionCount(ctx context.Context) (result GetConnectionCountResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "getconnectioncount", nil); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// GetNetTotalsResp holds the response to the GetNetTotals call.
//  {                                              (json object)
//    "totalbytesrecv" : n,                        (numeric) Total bytes received
//    "totalbytessent" : n,                        (numeric) Total bytes sent
//    "timemillis" : xxx,                          (numeric) Current UNIX epoch time in milliseconds
//    "uploadtarget" : {                           (json object)
//      "timeframe" : n,                           (numeric) Length of the measuring timeframe in seconds
//      "target" : n,                              (numeric) Target in bytes
//      "target_reached" : true|false,             (boolean) True if target is reached
//      "serve_historical_blocks" : true|false,    (boolean) True if serving historical blocks
//      "bytes_left_in_cycle" : n,                 (numeric) Bytes left in current time cycle
//      "time_left_in_cycle" : n                   (numeric) Seconds left in current time cycle
//    }
//  }
type GetNetTotalsResp struct {
	// Total bytes received
	TotalBytesRecv float64 `json:"totalbytesrecv"`

	// Total bytes sent
	TotalBytesSent float64 `json:"totalbytessent"`

	// Current UNIX epoch time in milliseconds
	TimeMillis float64 `json:"timemillis"`

	UploadTarget struct {
		// Length of the measuring timeframe in seconds
		TimeFrame float64 `json:"timeframe"`

		// Target in bytes
		Target float64 `json:"target"`

		// True if target is reached
		TargetReached bool `json:"target_reached"`

		// True if serving historical blocks
		ServeHistoricalBlocks bool `json:"serve_historical_blocks"`

		// Bytes left in current time cycle
		BytesLeftInCycle float64 `json:"bytes_left_in_cycle"`

		// Seconds left in current time cycle
		TimeLeftInCycle float64 `json:"time_left_in_cycle"`
	} `json:"uploadtarget"`
}

// GetNetTotals RPC method.
// Returns information about network traffic, including bytes in, bytes out,
// and current time.
func (bc *BitcoindClient) GetNetTotals(ctx context.Context) (result GetNetTotalsResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "getnettotals", nil); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// GetNetworkInfoResp holds the response to the GetNetworkInfo call.
//  {                                                    (json object)
//    "version" : n,                                     (numeric) the server version
//    "subversion" : "str",                              (string) the server subversion string
//    "protocolversion" : n,                             (numeric) the protocol version
//    "localservices" : "hex",                           (string) the services we offer to the network
//    "localservicesnames" : [                           (json array) the services we offer to the network, in human-readable form
//      "str",                                           (string) the service name
//      ...
//    ],
//    "localrelay" : true|false,                         (boolean) true if transaction relay is requested from peers
//    "timeoffset" : n,                                  (numeric) the time offset
//    "connections" : n,                                 (numeric) the total number of connections
//    "connections_in" : n,                              (numeric) the number of inbound connections
//    "connections_out" : n,                             (numeric) the number of outbound connections
//    "networkactive" : true|false,                      (boolean) whether p2p networking is enabled
//    "networks" : [                                     (json array) information per network
//      {                                                (json object)
//        "name" : "str",                                (string) network (ipv4, ipv6, onion, i2p)
//        "limited" : true|false,                        (boolean) is the network limited using -onlynet?
//        "reachable" : true|false,                      (boolean) is the network reachable?
//        "proxy" : "str",                               (string) ("host:port") the proxy that is used for this network, or empty if none
//        "proxy_randomize_credentials" : true|false     (boolean) Whether randomized credentials are used
//      },
//      ...
//    ],
//    "relayfee" : n,                                    (numeric) minimum relay fee rate for transactions in BTC/kvB
//    "incrementalfee" : n,                              (numeric) minimum fee rate increment for mempool limiting or BIP 125 replacement in BTC/kvB
//    "localaddresses" : [                               (json array) list of local addresses
//      {                                                (json object)
//        "address" : "str",                             (string) network address
//        "port" : n,                                    (numeric) network port
//        "score" : n                                    (numeric) relative score
//      },
//      ...
//    ],
//    "warnings" : "str"                                 (string) any network and blockchain warnings
//  }
type GetNetworkInfoResp struct {
	// the server version
	Version float64 `json:"version"`

	// the server subversion string
	SubVersion string `json:"subversion"`

	// the protocol version
	ProtocolVersion float64 `json:"protocolversion"`

	// the services we offer to the network
	LocalServices string `json:"localservices"`

	// the services we offer to the network, in human-readable form
	// Element: Str    the service name
	LocalServicesNames []string `json:"localservicesnames"`

	// true if transaction relay is requested from peers
	LocalRelay bool `json:"localrelay"`

	// the time offset
	TimeOffset float64 `json:"timeoffset"`

	// the total number of connections
	Connections float64 `json:"connections"`

	// the number of inbound connections
	ConnectionsIn float64 `json:"connections_in"`

	// the number of outbound connections
	ConnectionsOut float64 `json:"connections_out"`

	// whether p2p networking is enabled
	NetworkActive bool `json:"networkactive"`

	// information per network
	Networks []GetNetworkInfoRespNetworks `json:"networks"`

	// minimum relay fee rate for transactions in BTC/kvB
	RelayFee float64 `json:"relayfee"`

	// minimum fee rate increment for mempool limiting or BIP 125 replacement in BTC/kvB
	IncrementalFee float64 `json:"incrementalfee"`

	// list of local addresses
	LocalAddresses []GetNetworkInfoRespLocalAddresses `json:"localaddresses"`

	// any network and blockchain warnings
	Warnings string `json:"warnings"`
}

type GetNetworkInfoRespNetworks struct {
	// network (ipv4, ipv6, onion, i2p)
	Name string `json:"name"`

	// is the network limited using -onlynet?
	Limited bool `json:"limited"`

	// is the network reachable?
	Reachable bool `json:"reachable"`

	// ("host:port") the proxy that is used for this network, or empty if none
	Proxy string `json:"proxy"`

	// Whether randomized credentials are used
	ProxyRandomizeCredentials bool `json:"proxy_randomize_credentials"`
}

type GetNetworkInfoRespLocalAddresses struct {
	// network address
	Address string `json:"address"`

	// network port
	Port float64 `json:"port"`

	// relative score
	Score float64 `json:"score"`
}

// GetNetworkInfo RPC method.
// Returns an object containing various state info regarding P2P networking.
func (bc *BitcoindClient) GetNetworkInfo(ctx context.Context) (result GetNetworkInfoResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "getnetworkinfo", nil); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// GetNodeAddressesReq holds the arguments for the GetNodeAddresses call.
//  1. count      (numeric, optional, default=1) The maximum number of addresses to return. Specify 0 to return all known addresses.
//  2. network    (string, optional, default=all networks) Return only addresses of the specified network. Can be one of: ipv4, ipv6, onion, i2p.
type GetNodeAddressesReq struct {
	// The maximum number of addresses to return. Specify 0 to return all known addresses.
	// Default: 1
	Count *float64 `json:"count,omitempty"`

	// Return only addresses of the specified network. Can be one of: ipv4, ipv6, onion, i2p.
	// Default: all networks
	Network string `json:"network,omitempty"`
}

// GetNodeAddressesResp holds the response to the GetNodeAddresses call.
//  [                         (json array)
//    {                       (json object)
//      "time" : xxx,         (numeric) The UNIX epoch time when the node was last seen
//      "services" : n,       (numeric) The services offered by the node
//      "address" : "str",    (string) The address of the node
//      "port" : n,           (numeric) The port number of the node
//      "network" : "str"     (string) The network (ipv4, ipv6, onion, i2p) the node connected through
//    },
//    ...
//  ]
type GetNodeAddressesResp struct {
	Array []GetNodeAddressesRespElement
}

func (alts GetNodeAddressesResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.Array)
}

func (alts *GetNodeAddressesResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Array) == nil {
		return nil
	}
	alts.Array = reset.Array
	return &UnmarshalError{B: b, structName: "GetNodeAddressesResp"}
}

type GetNodeAddressesRespElement struct {
	// The UNIX epoch time when the node was last seen
	Time float64 `json:"time"`

	// The services offered by the node
	Services float64 `json:"services"`

	// The address of the node
	Address string `json:"address"`

	// The port number of the node
	Port float64 `json:"port"`

	// The network (ipv4, ipv6, onion, i2p) the node connected through
	Network string `json:"network"`
}

// GetNodeAddresses RPC method.
// Return known addresses, which can potentially be used to find new nodes in the network.
func (bc *BitcoindClient) GetNodeAddresses(ctx context.Context, args GetNodeAddressesReq) (result GetNodeAddressesResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "getnodeaddresses", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// GetPeerInfoResp holds the response to the GetPeerInfo call.
//  [                                     (json array)
//    {                                   (json object)
//      "id" : n,                         (numeric) Peer index
//      "addr" : "str",                   (string) (host:port) The IP address and port of the peer
//      "addrbind" : "str",               (string) (ip:port) Bind address of the connection to the peer
//      "addrlocal" : "str",              (string) (ip:port) Local address as reported by the peer
//      "network" : "str",                (string) Network (ipv4, ipv6, onion, i2p, not_publicly_routable)
//      "mapped_as" : n,                  (numeric) The AS in the BGP route to the peer used for diversifying
//                                        peer selection (only available if the asmap config flag is set)
//      "services" : "hex",               (string) The services offered
//      "servicesnames" : [               (json array) the services offered, in human-readable form
//        "str",                          (string) the service name if it is recognised
//        ...
//      ],
//      "relaytxes" : true|false,         (boolean) Whether peer has asked us to relay transactions to it
//      "lastsend" : xxx,                 (numeric) The UNIX epoch time of the last send
//      "lastrecv" : xxx,                 (numeric) The UNIX epoch time of the last receive
//      "last_transaction" : xxx,         (numeric) The UNIX epoch time of the last valid transaction received from this peer
//      "last_block" : xxx,               (numeric) The UNIX epoch time of the last block received from this peer
//      "bytessent" : n,                  (numeric) The total bytes sent
//      "bytesrecv" : n,                  (numeric) The total bytes received
//      "conntime" : xxx,                 (numeric) The UNIX epoch time of the connection
//      "timeoffset" : n,                 (numeric) The time offset in seconds
//      "pingtime" : n,                   (numeric) ping time (if available)
//      "minping" : n,                    (numeric) minimum observed ping time (if any at all)
//      "pingwait" : n,                   (numeric) ping wait (if non-zero)
//      "version" : n,                    (numeric) The peer version, such as 70001
//      "subver" : "str",                 (string) The string version
//      "inbound" : true|false,           (boolean) Inbound (true) or Outbound (false)
//      "bip152_hb_to" : true|false,      (boolean) Whether we selected peer as (compact blocks) high-bandwidth peer
//      "bip152_hb_from" : true|false,    (boolean) Whether peer selected us as (compact blocks) high-bandwidth peer
//      "startingheight" : n,             (numeric) The starting height (block) of the peer
//      "synced_headers" : n,             (numeric) The last header we have in common with this peer
//      "synced_blocks" : n,              (numeric) The last block we have in common with this peer
//      "inflight" : [                    (json array)
//        n,                              (numeric) The heights of blocks we're currently asking from this peer
//        ...
//      ],
//      "permissions" : [                 (json array) Any special permissions that have been granted to this peer
//        "str",                          (string) bloomfilter (allow requesting BIP37 filtered blocks and transactions),
//                                        noban (do not ban for misbehavior; implies download),
//                                        forcerelay (relay transactions that are already in the mempool; implies relay),
//                                        relay (relay even in -blocksonly mode, and unlimited transaction announcements),
//                                        mempool (allow requesting BIP35 mempool contents),
//                                        download (allow getheaders during IBD, no disconnect after maxuploadtarget limit),
//                                        addr (responses to GETADDR avoid hitting the cache and contain random records with the most up-to-date info).
type GetPeerInfoResp struct {
	Array []GetPeerInfoRespElement
}

func (alts GetPeerInfoResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.Array)
}

func (alts *GetPeerInfoResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Array) == nil {
		return nil
	}
	alts.Array = reset.Array
	return &UnmarshalError{B: b, structName: "GetPeerInfoResp"}
}

type GetPeerInfoRespElement struct {
	// Peer index
	ID float64 `json:"id"`

	// (host:port) The IP address and port of the peer
	Addr string `json:"addr"`

	// (ip:port) Bind address of the connection to the peer
	AddrBind string `json:"addrbind"`

	// (ip:port) Local address as reported by the peer
	AddrLocal string `json:"addrlocal"`

	// Network (ipv4, ipv6, onion, i2p, not_publicly_routable)
	Network string `json:"network"`

	// The AS in the BGP route to the peer used for diversifying
	// peer selection (only available if the asmap config flag is set)
	MappedAs float64 `json:"mapped_as"`

	// The services offered
	Services string `json:"services"`

	// the services offered, in human-readable form
	// Element: Str    the service name if it is recognised
	ServicesNames []string `json:"servicesnames"`

	// Whether peer has asked us to relay transactions to it
	RelayTxes bool `json:"relaytxes"`

	// The UNIX epoch time of the last send
	LastSend float64 `json:"lastsend"`

	// The UNIX epoch time of the last receive
	LastRecv float64 `json:"lastrecv"`

	// The UNIX epoch time of the last valid transaction received from this peer
	LastTransaction float64 `json:"last_transaction"`

	// The UNIX epoch time of the last block received from this peer
	LastBlock float64 `json:"last_block"`

	// The total bytes sent
	BytesSent float64 `json:"bytessent"`

	// The total bytes received
	BytesRecv float64 `json:"bytesrecv"`

	// The UNIX epoch time of the connection
	ConnTime float64 `json:"conntime"`

	// The time offset in seconds
	TimeOffset float64 `json:"timeoffset"`

	// ping time (if available)
	PingTime float64 `json:"pingtime"`

	// minimum observed ping time (if any at all)
	MinPing float64 `json:"minping"`

	// ping wait (if non-zero)
	PingWait float64 `json:"pingwait"`

	// The peer version, such as 70001
	Version float64 `json:"version"`

	// The string version
	SubVer string `json:"subver"`

	// Inbound (true) or Outbound (false)
	Inbound bool `json:"inbound"`

	// Whether we selected peer as (compact blocks) high-bandwidth peer
	BIP152HbTo bool `json:"bip152_hb_to"`

	// Whether peer selected us as (compact blocks) high-bandwidth peer
	BIP152HbFrom bool `json:"bip152_hb_from"`

	// The starting height (block) of the peer
	StartingHeight float64 `json:"startingheight"`

	// The last header we have in common with this peer
	SyncedHeaders float64 `json:"synced_headers"`

	// The last block we have in common with this peer
	SyncedBlocks float64 `json:"synced_blocks"`

	// Element: N    The heights of blocks we're currently asking from this peer
	InFlight []float64 `json:"inflight"`

	// Any special permissions that have been granted to this peer
	// Element: Str    bloomfilter (allow requesting BIP37 filtered blocks and transactions),
	// noban (do not ban for misbehavior; implies download),
	// forcerelay (relay transactions that are already in the mempool; implies relay),
	// relay (relay even in -blocksonly mode, and unlimited transaction announcements),
	// mempool (allow requesting BIP35 mempool contents),
	// download (allow getheaders during IBD, no disconnect after maxuploadtarget limit),
	// addr (responses to GETADDR avoid hitting the cache and contain random records with the most up-to-date info).
	Permissions []string `json:"permissions"`
}

// GetPeerInfo RPC method.
// Returns data about each connected network node as a json array of objects.
func (bc *BitcoindClient) GetPeerInfo(ctx context.Context) (result GetPeerInfoResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "getpeerinfo", nil); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// ListBannedResp holds the response to the ListBanned call.
//  [                              (json array)
//    {                            (json object)
//      "address" : "str",         (string) The IP/Subnet of the banned node
//      "ban_created" : xxx,       (numeric) The UNIX epoch time the ban was created
//      "banned_until" : xxx,      (numeric) The UNIX epoch time the ban expires
//      "ban_duration" : xxx,      (numeric) The ban duration, in seconds
//      "time_remaining" : xxx     (numeric) The time remaining until the ban expires, in seconds
//    },
//    ...
//  ]
type ListBannedResp struct {
	Array []ListBannedRespElement
}

func (alts ListBannedResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.Array)
}

func (alts *ListBannedResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.Array) == nil {
		return nil
	}
	alts.Array = reset.Array
	return &UnmarshalError{B: b, structName: "ListBannedResp"}
}

type ListBannedRespElement struct {
	// The IP/Subnet of the banned node
	Address string `json:"address"`

	// The UNIX epoch time the ban was created
	BanCreated float64 `json:"ban_created"`

	// The UNIX epoch time the ban expires
	BannedUntil float64 `json:"banned_until"`

	// The ban duration, in seconds
	BanDuration float64 `json:"ban_duration"`

	// The time remaining until the ban expires, in seconds
	TimeRemaining float64 `json:"time_remaining"`
}

// ListBanned RPC method.
// List all manually banned IPs/Subnets.
func (bc *BitcoindClient) ListBanned(ctx context.Context) (result ListBannedResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "listbanned", nil); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}

// Ping RPC method.
// Requests that a ping be sent to all other nodes, to measure ping time.
// Results provided in getpeerinfo, pingtime and pingwait fields are decimal seconds.
// Ping command is handled in queue with all other commands, so it measures processing backlog, not just network ping.
func (bc *BitcoindClient) Ping(ctx context.Context) (err error) {
	_, err = bc.sendRequest(ctx, "ping", nil)
	return
}

// SetBanReq holds the arguments for the SetBan call.
//  1. subnet      (string, required) The IP/Subnet (see getpeerinfo for nodes IP) with an optional netmask (default is /32 = single IP)
//  2. command     (string, required) 'add' to add an IP/Subnet to the list, 'remove' to remove an IP/Subnet from the list
//  3. bantime     (numeric, optional, default=0) time in seconds how long (or until when if [absolute] is set) the IP is banned (0 or empty means using the default time of 24h which can also be overwritten by the -bantime startup argument)
//  4. absolute    (boolean, optional, default=false) If set, the bantime must be an absolute timestamp expressed in UNIX epoch time
type SetBanReq struct {
	// The IP/Subnet (see getpeerinfo for nodes IP) with an optional netmask (default is /32 = single IP)
	SubNet string `json:"subnet"`

	// 'add' to add an IP/Subnet to the list, 'remove' to remove an IP/Subnet from the list
	Command string `json:"command"`

	// time in seconds how long (or until when if [absolute] is set) the IP is banned (0 or empty means using the default time of 24h which can also be overwritten by the -bantime startup argument)
	// Default: 0
	BanTime float64 `json:"bantime,omitempty"`

	// If set, the bantime must be an absolute timestamp expressed in UNIX epoch time
	// Default: false
	Absolute bool `json:"absolute,omitempty"`
}

// SetBan RPC method.
// Attempts to add or remove an IP/Subnet from the banned list.
func (bc *BitcoindClient) SetBan(ctx context.Context, args SetBanReq) (err error) {
	_, err = bc.sendRequest(ctx, "setban", args)
	return
}

// SetNetworkActiveReq holds the arguments for the SetNetworkActive call.
//  1. state    (boolean, required) true to enable networking, false to disable
type SetNetworkActiveReq struct {
	// true to enable networking, false to disable
	State bool `json:"state"`
}

// SetNetworkActiveResp holds the response to the SetNetworkActive call.
//  true|false    (boolean) The value that was passed in
type SetNetworkActiveResp struct {
	// The value that was passed in
	TrueOrFalse bool
}

func (alts SetNetworkActiveResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(alts.TrueOrFalse)
}

func (alts *SetNetworkActiveResp) UnmarshalJSON(b []byte) error {
	reset := *alts
	var decoder *json.Decoder
	decoder = json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&alts.TrueOrFalse) == nil {
		return nil
	}
	alts.TrueOrFalse = reset.TrueOrFalse
	return &UnmarshalError{B: b, structName: "SetNetworkActiveResp"}
}

// SetNetworkActive RPC method.
// Disable/enable all p2p network activity.
func (bc *BitcoindClient) SetNetworkActive(ctx context.Context, args SetNetworkActiveReq) (result SetNetworkActiveResp, err error) {
	var resultRaw json.RawMessage
	if resultRaw, err = bc.sendRequest(ctx, "setnetworkactive", args); err != nil {
		return
	}
	err = json.Unmarshal(resultRaw, &result)
	return
}
