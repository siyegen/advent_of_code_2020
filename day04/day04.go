package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func validBetween(val string, min, max int) bool {
	number, _ := strconv.Atoi(val)
	if number >= min && number <= max {
		return true
	}
	return false
}

func validByr(input string) bool {
	return validBetween(input, 1920, 2002)
}

func validIyr(input string) bool {
	return validBetween(input, 2010, 2020)
}

func validEyr(input string) bool {
	return validBetween(input, 2020, 2030)
}

func validHgt(input string) bool {
	if strings.Contains(input, "cm") {
		return validBetween(input[0:len(input)-2], 150, 193)
	} else if strings.Contains(input, "in") {
		return validBetween(input[0:len(input)-2], 59, 76)
	}
	return false
}

func validHcl(input string) bool {
	found, _ := regexp.MatchString("^#[a-f0-9]{6}$", input)
	return found
}

var validEye = map[string]int{"amb": 1, "blu": 1, "brn": 1, "gry": 1, "grn": 1, "hzl": 1, "oth": 1}

func validEcl(input string) bool {
	_, ok := validEye[input]
	return ok
}

func validPID(input string) bool {
	found, _ := regexp.MatchString("^[0-9]{9}$", input)
	return found
}

func main() {
	fmt.Println("Day4")
	validators := map[string]func(string) bool{
		"byr": validByr,
		"iyr": validIyr,
		"eyr": validEyr,
		"hgt": validHgt,
		"hcl": validHcl,
		"ecl": validEcl,
		"pid": validPID,
	}

	var fileName string
	flag.StringVar(&fileName, "input", "inputs/day01", "Filename to read")
	flag.Parse()

	fh, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fh)

	completeIds := []string{}
	var partialID string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			partialID = strings.TrimSpace(partialID)
			completeIds = append(completeIds, partialID)
			partialID = ""
			continue
		}

		partialID = partialID + line + " "
	}
	partialID = strings.TrimSpace(partialID)
	completeIds = append(completeIds, partialID)

	validA := 0
	validB := 0
	for _, id := range completeIds {
		parts := strings.Split(id, " ")
		validParts := true
		numValidFields := 0
		for _, part := range parts {
			index := strings.Index(part, ":")
			field, value := part[0:index], part[index+1:len(part)]
			if fieldValidator, ok := validators[field]; ok {
				numValidFields++
				if !fieldValidator(value) {
					validParts = false
				}
			}
		}

		if numValidFields == 7 {
			validA++
		}
		if validParts && numValidFields == 7 {
			validB++
		}
	}

	fmt.Printf("ValidA: %d ValidB: %d\n", validA, validB)

}
