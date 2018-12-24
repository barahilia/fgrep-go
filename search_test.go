package main

import "testing"

func compare(a, b []string) bool {
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

func check(
	t *testing.T,
	words []string,
	expected []string,
) {
	text := "abcde"
	res := Search(text, words...)

	if !compare(res, expected) {
		t.Errorf(
			"Search(%s, %s) returned %s expected %s",
			text, words, res, expected,
		)
	}
}

func TestInterface(t *testing.T) {
	check(t, []string{"a", "b", "c"}, []string{"a", "b", "c"})
}

func TestEmptyWords(t *testing.T) {
	check(t, []string{}, []string{})
}

func TestSingleWord(t *testing.T) {
	check(t, []string{"a"}, []string{"a"})
	check(t, []string{"cd"}, []string{"cd"})
	check(t, []string{"x"}, []string{})
	check(t, []string{"abd"}, []string{})
}
