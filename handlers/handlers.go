package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/x-ider/aviasales-test-task/anagrams"
	"io/ioutil"
	"net/http"
)

type AnagramDictionaryHandler struct {
	dictionary *anagrams.Dictionary
}

func NewAnagramDictionaryHandler(dictionary *anagrams.Dictionary) *AnagramDictionaryHandler {
	var adh = AnagramDictionaryHandler{dictionary}
	return &adh
}

func (adh *AnagramDictionaryHandler) HandleGetRequest(w http.ResponseWriter, r *http.Request) {
	wordParameters := r.URL.Query()["word"]
	var word string
	if len(wordParameters) > 0 {
		word = wordParameters[0]
	}
	searchResult := adh.dictionary.SearchAnagramsByWord(word)
	jsonString, _ := json.Marshal(searchResult)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonString)
}

func (adh *AnagramDictionaryHandler) HandleLoadRequest(w http.ResponseWriter, r *http.Request) {
	inputWords, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("An error occurred during reading request body! The dictionary hasn't been updated.")
		return
	}
	var words []string
	if err := json.Unmarshal(inputWords, &words); err != nil {
		fmt.Println("An error occurred during unmarshal process! The dictionary hasn't been updated.")
		return
	}
	adh.dictionary.PopulateWithNewWords(words)
	w.WriteHeader(http.StatusOK)
}
