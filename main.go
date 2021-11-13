package main

import (
	"fmt"
	"time"
)

//go is only available at main function running
func main() {
	go sexyCount("woodi")
	sexyCount("mircat")
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
