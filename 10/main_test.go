package main

import (
	"log"
	"testing"

	"github.com/giorgioprevitera/advent-of-code-2020/advent"
	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	i, ok := advent.GetInput("test-input.txt").(*[]int)
	if !ok {
		log.Fatal(ok)
	}
	input := Input(*i)
	input.normalise()

	expectedOutput := 220
	actualOutput := getPartOneAnswer(&input)
	assert.Equal(t, expectedOutput, actualOutput)
}

func TestPartTwo(t *testing.T) {
	i, ok := advent.GetInput("test-input.txt").(*[]int)
	if !ok {
		log.Fatal(ok)
	}
	input := Input(*i)
	input.normalise()

	expectedOutput := 19208
	actualOutput := getPartTwoAnswer(&input, 0)
	assert.Equal(t, expectedOutput, actualOutput)
}
