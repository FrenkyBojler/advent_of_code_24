package main

import (
	. "frenkybojler/adventofcode24/shared"
	"regexp"
	"strings"

	"github.com/samber/lo"
)

func checkLine(regex regexp.Regexp, line string) int {
	return len(regex.FindAllString(line, -1))
}

func getRow(line []string, backwards bool) string {
	lineString := strings.Join(line, "")

	if backwards {
		lineString = ReverseString(lineString)
	}

	return lineString
}

func getColumn(matrix [][]string, column int, backwards bool) string {
	columnString := ""

	for _, row := range matrix {
		columnString += row[column]
	}

	if backwards {
		columnString = ReverseString(columnString)
	}

	return columnString
}

func getDiagonal(matrix [][]string, x int, y int, length int, cross bool, backwards bool) string {
	diagonalString := ""

	if cross {
		for i := 0; i < length; i++ {
			if x+i >= len(matrix) || y-i < 0 {
				break
			}

			diagonalString += matrix[x+i][y-i]
		}
	} else {
		for i := 0; i < length; i++ {
			if x+i >= len(matrix) || y+i >= len(matrix[0]) {
				break
			}

			diagonalString += matrix[x+i][y+i]
		}
	}

	if backwards {
		diagonalString = ReverseString(diagonalString)
	}

	return diagonalString
}

func getMatrix(content string) [][]string {
	lines := strings.Split(content, "\n")
	matrixOfLetters := [][]string{}

	for y, line := range lines {
		matrixOfLetters = append(matrixOfLetters, []string{})
		for _, letter := range line {
			matrixOfLetters[y] = append(matrixOfLetters[y], string(letter))
		}
	}

	return matrixOfLetters
}

func part1() {
	regex, _ := regexp.Compile(`XMAS`)
	minDiagonalLength := 4

	fileContent, _ := ReadFileContent("input.txt")
	matrix := getMatrix(fileContent)

	stringsToTest := []string{}

	// get all lines
	for _, row := range matrix {
		stringsToTest = append(stringsToTest, getRow(row, false))
		stringsToTest = append(stringsToTest, getRow(row, true))
	}

	// get all columns
	for i := 0; i < len(matrix[0]); i++ {
		stringsToTest = append(stringsToTest, getColumn(matrix, i, false))
		stringsToTest = append(stringsToTest, getColumn(matrix, i, true))
	}

	// get all diagonals
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			stringsToTest = append(stringsToTest, getDiagonal(matrix, i, j, minDiagonalLength, false, false))
			stringsToTest = append(stringsToTest, getDiagonal(matrix, i, j, minDiagonalLength, false, true))
			stringsToTest = append(stringsToTest, getDiagonal(matrix, i, j, minDiagonalLength, true, false))
			stringsToTest = append(stringsToTest, getDiagonal(matrix, i, j, minDiagonalLength, true, true))
		}
	}

	stringsToTest = lo.Filter(stringsToTest, func(s string, _ int) bool {
		return s != ""
	})

	// count all occurences of XMAS
	occurences := lo.Sum(lo.Map(stringsToTest, func(s string, _ int) int {
		return checkLine(*regex, s)
	}))

	println(occurences)
}

func getMatrixSplitByLength(matrix [][]string, length int) [][][]string {
	matrix_split := [][][]string{}

	for i := 0; i < len(matrix); i++ {
		if i+length > len(matrix) {
			break
		}

		for j := 0; j < len(matrix[0]); j++ {

			if j+length > len(matrix[0]) {
				break
			}

			matrix_slice := [][]string{}

			for k := 0; k < length; k++ {
				list := []string{}

				for l := 0; l < length; l++ {
					list = append(list, matrix[i+k][j+l])
				}

				matrix_slice = append(matrix_slice, list)
			}

			matrix_split = append(matrix_split, matrix_slice)
		}
	}

	return matrix_split
}

func part2() {
	regex, _ := regexp.Compile(`MAS`)
	minDiagonalLength := 3

	fileContent, _ := ReadFileContent("input.txt")
	original_matrix := getMatrix(fileContent)

	matrix_split := getMatrixSplitByLength(original_matrix, 3)

	all_occurences := lo.Map(matrix_split, func(matrix [][]string, _ int) bool {
		stringsToTest := []string{}

		stringsToTest = append(stringsToTest, getDiagonal(matrix, 0, 0, minDiagonalLength, false, false))
		stringsToTest = append(stringsToTest, getDiagonal(matrix, 0, 0, minDiagonalLength, false, true))

		stringsToTest = append(stringsToTest, getDiagonal(matrix, 0, 2, minDiagonalLength, true, false))
		stringsToTest = append(stringsToTest, getDiagonal(matrix, 0, 2, minDiagonalLength, true, true))

		occurences := lo.Sum(lo.Map(stringsToTest, func(s string, _ int) int {
			return checkLine(*regex, s)
		}))

		return occurences == 2
	})

	println(lo.Sum(lo.Map(all_occurences, func(s bool, _ int) int {
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
