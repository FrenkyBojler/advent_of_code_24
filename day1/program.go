package main

import (
	"fmt"
	. "frenkybojler/adventofcode24/shared"
	"sort"
	"strings"

	"github.com/samber/lo"
)

func part1() {
	// test_input, _ := ReadFileContent("test_input.txt")
	input1, _ := ReadFileContent("input1.txt")
	without_spaces := strings.ReplaceAll(input1, "   ", "|")
	split := strings.Split(without_spaces, "\n")

	left_list := []int{}
	right_list := []int{}

	for _, s := range split {
		split := strings.Split(s, "|")
		left_list = append(left_list, ToInt(split[0]))
		right_list = append(right_list, ToInt(split[1]))
	}

	sort.Slice(left_list, func(i, j int) bool {
		return left_list[i] <= left_list[j]
	})

	sort.Slice(right_list, func(i, j int) bool {
		return right_list[i] <= right_list[j]
	})

	distances := lo.Map(left_list, func(s int, index int) int {
		return Abs(s - right_list[index])
	})

	sum := lo.Reduce(distances, func(acc int, curr int, index int) int {
		return acc + curr
	}, 0.0)

	fmt.Println(sum)
}

func part2() {
	//input1, _ := ReadFileContent("test_input.txt")
	input1, _ := ReadFileContent("input1.txt")
	without_spaces := strings.ReplaceAll(input1, "   ", "|")
	split := strings.Split(without_spaces, "\n")

	left_list := []int{}
	right_list := []int{}

	for _, s := range split {
		split := strings.Split(s, "|")
		left_list = append(left_list, ToInt(split[0]))
		right_list = append(right_list, ToInt(split[1]))
	}

	right_list_grouped := lo.GroupBy(right_list, func(s int) int {
		return s
	})

	lenghts := lo.MapEntries(right_list_grouped, func(key int, values []int) (int, int) {
		return key, len(values)
	})

	similarity_scores := lo.Map(left_list, func(s int, index int) int {
		return lenghts[s] * s
	})

	sum := lo.Reduce(similarity_scores, func(acc int, curr int, index int) int {
		return acc + curr
	}, 0)

	fmt.Println(sum)
}

func main() {
	part1()
	part2()
}
