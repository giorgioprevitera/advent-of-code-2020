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

	expectedOutput := 7
	actualOutput := countTrees(input, slopesPartOne)
	assert.Equal(t, expectedOutput, actualOutput)
}

func TestPartTwo(t *testing.T) {
	input, err := getInput("test-input.txt")
	if err != nil {
		t.Fatal(err)
	}

	expectedOutput := 336
	actualOutput := countTrees(input, slopesPartTwo)
	assert.Equal(t, expectedOutput, actualOutput)
}
