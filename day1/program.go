package main

import (
	"fmt"
	. "frenkybojler/adventofcode24/shared"
	"sort"
	"strings"

	"github.com/samber/lo"
)

func part1() {
	// testInput, _ := ReadFileContent("test_input.txt")
	input1, _ := ReadFileContent("input1.txt")
	withoutSpaces := strings.ReplaceAll(input1, "   ", "|")
	split := strings.Split(withoutSpaces, "\n")

	leftList := []int{}
	rightList := []int{}

	for _, s := range split {
		split := strings.Split(s, "|")
		leftList = append(leftList, ToInt(split[0]))
		rightList = append(rightList, ToInt(split[1]))
	}

	sort.Slice(leftList, func(i, j int) bool {
		return leftList[i] <= leftList[j]
	})

	sort.Slice(rightList, func(i, j int) bool {
		return rightList[i] <= rightList[j]
	})

	distances := lo.Map(leftList, func(s int, index int) int {
		return Abs(s - rightList[index])
	})

	sum := lo.Reduce(distances, func(acc int, curr int, index int) int {
		return acc + curr
	}, 0.0)

	fmt.Println(sum)
}

func part2() {
	//input1, _ := ReadFileContent("test_input.txt")
	input1, _ := ReadFileContent("input1.txt")
	withoutSpaces := strings.ReplaceAll(input1, "   ", "|")
	split := strings.Split(withoutSpaces, "\n")

	leftList := []int{}
	rightList := []int{}

	for _, s := range split {
		split := strings.Split(s, "|")
		leftList = append(leftList, ToInt(split[0]))
		rightList = append(rightList, ToInt(split[1]))
	}

	rightListGrouped := lo.GroupBy(rightList, func(s int) int {
		return s
	})

	lengths := lo.MapEntries(rightListGrouped, func(key int, values []int) (int, int) {
		return key, len(values)
	})

	similarityScores := lo.Map(leftList, func(s int, index int) int {
		return lengths[s] * s
	})

	sum := lo.Reduce(similarityScores, func(acc int, curr int, index int) int {
		return acc + curr
	}, 0)

	fmt.Println(sum)
}

func main() {
	part1()
	part2()
}
