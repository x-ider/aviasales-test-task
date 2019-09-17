package anagrams

import (
	"reflect"
	"testing"
)

func TestSortByChars(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"abba", "aabb"},
		{"foobar", "abfoor"},
		{"boofar", "abfoor"},
		{"roobaf", "abfoor"},
		{"камин", "аикмн"},
		{" ", " "},
		{"", ""},
	}
	for _, c := range cases {
		got := sortWordByChars(c.in)
		if got != c.want {
			t.Errorf("sortWordByChars(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestDictionary_Clear(t *testing.T) {
	cases := []Dictionary{
		{map[string][]string{"estt": {"test"}}},
		{map[string][]string{"art": {"art", "rat", "tar"}}},
		{map[string][]string{}},
	}
	for i := range cases {
		if cases[i].Clear(); cases[i].GetNumberOfKeys() != 0 {
			t.Errorf("Dictionary is not empty")
		}
	}
}

func TestDictionary_AddWords(t *testing.T) {
	got := Dictionary{map[string][]string{}}
	cases := []struct {
		in   []string
		want Dictionary
	}{
		{[]string{"art", "rat", "tar"}, Dictionary{map[string][]string{"art": {"art", "rat", "tar"}}}},
		{[]string{"test"}, Dictionary{map[string][]string{"estt": {"test"}}}},
		{[]string{}, Dictionary{map[string][]string{}}},
	}
	for _, c := range cases {
		got.PopulateWithNewWords(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("PopulateWithNewWords(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestDictionary_SearchAnagramsByWord(t *testing.T) {
	dictionary := Dictionary{map[string][]string{"art": {"art", "rat", "tar"}, "estt": {"test"}}}
	cases := []struct {
		in   string
		want []string
	}{
		{"RTA", []string{"art", "rat", "tar"}},
		{"test", []string{"test"}},
		{"pingas", nil},
		{" ", nil},
		{"", nil},
	}
	for _, c := range cases {
		got := dictionary.SearchAnagramsByWord(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("SearchAnagramsByWord(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
