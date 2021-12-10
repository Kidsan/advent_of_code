package main

import (
	"fmt"
	"io/ioutil"
	"os"
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
	fmt.Println(errors)
	for _, v := range errors {
		score := scoring[v]
		result += score
	}
	return result
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
}
