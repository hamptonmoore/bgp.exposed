package common

import "strconv"

type Packet struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"` // Either CreateRequest, UpdateRequest, RouteUpdate
}

type CreateRequest struct {
	PeerASN  uint32 `json:"peerASN"`
	PeerIP   string `json:"peerIP"`
	LocalASN uint32 `json:"localASN"`
}

func (c *CreateRequest) ToKey() string {
	return c.PeerIP + "|" + strconv.FormatUint(uint64(c.PeerASN), 10)
}

type UpdateRequest struct {
	FullTable   bool   `json:"fullTable"`
	AddPath     bool   `json:"addPath"`
	MD5Password string `json:"md5Password"`
}

type RouteData struct {
	Withdraws      []NLRI     `json:"withdraws"`
	Prefixes       []NLRI     `json:"prefixes"`
	AsPath         []uint32   `json:"asPath"`
	NextHop        string     `json:"nextHop"`
	Communities    [][]uint16 `json:"communities"`
	Origin         int        `json:"origin"`
	ExtCommunities [][]uint32 `json:"extCommunities"`
}

type NLRI struct {
	Prefix string `json:"prefix"`
	ID     uint32 `json:"id"`
}

type FSMUpdate struct {
	State          string `json:"state"`
	HoldTimer      uint   `json:"holdTimer"`
	KeepaliveTimer uint   `json:"keepaliveTimer"`
	SentKeepAlive  bool   `json:"sentKeepAlive"`
	LastUpdate     uint   `json:"lastUpdate"`
	LastKeepalive  uint   `json:"lastKeepalive"`
}
