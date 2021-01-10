package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/giorgioprevitera/advent-of-code-2020/advent"
)

func getPartOneAnswer(input *[]string) int {
	acc := 0
	currentActionIndex := 0
	executedAction := map[int]bool{}

	for {
		if executedAction[currentActionIndex] {
			break
		}

		executedAction[currentActionIndex] = true

		splitText := strings.Split((*input)[currentActionIndex], " ")
		action := splitText[0]
		value, _ := strconv.Atoi(splitText[1])

		switch action {
		case "acc":
			acc += value
			currentActionIndex++
		case "nop":
			currentActionIndex++
		case "jmp":
			currentActionIndex += value
		default:
			log.Fatalln("Action not supported")
		}
	}

	return acc
}

func getPartTwoAnswer(input *[]string) int {
	acc := 0
	currentActionIndex := 0
	executedAction := map[int]bool{}
	nopHasBeenSwapped := map[int]bool{}
	jmpHasBeenSwapped := map[int]bool{}
	dontSwap := false

	for {
		if executedAction[currentActionIndex] {
			currentActionIndex = 0
			acc = 0
			dontSwap = false
			executedAction = map[int]bool{}
		}

		executedAction[currentActionIndex] = true

		// parse line
		splitText := strings.Split((*input)[currentActionIndex], " ")
		action := splitText[0]
		value, _ := strconv.Atoi(splitText[1])

		switch action {
		case "acc":
			acc += value
			currentActionIndex++
		case "nop":
			if dontSwap || nopHasBeenSwapped[currentActionIndex] {
				// treat as nop
				currentActionIndex++
			} else {
				// treat as jmp and record the swap
				nopHasBeenSwapped[currentActionIndex] = true
				currentActionIndex += value
				dontSwap = true
			}
		case "jmp":
			if dontSwap || jmpHasBeenSwapped[currentActionIndex] {
				// treat as jmp
				currentActionIndex += value
			} else {
				// treat as nop and record the swap
				jmpHasBeenSwapped[currentActionIndex] = true
				currentActionIndex++
				dontSwap = true
			}
		default:
			log.Fatalln("Action not supported")
		}

		if currentActionIndex == len(*input) {
			break
		}
	}

	return acc
}

func main() {
	input, ok := advent.GetInput("input.txt").(*[]string)
	if !ok {
		log.Fatalln("Unable to make type assertion on input")
	}

	answerToPartOne := getPartOneAnswer(input)
	log.Printf("Answer to part one: %d\n", answerToPartOne)

	answerToPartTwo := getPartTwoAnswer(input)
	log.Printf("Answer to part two: %d\n", answerToPartTwo)
}
