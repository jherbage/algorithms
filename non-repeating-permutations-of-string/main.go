package main

import (
	"fmt"
)

/*
Return the number of total permutations of the provided string that don't have repeated consecutive letters.
Assume that all characters in the provided string are each unique.

For example, aab should return 2 because it has 6 total permutations (aab, aab, aba, aba, baa, baa),
but only 2 of them (aba and aba) don't have the same letter (in this case a) repeating.
*/
func do(a string) int {
	runes := []rune(a)
	output := permutations(runes)
	matched := 0

	// would do this with regex but golang regex doesnt support backrefs!
	for _, r := range output {
		if !includesRepeats(r) {
			fmt.Println("didnt match " + string(r))
			matched++
		}
	}
	return matched
}

func includesRepeats(a []rune) bool {
	for i, r := range a {
		if i == 0 {
			continue
		}
		if r == a[i-1] {
			return true
		}
	}
	return false
}

func permutations(arr []rune) [][]rune {
	var helper func([]rune, int)
	res := [][]rune{}

	helper = func(arr []rune, n int) {
		if n == 1 {
			tmp := make([]rune, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
