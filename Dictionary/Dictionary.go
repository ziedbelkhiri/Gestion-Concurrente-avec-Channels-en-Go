package Dictionary

// dictionary.go

import (
	"bufio"
	"os"
	"strings"
)

type Dictionary struct {
	filePath string
}

func NewDictionary(filePath string) *Dictionary {
	return &Dictionary{filePath: filePath}
}

func (d *Dictionary) Add(word, definition string) error {
	file, err := os.OpenFile(d.filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(word + ":" + definition + "\n")
	if err != nil {
		return err
	}

	return nil
}

func (d *Dictionary) Get(word string) (string, error) {
	file, err := os.Open(d.filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ":")
		if len(parts) == 2 && parts[0] == word {
			return parts[1], nil
		}
	}

	return "", nil // Word not found
}

func (d *Dictionary) Remove(word string) error {
	lines := []string{}

	file, err := os.Open(d.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ":")
		if len(parts) != 2 || parts[0] != word {
			lines = append(lines, scanner.Text())
		}
	}

	newFile, err := os.Create(d.filePath)
	if err != nil {
		return err
	}
	defer newFile.Close()

	writer := bufio.NewWriter(newFile)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	writer.Flush()

	return nil
}

func (d *Dictionary) List() (map[string]string, error) {
	entries := make(map[string]string)

	file, err := os.Open(d.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ":")
		if len(parts) == 2 {
			entries[parts[0]] = parts[1]
		}
	}

	return entries, nil
}
