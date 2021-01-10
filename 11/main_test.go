package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	input := getInput("test-input.txt")
	expectedOuput := 37
	actualOutput := getPartOneAnswer(input)
	assert.Equal(t, expectedOuput, actualOutput)
}

func TestPartTwo(t *testing.T) {
	input := getInput("test-input.txt")
	expectedOuput := 26
	actualOutput := getPartTwoAnswer(input)
	assert.Equal(t, expectedOuput, actualOutput)
}
