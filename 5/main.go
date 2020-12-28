package main

import (
	"bufio"
	"log"
	"os"
	"sort"
)

const (
	rowMovements    = "FB"
	columnMovements = "LR"
)

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

func getPosition(characters string, positions []int, movements string) int {
	if len(positions) == 1 {
		return positions[0]
	}

	l := len(positions) / 2

	remainingPositions := []int{}
	// Either front or Left
	if characters[0] == movements[0] {
		remainingPositions = positions[:l]
	}
	// Either back or right
	if characters[0] == movements[1] {
		remainingPositions = positions[l:]
	}

	return getPosition(characters[1:], remainingPositions, movements)
}

func getSeatID(s string) int {
	rowCharacters := s[:7]
	columnCharacters := s[7:10]

	rows := initialiseSlice(128)
	columns := initialiseSlice(8)

	row := getPosition(rowCharacters, rows, rowMovements)
	column := getPosition(columnCharacters, columns, columnMovements)
	seatID := (row * 8) + column

	return seatID
}

// initialiseSlice returns a slice of int of lenght l which elements are integers in the range 0..(l-1)
func initialiseSlice(l int) []int {
	rows := make([]int, l)
	for i := 0; i < l; i++ {
		rows[i] = i
	}
	return rows
}

func findMissingID(s *[]int) int {
	for i := range *s {
		if (*s)[i]+1 != (*s)[i+1] {
			return (*s)[i] + 1
		}
	}

	return -1
}

func main() {
	input, err := getInput()
	if err != nil {
		log.Fatalln(err)
	}

	seatIDs := &[]int{}
	for _, i := range input {
		*seatIDs = append(*seatIDs, getSeatID(i))
	}

	sort.Ints(*seatIDs)

	maxID := (*seatIDs)[len(*seatIDs)-1]
	missingID := findMissingID(seatIDs)

	log.Printf("Answer to part one: %d", maxID)
	log.Printf("Answer to part two: %d", missingID)
}
