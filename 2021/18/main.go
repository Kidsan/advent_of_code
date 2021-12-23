package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

type snailfishNumber struct {
	prev  *snailfishNumber
	next  *snailfishNumber
	value int
	depth int
}

func (n *snailfishNumber) reduce() *snailfishNumber {
	curr := n
	for curr != nil {
		if curr.depth >= 5 {
			pairRight := curr.next
			if pairRight.depth != curr.depth {
				panic(fmt.Sprintf("exploding pair should have same depth, got %d and %d", curr.depth, pairRight.depth))
			}

			replacementNode := &snailfishNumber{
				value: 0,
				depth: curr.depth - 1,
				prev:  curr.prev,
				next:  pairRight.next,
			}

			if curr.prev != nil {
				curr.prev.value += curr.value
				curr.prev.next = replacementNode
			}

			if pairRight.next != nil {
				pairRight.next.value += pairRight.value
				pairRight.next.prev = replacementNode
			}

			if n == curr {
				return replacementNode.reduce()
			}

			return n.reduce()
		}
		curr = curr.next
	}
	curr = n
	for curr != nil {
		if curr.value >= 10 {
			replacementLeft := &snailfishNumber{
				value: curr.value / 2,
				prev:  curr.prev,
				next:  nil,
				depth: curr.depth + 1,
			}

			rightValue := curr.value / 2
			if curr.value%2 == 1 {
				rightValue += 1
			}

			replacementRight := &snailfishNumber{
				value: rightValue,
				prev:  replacementLeft,
				next:  curr.next,
				depth: curr.depth + 1,
			}

			replacementLeft.next = replacementRight

			toLeft := curr.prev
			toRight := curr.next

			if toLeft != nil {
				toLeft.next = replacementLeft
			}
			if toRight != nil {
				toRight.prev = replacementRight
			}

			if n == curr {
				return replacementLeft.reduce()
			}

			return n.reduce()
		}
		curr = curr.next
	}
	return n
}

func (n1 *snailfishNumber) add(n2 *snailfishNumber) *snailfishNumber {
	n := n1
	for n != nil {
		n.depth++
		n = n.next
	}
	n = n2
	for n != nil {
		n.depth++
		n = n.next
	}

	lastNode1 := n1
	for lastNode1.next != nil {
		lastNode1 = lastNode1.next
	}
	lastNode1.next = n2
	n2.prev = lastNode1

	n1 = n1.reduce()
	return n1
}

func parseSnailfishNumber(input string) *snailfishNumber {
	var depth int

	var pointer *snailfishNumber
	var head *snailfishNumber

	for _, r := range input {
		switch r {
		case '[':
			depth++
		case ']':
			depth--
		case ',':

		default:
			parsedValue, _ := strconv.Atoi(string(r))
			newSnailfishNumber := &snailfishNumber{
				value: parsedValue,
				prev:  pointer,
				next:  nil,
				depth: depth,
			}
			if pointer == nil {
				head = newSnailfishNumber
			} else {
				pointer.next = newSnailfishNumber
			}
			pointer = newSnailfishNumber
		}
	}
	return head
}

func (n *snailfishNumber) magnitude() int {
	// copy number to avoid list
	var copy *snailfishNumber
	var head *snailfishNumber
	var last *snailfishNumber

	curr := n
	for curr != nil {
		cp := &snailfishNumber{
			value: curr.value,
			prev:  last,
			next:  nil,
			depth: curr.depth,
		}
		if head == nil {
			head = cp
			last = cp
		} else {
			last.next = cp
			last = cp
		}
		curr = curr.next
	}
	copy = head

	for depth := 4; depth > 0; depth-- {
		curr = copy
		for curr != nil {
			if curr.depth == depth && curr.next != nil && curr.next.depth == depth {
				left := curr
				right := curr.next
				newNode := snailfishNumber{
					value: 3*left.value + 2*right.value,
					prev:  left.prev,
					next:  right.next,
					depth: depth - 1,
				}
				if left == copy {
					copy = &newNode
				} else {
					left.prev.next = &newNode
				}
				if right.next != nil {
					right.next.prev = &newNode
				}
			}
			curr = curr.next
		}
	}

	return copy.value
}

func part1(numbers []*snailfishNumber) int {
	head := numbers[0]

	for i := 1; i < len(numbers); i++ {
		head = head.add(numbers[i])
	}
	return head.magnitude()
}

func copyLinkedList(n *snailfishNumber) *snailfishNumber {
	var head *snailfishNumber
	var last *snailfishNumber
	curr := n

	for curr != nil {
		cp := &snailfishNumber{
			value: curr.value,
			prev:  last,
			next:  nil,
			depth: curr.depth,
		}
		if head == nil {
			head = cp
			last = cp
		} else {
			last.next = cp
			last = cp
		}
		curr = curr.next
	}

	return head
}

func part2(numbers []*snailfishNumber) int {
	var result int

	for i, number := range numbers {
		for i2, v := range numbers {
			if i == i2 {
				continue
			}
			a := copyLinkedList(number)
			b := copyLinkedList(v)
			r := a.add(b)
			m := r.magnitude()
			if m > result {
				result = m
			}
		}
	}

	return result
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

	var snailfishNumbersPartOne []*snailfishNumber
	var snailfishNumbersPartTwo []*snailfishNumber

	for _, line := range strings.Split(string(content), "\n") {
		newNumber := parseSnailfishNumber(line)
		newNumberCopy := parseSnailfishNumber(line)
		snailfishNumbersPartOne = append(snailfishNumbersPartOne, newNumber)
		snailfishNumbersPartTwo = append(snailfishNumbersPartTwo, newNumberCopy)

	}

	fmt.Printf("Part One: %v (took %s)\n", part1(snailfishNumbersPartOne), time.Since(start))
	fmt.Printf("Part One: %v (took %s)\n", part2(snailfishNumbersPartTwo), time.Since(start))
}
