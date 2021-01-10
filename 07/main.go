package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/giorgioprevitera/advent-of-code-2020/advent"
)

type Input []string

type dep struct {
	number int
	bag    bag
}

type bag struct {
	color string
	deps  []dep
}

type bags []bag

func (i *Input) parse() bags {
	allBags := bags{}
	for _, l := range *i {

		c := strings.TrimSpace(strings.Split(l, "bags")[0])
		splitDeps := strings.Trim(strings.Split(l, "contain")[1], ".")
		deps := []dep{}

		splitComma := strings.Split(splitDeps, ",")

		if splitComma[0] != " no other bags" {

			for _, s := range splitComma {

				x := strings.Split(strings.TrimSpace(s), " ")

				num, _ := strconv.Atoi(x[0])
				d := dep{
					number: num,
					bag: bag{
						color: fmt.Sprintf("%s %s", x[1], x[2]),
					},
				}
				deps = append(deps, d)
			}
		}

		bag := bag{
			color: c,
			deps:  deps,
		}

		allBags = append(allBags, bag)
	}

	return allBags
}

func (bs *bags) contains(ba bag) bool {
	for _, b := range *bs {
		if ba.color == b.color {
			return true
		}
	}
	return false
}

func getPartOneAnswer(allBags *bags, color string) int {
	initialBags := bags{
		bag{
			color: color,
			deps:  []dep{},
		},
	}
	bagsContainingMyColor := bags{}
	foundBags := bags{}

	for {
		for _, bag := range initialBags {
			for _, b := range *allBags {
				if !foundBags.contains(b) && !bagsContainingMyColor.contains(b) {
					for _, dependency := range b.deps {
						if dependency.bag.color == bag.color {
							foundBags = append(foundBags, b)
							break
						}
					}
				}
			}
		}

		if len(foundBags) == 0 {
			break
		}

		bagsContainingMyColor = append(bagsContainingMyColor, foundBags...)
		initialBags = foundBags
		foundBags = bags{}
	}

	return len(bagsContainingMyColor)
}

func getPartTwoAnswer(bags *bags, color string) int {
	result := 0

	for _, outer := range *bags {
		for _, inner := range outer.deps {
			if outer.color == color {
				result += inner.number + inner.number*getPartTwoAnswer(bags, inner.bag.color)
			}
		}
	}
	return result
}

func main() {
	i, ok := advent.GetInput("input.txt").(*[]string)
	if !ok {
		log.Fatalln("Unable to make type assertion on input")
	}
	input := Input(*i)

	bags := input.parse()
	myColor := "shiny gold"

	answerPartOne := getPartOneAnswer(&bags, myColor)
	log.Printf("Answer to part one: %d", answerPartOne)

	answerPartTwo := getPartTwoAnswer(&bags, myColor)
	log.Printf("Answer to part two: %d", answerPartTwo)

}
