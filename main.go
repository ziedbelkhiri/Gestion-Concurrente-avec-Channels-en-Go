package main

import (
	"fmt"

	"./Dictionary"
)

func main() {
	myDictionary := Dictionary.Dictionary{}

	myDictionary.Add("word1", "definition1")
	myDictionary.Add("word2", "definition2")

	definition := myDictionary.Get("word1")
	fmt.Println("Definition:", definition)

	myDictionary.Remove("word2")

	wordList := myDictionary.List()
	for word, definition := range wordList {
		fmt.Printf("%s: %s\n", word, definition)
	}
}
