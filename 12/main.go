package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/giorgioprevitera/advent-of-code-2020/advent"
)

const (
	north   = 'N'
	south   = 'S'
	east    = 'E'
	west    = 'W'
	left    = 'L'
	right   = 'R'
	forward = 'F'
)

type Instruction struct {
	action rune
	value  int
}

type Instructions []Instruction

type Position struct {
	y int
	x int
}

type Navigator struct {
	position     Position
	facing       rune
	instructions *Instructions
	waypoint     Position
}

func newNavigator(i *[]string) *Navigator {
	instructions := make(Instructions, len(*i))
	for i, v := range *i {
		value, _ := strconv.Atoi(string(v[1:]))
		instructions[i].action = rune(v[0])
		instructions[i].value = value
	}
	return &Navigator{
		position:     Position{0, 0},
		facing:       east,
		instructions: &instructions,
		waypoint:     Position{1, 10},
	}
}

func (n *Navigator) rotate(direction rune, value int) {
	directions := "NESW"
	numberOfRotations := int(value / 90)

	currentDirectionIndex := 0
	for i, d := range directions {
		if n.facing == d {
			currentDirectionIndex = i
		}
	}

	newDirectionIndex := 0
	if direction == right {
		newDirectionIndex = (currentDirectionIndex + numberOfRotations) % len(directions)
	}
	if direction == left {
		newDirectionIndex = (currentDirectionIndex - numberOfRotations + len(directions)) % len(directions)
	}
	n.facing = rune(directions[newDirectionIndex])
}

func (n *Navigator) moveTowardsDirection(d rune, v int) {
	switch d {
	case north:
		n.position.y += v
	case south:
		n.position.y -= v
	case east:
		n.position.x += v
	case west:
		n.position.x -= v
	default:
	}
}

func (n *Navigator) moveWaypoint(d rune, v int) {
	switch d {
	case north:
		n.waypoint.y += v
	case south:
		n.waypoint.y -= v
	case east:
		n.waypoint.x += v
	case west:
		n.waypoint.x -= v
	default:
	}
}

func (n *Navigator) moveTowardsWaypoint(v int) {
	n.position.x += v * n.waypoint.x
	n.position.y += v * n.waypoint.y
}

func (n *Navigator) rotateWaypoint(direction rune, v int) {
	numberOfRotations := int(v / 90)

	if direction == right {
		for i := 0; i < numberOfRotations; i++ {
			n.waypoint.y, n.waypoint.x = -n.waypoint.x, n.waypoint.y
		}
	}
	if direction == left {
		for i := 0; i < numberOfRotations; i++ {
			n.waypoint.y, n.waypoint.x = n.waypoint.x, -n.waypoint.y
		}
	}

}

func (n *Navigator) navigate() {
	for _, instruction := range *n.instructions {
		switch instruction.action {
		case north, south, east, west:
			n.moveTowardsDirection(instruction.action, instruction.value)
		case forward:
			n.moveTowardsDirection(n.facing, instruction.value)
		case left, right:
			n.rotate(instruction.action, instruction.value)
		default:
		}
	}
}

func (n *Navigator) navigateWaypoint() {
	for _, instruction := range *n.instructions {
		switch instruction.action {
		case north, south, east, west:
			n.moveWaypoint(instruction.action, instruction.value)
		case forward:
			n.moveTowardsWaypoint(instruction.value)
		case left, right:
			n.rotateWaypoint(instruction.action, instruction.value)
		default:
		}
	}
}

func (n *Navigator) print() {
	fmt.Printf("x:%d, y:%d, f:%c\n", n.position.x, n.position.y, n.facing)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getPartOneAnswer(input *[]string) int {
	navigator := newNavigator(input)
	navigator.navigate()
	return abs(navigator.position.x) + abs(navigator.position.y)
}

func getPartTwoAnswer(input *[]string) int {
	navigator := newNavigator(input)
	navigator.navigateWaypoint()
	log.Printf("x:%d, y:%d\n", navigator.waypoint.x, navigator.waypoint.y)
	return abs(navigator.position.x) + abs(navigator.position.y)
}

func main() {
	input, ok := advent.GetInput("input.txt").(*[]string)
	if !ok {
		log.Fatal(ok)
	}

	answerToPartOne := getPartOneAnswer(input)
	log.Printf("Answer to part one: %d\n", answerToPartOne)

	answerToPartTwo := getPartTwoAnswer(input)
	log.Printf("Answer to part two: %d\n", answerToPartTwo)

}
