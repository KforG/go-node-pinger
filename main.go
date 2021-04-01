package main

import (
	"fmt"
	"time"

	"github.com/go-ping/ping"
)

var nodes = [4]string{"fr1.vtconline.org", "p2proxy.vertcoin.org", "p2p-usa.xyz", "p2p-ekb.xyz"}
var results = [len(nodes)]time.Duration{}

func main() {
	pingNode()
	closestNode()
}

func pingNode() {
	for i := 0; i < len(nodes); i++ {
		pinger, err := ping.NewPinger(nodes[i])
		pinger.SetPrivileged(true) //This line is needed for windows because of ICMP
		if err != nil {
			panic(err)
		}
		pinger.Count = 3
		err = pinger.Run()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s: %v \n", nodes[i], pinger.Statistics().AvgRtt)
		results[i] = pinger.Statistics().AvgRtt
	}
}

func closestNode() string {
	var bestNode string
	lowest := results[0]

	for b := 0; b < len(results); b++ {
		if results[b] <= lowest {
			bestNode = nodes[b]
			lowest = results[b]
		}
	}
	fmt.Printf("Selected node: %s ", bestNode)
	return bestNode
}
