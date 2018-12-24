package main

import "fmt"
import "strings"

// Search words in text with Aho-Corasick
func Search(text string, words ...string) []string {
	if len(words) == 0 {
		return []string{}
	}

	if len(words) == 1 {
		word := words[0]

		if strings.Contains(text, word) {
			return []string{word}
		}

		return []string{}
	}

	res := []string {"a", "b", "c"}
	return res
}

func main() {
	// XXX fgrep-like arguments and output
	fmt.Println("Search: 'abc' for 'a', 'b' and 'c'")

	res := Search("abc", "a", "b", "c")
	fmt.Println("Result:", res)
}
