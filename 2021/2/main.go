package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	direction string
	distance  int
}

func FinalPosition(moves ...Instruction) int {
	var horizontalPosition int
	var verticalPosition int

	for _, value := range moves {
		switch value.direction {
		case "forward":
			horizontalPosition += value.distance
		case "up":
			verticalPosition -= value.distance
		case "down":
			verticalPosition += value.distance
		}
	}
	return horizontalPosition * verticalPosition
}

func AimingSubmarine(moves ...Instruction) int {
	var aim int
	var horizontalPosition int
	var verticalPosition int

	for _, value := range moves {
		switch value.direction {
		case "forward":
			horizontalPosition += value.distance
			verticalPosition += value.distance * aim
		case "up":
			aim -= value.distance
		case "down":
			aim += value.distance
		}
	}
	return horizontalPosition * verticalPosition
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

	var instructions []Instruction
	for _, input := range inputLists {
		movement := strings.Split(input, " ")
		parsedDistance, err := strconv.Atoi(movement[1])
		if err != nil {
			continue
		}

		instructions = append(instructions, Instruction{
			direction: movement[0],
			distance:  parsedDistance,
		})
	}
	fmt.Println(FinalPosition(instructions...))
	fmt.Println(AimingSubmarine(instructions...))
}
