package main

import (
	"log"

	"github.com/giorgioprevitera/advent-of-code-2020/advent"
)

type Input []string

type group struct {
	linesCount     int
	characterCount map[rune]int
}

type Groups map[int]*group

func countchars(s string) int {
	return 1
}

func (i *Input) parse() Groups {
	groupID := 0
	groups := Groups{}

	for _, line := range *i {

		if groups[groupID] == nil {
			groups[groupID] = &group{
				linesCount:     0,
				characterCount: make(map[rune]int),
			}
		}

		if line == "" {
			groupID++
			continue
		}

		groups[groupID].linesCount++

		for _, c := range line {
			groups[groupID].characterCount[c]++
		}
	}

	return groups
}

func getAnswerPartOne(i *Input) int {
	answer := 0
	groups := i.parse()
	for _, g := range groups {
		answer += len(g.characterCount)
	}

	return answer
}

func getAnswerPartTwo(i *Input) int {
	answer := 0
	groups := i.parse()
	for _, g := range groups {
		for _, c := range g.characterCount {
			if c == g.linesCount {
				answer++
			}
		}
	}

	return answer
}

func main() {
	i, ok := advent.GetInput("input.txt").(*[]string)
	if !ok {
		log.Fatalln("Unable to make type assertion on input")
	}
	input := Input(*i)

	answerPartOne := getAnswerPartOne(&input)
	answerPartTwo := getAnswerPartTwo(&input)

	log.Printf("Answer to part one: %d", answerPartOne)
	log.Printf("Answer to part two: %d", answerPartTwo)
}
