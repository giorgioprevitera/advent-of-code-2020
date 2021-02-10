package main

import (
	"log"
	"strconv"
	"strings"
)

type number struct {
	value  int
	spoken bool
	turn   int
}

func parseInput(input string) map[int]number {
	output := map[int]number{}
	for i, v := range strings.Split(input, ",") {
		n, _ := strconv.Atoi(v)

		output[n] = number{value: n, spoken: true, turn: i + 1}
	}

	if !output[0].spoken {
		output[0] = number{value: 0, spoken: true, turn: len(output) + 1}
	}

	return output
}

func getAnswer(input map[int]number, limit int) int {
	current := input[0]

	for turn := len(input) + 1; turn < limit; turn++ {
		nextNumber := turn - current.turn

		current.spoken = true
		current.turn = turn
		input[current.value] = current

		if !current.spoken {
			current = input[0]
		} else {
			if !input[nextNumber].spoken {
				input[nextNumber] = number{value: nextNumber, spoken: false, turn: turn + 1}
			}
			current = input[nextNumber]
		}
	}
	return current.value
}

func main() {
	input := parseInput("20,0,1,11,6,3")

	answerToPartOne := getAnswer(input, 2020)
	log.Println("Answer to part one:", answerToPartOne)

	input = parseInput("20,0,1,11,6,3")

	answerToPartTwo := getAnswer(input, 30000000)
	log.Println("Answer to part two:", answerToPartTwo)
}
