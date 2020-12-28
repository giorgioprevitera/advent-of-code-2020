package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {

	testData := []struct {
		Input          string
		ExpectedOutput int
	}{
		{"BFFFBBFRRR", 567},
		{"FFFBBBFRRR", 119},
		{"BBFFBBFRLL", 820},
	}

	for _, d := range testData {
		t.Run("Testing", func(t *testing.T) {
			t.Logf("%s: %d", d.Input, d.ExpectedOutput)
			assert.Equal(t, getSeatID(d.Input), d.ExpectedOutput)
		})
	}

}

func TestPartTwo(t *testing.T) {
	input := []int{1, 2, 4, 5}
	expectedOutput := 3
	actualOutput := findMissingID(&input)

	assert.Equal(t, expectedOutput, actualOutput)
}
