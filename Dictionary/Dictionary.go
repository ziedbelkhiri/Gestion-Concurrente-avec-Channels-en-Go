// dictionary.go
package dictionary

import (
	"bufio"
	"os"
	"strings"
)

type Dictionary struct {
	filePath   string
	addChan    chan entry
	removeChan chan string
}

type entry struct {
	word       string
	definition string
}

func NewDictionary(filePath string) *Dictionary {
	dict := &Dictionary{
		filePath:   filePath,
		addChan:    make(chan entry),
		removeChan: make(chan string),
	}

	go dict.handleAdditions()
	go dict.handleRemovals()

	return dict
}

func (d *Dictionary) handleAdditions() {
	file, err := os.OpenFile(d.filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for {
		entry := <-d.addChan
		_, err := file.WriteString(entry.word + ":" + entry.definition + "\n")
		if err != nil {
			panic(err)
		}
	}
}

func (d *Dictionary) handleRemovals() {
	for {
		word := <-d.removeChan
		lines := []string{}

		file, err := os.Open(d.filePath)
		if err != nil {
			panic(err)
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			parts := strings.Split(scanner.Text(), ":")
			if len(parts) != 2 || parts[0] != word {
				lines = append(lines, scanner.Text())
			}
		}
		file.Close()

		newFile, err := os.Create(d.filePath)
		if err != nil {
			panic(err)
		}

		writer := bufio.NewWriter(newFile)
		for _, line := range lines {
			_, err := writer.WriteString(line + "\n")
			if err != nil {
				panic(err)
			}
		}
		writer.Flush()
		newFile.Close()
	}
}

func (d *Dictionary) Add(word, definition string) {
	d.addChan <- entry{word: word, definition: definition}
}

func (d *Dictionary) Remove(word string) {
	d.removeChan <- word
}
