package Dictionary

type Dictionary map[string]string

func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}

func (d Dictionary) Get(word string) string {
	return d[word]
}

func (d Dictionary) Remove(word string) {
	delete(d, word)
}

func (d Dictionary) List() map[string]string {
	return d
}
