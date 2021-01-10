package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
)

type grid [][]rune

const (
	floor    = '.'
	empty    = 'L'
	occupied = '#'
)

func getInput(filename string) *grid {
	output := make(grid, 0)

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		s := make([]rune, len(line))
		for j, r := range line {
			s[j] = r
		}
		output = append(output, s)
	}
	return &output
}

func (g *grid) copy() *grid {
	c := make(grid, len(*g))
	for i, v := range *g {
		c[i] = make([]rune, len(v))
		copy(c[i], v)
	}
	return &c
}

func (g *grid) print() {
	fmt.Println("--------START---------")
	for _, line := range *g {
		for _, seat := range line {
			fmt.Printf("%c", seat)
		}
		fmt.Println()
	}
	fmt.Println("----------END---------")
}

func (g *grid) count(t rune) int {
	count := 0
	for _, line := range *g {
		for _, seat := range line {
			if seat == t {
				count++
			}
		}
	}
	return count

}

var deltas = [][]int{
	{-1, -1}, {-1, 0}, {-1, +1},
	{0, -1} /*{0,0}*/, {0, +1},
	{+1, -1}, {+1, 0}, {+1, +1},
}

// Int is a wrapper around the basic type `int` so we can add the `between` method to it
type Int int

func (i Int) between(a, b int) bool {
	if int(i) >= a && int(i) <= b {
		return true
	}
	return false
}

func (g *grid) applyRules(maxOccupied int, rules func(y, x, maxY, maxX int, d []int, g *grid, seatCanBeTaken bool, occupiedCount int) (bool, int)) *grid {
	new := g.copy()
	maxY := len(*g) - 1
	maxX := len((*g)[0]) - 1

	for y, line := range *g {
		for x, seat := range line {
			if seat == floor {
				continue
			}

			seatCanBeTaken := true
			occupiedCount := 0
			for _, d := range deltas {
				seatCanBeTaken, occupiedCount = rules(y, x, maxY, maxX, d, g, seatCanBeTaken, occupiedCount)
			}

			if seat == empty && seatCanBeTaken {
				(*new)[y][x] = occupied
			}
			if seat == occupied && occupiedCount >= maxOccupied {
				(*new)[y][x] = empty
			}
		}
	}
	return new
}

func partOneRules(y, x, maxY, maxX int, d []int, g *grid, seatCanBeTaken bool, occupiedCount int) (bool, int) {
	nextY := Int(y + d[0])
	nextX := Int(x + d[1])
	if nextY.between(0, maxY) && nextX.between(0, maxX) {
		nextSeat := (*g)[nextY][nextX]
		if nextSeat == occupied {
			seatCanBeTaken = false
			occupiedCount++
		}
	}
	return seatCanBeTaken, occupiedCount
}

func partTwoRules(y, x, maxY, maxX int, d []int, g *grid, seatCanBeTaken bool, occupiedCount int) (bool, int) {
	nextY := Int(y + d[0])
	nextX := Int(x + d[1])
	for nextY.between(0, maxY) && nextX.between(0, maxX) {
		nextSeat := (*g)[nextY][nextX]
		if nextSeat == occupied {
			seatCanBeTaken = false
			occupiedCount++
			break
		}
		if nextSeat == empty || occupiedCount > 5 {
			break
		}
		nextY += Int(d[0])
		nextX += Int(d[1])
	}
	return seatCanBeTaken, occupiedCount
}

func getPartOneAnswer(old *grid) int {
	maxOccupied := 4
	for {
		new := old.applyRules(maxOccupied, partOneRules)
		if reflect.DeepEqual(old, new) {
			break
		}
		old = new
	}
	return old.count(occupied)
}

func getPartTwoAnswer(old *grid) int {
	maxOccupied := 5
	for {
		new := old.applyRules(maxOccupied, partTwoRules)
		if reflect.DeepEqual(old, new) {
			break
		}
		old = new
	}
	return old.count(occupied)
}

func main() {
	input := getInput("input.txt")

	answerPartOne := getPartOneAnswer(input)
	fmt.Printf("Answer to part one: %d\n", answerPartOne)

	answerPartTwo := getPartTwoAnswer(input)
	fmt.Printf("Answer to part two: %d\n", answerPartTwo)
}
