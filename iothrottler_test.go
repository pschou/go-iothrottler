package iothrottler_test

import (
	"fmt"
	"time"

	"github.com/pschou/go-iothrottler"
)

func ExampleNewLimit() {
	fmt.Println("making new throttler")
	throttler := iothrottler.NewLimit(10000000, 1500, 26+12)

	tick := time.Now()

	// Send 100 packets
	<-throttler.C
	for i := 0; i < 812; i++ {
		<-throttler.C
	}

	fmt.Println("812 packets in time:", time.Now().Sub(tick).Round(time.Second/100))
	// Output:
	// making new throttler
	// 812 packets in time: 1s
}
