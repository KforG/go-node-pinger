package main

import (
	"fmt"
	"time"

	"github.com/go-ping/ping"
)

var nodes = [4]string{"fr1.vtconline.org", "p2proxy.vertcoin.org", "p2p-usa.xyz", "p2p-ekb.xyz"}
var results = [len(nodes)]time.Duration{}

func main() {
	closestNode()
}

func pingNode() {
	for i := 0; i < len(nodes); i++ {
		pinger, err := ping.NewPinger(nodes[i])
		pinger.SetPrivileged(true)  //This line is needed for windows because of ICMP
		pinger.Timeout = 5000000000 //If response time is longer than 5 seconds, the pinger will exit regardless of how many packets have been recieved
		if err != nil {
			fmt.Println("Error: Check if you are connected to the internet")
			panic(err)
		}
		pinger.Count = 3
		err = pinger.Run()
		if err != nil {
			fmt.Println("Error: Check if you are connected to the internet")
			panic(err)
		}
		results[i] = pinger.Statistics().AvgRtt
		if results[i] == 0 {
			results[i] = 5000000000
		}
		fmt.Printf("%s: %v \n", nodes[i], results[i])
	}
}

func closestNode() string {
	pingNode()

	bestNode := nodes[0]
	lowest := results[0]

	for i := 1; i < len(results); i++ {
		if results[i] <= lowest {
			bestNode = nodes[i]
			lowest = results[i]
		}
	}
	if bestNode == "p2proxy.vertcoin.org" { //Currently this node uses port 9172 instead of 9171, if this changes this statement can be removed and 9171 can be added to all nodes
		bestNode += ":9172"
	} else {
		bestNode += ":9171"
	}

	fmt.Printf("Selected node: %s\n", bestNode)
	return bestNode
}
