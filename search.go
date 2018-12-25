package main

import "fmt"

// Node of Trie data structure
type Node struct {
	children map[rune]*Node
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
	}

	return root
}

// Search words positions in text with Aho-Corasick
func Search(text string, words ...string) []int {
	trie := BuildTrie(words...)
	matches := []int{}

	depth, node := 0, trie

	for position, char := range text {
		child, present := node.children[char]

		if present {
			if len(child.children) == 0 {
				matches = append(matches, position - depth)
				depth, node = 0, trie
			} else {
				depth, node = depth + 1, child
			}
		} else {
			depth, node = 0, trie
		}
	}

	return matches
}

func main() {
	// XXX fgrep-like arguments and output
	fmt.Println("Search: 'abc' for 'a', 'b' and 'c'")

	res := Search("abc", "a", "b", "c")
	fmt.Println("Result:", res)
}
