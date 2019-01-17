package main

import "fmt"

// Node of Trie data structure
type Node struct {
	children map[rune]*Node
	inDictionary bool
	suffix *Node
	dictionarySuffix *Node
	depth int
}

func initNode() *Node {
	return &Node {
		children: make(map[rune]*Node),
	}
}

// BuildTrie builds a Trie data structure
func buildTrie(words ...string) *Node {
	root := initNode()

	for _, word := range words {
		node := root

		for i, char := range word {
			child, present := node.children[char]

			if !present {
				child = initNode()
				child.depth = i + 1
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
			suffix, present := trie.children[char]

			if present {
				child.suffix = suffix
			} else {
				child.suffix = trie
			}
		}

		addSuffixesInside(trie, child)
	}
}


// AddSuffixes pointer to the longest strict suffix branch from every node
func addSuffixes(trie *Node) {
	for _, child := range trie.children {
		child.suffix = trie
	}

	for _, child := range trie.children {
		addSuffixesInside(trie, child)
	}
}

func addDictionarySuffixesInside(current *Node) {
	for _, child := range current.children {
		next := child.suffix

		for next != nil {
			if next.inDictionary {
				child.dictionarySuffix = next
				break
			}
			next = next.suffix
		}

		addDictionarySuffixesInside(child)
	}
}

// AddDictionarySuffixes pointer to the largest suffix in dictionary
func addDictionarySuffixes(trie *Node) {
	for _, child := range trie.children {
		child.dictionarySuffix = nil
		addDictionarySuffixesInside(child)
	}
}

// Compile words into a Aho-Corasick trie with suffixes links
func compile(words ...string) *Node {
	trie := buildTrie(words...)

	addSuffixes(trie)
	addDictionarySuffixes(trie)

	return trie
}

func nextNode(node *Node, char rune) *Node {
	for {
		child, present := node.children[char]

		if present {
			return child
		}

		if node.suffix == nil {
			return nil
		}

		node = node.suffix
	}
}

// Search words positions in text with Aho-Corasick
func Search(text string, words ...string) []int {
	trie := compile(words...)

	matches := []int{}
	node := trie

	for position, char := range text {
		node = nextNode(node, char)

		if node == nil {
			node = trie
			continue
		}

		// Current node
		if node.inDictionary {
			matches = append(matches, position - node.depth + 1)
		}

		// Dictionary suffixes
		for suffix := node.dictionarySuffix; suffix != nil; {
			matches = append(matches, position - suffix.depth + 1)
			suffix = suffix.dictionarySuffix
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
