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

	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			result += 1
		}
	}

	return result
}

func SlidingWindowComparison(depths ...int) int {
	var result int
	slidingWindowDepths := make([]int, 0)

	for i := 2; i < len(depths); i++ {
		slidingWindowDepths = append(slidingWindowDepths, depths[i]+depths[i-1]+depths[i-2])
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
