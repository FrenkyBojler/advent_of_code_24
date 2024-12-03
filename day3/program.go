package main

import (
	"fmt"
	. "frenkybojler/adventofcode24/shared"
	"regexp"

	"github.com/samber/lo"
)

func get_numbers(s string) (int, int) {
	var a, b int
	fmt.Sscanf(s, "mul(%d,%d)", &a, &b)
	return a, b
}

func part1() {
	//file_content, _ := ReadFileContent("test_input.txt")
	file_content, _ := ReadFileContent("input.txt")
	pattern := `mul\(\d{1,3},\d{1,3}\)`
	regex, _ := regexp.Compile(pattern)

	muls := lo.Map(regex.FindAllString(file_content, -1), func(s string, index int) int {
		a, b := get_numbers(s)
		return a * b
	})

	fmt.Println(lo.Sum(muls))
}

func part2() {
	//file_content, _ := ReadFileContent("test_input.txt")
	file_content, _ := ReadFileContent("input.txt")
	mul_pattern := `do\(\)|don't\(\)|mul\(\d{1,3},\d{1,3}\)`
	mul_regex, _ := regexp.Compile(mul_pattern)

	allowed_commands := mul_regex.FindAllString(file_content, -1)

	can_mul := true
	sum := 0

	for _, command := range allowed_commands {
		if command == "do()" {
			can_mul = true
		} else if command == "don't()" {
			can_mul = false
		} else {
			if can_mul {
				a, b := get_numbers(command)
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
