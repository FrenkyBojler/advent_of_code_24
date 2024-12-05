package main

import (
	"fmt"
	. "frenkybojler/adventofcode24/shared"
	"strings"

	"github.com/samber/lo"
)

type Rule struct {
	pre   int
	after int
}

func isTheSame(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, val := range a {
		if val != b[i] {
			return false
		}
	}

	return true
}

func isGreater(a, b int, rules []Rule) bool {
	for _, rule := range rules {
		if rule.pre == b && rule.after == a {
			return false
		}
	}

	return true
}

func bubbleSort(list []int, rules []Rule) []int {
	n := len(list)
	result := make([]int, n)
	copy(result, list)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if !isGreater(result[j], result[j+1], rules) {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	return result
}

func main() {
	rulesFileContent, _ := ReadFileContent("input_rules.txt")
	pagesFileContent, _ := ReadFileContent("input_pages.txt")
	pages := strings.Split(pagesFileContent, "\n")

	rules := lo.Map(strings.Split(rulesFileContent, "\n"), func(rule string, _ int) Rule {
		rule_parts := strings.Split(rule, "|")
		return Rule{
			pre:   ToInt(rule_parts[0]),
			after: ToInt(rule_parts[1]),
		}
	})

	pagesLists := lo.Map(pages, func(page string, _ int) []int {
		return lo.Map(strings.Split(page, ","), func(page_number string, _ int) int {
			return ToInt(page_number)
		})
	})

	sorted := lo.Map(pagesLists, func(page_list []int, _ int) []int {
		return bubbleSort(page_list, rules)
	})

	correctlySorted := lo.Filter(sorted, func(page_list []int, i int) bool {
		return isTheSame(page_list, pagesLists[i])
	})

	wronglySorted := lo.Filter(sorted, func(page_list []int, i int) bool {
		return !isTheSame(page_list, pagesLists[i])
	})

	sum1 := lo.Reduce(lo.Map(correctlySorted, func(page_list []int, _ int) int {
		return page_list[len(page_list)/2]
	}), func(acc, val int, _ int) int {
		return acc + val
	}, 0)

	sum2 := lo.Reduce(lo.Map(wronglySorted, func(page_list []int, _ int) int {
		return page_list[len(page_list)/2]
	}), func(acc, val int, _ int) int {
		return acc + val
	}, 0)

	fmt.Println(sum1)
	fmt.Println(sum2)
}
