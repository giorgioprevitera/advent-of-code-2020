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

	expectedOutput := 11
	actualOutput := getAnswerPartOne(input)

	assert.Equal(t, expectedOutput, actualOutput)
}

func TestPartTwo(t *testing.T) {
	input, err := getInput("test-input.txt")
	if err != nil {
		t.Fatal(err)
	}

	expectedOutput := 6
	actualOutput := getAnswerPartTwo(input)

	assert.Equal(t, expectedOutput, actualOutput)
}
