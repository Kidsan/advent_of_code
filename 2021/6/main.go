package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Fish struct {
	ocean *Ocean
	timer int
}

func (f *Fish) notify() *Fish {
	f.timer--
	if f.timer < 0 {
		f.timer = 6
		newFish := &Fish{
			timer: 8,
			ocean: f.ocean,
		}
		return newFish
	}
	return nil
}

type Ocean struct {
	fish []*Fish
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

	var ocean Ocean
	for _, v := range inputLists {
		num, _ := strconv.Atoi(v)
		fish := &Fish{
			timer: num,
			ocean: &ocean,
		}
		ocean.fish = append(ocean.fish, fish)
	}

	// why doesn't this work?
	// for i := 0; i < 80; i++ {
	// 	for j := 0; j < len(ocean.fish); j++ {
	// 		ocean.fish[j].notify()
	// 	}
	// }

	for i := 0; i < 80; i++ {
		var catchOfTheDay []*Fish
		for _, f := range ocean.fish {
			newFish := f.notify()
			if newFish != nil {
				catchOfTheDay = append(catchOfTheDay, newFish)
			}
		}
		ocean.fish = append(ocean.fish, catchOfTheDay...)
	}
	fmt.Println(len(ocean.fish))
}
