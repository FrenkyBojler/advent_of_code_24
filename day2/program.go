package main

import (
	"fmt"
	. "frenkybojler/adventofcode24/shared"
	"strings"

	"github.com/samber/lo"
)

func part1() {
	//fileContent, _ := ReadFileContent("test_input.txt")
	fileContent, _ := ReadFileContent("input.txt")
	reports := lo.Map(strings.Split(fileContent, "\n"), func(s string, index int) []int {
		return ToIntSlice(strings.Split(s, " "))
	})

	safetyResult := lo.Map(reports, func(report []int, index int) bool {
		isDecreasing := report[0] > report[1]

		for index, i := range report {

			if index == len(report)-1 {
				break
			}

			if i == report[index+1] {
				return false
			}

			if isDecreasing && i > report[index+1] && Abs(i-report[index+1]) <= 3 {
				continue
			}

			if !isDecreasing && i < report[index+1] && Abs(i-report[index+1]) <= 3 {
				continue
			}

			return false
		}

		return true
	})

	fmt.Println(lo.Sum(lo.Map(safetyResult, func(s bool, index int) int {
		if s {
			return 1
		}
		return 0
	})))
}

func checkReport(report []int, returnOnFail bool) (bool, []int) {
	isDecreasing := report[0] > report[1]
	issuesIndexes := []int{}

	for index, i := range report {
		if returnOnFail && len(issuesIndexes) > 1 {
			return false, issuesIndexes
		}

		if index == len(report)-1 {
			break
		}

		if i == report[index+1] {
			issuesIndexes = append(issuesIndexes, index)
			continue
		}

		if isDecreasing && i > report[index+1] && Abs(i-report[index+1]) <= 3 {
			continue
		}

		if !isDecreasing && i < report[index+1] && Abs(i-report[index+1]) <= 3 {
			continue
		}

		issuesIndexes = append(issuesIndexes, index)
	}

	return len(issuesIndexes) == 0, issuesIndexes
}

func avgLevel(report []int) int {
	avg := 0
	for _, i := range report {
		avg += i
	}

	return avg / len(report)
}

func part2() {
	//fileContent, _ := ReadFileContent("test_input.txt")
	fileContent, _ := ReadFileContent("input.txt")
	reports := lo.Map(strings.Split(fileContent, "\n"), func(s string, index int) []int {
		return ToIntSlice(strings.Split(s, " "))
	})

	safetyResult := lo.Map(reports, func(report []int, index int) bool {
		levelValueIssue := []int{}

		for levelIndex, i := range report {
			if abs := Abs(i - avgLevel(report)); abs >= 4 {
				levelValueIssue = append(levelValueIssue, levelIndex)
			}
		}

		result, directionIssues := checkReport(report, false)

		if !result {
			issuesIndexes := lo.Uniq(append(directionIssues, levelValueIssue...))
			issuesIndexes = lo.Uniq(append(issuesIndexes, lo.FlatMap(issuesIndexes, func(i int, index int) []int {
				return []int{lo.Min([]int{i + 1, len(report) - 1}), lo.Max([]int{i - 1, 0})}
			})...))

			for _, i := range issuesIndexes {
				reportToCheck := lo.DropByIndex(report, i)
				currentResult, _ := checkReport(reportToCheck, true)

				if currentResult {
					return true
				}
			}

			return false
		}

		return true
	})

	fmt.Println(lo.Sum(lo.Map(safetyResult, func(s bool, index int) int {
		if s {
			return 1
		}
		return 0
	})))
}

func main() {
	part1()
	part2()
}
