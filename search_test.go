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

func verify(
	t *testing.T,
	words []string,
	expected []int,
) {
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