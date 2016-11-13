//
// From http://stackoverflow.com/questions/6395076/in-golang-using-reflect-how-do-you-set-the-value-of-a-struct-field
//

package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	Number int
	Text string
}

func main() {
	foo := Foo{123, "Hello"}
	fmt.Println(foo)
	reflect.ValueOf(&foo).Elem().Field(0).SetInt(321)
	fmt.Println(foo)
}