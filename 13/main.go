package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseInput(filename string) (int, *[]int) {
	var intInput []int
	var target int

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	scanner.Scan()
	target, _ = strconv.Atoi(scanner.Text())

	scanner.Scan()
	toParse := scanner.Text()
	x := strings.Split(toParse, ",")

	for _, c := range x {
		if c == "x" {
			continue
		}
		v, _ := strconv.Atoi(c)
		intInput = append(intInput, v)
	}
	return target, &intInput
}

func getIndexOfMinValue(d []int) int {
	index := 0
	for i, v := range d {
		if v < d[index] {
			index = i
		}
	}
	return index
}

func search(input *[]int, target int) (int, int) {
	busIDs := make([]int, len(*input))
	copy(busIDs, *input)

	for i := range busIDs {
		for busIDs[i] < target {
			busIDs[i] = busIDs[i] + (*input)[i]
		}
	}

	indexOfMinValue := getIndexOfMinValue(busIDs)
	return busIDs[indexOfMinValue], (*input)[indexOfMinValue]
}

func getPartOneAnswer(input *[]int, target int) int {
	timestamp, busID := search(input, target)
	return (timestamp - target) * busID
}

func main() {
	target, input := parseInput("input.txt")

	answerToPartOne := getPartOneAnswer(input, target)
	log.Printf("Answer to part one: %d\n", answerToPartOne)
}
