package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	input := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}
	expectedOutput := 514579

	actualOutput, err := searchAnswerPartOne(input, 2020)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expectedOutput, actualOutput)
}

func TestPartTwo(t *testing.T) {
	input := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}
	expectedOutput := 241861950

	actualOutput, err := searchAnswerPartTwo(input, 2020)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expectedOutput, actualOutput)
}
