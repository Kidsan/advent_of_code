package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type probe struct {
	posX      int
	posY      int
	velocityX int
	velocityY int
}

func (p *probe) ApplyVelocity() {
	// fmt.Println("before:", p.velocityX, p.velocityY)
	p.posX += p.velocityX
	p.posY += p.velocityY

	if p.velocityX > 0 {
		p.velocityX -= 1
	}

	if p.velocityX < 0 {
		p.velocityX += 1
	}
	p.velocityY -= 1
}

func (p *probe) isWithinTarget(targetCoordinates map[string]int) bool {
	key := fmt.Sprintf("%d,%d", p.posX, p.posY)
	// fmt.Println(targetCoordinates)
	// os.Exit(0)
	_, ok := targetCoordinates[key]
	return ok
}

func getTargetCoordinates(xRange, yRange []int) map[string]int {
	result := make(map[string]int)
	for _, xValue := range xRange {
		for _, yValue := range yRange {
			key := fmt.Sprintf("%d,%d", xValue, yValue)
			result[key] = 1
		}
	}
	return result
}

func getPossibleVelocities(xRange, yRange []int) map[string][]string {
	left := xRange[0]              //  left
	right := xRange[len(xRange)-1] // right
	bottom := yRange[0]            // bottom
	top := yRange[len(yRange)-1]   //  top

	targetCoordinates := getTargetCoordinates(xRange, yRange)
	fmt.Println(left, right, bottom, top)
	//os.Exit(0)
	options := make(map[string][]string)

	for i := 0; i <= right; i++ {
		for y := bottom; y < 1000; y++ {
			p := probe{
				posX:      0,
				posY:      0,
				velocityX: i,
				velocityY: y,
			}

			reachedPositions := make([]string, 0)
			for p.posX <= right && p.posY >= bottom {
				p.ApplyVelocity()
				position := fmt.Sprintf("%d,%d", p.posX, p.posY)
				reachedPositions = append(reachedPositions, position)
				if p.isWithinTarget(targetCoordinates) {
					key := fmt.Sprintf("%d,%d", i, y)
					options[key] = reachedPositions
				}
			}
		}

	}

	return options
}

func part1(xRange, yRange []int) int {
	validVelocities := getPossibleVelocities(xRange, yRange)

	peak := 0
	for _, positions := range validVelocities {
		for _, position := range positions {
			parsed := strings.Split(position, ",")
			y, _ := strconv.Atoi(parsed[1])
			if y > peak {
				peak = y
			}
		}
	}

	fmt.Println("Part Two: ", len(validVelocities))
	return peak
}

func main() {
	start := time.Now()
	xRange := make([]int, 0)
	for i := 265; i <= 287; i++ {
		xRange = append(xRange, i)
	}

	yRange := make([]int, 0)
	for i := -103; i <= -58; i++ {
		yRange = append(yRange, i)
	}

	fmt.Printf("Part One: %v (took %s)\n", part1(xRange, yRange), time.Since(start))
}
