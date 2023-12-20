package main

import (
	dictionary "myproject/Dictionary"
)

func main() {
	filePath := "path/to/your/dictionary/file.txt"

	myDictionary := dictionary.NewDictionary(filePath)

	go myDictionary.Add("apple", "a fruit")
	go myDictionary.Add("golang", "a programming language")

	go myDictionary.Remove("car")
	go myDictionary.Remove("apple")

	// var input string
	// fmt.Scanln(&input)
}
