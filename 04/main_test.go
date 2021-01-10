package main

import (
	"log"
	"testing"

	"github.com/giorgioprevitera/advent-of-code-2020/advent"
	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	i, ok := advent.GetInput("test-part-one-input.txt").(*[]string)
	if !ok {
		log.Fatalln("Unable to make type assertion on input")
	}

	input := Input(*i)

	expectedOutput := 2

	passports := input.parse()
	actualOutput := 0
	for _, p := range *passports {
		if p.haveValidKeys() {
			actualOutput++
		}
	}

	assert.Equal(t, expectedOutput, actualOutput)
}

func TestPartTwo(t *testing.T) {
	t.Run("Test invalid passports", func(t *testing.T) {
		i, ok := advent.GetInput("./test-part-two-invalid-input.txt").(*[]string)
		if !ok {
			log.Fatalln("Unable to make type assertion on input")
		}
		input := Input(*i)

		expectedOutput := 0

		passports := input.parse()
		actualOutput := 0
		for _, p := range *passports {
			if p.isValid() {
				actualOutput++
			}
		}

		assert.Equal(t, expectedOutput, actualOutput)

	})
	t.Run("Test valid passports", func(t *testing.T) {
		i, ok := advent.GetInput("./test-part-two-valid-input.txt").(*[]string)
		if !ok {
			log.Fatalln("Unable to make type assertion on input")
		}
		input := Input(*i)

		expectedOutput := 4

		passports := input.parse()
		actualOutput := 0
		for _, p := range *passports {
			if p.isValid() {
				actualOutput++
			}
		}

		assert.Equal(t, expectedOutput, actualOutput)

	})
}
