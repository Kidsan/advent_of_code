package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func findFirstError(input string) string {
	result := ""
	builder := ""
	for _, letter := range input {
		parsed := string(letter)
		switch parsed {
		case "[", "(", "{", "<":
			builder += parsed
		case "}":
			last := builder[len(builder)-1]
			if !(string(last) == "{") {
				result = parsed
				return result
			}
			builder = builder[:len(builder)-1]
		case "]":
			last := builder[len(builder)-1]
			if !(string(last) == "[") {
				result = parsed
				return result
			}
			builder = builder[:len(builder)-1]
		case ">":
			last := builder[len(builder)-1]
			if !(string(last) == "<") {
				result = parsed
				return result
			}
			builder = builder[:len(builder)-1]
		case ")":
			last := builder[len(builder)-1]
			if !(string(last) == "(") {
				result = parsed
				return result
			}
			builder = builder[:len(builder)-1]
		}
	}

	return result
}

func part1(input []string) int {
	errors := make([]string, 0)
	scoring := map[string]int{")": 3, "]": 57, "}": 1197, ">": 25137}
	result := 0
	for _, line := range input {
		if errorInLine := findFirstError(line); errorInLine != "" {
			errors = append(errors, errorInLine)
		}
	}

	for _, v := range errors {
		score := scoring[v]
		result += score
	}
	return result
}

func filterOutInvalidLines(input []string) []string {
	result := make([]string, 0)
	for _, line := range input {
		if errorInLine := findFirstError(line); errorInLine == "" {
			result = append(result, line)
		}
	}
	return result
}

func findMissingLineEnding(input string) string {
	ending := make([]string, 0)
	for _, letter := range input {
		parsed := string(letter)
		switch parsed {
		case "[":
			ending = append([]string{"]"}, ending...)
		case "(":
			ending = append([]string{")"}, ending...)
		case "{":
			ending = append([]string{"}"}, ending...)
		case "<":
			ending = append([]string{">"}, ending...)
		case "}":
			ending = ending[1:]
		case "]":
			ending = ending[1:]
		case ">":
			ending = ending[1:]
		case ")":
			ending = ending[1:]
		}
	}

	return strings.Join(ending, "")
}

func getEndingScore(input string) int {
	result := 0
	scoring := map[string]int{")": 1, "]": 2, "}": 3, ">": 4}

	for _, v := range input {
		letter := string(v)
		result = (result * 5) + scoring[letter]
	}

	return result

}

func part2(input []string) int {
	lines := filterOutInvalidLines(input)
	scores := make([]int, 0)

	for _, line := range lines {
		lineEnding := findMissingLineEnding(line)
		scores = append(scores, getEndingScore(lineEnding))
	}

	sort.Ints(scores)
	return scores[len(scores)/2]
}

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	content, err := ioutil.ReadAll(input)
	if err != nil {
		panic(err)
	}

	inputLists := strings.Split(string(content), "\n")

	fmt.Println(part1(inputLists))
	fmt.Println(part2(inputLists))
}
