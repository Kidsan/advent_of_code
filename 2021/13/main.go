package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func foldY(input [][]int, foldLine int) [][]int {
	result := make([][]int, len(input)-foldLine-1)

	for i := range result {
		result[i] = append(result[i], input[i]...)
	}

	for i := 1; i+foldLine < len(input); i++ {
		rowAboveY := foldLine - i
		rowBelowY := foldLine + i

		for index, value := range input[rowBelowY] {
			if value == 1 {
				result[rowAboveY][index] = value
			} else {
				result[rowAboveY][index] = input[rowAboveY][index]
			}
		}
	}
	return result
}

func foldX(input [][]int, foldLine int) [][]int {
	result := make([][]int, len(input))

	for i := range result {
		result[i] = append(result[i], input[i][:foldLine]...)
	}

	for i := 1; i+foldLine < len(input[0]); i++ {
		cellLeftX := foldLine - i
		cellRightX := foldLine + i

		for index, line := range input {
			if line[cellRightX] == 1 {
				result[index][cellLeftX] = line[cellRightX]
			}
		}
	}
	return result
}

func part1(input [][]int, foldLine fold) int {
	var resultFromFold [][]int

	switch foldLine.dimension {
	case "x":
		resultFromFold = foldX(input, foldLine.line)
	case "y":
		resultFromFold = foldY(input, foldLine.line)
	}

	count := 0

	for _, line := range resultFromFold {
		for _, value := range line {
			count += value
		}
	}
	return count
}

func part2(input [][]int, folds []fold) {
	resultFromFold := make([][]int, len(input))
	for i := range input {
		resultFromFold[i] = make([]int, len(input[i]))
		copy(resultFromFold[i], input[i])
	}

	for _, fold := range folds {
		switch fold.dimension {
		case "x":
			resultFromFold = foldX(resultFromFold, fold.line)
		case "y":
			resultFromFold = foldY(resultFromFold, fold.line)
		}
	}

	for _, line := range resultFromFold {
		outString := ""
		for _, value := range line {
			if value == 1 {
				outString += "#"
			} else {
				outString += " "
			}
		}
		fmt.Println(outString)
	}
}

type fold struct {
	dimension string
	line      int
}

type point struct {
	x int
	y int
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

	folds := []fold{}
	gridXSize := 0
	gridYSize := 0
	markedPoints := []point{}

	inputLists := strings.Split(string(content), "\n")
	for _, line := range inputLists {
		switch {
		case len(line) < 12 && len(line) > 0:
			values := strings.Split(line, ",")

			xValue, _ := strconv.Atoi(values[0])
			yValue, _ := strconv.Atoi(values[1])
			if xValue > gridXSize {
				gridXSize = xValue
			}
			if yValue > gridYSize {
				gridYSize = yValue
			}
			markedPoints = append(markedPoints, point{x: xValue, y: yValue})
		case len(line) >= 12:
			line = strings.Trim(line, "fold along ")
			parsed := strings.Split(line, "=")
			foldDimension := parsed[0]
			foldLine, _ := strconv.Atoi(parsed[1])
			folds = append(folds, fold{dimension: foldDimension, line: foldLine})
		}
	}

	grid := make([][]int, gridYSize+1)
	for index := range grid {
		grid[index] = make([]int, gridXSize+1)
	}

	for _, point := range markedPoints {
		grid[point.y][point.x] = 1
	}

	fmt.Println(part1(grid, folds[0]))
	part2(grid, folds)
}
