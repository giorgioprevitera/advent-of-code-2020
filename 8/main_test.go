package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	input, err := getInput("test-input.txt")
	if err != nil {
		log.Fatal(err)
	}

	expectedOutput := 5
	actualOutput := getPartOneAnswer(&input)

	assert.Equal(t, expectedOutput, actualOutput)

}
