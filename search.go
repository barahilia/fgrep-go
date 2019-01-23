package main

// Node of Trie data structure
type Node struct {
	children map[rune]*Node
	inDictionary bool
	suffix *Node
	dictionarySuffix *Node
	depth int
}

func newNode(depth int) *Node {
	return &Node {
		children: make(map[rune]*Node),
		depth: depth,
	}
}

// buildTrie builds a Trie data structure
func buildTrie(words ...string) *Node {
	root := newNode(0)

	for _, word := range words {
		node := root

		for i, char := range word {
			child, present := node.children[char]

			if !present {
				child = newNode(i + 1)
				node.children[char] = child
			}

			node = child
		}

		node.inDictionary = true
	}

	return root
}

// addSuffixes pointer to the longest strict suffix branch from every node
func addSuffixes(trie *Node, current *Node) {
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

		addSuffixes(trie, child)
	}
}

// addDictionarySuffixes pointer to the largest suffix in dictionary
func addDictionarySuffixes(current *Node) {
	for _, child := range current.children {
		next := child.suffix

		for next != nil {
			if next.inDictionary {
				child.dictionarySuffix = next
				break
			}
			next = next.suffix
		}

		addDictionarySuffixes(child)
	}
}

// compile words into a Aho-Corasick trie with suffixes links
func compile(words ...string) *Node {
	trie := buildTrie(words...)

	for _, child := range trie.children {
		child.suffix = trie
		addSuffixes(trie, child)
	}

	for _, child := range trie.children {
		child.dictionarySuffix = nil
		addDictionarySuffixes(child)
	}

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

// Match is returned by Search for words in text
type Match struct {
	word string
	start int
	end int
}

func newMatch(text string, start int, end int) Match {
	return Match{text[start: end], start, end}
}

func searchImpl(text string, trie *Node, matches chan<- Match) {
	node := trie

	for position, char := range text {
		node = nextNode(node, char)

		if node == nil {
			node = trie
			continue
		}

		if node.inDictionary {
			matches <- newMatch(text, position - node.depth + 1, position + 1)
		}

		for suffix := node.dictionarySuffix; suffix != nil; {
			matches <- newMatch(text, position - suffix.depth + 1, position + 1)
			suffix = suffix.dictionarySuffix
		}
	}

	close(matches)
}

// Search for words in text with Aho-Corasick
func Search(text string, words ...string) []Match {
	trie := compile(words...)

	result := []Match{}
	matches := make(chan Match)

	go searchImpl(text, trie, matches)

	for match := range matches {
		result = append(result, match)
	}

	return result
}
