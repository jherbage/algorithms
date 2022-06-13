package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		current map[string]int
		new     map[string]int
		output  map[string]int
	}{
		{
			current: map[string]int{},
			new:     map[string]int{},
			output:  map[string]int{},
		},
		{
			current: map[string]int{
				"a ball": 1,
			},
			new: map[string]int{
				"a ball": 1,
			},
			output: map[string]int{
				"a ball": 2,
			},
		},
		{
			current: map[string]int{
				"a ball": 1,
			},
			new: map[string]int{
				"another ball": 1,
			},
			output: map[string]int{
				"a ball":       1,
				"another ball": 1,
			},
		},
		{
			current: map[string]int{
				"a ball":     2,
				"and string": 2,
			},
			new: map[string]int{
				"another ball": 1,
				"a ball":       1,
			},
			output: map[string]int{
				"a ball":       3,
				"another ball": 1,
				"and string":   2,
			},
		},
	}

	for _, testCase := range testCases {
		output := do(testCase.current, testCase.new)
		assert.Equal(t, testCase.output, output, "output is not correct")
	}
}
