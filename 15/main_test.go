package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestPartOne(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{input: "0,3,6", expected: 436},
		{input: "1,3,2", expected: 1},
		{input: "2,1,3", expected: 10},
		{input: "1,2,3", expected: 27},
		{input: "2,3,1", expected: 78},
		{input: "3,2,1", expected: 438},
		{input: "3,1,2", expected: 1836},
	}

	for _, tc := range testCases {
		input := parseInput(tc.input)
		actual := getAnswer(input, 2020)
		assert.Equal(t, tc.expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{input: "0,3,6", expected: 175594},
		{input: "1,3,2", expected: 2578},
		{input: "2,1,3", expected: 3544142},
		{input: "1,2,3", expected: 261214},
		{input: "2,3,1", expected: 6895259},
		{input: "3,2,1", expected: 18},
		{input: "3,1,2", expected: 362},
	}

	for _, tc := range testCases {
		input := parseInput(tc.input)
		actual := getAnswer(input, 30000000)
		assert.Equal(t, tc.expected, actual)
	}
}
