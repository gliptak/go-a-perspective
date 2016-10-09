//
// A solution to https://tour.golang.org/moretypes/26
//
package main

import "fmt"

func fibonacci() func() int {
	current := 0
	next := 1
	return func() int {
		current, next = next, current + next
		return current
	}
}

func main() {
	f := fibonacci()
	for i:=0; i < 10; i++ {
		fmt.Println(f(),)
	}
}