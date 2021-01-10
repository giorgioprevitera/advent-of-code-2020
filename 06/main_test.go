package main

import (
	"log"
	"testing"

	"github.com/giorgioprevitera/advent-of-code-2020/advent"
	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	i, ok := advent.GetInput("test-input.txt").(*[]string)
	if !ok {
		log.Fatalln("Unable to make type assertion on input")
	}
	input := Input(*i)

	expectedOutput := 11
	actualOutput := getAnswerPartOne(&input)

	assert.Equal(t, expectedOutput, actualOutput)
}

func TestPartTwo(t *testing.T) {
	i, ok := advent.GetInput("test-input.txt").(*[]string)
	if !ok {
		log.Fatalln("Unable to make type assertion on input")
	}
	input := Input(*i)

	expectedOutput := 6
	actualOutput := getAnswerPartTwo(&input)

	assert.Equal(t, expectedOutput, actualOutput)
}
