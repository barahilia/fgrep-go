package main

import (
	"fmt"
	"sort"
)

// Node of Trie data structure
type Node struct {
	children map[rune]*Node
	inDictionary bool
	suffix *Node
}

func initNode() *Node {
	return &Node {
		children: make(map[rune]*Node),
	}
}

// BuildTrie builds a Trie data structure
func BuildTrie(words ...string) *Node {
	root := initNode()

	for _, word := range words {
		node := root

		for _, char := range word {
			child, present := node.children[char]

			if !present {
				child = initNode()
				node.children[char] = child
			}

			node = child
		}

		node.inDictionary = true
	}

	return root
}

func addSuffixesInside(trie *Node, current *Node) {
	for char, child := range current.children {
		suffix, present := current.suffix.children[char]

		if present {
			child.suffix = suffix
		} else {
			child.suffix = trie
		}

		addSuffixesInside(trie, child)
	}
}


// AddSuffixes pointer to the longest strict suffix branch from every node
func AddSuffixes(trie *Node) {
	for _, child := range trie.children {
		child.suffix = trie
	}

	for _, child := range trie.children {
		addSuffixesInside(trie, child)
	}
}

func sortedUnique(arr []int) []int {
	set := make(map[int]struct{})
	unique := []int{}

	for _, i := range arr {
		_, present := set[i]

		if !present {
			set[i] = struct{}{}
			unique = append(unique, i)
		}
	}

	sort.Ints(unique)

	return unique
}

// Search words positions in text with Aho-Corasick
func Search(text string, words ...string) []int {
	trie := BuildTrie(words...)
	matches := []int{}

	depth, node := 0, trie

	for position, char := range text {
		child, present := node.children[char]

		if present {
			if child.inDictionary {
				matches = append(matches, position - depth)
			}

			if len(child.children) == 0 {
				depth, node = 0, trie
			} else {
				depth, node = depth + 1, child
			}
		} else {
			depth, node = 0, trie
		}
	}

	return sortedUnique(matches)
}

func main() {
	// XXX fgrep-like arguments and output
	fmt.Println("Search: 'abc' for 'a', 'b' and 'c'")

	res := Search("abc", "a", "b", "c")
	fmt.Println("Result:", res)
}
