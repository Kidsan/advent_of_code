package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func DepthComparison(depths ...int) int {
	var result int

	for index, value := range depths {
		if index == 0 {
			continue
		}
		if value > depths[index-1] {
			result += 1
		}
	}

	return result
}

func SlidingWindowComparison(depths ...int) int {
	var result int
	slidingWindowDepths := make([]int, 0)

	for index, value := range depths {
		if index > 1 {
			slidingWindowDepths = append(slidingWindowDepths, value+depths[index-1]+depths[index-2])
		}
	}
	result = DepthComparison(slidingWindowDepths...)
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

	var numberInputs []int
	for _, input := range inputLists {
		num, err := strconv.Atoi(input)
		if err != nil {
			continue
		}

		numberInputs = append(numberInputs, num)
	}

	fmt.Println(DepthComparison(numberInputs...))
	fmt.Println(SlidingWindowComparison(numberInputs...))
}
