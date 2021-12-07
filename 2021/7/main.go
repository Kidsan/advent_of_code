package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part1(values []int, floorSize int) int {
	floor := make([]int, floorSize)
	for index := range floor {
		var distanceToOtherCrabs int
		for _, crabPosition := range values {
			if index != crabPosition {
				distance := crabPosition - index
				if distance < 0 {
					distance = distance * -1
				}
				distanceToOtherCrabs += distance
			}

		}
		floor[index] = distanceToOtherCrabs
	}
	sort.Ints(floor)
	return floor[0]
}

func getScaledFuelLoss(value int) int {
	result := 0
	for value > 0 {
		result += value
		value--
	}
	return result
}

func part2(values []int, floorSize int) int {
	floor := make([]int, floorSize)
	for index := range floor {
		var distanceToOtherCrabs int
		for _, crabPosition := range values {
			if index != crabPosition {
				distance := crabPosition - index
				if distance < 0 {
					distance = distance * -1
				}
				distance = getScaledFuelLoss(distance)
				distanceToOtherCrabs += distance
			}

		}
		floor[index] = distanceToOtherCrabs
	}
	sort.Ints(floor)
	return floor[0]
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

	inputLists := strings.Split(string(content), ",")
	var crabs []int
	var sizeOfFloor int
	for _, value := range inputLists {
		parsed, _ := strconv.Atoi(value)
		if parsed > sizeOfFloor {
			sizeOfFloor = parsed
		}
		crabs = append(crabs, parsed)

	}

	fmt.Println(part1(crabs, sizeOfFloor))
	fmt.Println(part2(crabs, sizeOfFloor))
}
