package main

import (
	"fmt"
	"myproject/Dictionary"
)

func main() {

	filePath := "C:/Users/MHD/OneDrive/Desktop/goo/Dictionnaires.txt"

	myDictionary := Dictionary.NewDictionary(filePath)

	// Ajouter des mots et des définitions
	myDictionary.Add("apple", "a fruit")
	myDictionary.Add("golang", "a programming language")
	myDictionary.Add("car", "a vehicle")

	// Obtenir une définition
	definition, _ := myDictionary.Get("apple")
	fmt.Println("Definition of 'apple':", definition)

	// Supprimer un mot
	myDictionary.Remove("car")

	// Liste de tous les mots
	words, _ := myDictionary.List()
	fmt.Println("All words in the dictionary:", words)
}
