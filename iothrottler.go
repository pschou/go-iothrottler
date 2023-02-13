// Copyright 2021 github.com/pschou/go-iothrottler
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package iothrottler

import (
	"time"
)

type IOThrottler struct {
	c chan int64
}

func (t IOThrottler) SendN(n int64) {
	t.c <- n
}

// Create a new IOThrottler with specified bandwith limitation
func NewIOThrottler(bandwidth int64) (t *IOThrottler) {
	sec := int64(time.Second)
	t = &IOThrottler{c: make(chan int64)}
	go func() {
		n := <-t.c // Wait for first hit
		time.Sleep(time.Duration(n * sec / bandwidth))
		for {
			n = <-t.c // Get next request
			time.Sleep(time.Duration(n * sec / bandwidth))
		}
	}()
	return
}
