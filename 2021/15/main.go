package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func part1(x int, y int, rows [][]int, seen map[string]int) int {
	key := fmt.Sprintf("%d,%d", x, y)
	value, ok := seen[key]
	if ok {
		return value
	}
	seen[key] = 1
	if y < 0 || y >= len(rows) || x < 0 || x >= len(rows[y]) {
		return math.MaxInt
	}
	if y == len(rows)-1 && x == len(rows[y])-1 {
		return rows[y][x]
	}
	right := float64(part1(x+1, y, rows, seen))
	down := float64(part1(x, y+1, rows, seen))

	cheapest := math.Min(down, right)
	ans := rows[y][x] + int(cheapest)
	seen[key] = ans
	return ans

}

func main() {
	start := time.Now()
	input, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	content, err := ioutil.ReadAll(input)
	if err != nil {
		panic(err)
	}

	inputLists := strings.Split(string(content), "\n")
	data := make([][]int, 0)
	seen := make(map[string]int)

	for _, line := range inputLists {
		values := []int{}
		for _, r := range line {
			num, _ := strconv.Atoi(string(r))
			values = append(values, num)
		}
		data = append(data, values)
	}

	fmt.Printf("Part One: %v (took %s)\n", part1(0, 0, data, seen), time.Since(start))
}
