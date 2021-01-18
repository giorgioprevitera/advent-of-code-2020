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
