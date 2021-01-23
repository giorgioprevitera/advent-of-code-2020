package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const ignoredBus = -1

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
		v, _ := strconv.Atoi(c)
		if c == "x" {
			v = ignoredBus
		}
		intInput = append(intInput, v)
	}
	return target, &intInput
}

func getIndexOfMinValue(d []int) int {
	index := 0
	for i, v := range d {
		if v < d[index] && v != ignoredBus {
			index = i
		}
	}
	return index
}

func search(input *[]int, target int) (int, int) {
	busIDs := make([]int, len(*input))
	copy(busIDs, *input)

	for i := range busIDs {
		if busIDs[i] == ignoredBus {
			continue
		}
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

/*
```
time+0 % 3  == 0
time+1 % 5  == 0
time+2 % 13 == 0
```

3 is an answer for 1 -> inc := 3
start from 3: 9 is an answer for 1 and 2 (9+0%3 and 9+1%5 are equal to 0) -> inc := inc*5 = 15
start from 9: 24 is an answer for 1, 2 and 3 (24+0%3 and 24+1%5 and 24+2%13 are equal to 0) -> inc := inc*13 = 195
*/

func getPartTwoAnswer(input *[]int) int {

	increment := 1
	timestamp := (*input)[0]

	for i, b := range *input {
		if b == ignoredBus {
			continue
		}
		for (timestamp+i)%b != 0 {
			timestamp += increment
		}
		increment *= b
	}
	return timestamp
}

func main() {
	target, input := parseInput("input.txt")

	answerToPartOne := getPartOneAnswer(input, target)
	log.Printf("Answer to part one: %d\n", answerToPartOne)

	answertoPartTwo := getPartTwoAnswer(input)
	log.Printf("Answer to part two: %d\n", answertoPartTwo)

}
