package main

import (
	"testing"

	"github.com/giorgioprevitera/advent-of-code-2020/advent"
	"gotest.tools/assert"
)

func TestPartOne(t *testing.T) {
	input, ok := advent.GetInput("test-input.txt").(*[]string)
	if !ok {
		t.Fatal(ok)
	}

	expectedOutput := 25
	actualOutput := getPartOneAnswer(input)
	assert.Equal(t, expectedOutput, actualOutput)
}

func TestPartTwo(t *testing.T) {
	input, ok := advent.GetInput("test-input.txt").(*[]string)
	if !ok {
		t.Fatal(ok)
	}

	expectedOutput := 286
	actualOutput := getPartTwoAnswer(input)
	assert.Equal(t, expectedOutput, actualOutput)
}
