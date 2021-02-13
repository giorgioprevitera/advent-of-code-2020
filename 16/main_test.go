package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOne(t *testing.T) {
	rules, _, otherTickets := parseInput("./test-input.txt")

	expected := 71
	actual := getPartOnwAnswer(rules, otherTickets)

	assert.Equal(t, expected, actual)
}

func TestPartTwo(t *testing.T) {
	rules, myTicket, otherTickets := parseInput("./input.txt")

	expected := 1053686852011
	actual := getPartTwoAnswer(rules, myTicket, otherTickets)

	assert.Equal(t, expected, actual)
}
