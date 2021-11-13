package main

import (
	"fmt"
)

func multiply(a, b int) int{
	return a * b
}

func superAdd(numbers ...int) int {
	total :=0
	for _, num := range numbers{
		total += num
	}
	return total
}

func canIPlay(age int) bool {
	switch koreanAge:= age +2; koreanAge{
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func repeatInfinitely(words ...string) {
	fmt.Println(words)
}

//Entry Point
//fmt=formatting
//auto type inference works inside func(:=)
func main() {
	a := 2
	b := &a
	*b = 10
	fmt.Println(*b, a)
}