//
// Inspired by https://gobyexample.com/timeouts
//
package main

import (
	"time"
	"fmt"
	"math/rand"
)

func sleeper(i int, c chan int) {
	value := rand.Intn(20 - 1) + 1
	fmt.Printf("%d sleeping for %d\n", i, value)
	time.Sleep(time.Duration(time.Duration(value) * time.Second))
	c <- i
}

func main() {
	rand.Seed(42)
	const count = 5
	c := make(chan int)
	for i:=0; i<count; i++ {
		go sleeper(i, c)
	}
	timeout := time.After(10 * time.Second)
	for i:=0; i<count; i++ {
		select {
		case id := <- c:
			fmt.Printf("Completed %d\n", id)
		case <- timeout:
			fmt.Println("Done processing")
		}
	}
}