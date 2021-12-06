package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Ocean struct {
	fish map[int]int
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

	ocean := Ocean{
		fish: map[int]int{8: 0, 7: 0, 6: 0, 5: 0, 4: 0, 3: 0, 2: 0, 1: 0, 0: 0},
	}
	for _, v := range inputLists {
		num, _ := strconv.Atoi(v)
		ocean.fish[num] += 1
	}

	for i := 0; i < 256; i++ {
		nextState := make(map[int]int)
		for k := range ocean.fish {
			switch k {
			case 0:
				nextState[8] += ocean.fish[0]
				nextState[6] += ocean.fish[0]
			default:
				nextState[k-1] += ocean.fish[k]
			}
		}
		ocean.fish = nextState

	}
	total := 0
	for _, count := range ocean.fish {
		total += count
	}
	fmt.Println(total)
}
