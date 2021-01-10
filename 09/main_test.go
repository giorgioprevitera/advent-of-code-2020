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
		log.Fatalln("Unable to make type assertion on input")
	}
	input := Input(*i)

	preamble := 5
	expectedOutput := 127
	actualOutput := getPartOneAnswer(&input, preamble)
	assert.Equal(t, expectedOutput, actualOutput)
}

func TestPartTwo(t *testing.T) {
	i, ok := advent.GetInput("test-input.txt").(*[]int)
	if !ok {
		log.Fatalln("Unable to make type assertion on input")
	}
	input := Input(*i)

	answerPartOne := 127
	expectedOutput := 62
	actualOutput := getPartTwoAnswer(&input, answerPartOne)
	assert.Equal(t, expectedOutput, actualOutput)
}
