package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	input, err := getInput("test-input.txt")
	if err != nil {
		t.Fatal(err)
	}

	preamble := 5
	expectedOutput := 127
	actualOutput := getPartOneAnswer(input, preamble)
	assert.Equal(t, expectedOutput, actualOutput)
}

func TestPartTwo(t *testing.T) {
	input, err := getInput("test-input.txt")
	if err != nil {
		t.Fatal(err)
	}

	answerPartOne := 127
	expectedOutput := 62
	actualOutput := getPartTwoAnswer(input, answerPartOne)
	assert.Equal(t, expectedOutput, actualOutput)
}
