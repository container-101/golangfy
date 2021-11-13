package main

import "fmt"

type person struct {
	name string
	age int
	favFood []string
}

//Entry Point
//fmt=formatting
//auto type inference works inside func(:=)
func main() {
	favFood := []string{"ramen"}
	woodi :=person{
		name:"woodi",age:18,favFood: favFood,
	}
	fmt.Println(woodi.name)
}