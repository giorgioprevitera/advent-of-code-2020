package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Input []int

func getInput(filename string) (*Input, error) {
	f, err := os.Open(filename)
	if err != nil {
		return &Input{}, err
	}
	defer f.Close()

	input := Input{}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		input = append(input, i)
	}

	return &input, nil
}

func getPartOneAnswer(input *Input, preamble int) int {
	answer := -1
	IsSumOfPreviousTwo := make(map[int]bool, len(*input)-preamble)

	for i, n := range *input {

		// skip preamble numbers
		if i < preamble {
			continue
		}

		for _, x := range (*input)[i-preamble : i] {

			IsSumOfPreviousTwo[n] = false
			for _, y := range (*input)[i-preamble+1 : i] {
				if x+y == n {
					IsSumOfPreviousTwo[n] = true
					break
				}
			}

			if IsSumOfPreviousTwo[n] {
				break
			}
		}
	}

	for k, v := range IsSumOfPreviousTwo {
		if !v {
			answer = k
		}
	}
	return answer
}

func min(i Input) int {
	min := i[0]
	for _, n := range i {
		if n < min {
			min = n
		}
	}

	return min
}

func max(i Input) int {
	max := i[0]
	for _, n := range i {
		if n > max {
			max = n
		}
	}

	return max
}

func getPartTwoAnswer(input *Input, answerPartOne int) int {
	minValue, maxValue := 0, 0

	for i, x := range *input {
		acc := x
		for j, y := range (*input)[i+1:] {

			acc += y
			if acc > answerPartOne {
				break
			}

			if acc == answerPartOne {
				minValue = min((*input)[i : i+j+2])
				maxValue = max((*input)[i : i+j+2])
				return minValue + maxValue
			}
		}
	}

	return -1
}

func main() {
	input, err := getInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	preamble := 25

	answerPartOne := getPartOneAnswer(input, preamble)
	fmt.Printf("Answer to part one: %d\n", answerPartOne)

	answerPartTwo := getPartTwoAnswer(input, answerPartOne)
	fmt.Printf("Answer to part two: %d\n", answerPartTwo)

}
