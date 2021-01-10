package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/giorgioprevitera/advent-of-code-2020/advent"
)

type Input []string

type Passport map[string]string

type Passports map[int]Passport

type ValidPassportKeys []string

var validPassportKeys = ValidPassportKeys{
	"byr",
	"ecl",
	"eyr",
	"hcl",
	"hgt",
	"iyr",
	"pid",
}

var validEyeColors = []string{
	"amb",
	"blu",
	"brn",
	"gry",
	"grn",
	"hzl",
	"oth",
}

// contains returns true if string s is contained in []string v
func contains(v *[]string, s string) bool {
	for _, k := range *v {
		if k == s {
			return true
		}
	}

	return false
}

func (i *Input) parse() *Passports {
	passportID := 0
	passports := make(Passports)

	for _, line := range *i {
		if passports[passportID] == nil {
			passports[passportID] = make(Passport)
		}

		if line == "" {
			passportID++
			continue
		}

		fields := strings.Split(line, " ")
		for _, field := range fields {
			splitFields := strings.Split(field, ":")
			passports[passportID][splitFields[0]] = splitFields[1]
		}
	}

	return &passports
}

func (p *Passport) getKeys() *[]string {
	keys := &[]string{}
	for k := range *p {
		*keys = append(*keys, k)
	}
	return keys
}

func (p *Passport) haveValidKeys() bool {
	for _, k := range validPassportKeys {
		if !contains(p.getKeys(), k) {
			return false
		}
	}
	return true
}

func (p *Passport) isValid() bool {

	if !p.haveValidKeys() {
		return false
	}

	// Validate byr
	byr, _ := strconv.Atoi((*p)["byr"])
	if byr < 1920 || byr > 2002 {
		return false
	}

	// Validate iyr
	iyr, _ := strconv.Atoi((*p)["iyr"])
	if iyr < 2010 || iyr > 2020 {
		return false
	}

	// Validate eyr
	eyr, _ := strconv.Atoi((*p)["eyr"])
	if eyr < 2020 || eyr > 2030 {
		return false
	}

	// Validate hgt
	hgt := (*p)["hgt"]
	if strings.HasSuffix(hgt, "cm") {
		h, _ := strconv.Atoi(strings.TrimSuffix(hgt, "cm"))
		if h < 150 || h > 193 {
			return false
		}
	} else {
		if strings.HasSuffix(hgt, "in") {
			h, _ := strconv.Atoi(strings.TrimSuffix(hgt, "in"))
			if h < 59 || h > 76 {
				return false
			}
		} else {
			return false
		}
	}

	// Validate hcl
	hcl := (*p)["hcl"]
	m, err := regexp.MatchString("^#[0-9a-f]{6}$", hcl)
	if err != nil {
		log.Fatal(err)
	}
	if !m {
		return false
	}

	// Validate ecl
	ecl := (*p)["ecl"]
	if !contains(&validEyeColors, ecl) {
		return false
	}

	// Validate pid
	pid := (*p)["pid"]
	m, err = regexp.MatchString("^[0-9]{9}$", pid)
	if err != nil {
		log.Fatal(err)
	}
	if !m {
		return false
	}

	return true
}

func main() {
	i, ok := advent.GetInput("input.txt").(*[]string)
	if !ok {
		log.Fatalln("Unable to make type assertion on input")
	}
	input := Input(*i)

	passports := input.parse()

	validPassportsPartOne := 0
	validPassportsPartTwo := 0
	for _, passport := range *passports {
		if passport.haveValidKeys() {
			validPassportsPartOne++
		}
		if passport.isValid() {
			validPassportsPartTwo++
		}
	}

	log.Printf("Answer part one: %d\n", validPassportsPartOne)
	log.Printf("Answer part two: %d\n", validPassportsPartTwo)
}
