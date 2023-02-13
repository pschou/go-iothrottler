package iothrottler_test

import (
	"fmt"
	"time"

	"github.com/pschou/go-iothrottler"
)

func ExampleNewIOThrottler() {
	throttler := iothrottler.NewIOThrottler(100 << 20) // simulate a 100Mbps
	tick, now := time.Now(), time.Time{}
	throttler.SendN(50 << 20) // Send 100Mb

	now = time.Now()
	fmt.Println("0s Tock:", now.Sub(tick).Round(time.Second/10))
	tick = now

	throttler.SendN(200 << 20) // Send 200Mb

	now = time.Now()
	fmt.Println("0.5s Tock:", now.Sub(tick).Round(time.Second/10))
	tick = now

	throttler.SendN(10 << 20) // Send 10Mb

	now = time.Now()
	fmt.Println("2s Tock:", now.Sub(tick).Round(time.Second/10))
	tick = now

	throttler.SendN(0) // Must send again to hit delay

	now = time.Now()
	fmt.Println("0.1s Tock:", now.Sub(tick).Round(time.Second/10))
	// Output:
	// 0s Tock: 0s
	// 0.5s Tock: 500ms
	// 2s Tock: 2s
	// 0.1s Tock: 100ms
}
