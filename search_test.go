package main

import "testing"

func compare(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i:=0; i<len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// Text and indexes
// a b x a b c d e f a b
// 0 1 2 3 4 5 6 7 8 9 0
func verify(t *testing.T, words []string, expected []int) {
	text := "abxabcdefab"
	res := Search(text, words...)

	if !compare(res, expected) {
		t.Errorf(
			"Search(%s, %s) returned %d expected %d",
			text, words, res, expected,
		)
	}
}


func TestSearchSinglePosition(t *testing.T) {
	verify(t, []string{"c"}, []int{5})
}

func TestSearchSingleWord(t *testing.T) {
	verify(t, []string{"abc"}, []int{3})
}

func TestSearchManyWords(t *testing.T) {
	verify(t, []string{"aac", "ac", "abc", "bca"}, []int{3})
}

func TestSearchManyMatches(t *testing.T) {
	verify(t, []string{"ab"}, []int{0, 3, 9})

	verify(t, []string{"abc", "fab"}, []int{3, 8})

	verify(t, []string{"ab", "cde"}, []int{0, 3, 5, 9})
}

func TestPrefixWord(t *testing.T) {
	verify(t, []string{"ab", "abcd"}, []int{0, 3, 9})

	verify(t, []string{"abcd", "ab"}, []int{0, 3, 9})
}

func TestPrefixTwice(t *testing.T) {
	verify(t, []string{"ab", "abc", "abcd"}, []int{0, 3, 9})

	verify(t, []string{"abcd", "abc", "ab"}, []int{0, 3, 9})
}

func TestSuffixWord(t *testing.T) {
	verify(t, []string{"ab", "b"}, []int{0, 1, 3, 4, 9, 10})

	verify(t, []string{"abx", "bx"}, []int{0, 1})
}

func TestSuffixTwice(t *testing.T) {
	verify(t, []string{"abx", "bx", "x"}, []int{0, 1, 2})
}

func TestMoveToSuffix(t *testing.T) {
	verify(t, []string{"abxb", "bxa"}, []int{1})
}

func TestMoveToSuffixTwice(t *testing.T) {
	verify(t, []string{"abxb", "bxb", "xa"}, []int{2})
}

func TestWikipediaExample(t *testing.T) {
	// From https://en.wikipedia.org/wiki/Aho%E2%80%93Corasick_algorithm

	res := Search(
		"abccab",
		"a", "ab", "bab", "bc", "bca", "c", "caa")

	expected := []int{0, 1, 2, 3, 4}

	if !compare(res, expected) {
		t.Errorf("Got result %d expected %d", res, expected)
	}
}

func TestSearch2Interface(t *testing.T) {
	res := Search2("abc", "b")
	expected := Match{"b", 1, 2}

	if len(res) != 1 || res[0] != expected {
		t.Fail()
	}
}
