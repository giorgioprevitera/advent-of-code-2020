package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	target, input := parseInput("test-input.txt")
	expectedValue := 295
	actualValue := getPartOneAnswer(input, target)

	assert.Equal(t, expectedValue, actualValue)
}

func TestPartTwo(t *testing.T) {
	_, input := parseInput("test-input.txt")
	expectedValue := 1068781
	actualValue := getPartTwoAnswer(input)

	assert.Equal(t, expectedValue, actualValue)
}
