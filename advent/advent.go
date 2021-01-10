package advent

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func GetInput(filename string) interface{} {
	stringInput := &[]string{}
	intInput := &[]int{}

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	isInt := true
	for scanner.Scan() {
		v := scanner.Text()
		i, err := strconv.Atoi(v)
		if err != nil {
			isInt = false
			*stringInput = append(*stringInput, v)
		} else {
			*intInput = append(*intInput, i)
		}
	}

	if isInt {
		return intInput
	}

	return stringInput
}
