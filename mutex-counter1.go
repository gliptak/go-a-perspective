//
// Inspired by https://github.com/golang/tour/blob/master/content/concurrency/mutex-counter.go
//

package main

import (
	"fmt"
	"sync"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mux sync.Mutex
	v   map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	defer wg.Done()
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	return c.v[key]
}

var wg sync.WaitGroup

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go c.Inc("somekey")
	}
	wg.Wait()
	fmt.Println(c.Value("somekey"))
}