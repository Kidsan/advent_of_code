package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Line struct {
	points []point
}

type point struct {
	x int
	y int
}

func countOverlappingPoints(values []Line) int {
	var coveredPoints = make(map[string]int)

	for _, value := range values {
		for _, point := range value.points {
			xString := strconv.Itoa(point.x)
			yString := strconv.Itoa(point.y)
			pointString := xString + "," + yString
			_, ok := coveredPoints[pointString]
			if ok {
				coveredPoints[pointString]++
			} else {
				coveredPoints[pointString] = 1
			}

		}
	}

	result := 0
	for _, value := range coveredPoints {
		if value > 1 {
			result += 1
		}
	}
	return result
}

func parsePoints(input string, diagonals bool) []point {
	var result []point
	pointSpecification := strings.Split(input, " -> ")
	startingPointSpec := strings.Split(pointSpecification[0], ",")
	endPointSpec := strings.Split(pointSpecification[1], ",")

	startingPointX, _ := strconv.Atoi(startingPointSpec[0])
	startingPointY, _ := strconv.Atoi(startingPointSpec[1])
	endPointX, _ := strconv.Atoi(endPointSpec[0])
	endPointY, _ := strconv.Atoi(endPointSpec[1])

	// shitshow
	switch {
	case startingPointX == endPointX:
		switch startingPointY < endPointY {
		case true:
			for i := startingPointY; i <= endPointY; i++ {
				result = append(result, point{
					x: startingPointX,
					y: i,
				})
			}
		default:
			for i := endPointY; i <= startingPointY; i++ {
				result = append(result, point{
					x: startingPointX,
					y: i,
				})
			}
		}
	case startingPointY == endPointY:
		switch startingPointX < endPointX {
		case true:
			for i := startingPointX; i <= endPointX; i++ {
				result = append(result, point{
					x: i,
					y: startingPointY,
				})
			}
		default:
			for i := endPointX; i <= startingPointX; i++ {
				result = append(result, point{
					x: i,
					y: startingPointY,
				})
			}
		}
	default:
		if diagonals {
			switch {
			case startingPointX > endPointX:
				if endPointY > startingPointY {
					tmp := endPointY
					for i := endPointX; i <= startingPointX; i++ {
						result = append(result, point{x: i, y: tmp})
						tmp--
					}
				} else {
					tmp := endPointY
					for i := endPointX; i <= startingPointX; i++ {
						result = append(result, point{x: i, y: tmp})
						tmp++
					}
				}
			default:
				if endPointY >= startingPointY {
					tmp := startingPointY
					for i := startingPointX; i <= endPointX; i++ {
						result = append(result, point{x: i, y: tmp})
						tmp++
					}
				} else {
					tmp := startingPointY
					for i := startingPointX; i <= endPointX; i++ {
						result = append(result, point{x: i, y: tmp})
						tmp--
					}
				}
			}
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

	var linesWithoutDiagonals []Line
	var linesWithDiagonals []Line

	for _, lineSpec := range inputLists {
		parsedPoints := parsePoints(lineSpec, false)
		linesWithoutDiagonals = append(linesWithoutDiagonals, Line{points: parsedPoints})
	}

	for _, lineSpec := range inputLists {
		parsedPoints := parsePoints(lineSpec, true)
		linesWithDiagonals = append(linesWithDiagonals, Line{points: parsedPoints})
	}

	fmt.Println(countOverlappingPoints(linesWithoutDiagonals)) // 5373
	fmt.Println(countOverlappingPoints(linesWithDiagonals))    // 21514
}
