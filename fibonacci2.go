//
// https://tour.golang.org/concurrency/4
//
package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	n := 10
	c := make(chan int)
	go fibonacci(n, c)
	for i := range c {
		fmt.Println(i)
	}
}
