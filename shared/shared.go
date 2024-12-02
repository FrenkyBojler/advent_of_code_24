package shared

import (
	"io"
	"os"
	"strconv"
)

func ReadFileContent(path string) (string, error) {
	file_content, err := os.Open(path)

	if err != nil {
		return "", err
	}

	file_content_string, err := io.ReadAll(file_content)
	defer file_content.Close()

	return string(file_content_string), err
}

func ToIntSlice(s []string) []int {
	var result []int
	for _, v := range s {
		result = append(result, ToInt(v))
	}
	return result
}

func ToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
