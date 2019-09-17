package anagrams

import (
	"sort"
	"strings"
)

type Dictionary struct {
	d map[string][]string
}

func NewEmptyDictionary() *Dictionary {
	var dictionary Dictionary
	dictionary.d = make(map[string][]string)
	return &dictionary
}

func (dictionary *Dictionary) GetNumberOfKeys() int {
	return len(dictionary.d)
}

func sortWordByChars(word string) string {
	s := strings.Split(word, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func (dictionary *Dictionary) SearchAnagramsByWord(word string) []string {
	word = strings.ToLower(word)
	word = sortWordByChars(word)
	return dictionary.d[word]
}

func (dictionary *Dictionary) PopulateWithNewWords(words []string) {
	if dictionary.GetNumberOfKeys() > 0 {
		dictionary.Clear()
	}
	for _, word := range words {
		word = strings.ToLower(word)
		key := sortWordByChars(word)
		dictionary.d[key] = append(dictionary.d[key], word)
	}
}

func (dictionary *Dictionary) Clear() {
	newDictionary := NewEmptyDictionary()
	*dictionary = *newDictionary
}
