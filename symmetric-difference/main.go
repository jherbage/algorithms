package main

import (
	"errors"
	"sort"
)

func do(lists [][]int) ([]int, error) {

	if len(lists) < 2 {
		return nil, errors.New("Min number of lists is 2")
	}

	symmetric := compare(nil, lists)
	// map back to list and sort so tests are easier
	keys := make([]int, 0, len(symmetric))
	for k := range symmetric {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	return keys, nil
}

func createUniqueMap(list []int) map[int]bool {
	symmetric := map[int]bool{}
	for _, i := range list {
		if _, ok := symmetric[i]; !ok {
			symmetric[i] = true
		}
	}
	return symmetric
}

// recursive symmetric comparison of int lists
func compare(symmetric map[int]bool, lists [][]int) map[int]bool {

	var newlists [][]int
	var lastlist map[int]bool
	// pop last element from list
	if len(lists) > 0 {
		lastlist = createUniqueMap(lists[len(lists)-1])
		newlists = lists[0 : len(lists)-1]
	}

	if symmetric == nil {
		return compare(lastlist, newlists)
	} else if newlists != nil {
		for i, _ := range symmetric {
			if ok := lastlist[i]; !ok {
				delete(symmetric, i)
			}
		}
		return compare(symmetric, newlists)
	} else {
		return symmetric
	}

	return symmetric
}
