package utils

import (
	"bufio"
	"os"
)

func Readfile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	return file
}

func ReadFileIntoLines(path string) []string {
	file := Readfile(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func ReadFileIntoString(path string) string {
	file, err := os.ReadFile(path)
	CheckError(err)

	return string(file)
}
