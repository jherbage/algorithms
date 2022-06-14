package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		input  string
		output int
	}{
		{
			input:  "",
			output: 0,
		},
		{
			input:  "aabb",
			output: 8,
		},
		{
			input:  "abcdefa",
			output: 3600,
		},
	}

	for _, testCase := range testCases {
		output := do(testCase.input)
		assert.Equal(t, testCase.output, output, "output is not correct")
	}
}

func Test_includesRepeats(t *testing.T) {
	testCases := []struct {
		input  string
		output bool
	}{
		{
			input:  "",
			output: false,
		},
		{
			input:  "aab",
			output: true,
		},
		{
			input:  "abc",
			output: false,
		},
	}

	for _, testCase := range testCases {
		r := []rune(testCase.input)
		output := includesRepeats(r)
		assert.Equal(t, testCase.output, output, "output is not correct")
	}
}
