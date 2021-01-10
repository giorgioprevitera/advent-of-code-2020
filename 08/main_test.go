package main

import (
	"log"
	"testing"

	"github.com/giorgioprevitera/advent-of-code-2020/advent"
	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	input, ok := advent.GetInput("test-input.txt").(*[]string)
	if !ok {
		log.Fatalln("Unable to make type assertion on input")
	}

	expectedOutput := 5
	actualOutput := getPartOneAnswer(input)

	assert.Equal(t, expectedOutput, actualOutput)

}

func TestPartTwo(t *testing.T) {
	input, ok := advent.GetInput("test-input.txt").(*[]string)
	if !ok {
		log.Fatalln("Unable to make type assertion on input")
	}

	expectedOutput := 8
	actualOutput := getPartTwoAnswer(input)

	assert.Equal(t, expectedOutput, actualOutput)

}
