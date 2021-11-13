package main

import (
	"fmt"
	"golwee/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	err := dictionary.Delete(baseWord)
	if err != nil {
		fmt.Println(err)
	}
	word, err2 := dictionary.Search(baseWord)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(word)
}
