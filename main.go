package main

import (
	"fmt"
	"github.com/x-ider/aviasales-test-task/anagrams"
	"github.com/x-ider/aviasales-test-task/handlers"
	"net/http"
)

func main() {
	dictionary := anagrams.NewEmptyDictionary()
	handler := handlers.NewAnagramDictionaryHandler(dictionary)
	http.HandleFunc("/get", handler.HandleGetRequest)
	http.HandleFunc("/load", handler.HandleLoadRequest)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("An error occurred during server running!")
	}
}
