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
	trie := buildTrie("")

	if len(trie.children) != 0 {
		t.Fail()
	}
}

func TestSingleWordTrie(t *testing.T) {
	trie := buildTrie("abc")

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
	trie := buildTrie("a", "cba", "cbx", "abc")

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

	trie := buildTrie("a", "ab", "bab", "bc", "bca", "c", "caa")

	correct := (
		isEmptyNode(trie.children['a'].children['b']) &&

		isEmptyNode(trie.children['b'].children['a'].children['b']) &&
		isEmptyNode(trie.children['b'].children['c'].children['a']) &&

		isEmptyNode(trie.children['c'].children['a'].children['a']))

	if !correct {
		t.Fail()
	}
}

func TestSuffixTwoLetters(t *testing.T) {
	trie := compile("ab", "b")

	bNode := trie.children['b']
	aNode := trie.children['a']
	abNode := aNode.children['b']

	correct := (
		aNode.suffix == trie &&
		bNode.suffix == trie &&
		abNode.suffix == bNode)

	if !correct {
		t.Fail()
	}
}

func TestSuffixDeep(t *testing.T) {
	trie := compile("abcd", "xyzabc", "abx")

	abcNode := trie.children['a'].children['b'].children['c']
	xyzNode := trie.children['x'].children['y'].children['z']
	deepNode := xyzNode.children['a'].children['b'].children['c']

	if deepNode.suffix != abcNode {
		t.Fail()
	}
}

func getNode(trie *Node, s string) *Node {
	node := trie

	for _, char := range s {
		node = node.children[char]
	}

	return node
}

func TestWikipediaSuffix(t *testing.T) {
	// From https://en.wikipedia.org/wiki/Aho%E2%80%93Corasick_algorithm

	trie := compile("a", "ab", "bab", "bc", "bca", "c", "caa")

	verifySuffix := func (s string, suffix string) {
		node := getNode(trie, s)
		suffixNode := getNode(trie, suffix)

		if node.suffix != suffixNode {
			t.Errorf("Suffix of '%s' is not '%s'", s, suffix)
		}
	}

	if trie.suffix != nil {
		t.Error("Root suffix is not nil")
	}

	verifySuffix("a", "")
	verifySuffix("ab", "b")
	verifySuffix("b", "")
	verifySuffix("ba", "a")
	verifySuffix("bab", "ab")
	verifySuffix("bc", "c")
	verifySuffix("bca", "ca")
	verifySuffix("c", "")
	verifySuffix("ca", "a")
	verifySuffix("caa", "a")
}

func TestDictionarySuffix(t *testing.T) {
	trie := compile("ab", "b")

	bNode := trie.children['b']
	aNode := trie.children['a']
	abNode := aNode.children['b']

	correct := (
		aNode.dictionarySuffix == nil &&
		bNode.dictionarySuffix == nil &&
		abNode.dictionarySuffix == bNode)

	if !correct {
		t.Fail()
	}
}

func TestWikipediaDictionarySuffix(t *testing.T) {
	// From https://en.wikipedia.org/wiki/Aho%E2%80%93Corasick_algorithm

	trie := compile("a", "ab", "bab", "bc", "bca", "c", "caa")

	verifyEmpty := func (s string) {
		node := getNode(trie, s)

		if node.dictionarySuffix != nil {
			t.Errorf("Dictionary suffix of %s is not empty", s)
		}
	}

	verifyDictionarySuffix := func (s string, suffix string) {
		node := getNode(trie, s)
		suffixNode := getNode(trie, suffix)

		if node.dictionarySuffix != suffixNode {
			t.Errorf("Dictionary suffix of %s is not %s", s, suffix)
		}
	}

	verifyEmpty("")
	verifyEmpty("a")
	verifyEmpty("ab")
	verifyEmpty("b")
	verifyDictionarySuffix("ba", "a")
	verifyDictionarySuffix("bab", "ab")
	verifyDictionarySuffix("bc", "c")
	verifyDictionarySuffix("bca", "a")
	verifyEmpty("c")
	verifyDictionarySuffix("ca", "a")
	verifyDictionarySuffix("caa", "a")
}
