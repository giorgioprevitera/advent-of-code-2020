package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/giorgioprevitera/advent-of-code-2020/advent"
	"github.com/golang-collections/collections/stack"
)

type bitmask string
type memory map[int64]int64

func prettyPrint(n int) {
	fmt.Printf("%036b - %d\n", n, n)
}

func (b bitmask) applyPartOne(n int64) int64 {
	l := len(b) - 1

	for i, v := range b {
		position := l - i
		switch v {
		case '0':
			n = setBitToZero(n, position)
		case '1':
			n = setBitToOne(n, position)
		default:
		}
	}
	return n
}

func (b bitmask) applyPartTwo(r int64) string {
	startRegister := []byte(fmt.Sprintf("%036b", r))
	for i, v := range b {
		switch v {
		case '1':
			startRegister[i] = '1'
		case 'X':
			startRegister[i] = 'X'
		}
	}
	return string(startRegister)
}

func setBitToOne(n int64, p int) int64 {
	return n | (1 << p)
}

func setBitToZero(n int64, p int) int64 {
	return n & (n ^ (1 << p))
}

func getMask(m string) bitmask {
	return bitmask(strings.Split(m, " = ")[1])
}

func parseInstruction(i string) (register, value int64) {
	splitStrings := strings.Split(i, " = ")

	v, err := strconv.Atoi(splitStrings[1])
	if err != nil {
		log.Fatal(err)
	}

	registerString := strings.TrimSuffix(strings.TrimPrefix(splitStrings[0], "mem["), "]")

	r, err := strconv.Atoi(registerString)
	if err != nil {
		log.Fatal(err)
	}

	value = int64(v)
	register = int64(r)

	return
}

func (m *memory) sum() (sum int64) {
	for _, v := range *m {
		sum += v
	}
	return
}

func getPartOneAnswer(input *[]string) int64 {
	var mask bitmask
	mem := make(memory, 0)

	for _, line := range *input {
		if strings.HasPrefix(line, "mask") {
			mask = getMask(line)
			continue
		}
		if strings.HasPrefix(line, "mem") {
			register, value := parseInstruction(line)
			mem[register] = mask.applyPartOne(value)
		}
	}

	answer := mem.sum()
	return answer
}

func (m *memory) writeFloating(address, value int64, b bitmask) {
	startAddress := b.applyPartTwo(address)

	s := stack.New()
	s.Push(startAddress)

	for s.Len() > 0 {
		a := s.Pop().(string)

		if strings.Contains(a, "X") {
			s.Push(strings.Replace(a, "X", "1", 1))
			s.Push(strings.Replace(a, "X", "0", 1))
		} else {
			reg, _ := strconv.ParseInt(a, 2, 64)
			(*m)[reg] = value
		}
	}
}

func getPartTwoAnswer(input *[]string) int64 {
	var mask bitmask
	mem := make(memory, 0)

	for _, line := range *input {
		if strings.HasPrefix(line, "mask") {
			mask = getMask(line)
			continue
		}
		if strings.HasPrefix(line, "mem") {
			address, value := parseInstruction(line)
			mem.writeFloating(address, value, mask)
		}
	}

	answer := mem.sum()
	return answer
}

func main() {
	input, ok := advent.GetInput("input.txt").(*[]string)
	if !ok {
		log.Fatal("Unable to make type assertion on input")
	}

	answerToPartOne := getPartOneAnswer(input)
	log.Printf("Answer to part one: %d\n", answerToPartOne)

	answerToPartTwo := getPartTwoAnswer(input)
	log.Printf("Answer to part two: %d\n", answerToPartTwo)

}
