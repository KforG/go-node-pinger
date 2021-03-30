package main

import (
	"fmt"
	"time"

	"github.com/go-ping/ping"
)

func main() {
	nodes := [4]string{"fr1.vtconline.org", "p2proxy.vertcoin.org", "p2p-usa.xyz", "p2p-ekb.xyz"}
	results := [len(nodes)]time.Duration{}

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

		results[i] = pinger.Statistics().AvgRtt
	}
	fmt.Println(results)
}
