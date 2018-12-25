package main

import "testing"

func TestNodeCreation(t *testing.T) {
	node := &Node{}

	if node == nil {
		t.Fail()
	}
}

func TestChildren(t *testing.T) {
	root := &Node{
		children: make(map[rune]*Node),
	}
	root.children['a'] = &Node{}

	_, ok := root.children['a']

	if !ok {
		t.Fail()
	}
}

func TestEmptyWordTrie(t *testing.T) {
	trie := BuildTrie("")

	if len(trie.children) != 0 {
		t.Fail()
	}
}

func TestSingleWordTrie(t *testing.T) {
	trie := BuildTrie("abc")

	for _, char := range "abc" {
		if len(trie.children) != 1 {
			t.FailNow()
		}

		trie = trie.children[char]
	}

	if len(trie.children) != 0 {
		t.Fail()
	}
}

func isEmptyNode(node *Node) bool {
	return len(node.children) == 0
}

func TestMultipleWords(t *testing.T) {
	trie := BuildTrie("a", "cba", "cbx", "abc")

	// a -> b -> c
	// c -> b -> a
	// 			 x

	correct := (
		isEmptyNode(trie.children['a'].children['b'].children['c']) &&
		isEmptyNode(trie.children['c'].children['b'].children['a']) &&
		isEmptyNode(trie.children['c'].children['b'].children['x']))

	if !correct {
		t.Fail()
	}
}

func TestWikipediaTrie(t *testing.T) {
	// From https://en.wikipedia.org/wiki/Aho%E2%80%93Corasick_algorithm

	trie := BuildTrie("a", "ab", "bab", "bc", "bca", "c", "caa")

	correct := (
		isEmptyNode(trie.children['a'].children['b']) &&

		isEmptyNode(trie.children['b'].children['a'].children['b']) &&
		isEmptyNode(trie.children['b'].children['c'].children['a']) &&

		isEmptyNode(trie.children['c'].children['a'].children['a']))

	if !correct {
		t.Fail()
	}
}
