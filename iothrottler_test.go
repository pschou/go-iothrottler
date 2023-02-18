package iothrottler_test

import (
	"fmt"
	"time"

	"github.com/pschou/go-iothrottler"
)

func ExampleNewIOThrottler() {
	fmt.Println("making new throttler")
	throttler := iothrottler.NewIOThrottler(10000000, 1500, 12)
	for j := 0; j < 5; j++ {
		fmt.Println("second", j)
		tick := time.Now()

		// Send 100 packets
		for i := 0; i < 812; i++ {
			<-throttler.C
		}

		fmt.Println("packets:", time.Now().Sub(tick))
	}
	// Output:
	// 0s Tock: 0s
	// 0.5s Tock: 500ms
	// 2s Tock: 2s
	// 0.1s Tock: 100ms
}
