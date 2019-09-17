package handlers

import (
	"bytes"
	"github.com/x-ider/aviasales-test-task/anagrams"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

func TestAnagramDictionaryHandler_HandleGetRequest(t *testing.T) {
	dictionary := anagrams.NewEmptyDictionary()
	dictionary.PopulateWithNewWords([]string{"foObar", "AaBb", "baba", "booFAR", "test", "Art", "rat", "TAR"})
	agh := NewAnagramDictionaryHandler(dictionary)

	cases := []struct {
		in, want string
	}{
		{"Abba", `["aabb","baba"]`},
		{"rabOOf", `["foobar","boofar"]`},
		{"unknown", "null"},
		{"test", `["test"]`},
		{"RTA", `["art","rat","tar"]`},
		{"123", "null"},
		{" ", "null"},
		{"", "null"},
	}

	for i := range cases {
		request := httptest.NewRequest("GET", "/get?word="+url.QueryEscape(cases[i].in), nil)
		response := httptest.NewRecorder()

		handler := http.HandlerFunc(agh.HandleGetRequest)
		handler.ServeHTTP(response, request)

		if status := response.Code; status != http.StatusOK {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		got, err := ioutil.ReadAll(response.Body)
		if err != nil {
			t.Error("An error occurred during reading response body!")
			return
		}
		if string(got) != cases[i].want {
			t.Errorf("\ninput: %q, type %T\ngot %q,type %T\nwant %q, type %T", cases[i].in, cases[i].in,
				got, got, cases[i].want, cases[i].want)
		}
	}

}

func TestAnagramDictionaryHandler_HandleLoadRequest(t *testing.T) {
	dictionary := anagrams.NewEmptyDictionary()
	adh := NewAnagramDictionaryHandler(dictionary)

	emptyDictionary := anagrams.NewEmptyDictionary()

	dictionaryAABB := anagrams.NewEmptyDictionary()
	dictionaryAABB.PopulateWithNewWords([]string{"Aabb", "baba"})

	dictionaryOneWord := anagrams.NewEmptyDictionary()
	dictionaryOneWord.PopulateWithNewWords([]string{"oneWord"})

	cases := []struct {
		in   string
		want *anagrams.Dictionary
	}{
		{`["Aabb", "baba"]`, dictionaryAABB},
		{`["oneWord"]`, dictionaryOneWord},
		{`[lostQuote"]`, emptyDictionary},
		{" ", emptyDictionary},
		{"", emptyDictionary},
	}

	for i := range cases {
		request := httptest.NewRequest("POST", "/load", bytes.NewReader([]byte(cases[i].in)))
		response := httptest.NewRecorder()

		handler := http.HandlerFunc(adh.HandleLoadRequest)
		handler.ServeHTTP(response, request)

		if status := response.Code; status != http.StatusOK {
			t.Errorf("Handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		if !reflect.DeepEqual(adh.dictionary, cases[i].want) {
			t.Errorf("\ninput: %q, type %T\ngot %q,type %T\nwant %q, type %T", cases[i].in, cases[i].in,
				adh.dictionary, adh.dictionary, cases[i].want, cases[i].want)
		}
		adh.dictionary.Clear()
	}
}
