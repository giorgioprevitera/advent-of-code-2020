package main

import (
	"bufio"
	"log"
	"os"
)

type Slope [2]int

type Slopes []Slope

var slopesPartOne = Slopes{
	{3, 1},
}

var slopesPartTwo = Slopes{
	{1, 1},
	{3, 1},
	{5, 1},
	{7, 1},
	{1, 2},
}

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

func updateCurrentPosition(p *int, i, l int) {
	if *p+i < l {
		*p += i
	} else {
		*p = i - (l - *p)
	}
}

func countTrees(input []string, slopes Slopes) int {
	maxLength := len(input[0])
	total := 1

	for _, slope := range slopes {
		currentPosition := 0
		numberOfTrees := 0
		for i, l := range input {
			if i%slope[1] != 0 {
				continue
			}
			if l[currentPosition] == '#' {
				numberOfTrees++
			}
			updateCurrentPosition(&currentPosition, slope[0], maxLength)
		}
		total = total * numberOfTrees
	}

	return total
}

func main() {
	input, err := getInput("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	answerPartOne := countTrees(input, slopesPartOne)
	answerPartTwo := countTrees(input, slopesPartTwo)

	log.Printf("Answer part one: %d\n", answerPartOne)
	log.Printf("Answer part two: %d\n", answerPartTwo)
}
