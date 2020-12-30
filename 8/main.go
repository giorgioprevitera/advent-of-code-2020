package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func getInput(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}
	defer f.Close()

	input := []string{}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		i := scanner.Text()
		input = append(input, i)
	}

	return input, nil
}

func getPartOneAnswer(input *[]string) int {

	acc := 0
	indexToExecute := 0
	executedInstructions := map[int]bool{}

	for {
		if executedInstructions[indexToExecute] {
			break
		}

		executedInstructions[indexToExecute] = true

		splitText := strings.Split((*input)[indexToExecute], " ")
		action := splitText[0]
		signedDigitValue, _ := strconv.Atoi(splitText[1])

		switch action {
		case "acc":
			acc += signedDigitValue
			indexToExecute++
		case "nop":
			indexToExecute++
		case "jmp":
			indexToExecute += signedDigitValue
		default:
			log.Fatalln("Action not supported")
		}
	}

	return acc
}

func main() {

	input, err := getInput("input.txt")
	// input, err := getInput("test-input.txt")

	if err != nil {
		log.Fatal(err)
	}

	answerToPartOne := getPartOneAnswer(&input)

	log.Printf("Answer to part one: %d\n", answerToPartOne)
}
