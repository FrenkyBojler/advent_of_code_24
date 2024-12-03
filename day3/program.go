package main

import (
	"fmt"
	. "frenkybojler/adventofcode24/shared"
	"regexp"

	"github.com/samber/lo"
)

func getNumbers(s string) (int, int) {
	var a, b int
	fmt.Sscanf(s, "mul(%d,%d)", &a, &b)
	return a, b
}

func part1() {
	//fileContent, _ := ReadFileContent("test_input.txt")
	fileContent, _ := ReadFileContent("input.txt")
	pattern := `mul\(\d{1,3},\d{1,3}\)`
	regex, _ := regexp.Compile(pattern)

	muls := lo.Map(regex.FindAllString(fileContent, -1), func(s string, index int) int {
		a, b := getNumbers(s)
		return a * b
	})

	fmt.Println(lo.Sum(muls))
}

func part2() {
	//fileContent, _ := ReadFileContent("test_input.txt")
	fileContent, _ := ReadFileContent("input.txt")
	mulPattern := `do\(\)|don't\(\)|mul\(\d{1,3},\d{1,3}\)`
	mulRegex, _ := regexp.Compile(mulPattern)

	allowedCommands := mulRegex.FindAllString(fileContent, -1)

	canMul := true
	sum := 0

	for _, command := range allowedCommands {
		if command == "do()" {
			canMul = true
		} else if command == "don't()" {
			canMul = false
		} else {
			if canMul {
				a, b := getNumbers(command)
				sum += a * b
			}
		}
	}

	fmt.Println(sum)
}

func main() {
	part1()
	part2()
}
