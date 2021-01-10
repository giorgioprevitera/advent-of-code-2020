package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	input := []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}
	expectedOutput := 2
	actualOutput := countValidPasswords(&input, partOne)
	assert.Equal(t, expectedOutput, actualOutput)
}

func TestPartTwo(t *testing.T) {
	input := []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}
	expectedOutput := 1
	actualOutput := countValidPasswords(&input, partTwo)
	assert.Equal(t, expectedOutput, actualOutput)
}
