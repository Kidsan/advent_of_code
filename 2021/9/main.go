package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getLowestPoints(numbers [][]int) [][]int {
	lowestPoints := make([][]int, 0)

	for rowIndex, row := range numbers {
		for colIndex, number := range row {
			if colIndex != 0 {
				if number >= row[colIndex-1] {
					continue
				}
			}

			if rowIndex != 0 {
				if number >= numbers[rowIndex-1][colIndex] {
					continue
				}
			}

			if colIndex != len(row)-1 {
				if number >= row[colIndex+1] {
					continue
				}
			}

			if rowIndex != len(numbers)-1 {
				if number >= numbers[rowIndex+1][colIndex] {
					continue
				}
			}

			lowestPoints = append(lowestPoints, []int{colIndex, rowIndex, number})
		}
	}

	return lowestPoints
}

func part1(input [][]int) int {
	result := 0
	lowestPoints := getLowestPoints(input)

	for _, v := range lowestPoints {
		result += 1 + v[2]
	}

	return result
}

func part2(numbers [][]int) int {
	lowestPoints := getLowestPoints(numbers)
	basins := make([]int, 0)

	for _, pos := range lowestPoints {
		key := fmt.Sprintf("%d,%d", pos[0], pos[1])

		basinPoints := map[string]int{
			key: pos[2],
		}
		positions := getBasinSize(pos[0], pos[1], pos[2], numbers, basinPoints)

		// get the unique poistions
		basins = append(basins, len(positions))
	}

	sort.Ints(basins)
	largest := basins[len(basins)-3:]
	return largest[0] * largest[1] * largest[2]
}

func belongsToBasin(currentValue int, rowIndex int, colIndex int, visitedNodes map[string]int, numbers [][]int) map[string]int {
	key := fmt.Sprintf("%d,%d", rowIndex, colIndex)

	_, ok := visitedNodes[key]

	if ok {
		return visitedNodes
	}
	visitedNodes[key] = currentValue
	return getBasinSize(colIndex, rowIndex, currentValue, numbers, visitedNodes)
}

func getBasinSize(colIndex int, rowIndex int, number int, numbers [][]int, visitedNodes map[string]int) map[string]int {
	row := numbers[rowIndex]

	if colIndex != 0 {
		num := row[colIndex-1]
		if num-number >= 0 && num != 9 {
			visitedNodes = belongsToBasin(num, rowIndex, colIndex-1, visitedNodes, numbers)
		}
	}

	if rowIndex != 0 {
		num := numbers[rowIndex-1][colIndex]
		if num-number >= 0 && num != 9 {
			visitedNodes = belongsToBasin(num, rowIndex-1, colIndex, visitedNodes, numbers)
		}
	}

	if colIndex != len(row)-1 {
		num := row[colIndex+1]
		if num-number >= 0 && num != 9 {
			visitedNodes = belongsToBasin(num, rowIndex, colIndex+1, visitedNodes, numbers)
		}
	}

	if rowIndex != len(numbers)-1 {
		num := numbers[rowIndex+1][colIndex]
		if num-number >= 0 && num != 9 {
			visitedNodes = belongsToBasin(num, rowIndex+1, colIndex, visitedNodes, numbers)
		}

	}
	return visitedNodes
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

	var numbers [][]int
	for lineIndex, line := range inputLists {
		parts := strings.Split(line, "")
		numbers = append(numbers, make([]int, 0))

		for _, part := range parts {
			num, _ := strconv.Atoi(part)
			numbers[lineIndex] = append(numbers[lineIndex], num)
		}
	}

	fmt.Println(part1(numbers))
	fmt.Println(part2(numbers))
}
