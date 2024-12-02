package main

import (
	"fmt"
	. "frenkybojler/adventofcode24/shared"
	"strings"

	"github.com/samber/lo"
)

func part1() {
	//file_content, _ := ReadFileContent("test_input.txt")
	file_content, _ := ReadFileContent("input.txt")
	reports := lo.Map(strings.Split(file_content, "\n"), func(s string, index int) []int {
		return ToIntSlice(strings.Split(s, " "))
	})

	safety_result := lo.Map(reports, func(report []int, index int) bool {
		is_decreasing := report[0] > report[1]

		for index, i := range report {

			if index == len(report)-1 {
				break
			}

			if i == report[index+1] {
				return false
			}

			if is_decreasing && i > report[index+1] && Abs(i-report[index+1]) <= 3 {
				continue
			}

			if !is_decreasing && i < report[index+1] && Abs(i-report[index+1]) <= 3 {
				continue
			}

			return false
		}

		return true
	})

	fmt.Println(lo.Sum(lo.Map(safety_result, func(s bool, index int) int {
		if s {
			return 1
		}
		return 0
	})))
}

func check_report(report []int, return_on_fail bool) (bool, []int) {
	is_decreasing := report[0] > report[1]
	issues_indexes := []int{}

	for index, i := range report {
		if return_on_fail && len(issues_indexes) > 1 {
			return false, issues_indexes
		}

		if index == len(report)-1 {
			break
		}

		if i == report[index+1] {
			issues_indexes = append(issues_indexes, index)
			continue
		}

		if is_decreasing && i > report[index+1] && Abs(i-report[index+1]) <= 3 {
			continue
		}

		if !is_decreasing && i < report[index+1] && Abs(i-report[index+1]) <= 3 {
			continue
		}

		issues_indexes = append(issues_indexes, index)
	}

	return len(issues_indexes) == 0, issues_indexes
}

func avg_level(report []int) int {
	avg := 0
	for _, i := range report {
		avg += i
	}

	return avg / len(report)
}

func part2() {
	//file_content, _ := ReadFileContent("test_input.txt")
	file_content, _ := ReadFileContent("input.txt")
	reports := lo.Map(strings.Split(file_content, "\n"), func(s string, index int) []int {
		return ToIntSlice(strings.Split(s, " "))
	})

	safety_result := lo.Map(reports, func(report []int, index int) bool {
		level_value_issue := []int{}

		for level_index, i := range report {
			if abs := Abs(i - avg_level(report)); abs >= 4 {
				level_value_issue = append(level_value_issue, level_index)
			}
		}

		result, direction_issues := check_report(report, false)

		if !result {
			issues_indexes := lo.Uniq(append(direction_issues, level_value_issue...))
			issues_indexes = lo.Uniq(append(issues_indexes, lo.FlatMap(issues_indexes, func(i int, index int) []int {
				return []int{lo.Min([]int{i + 1, len(report) - 1}), lo.Max([]int{i - 1, 0})}
			})...))

			for _, i := range issues_indexes {
				report_to_check := lo.DropByIndex(report, i)
				current_result, _ := check_report(report_to_check, true)

				if current_result {
					return true
				}
			}

			return false
		}

		return true
	})

	fmt.Println(lo.Sum(lo.Map(safety_result, func(s bool, index int) int {
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
