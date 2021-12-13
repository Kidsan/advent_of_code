package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type caveSystem map[string][]string

func alreadyVisited(s []string, value string) bool {
	for _, v := range s {
		if v == value {
			return true
		}
	}
	return false
}

func getPaths(node string, path []string, graph caveSystem, paths *[][]string, usedFreeSmallCave bool) {
	if node == "end" {
		c := make([]string, len(path))
		copy(c, path)
		*paths = append(*paths, c)
		return
	}
	for _, nextNode := range graph[node] {
		if nextNode == "start" {
			continue
		}

		if strings.ToLower(string(nextNode)) == string(nextNode) && alreadyVisited(path, nextNode) {
			if !usedFreeSmallCave {
				path = append(path, nextNode)
				getPaths(nextNode, path, graph, paths, true)
			}
			continue
		}
		path = append(path, nextNode)
		getPaths(nextNode, path, graph, paths, usedFreeSmallCave)
		path = path[:len(path)-1]
	}
}

func part1(input caveSystem) int {
	paths := [][]string{}
	path := []string{"start"}
	getPaths("start", path, input, &paths, true)

	return len(paths)
}

func part2(input caveSystem) int {
	paths := [][]string{}
	path := []string{"start"}
	getPaths("start", path, input, &paths, false)
	return len(paths)
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	content, err := ioutil.ReadAll(input)
	if err != nil {
		panic(err)
	}

	graph := make(caveSystem)
	inputLists := strings.Split(string(content), "\n")

	for _, line := range inputLists {
		parsed := strings.Split(line, "-")
		from := parsed[0]
		to := parsed[1]
		if _, ok := graph[from]; !ok {
			graph[from] = []string{}
		}
		if _, ok := graph[to]; !ok {
			graph[to] = []string{}
		}
		graph[from] = append(graph[from], to)
		graph[to] = append(graph[to], from)
	}

	fmt.Println(part1(graph))
	fmt.Println(part2(graph))

}
