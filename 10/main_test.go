package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	input := getInput("test-input.txt")
	expectedOutput := 220
	actualOutput := getPartOneAnswer(input)
	assert.Equal(t, expectedOutput, actualOutput)
}

func TestPartTwo(t *testing.T) {
	input := getInput("test-input.txt")
	expectedOutput := 19208
	actualOutput := getPartTwoAnswer(input, 0)
	assert.Equal(t, expectedOutput, actualOutput)
}
