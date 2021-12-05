package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Cell struct {
	value  int
	marked bool
}

type Board [][]Cell

func (b Board) markNumber(toMark int) {
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if b[x][y].value == toMark {
				b[x][y].marked = true
			}
		}
	}
}

func (b Board) isComplete() bool {
	for i := 0; i < 5; i++ {
		countMarkedInRow := 0
		countMarkedInColumn := 0
		for j := 0; j < 5; j++ {
			if b[i][j].marked {
				countMarkedInRow++
			}
			if b[j][i].marked {
				countMarkedInColumn++
			}
		}
		if countMarkedInRow == 5 || countMarkedInColumn == 5 {
			return true
		}
	}
	return false
}

func (b Board) sumUnMarkedCells() int {
	var sum int
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if !b[x][y].marked {
				sum += b[x][y].value
			}
		}
	}
	return sum
}

func (b Board) prettyPrint() {
	var result string
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if b[x][y].marked {
				result += "*" + fmt.Sprint(b[x][y].value)
			} else {
				result += fmt.Sprint(b[x][y].value)
			}
			result += " "
		}
		result += "\n"
	}
	fmt.Println(result)
}

func NewBoard(boardSpec string) *Board {
	rows := strings.Split(boardSpec, "\n")
	board := Board{{}, {}, {}, {}, {}}

	for x := 0; x < 5; x++ {
		board[x] = make([]Cell, 5)
	}

	for y, row := range rows {
		row = strings.ReplaceAll(row, "  ", " ")
		row = strings.Trim(row, " ")
		rowValues := strings.Split(row, " ")

		for x, numericText := range rowValues {
			if numericText == "" || numericText == " " {
				continue
			}
			newValue, _ := strconv.Atoi(numericText)
			newCell := Cell{value: newValue}
			board[x][y] = newCell
		}
	}
	return &board
}

func Part1(drawnValues []int, boards []Board) int {
	for _, v := range drawnValues {
		for _, b := range boards {
			b.markNumber(v)
			if b.isComplete() {
				unmarkedCellSum := b.sumUnMarkedCells()
				return v * unmarkedCellSum
			}
		}

	}
	fmt.Println("No Winner Was Found!")
	return 0
}

// TODO: fix this function
func Part2(drawnValues []int, boards []Board) int {
	var completeBoards []*Board
	boardsToLetWin := len(boards) - 1
	for _, v := range drawnValues {
		for _, b := range boards {
			b.prettyPrint()
			if !b.isComplete() {
				b.markNumber(v)
				if b.isComplete() {
					completeBoards = append(completeBoards, &b)
					if len(completeBoards) == boardsToLetWin {
						unmarkedCellSum := b.sumUnMarkedCells()

						return v * unmarkedCellSum
					}
				}
			}
		}

	}
	fmt.Println(boardsToLetWin, len(completeBoards))
	fmt.Println("No Winner Was Found!")
	return 0
}

func temp(inputLists []string) int {
	drawValues := strings.Split(inputLists[0], ",")
	boardStrings := inputLists[1:]
	boards := make([]Board, 0)
	drawNumbers := make([]int, 0)
	for _, board := range boardStrings {
		// fmt.Println(board, "\n")
		boards = append(boards, *NewBoard(board))
	}

	for _, value := range drawValues {
		v, _ := strconv.Atoi(string(value))
		drawNumbers = append(drawNumbers, v)
	}

	return Part1(drawNumbers, boards)
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

	inputLists := strings.Split(string(content), "\n\n")
	drawValues := strings.Split(inputLists[0], ",")
	boardStrings := inputLists[1:]
	boards := make([]Board, 0)
	drawNumbers := make([]int, 0)
	for _, board := range boardStrings {
		// fmt.Println(board, "\n")
		boards = append(boards, *NewBoard(board))
	}

	for _, value := range drawValues {
		v, _ := strconv.Atoi(string(value))
		drawNumbers = append(drawNumbers, v)
	}

	fmt.Println(Part1(drawNumbers, boards))
	fmt.Println(Part2(drawNumbers, boards))
}
