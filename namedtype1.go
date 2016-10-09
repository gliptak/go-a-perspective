package main

import "fmt"

type Name string
type City string

func hellofrom(name Name, city City) {
	fmt.Printf("Hello %s from %s!", name, city)
}

func main() {
	const name = Name("James")
	const city = City("New York")
	hellofrom(name, city)
	//hellofrom(city, name) this call doesn't compile
}