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

	bags := input.parse()

	expectedOutput := 4
	actualOutput := getPartOneAnswer(&bags, "shiny gold")

	assert.Equal(t, expectedOutput, actualOutput)
}

func TestPartTwo(t *testing.T) {
	input, err := getInput("test-input.txt")
	if err != nil {
		t.Fatal(err)
	}

	bags := input.parse()

	expectedOutput := 32
	actualOutput := getPartTwoAnswer(&bags, "shiny gold")

	assert.Equal(t, expectedOutput, actualOutput)
}
