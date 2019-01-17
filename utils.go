package main

import "sort"

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
