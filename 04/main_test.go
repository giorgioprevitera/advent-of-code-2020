package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	input, err := getInput("test-part-one-input.txt")
	if err != nil {
		t.Fatal(err)
	}

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
		input, err := getInput("test-part-two-invalid-input.txt")
		if err != nil {
			t.Fatal(err)
		}

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
		input, err := getInput("test-part-two-valid-input.txt")
		if err != nil {
			t.Fatal(err)
		}

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
