package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func Part1(values []string) int64 {
	var gammaResult string
	var epsilonResult string

	for index := range values[0] {
		var countOfOnes int
		var countOfZeroes int
		for _, value := range values {
			switch string(value[index]) {
			case "0":
				countOfZeroes++
			case "1":
				countOfOnes++
			}
		}

		switch countOfOnes > countOfZeroes {
		case true:
			gammaResult += "1"
			epsilonResult += "0"
		case false:
			gammaResult += "0"
			epsilonResult += "1"
		default:
			continue
		}
	}

	gammaDecimal, err := strconv.ParseInt(gammaResult, 2, 64)
	if err != nil {
		panic(err)
	}

	epsilonDecimal, err := strconv.ParseInt(epsilonResult, 2, 64)
	if err != nil {
		panic(err)
	}

	return gammaDecimal * epsilonDecimal
}

func getOxygenGeneratorRating(values []string, index int) string {
	if len(values) == 1 {
		return values[0]
	}

	filteredElements := make([]string, 0)
	var countOfOnes []string
	var countOfZeroes []string

	for _, value := range values {
		switch string(value[index]) {
		case "0":
			countOfZeroes = append(countOfZeroes, value)
		case "1":
			countOfOnes = append(countOfOnes, value)
		}
	}

	newIndex := index + 1

	switch {
	case len(countOfOnes) == len(countOfZeroes):
		filteredElements = append(filteredElements, countOfOnes...)
	case len(countOfOnes) > len(countOfZeroes):
		filteredElements = append(filteredElements, countOfOnes...)
	case len(countOfOnes) < len(countOfZeroes):
		filteredElements = append(filteredElements, countOfZeroes...)
	}

	return getOxygenGeneratorRating(filteredElements, newIndex)
}

func getCO2ScrubberRating(values []string, index int) string {
	if len(values) == 1 {
		return values[0]
	}

	filteredElements := make([]string, 0)
	var countOfOnes []string
	var countOfZeroes []string

	for _, value := range values {
		switch string(value[index]) {
		case "0":
			countOfZeroes = append(countOfZeroes, value)
		case "1":
			countOfOnes = append(countOfOnes, value)
		}
	}

	newIndex := index + 1

	switch {
	case len(countOfOnes) == len(countOfZeroes):
		filteredElements = append(filteredElements, countOfZeroes...)
	case len(countOfOnes) > len(countOfZeroes):
		filteredElements = append(filteredElements, countOfZeroes...)
	case len(countOfOnes) < len(countOfZeroes):
		filteredElements = append(filteredElements, countOfOnes...)
	}

	return getCO2ScrubberRating(filteredElements, newIndex)
}

func Part2(values []string) int64 {
	fistNum := getOxygenGeneratorRating(values, 0)
	secondNum := getCO2ScrubberRating(values, 0)

	fistNumDecimal, _ := strconv.ParseInt(fistNum, 2, 64)
	secondNumDecimal, _ := strconv.ParseInt(secondNum, 2, 64)
	return fistNumDecimal * secondNumDecimal
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

	fmt.Println(Part1(inputLists))
	fmt.Println(Part2(inputLists))
}
