package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		inputs      [][]int
		output      []int
		shouldError bool
	}{
		{
			inputs: [][]int{
				{1, 2, 3},
			},
			output:      nil,
			shouldError: true,
		},
		{
			inputs:      [][]int{},
			output:      nil,
			shouldError: true,
		},
		{
			inputs:      nil,
			output:      nil,
			shouldError: true,
		},
		{
			inputs: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			output:      []int{},
			shouldError: false,
		},
		{
			inputs: [][]int{
				{1, 2, 3},
				{3, 4, 5, 6},
			},
			output:      []int{3},
			shouldError: false,
		},
		{
			inputs: [][]int{
				{1, 2, 3},
				{1, 2, 3, 4, 5, 6},
			},
			output:      []int{1, 2, 3},
			shouldError: false,
		},
		{
			inputs: [][]int{
				{1, 2, 3, 4},
				{1, 2, 3, 4, 5, 6},
				{3, 4},
			},
			output:      []int{3, 4},
			shouldError: false,
		},
		{
			inputs: [][]int{
				{1, 2, 3, 4, 5},
				{1, 2, 3, 4, 5, 6},
				{3, 4, 5},
				{4, 5},
			},
			output:      []int{4, 5},
			shouldError: false,
		},
	}

	for _, testCase := range testCases {
		output, err := do(testCase.inputs)
		if testCase.shouldError {
			assert.NotNil(t, err, "Test should error")
		}
		assert.Equal(t, testCase.output, output, "output is not correct")
	}
}
