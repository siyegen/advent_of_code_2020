package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day1")

	var fileName string
	flag.StringVar(&fileName, "input", "inputs/day01", "Filename to read")
	flag.Parse()

	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	data := string(bytes)
	lines := strings.Split(data, "\n")

	partA(lines)
	partB(lines)
}

func partA(lines []string) {
	numbers := map[int]struct{}{}
	for _, line := range lines {
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		numbers[num] = struct{}{}
		if _, ok := numbers[2020-num]; ok {
			fmt.Println("Part A", num*(2020-num))
		}
	}
}

func partB(lines []string) {
	part1Numbers := map[int]struct{}{}
	part2Numbers := map[int]int{}
	for _, line := range lines {
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		part1Numbers[num] = struct{}{}
		possibleNum := 2020 - num
		part2Numbers[possibleNum] = num

		for part1 := range part1Numbers {
			toFind := part1 + num
			if val, ok := part2Numbers[toFind]; ok {
				fmt.Println("Part B", num*part1*val)
				return
			}
		}
	}
}
