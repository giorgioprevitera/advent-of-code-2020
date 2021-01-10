package main

import (
	"errors"
	"log"

	"github.com/giorgioprevitera/advent-of-code-2020/advent"
)

func searchAnswerPartOne(input *[]int, sum int) (int, error) {
	for i, x := range *input {
		for _, y := range (*input)[i:] {
			if x+y == sum {
				return x * y, nil
			}
		}
	}

	return -1, errors.New("Answer not found")
}

func searchAnswerPartTwo(input *[]int, sum int) (int, error) {
	for i, x := range *input {
		for j, y := range (*input)[i:] {
			for _, z := range (*input)[j:] {
				if x+y+z == sum {
					return x * y * z, nil
				}
			}
		}
	}

	return -1, errors.New("Answer not found")
}

func main() {
	input, ok := advent.GetInput("input.txt").(*[]int)
	if !ok {
		log.Fatalln("Unable to make type assertion on input")
	}

	answerPartOne, err := searchAnswerPartOne(input, 2020)
	answerPartTwo, err := searchAnswerPartTwo(input, 2020)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Answer to part one is %d ", answerPartOne)
	log.Printf("Answer to part two is %d ", answerPartTwo)
}
