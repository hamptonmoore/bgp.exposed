package main

import (
	"log"

	"github.com/hamptonmoore/bgp.exposed/backend/bgp"
	"github.com/hamptonmoore/bgp.exposed/backend/common"
	gobgp "github.com/osrg/gobgp/v3/api"
)

func main() {

	server, err := bgp.CreateGoBGPServer("127.0.0.1", "50051", 2000, "1.1.1.1", 1001)
	if err != nil {
		log.Fatal(err)
	}

	create := common.CreateRequest{
		LocalASN: 1001,
		PeerASN:  1002,
		PeerIP:   "127.0.0.1",
	}
	down := bgp.CreateDownstream(create, server)

	peerUpdate := make(chan *gobgp.Peer, 12)
	go down.SubscribeToPeer(peerUpdate)
	if err != nil {
		log.Fatal("Subscribe failed", err)
	}
	log.Println(<-peerUpdate)
}
