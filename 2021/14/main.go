package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type rule struct {
	matches string
	inserts string
}

type insertion struct {
	value    string
	replaces string
	number   int
}

func insertionPhase(input map[string]int, rules []rule) {
	insertions := make([]insertion, 0)
	for _, r := range rules {
		count, ok := input[r.matches]
		if !ok {
			continue
		}
		insertions = append(insertions, insertion{
			value:    fmt.Sprintf("%s%s", string(r.matches[0]), r.inserts),
			replaces: r.matches,
			number:   count,
		})
		insertions = append(insertions, insertion{
			value:    fmt.Sprintf("%s%s", r.inserts, string(r.matches[1])),
			replaces: r.matches,
			number:   count,
		})
		delete(input, r.matches)
	}

	for _, v := range insertions {
		input[v.value] += v.number
	}

}

func insertions(input map[string]int, rules []rule, iterations int) map[string]int {
	for i := 0; i < iterations; i++ {
		insertionPhase(input, rules)
	}
	return input
}

func applyRules(input string, rules []rule, iterations int) int {
	chars := make(map[string]int)

	for i := 0; i+1 < len(input); i++ {
		newEntry := string(input[i]) + string(input[i+1])
		chars[newEntry]++
	}
	chars[string(input[len(input)-1])] = 1

	insertions(chars, rules, iterations)

	values := make(map[string]int)

	for key, count := range chars {
		letters := strings.Split(key, "")
		values[letters[0]] += count
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

	fmt.Println(applyRules(inputText, rules, 10))
	fmt.Printf("part one took %s\n", time.Since(start))
	fmt.Println(applyRules(inputText, rules, 40))
	fmt.Printf("part two took %s\n", time.Since(start))
}
