package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

type entry struct {
	signal []string
	output []string
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func part1(input []entry) int {
	result := 0

	for _, entry := range input {
		// fmt.Println(entry)
		// os.Exit(0)
		for _, element := range entry.output {
			switch len(element) {
			case 2, 3, 4, 7:
				result++
			}
		}
	}
	return result
}

func subtractWord(root, toSubtract string) string {
	result := ""
	for _, letter := range root {
		if !strings.Contains(toSubtract, string(letter)) {
			result += string(letter)
		}

	}
	return SortString(result)
}

func contains(word, toCheck string) bool {
	result := true
	for _, letter := range toCheck {
		if !strings.Contains(word, string(letter)) {
			result = false
		}

	}
	return result
}

func parseSignals(values []string) map[string]string {
	numbers := make([]string, 10)
	result := make(map[string]string)
	result["abcdefg"] = "8"

	for _, v := range values {
		switch len(v) {
		case 2:
			numbers[1] = SortString(v)
			result[SortString(v)] = "1"
		case 3:
			numbers[7] = SortString(v)
			result[SortString(v)] = "7"
		case 4:
			numbers[4] = SortString(v)
			result[SortString(v)] = "4"
		}
	}

	L := subtractWord(numbers[4], numbers[1])

	for _, word := range values {
		sortedWord := SortString(word)
		switch len(sortedWord) {
		case 5:
			switch {
			case contains(sortedWord, numbers[1]):
				result[sortedWord] = "3"
			case contains(sortedWord, L):
				result[sortedWord] = "5"
			default:
				result[sortedWord] = "2"
			}
		case 6:
			switch {
			case contains(sortedWord, numbers[4]):
				result[sortedWord] = "9"
			case contains(sortedWord, L):
				result[sortedWord] = "6"
			default:
				result[sortedWord] = "0"
			}
		}
	}

	return result
}

func part2(input []entry) int {
	outputs := make([]int, 0)

	for _, entry := range input {
		decodeMap := parseSignals(entry.signal)
		entryResultWord := ""
		for _, word := range entry.output {
			number, ok := decodeMap[word]
			if !ok {
				panic("nope")
			}
			entryResultWord += number
		}
		parsed, err := strconv.Atoi(entryResultWord)
		if err != nil {
			panic("NaN")
		}
		outputs = append(outputs, parsed)
	}

	result := 0
	for _, output := range outputs {
		result += output
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
	var entries []entry
	for _, value := range inputLists {
		parsed := strings.Split(value, " | ")
		newEntry := entry{}

		signals := strings.Split(parsed[0], " ")
		outputSignals := strings.Split(parsed[1], " ")

		for _, inputString := range signals {
			newEntry.signal = append(newEntry.signal, SortString(strings.TrimSpace(inputString)))
		}
		for _, inputString := range outputSignals {
			newEntry.output = append(newEntry.output, SortString(strings.TrimSpace(inputString)))
		}

		entries = append(entries, newEntry)

	}

	fmt.Println(part1(entries))
	fmt.Println(part2(entries))
}
