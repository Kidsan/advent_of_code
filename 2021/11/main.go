package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func visitCell(x, y int, input [][]int, flashed map[string]int) map[string]int {
	input[y][x]++
	if input[y][x] > 9 {
		key := fmt.Sprintf("%d,%d", x, y)
		_, ok := flashed[key]
		if !ok {
			flashed[key] = 1

			if y+1 < len(input) {
				flashed = visitCell(x, y+1, input, flashed)
				if x+1 < len(input[y]) {
					flashed = visitCell(x+1, y+1, input, flashed)
				}
				if x-1 >= 0 {
					flashed = visitCell(x-1, y+1, input, flashed)
				}
			}

			if x-1 >= 0 {
				flashed = visitCell(x-1, y, input, flashed)
			}

			if x+1 < len(input[y]) {
				flashed = visitCell(x+1, y, input, flashed)
			}

			if y-1 >= 0 {
				flashed = visitCell(x, y-1, input, flashed)
				if x+1 < len(input[y]) {
					flashed = visitCell(x+1, y-1, input, flashed)
				}
				if x-1 >= 0 {
					flashed = visitCell(x-1, y-1, input, flashed)
				}
			}
		}

	}
	return flashed
}

func tick(input [][]int) int {
	flashed := make(map[string]int)
	for y, row := range input {
		for x := range row {
			flashed = visitCell(x, y, input, flashed)
		}
	}

	for key := range flashed {
		coordinates := strings.Split(key, ",")
		xValue, _ := strconv.Atoi(coordinates[0])
		yValue, _ := strconv.Atoi(coordinates[1])
		input[yValue][xValue] = 0
	}

	return len(flashed)
}

func part1(input [][]int) int {
	duplicate := make([][]int, len(input))
	for i := range input {
		duplicate[i] = make([]int, len(input[i]))
		copy(duplicate[i], input[i])
	}

	result := 0

	for i := 0; i < 100; i++ {
		result += tick(duplicate)
	}

	return result
}

func part2(input [][]int) int {
	result := 0
	countOfCells := len(input) * len(input[0])

	for i := 1; result < 1; i++ {
		flashedInCycle := tick(input)
		if flashedInCycle == countOfCells {
			result = i
		}
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
	cells := make([][]int, 0)
	for y, line := range inputLists {
		newArray := make([]int, len(line))
		cells = append(cells, newArray)
		for x, value := range line {
			cells[y][x], _ = strconv.Atoi(string(value))
		}
	}

	fmt.Println(part1(cells))
	fmt.Println(part2(cells))
}
