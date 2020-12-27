package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	partOne = 1
	partTwo = 2
)

var functionsMap = map[int]func(string) bool{
	partOne: inputContainsValidPasswordPartOne,
	partTwo: inputContainsValidPasswordPartTwo,
}

func getInput() ([]string, error) {
	f, err := os.Open("input.txt")
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

func inputContainsValidPasswordPartOne(i string) bool {
	data := strings.Split(i, " ")

	digits := strings.Split(data[0], "-")
	character := string(data[1][0])
	password := data[2]

	minDigit, _ := strconv.Atoi(digits[0])
	maxDigit, _ := strconv.Atoi(digits[1])

	characterCount := strings.Count(password, character)
	if minDigit <= characterCount && characterCount <= maxDigit {
		return true
	}

	return false
}

func inputContainsValidPasswordPartTwo(i string) bool {
	data := strings.Split(i, " ")

	digits := strings.Split(data[0], "-")
	character := data[1][0]
	password := data[2]

	minDigit, _ := strconv.Atoi(digits[0])
	maxDigit, _ := strconv.Atoi(digits[1])

	if password[minDigit-1] == character && password[maxDigit-1] == character {
		return false
	}

	if password[minDigit-1] != character && password[maxDigit-1] != character {
		return false
	}

	return true
}

func countValidPasswords(input []string, part int) int {
	validPasswordsCount := 0

	for _, i := range input {
		if functionsMap[part](i) {
			validPasswordsCount++
		}
	}

	return validPasswordsCount
}

func main() {
	input, err := getInput()
	if err != nil {
		log.Fatalln(err)
	}

	answerPartOne := countValidPasswords(input, partOne)
	answerPartTwo := countValidPasswords(input, partTwo)

	log.Printf("Answer part one: %d\n", answerPartOne)
	log.Printf("Answer part two: %d\n", answerPartTwo)
}
