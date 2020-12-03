package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// type policy struct {
// 	min    int
// 	max    int
// 	letter string
// }

func main() {
	fmt.Println("Day2")

	var fileName string
	flag.StringVar(&fileName, "input", "inputs/day02", "Filename to read")
	flag.Parse()

	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	data := string(bytes)

	lines := strings.Split(data, "\n")

	validAPasswords := 0
	validBPasswords := 0
	for _, line := range lines {
		input := strings.Split(line, ": ")
		policyPart, password := strings.Trim(input[0], " "), strings.Trim(input[1], " ")
		policyIndex := strings.Index(policyPart, " ")
		policyStr, letter := policyPart[0:policyIndex], policyPart[policyIndex+1:len(policyPart)]

		hyIndex := strings.Index(policyStr, "-")
		firstStr, secondStr := policyStr[0:hyIndex], policyStr[hyIndex+1:len(policyStr)]
		firstNum, _ := strconv.Atoi(firstStr)
		secondNum, _ := strconv.Atoi(secondStr)
		if partA(password, letter, firstNum, secondNum) {
			validAPasswords++
		}
		if partB(password, letter, firstNum, secondNum) {
			validBPasswords++
		}
	}

	fmt.Println("Part A: ", validAPasswords)
	fmt.Println("Part B: ", validBPasswords)
}

func partA(password, letter string, min, max int) bool {
	count := strings.Count(password, letter)
	if count >= min && count <= max {
		return true
	}
	return false
}

func partB(password, letter string, firstPos, secondPos int) bool {
	if (string(password[firstPos-1]) == letter) != (string(password[secondPos-1]) == letter) {
		return true
	}
	return false
}
