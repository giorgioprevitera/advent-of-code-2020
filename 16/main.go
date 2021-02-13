package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type rule struct {
	name string
	min1 value
	max1 value
	min2 value
	max2 value
}

type rules []rule

type value int

type ticket []value

type tickets []ticket

func parseInput(filename string) (rules, ticket, tickets) {
	var rules rules
	var otherTickets tickets
	var myTicket ticket

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	isRule := true
	isMyTicket := false
	isOtherTickets := false

	for scanner.Scan() {
		t := scanner.Text()

		if t == "your ticket:" || t == "" {
			isRule = false
			isMyTicket = true
			continue
		}
		if t == "nearby tickets:" {
			isMyTicket = false
			isOtherTickets = true
			continue
		}

		if isRule {
			rules.parseLineAsRule(t)
		}
		if isMyTicket {
			myTicket = parseLineAsTicket(t)
		}
		if isOtherTickets {
			otherTickets.parseLineAsOtherTickets(t)
		}
	}

	return rules, myTicket, otherTickets
}

func (r *rules) parseLineAsRule(line string) {
	name := strings.Split(line, ": ")[0]
	values := strings.Split(line, ": ")[1]

	v1 := strings.Split(values, " or ")[0]
	v1min, _ := strconv.Atoi(strings.Split(v1, "-")[0])
	v1max, _ := strconv.Atoi(strings.Split(v1, "-")[1])

	v2 := strings.Split(values, " or ")[1]
	v2min, _ := strconv.Atoi(strings.Split(v2, "-")[0])
	v2max, _ := strconv.Atoi(strings.Split(v2, "-")[1])

	*r = append(*r, rule{name, value(v1min), value(v1max), value(v2min), value(v2max)})
}
func parseLineAsTicket(line string) ticket {
	values := strings.Split(line, ",")

	ticket := make(ticket, len(values))

	for i, v := range values {
		n, _ := strconv.Atoi(v)
		ticket[i] = value(n)
	}
	return ticket
}

func (t *tickets) parseLineAsOtherTickets(line string) {
	values := strings.Split(line, ",")

	ticket := make(ticket, len(values))

	for i, v := range values {
		n, _ := strconv.Atoi(v)
		ticket[i] = value(n)
	}

	*t = append(*t, ticket)
}

func (v value) isValid(r rule) bool {
	return (v >= r.min1 && v <= r.max1) || (v >= r.min2 && v <= r.max2)
}

func getValidTickets(t tickets, rules rules) tickets {
	validTickets := tickets{}
	for _, ticket := range t {
		ticketValid := true
		for _, v := range ticket {
			valueValid := false
			for _, r := range rules {
				if v.isValid(r) {
					valueValid = true
					break
				}
			}
			if !valueValid {
				ticketValid = false
			}
		}
		if ticketValid {
			validTickets = append(validTickets, ticket)
		}
	}
	return validTickets
}

type possibleFields map[string][]int

func getPossibleFields(validTickets tickets, rules rules) possibleFields {
	// 0: possible
	// 1: impossible
	p := possibleFields{}
	for _, r := range rules {
		p[r.name] = make([]int, len(validTickets[0]))
	}

	for _, ticket := range validTickets {
		for i, v := range ticket {
			for _, r := range rules {
				if !v.isValid(r) {
					p[r.name][i] = 1
				}
			}
		}
	}

	return p
}

func (p possibleFields) deletePosition(i int) {
	for n := range p {
		p[n][i] = 1
	}
}

// Returns the first, and only, element with a single possible field
// 0: possible
// 1: impossible
func (p possibleFields) getSingleValidElement() (string, int) {
	for k, v := range p {
		count := 0
		index := 0
		for i, n := range v {
			if n == 0 {
				count++
				index = i
			}
		}
		if count == 1 {
			return k, index
		}
	}

	return "", -1
}

func getPartOnwAnswer(rules rules, otherTickets tickets) int {
	invalidValues := []value{}

	for _, ticket := range otherTickets {
		for _, value := range ticket {
			invalid := true
			for _, rule := range rules {
				if value.isValid(rule) {
					invalid = false
					break
				}
			}
			if invalid {
				invalidValues = append(invalidValues, value)
			}
		}
	}

	answer := 0
	for _, v := range invalidValues {
		answer += int(v)
	}

	return answer
}

func getPartTwoAnswer(rules rules, myTicket ticket, otherTickets tickets) int {
	answer := 1

	validTickets := getValidTickets(otherTickets, rules)
	fields := getPossibleFields(validTickets, rules)

	numberOfFields := len(validTickets[0])
	for numberOfFields > 0 {
		currentName, currentPosition := fields.getSingleValidElement()
		if strings.HasPrefix(currentName, "departure") {
			answer *= int(myTicket[currentPosition])
		}
		fields.deletePosition(currentPosition)
		numberOfFields--
	}

	return answer
}

func main() {
	rules, myTicket, otherTickets := parseInput("input.txt")

	answerToPartOne := getPartOnwAnswer(rules, otherTickets)
	log.Println("Answer to part one:", answerToPartOne)

	answerToPartTwo := getPartTwoAnswer(rules, myTicket, otherTickets)
	log.Println("Answer to part two:", answerToPartTwo)
}
