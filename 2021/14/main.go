package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type rule struct {
	matches string
	inserts string
}

type Letter struct {
	next  *Letter
	value string
}

type InputString struct {
	start  *Letter
	length int
}

type insertion struct {
	pointer *Letter
	new     *Letter
}

func (i *InputString) add(newLetter string) {
	new := &Letter{value: newLetter, next: i.start}
	i.start = new
	i.length++
}

func (i *InputString) toString() string {
	result := ""

	iter := i.start
	for iter != nil {
		result += iter.value
		iter = iter.next
	}

	return result
}

func insertionPhase(input *InputString, rules []rule) {
	insertions := make([]insertion, 0)
	for _, r := range rules {
		current := input.start
		for current != nil {
			oldNext := current.next
			if oldNext != nil {
				if current.value == string(r.matches[0]) && oldNext.value == string(r.matches[1]) {
					newInsertion := &Letter{value: r.inserts, next: oldNext}
					insertions = append(insertions, insertion{pointer: current, new: newInsertion})
					input.length++

				}
			}
			current = oldNext
		}
	}

	for _, i := range insertions {
		l := i.pointer
		l.next = i.new
	}
}

func insertions(input *InputString, rules []rule, iterations int) string {
	for i := 0; i < iterations; i++ {
		insertionPhase(input, rules)
	}
	return input.toString()
}

func part1(input string, rules []rule) int {
	iterations := 10
	inputList := InputString{}
	for i := len(input) - 1; i >= 0; i-- {
		inputList.add(string(input[i]))
	}

	insertions(&inputList, rules, iterations)

	values := make(map[string]int)
	curr := inputList.start
	for curr != nil {
		values[curr.value]++
		curr = curr.next
	}

	most := 0
	least := 0
	for _, value := range values {
		if least == 0 || value < least {
			least = value
		}
		if value > most {
			most = value
		}
	}

	return most - least
}

func part2(input string, rules []rule) int {
	iterations := 40
	inputList := InputString{}
	for i := len(input) - 1; i >= 0; i-- {
		inputList.add(string(input[i]))
	}

	insertions(&inputList, rules, iterations)

	values := make(map[string]int)
	curr := inputList.start
	for curr != nil {
		values[curr.value]++
		curr = curr.next
	}

	most := 0
	least := 0
	for _, value := range values {
		if least == 0 || value < least {
			least = value
		}
		if value > most {
			most = value
		}
	}

	return most - least
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
	inputText := inputLists[0]
	rulesList := inputLists[2:]
	rules := make([]rule, 0)
	for _, line := range rulesList {
		newRule := strings.Split(line, " -> ")
		rules = append(rules, rule{
			matches: newRule[0],
			inserts: newRule[1],
		})
	}

	fmt.Println(part1(inputText, rules))
	fmt.Println(part2(inputText, rules))
}
