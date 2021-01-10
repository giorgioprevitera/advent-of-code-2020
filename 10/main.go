package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

type Input []int

func getInput(filename string) *Input {
	var input Input

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	input = append(input, 0)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		input = append(input, i)
	}

	sort.Ints(input)
	input = append(input, input[len(input)-1]+3)
	return &input
}

func getPartOneAnswer(input *Input) int {
	oneJoltDiffs := 0
	threeJoltDiffs := 0

	for i := 0; i < len(*input)-1; i++ {
		if (*input)[i+1]-(*input)[i] == 1 {
			oneJoltDiffs++
		}
		if (*input)[i+1]-(*input)[i] == 3 {
			threeJoltDiffs++
		}
	}

	return oneJoltDiffs * threeJoltDiffs
}

var DP = map[int]int{}

func getPartTwoAnswer(input *Input, index int) int {
	if index == len(*input)-1 {
		return 1
	}
	if DP[index] != 0 {
		return DP[index]
	}
	answer := 0
	for j := index + 1; j < len(*input); j++ {
		if (*input)[j]-(*input)[index] <= 3 {
			answer += getPartTwoAnswer(input, j)
		}
	}
	DP[index] = answer
	return answer
}

func main() {
	input := getInput("input.txt")

	answerPartOne := getPartOneAnswer(input)
	log.Printf("Answer to part one: %d\n", answerPartOne)

	answerPartTwo := getPartTwoAnswer(input, 0)
	log.Printf("Answer to part two: %d\n", answerPartTwo)
}
