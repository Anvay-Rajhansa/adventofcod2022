package day5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack []string

func Execute() {
	file, err := os.Open("day5/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	/*
		[T]     [D]         [L]
		[R]     [S] [G]     [P]         [H]
		[G]     [H] [W]     [R] [L]     [P]
		[W]     [G] [F] [H] [S] [M]     [L]
		[Q]     [V] [B] [J] [H] [N] [R] [N]
		[M] [R] [R] [P] [M] [T] [H] [Q] [C]
		[F] [F] [Z] [H] [S] [Z] [T] [D] [S]
		[P] [H] [P] [Q] [P] [M] [P] [F] [D]
		 1   2   3   4   5   6   7   8   9
	*/

	var stack1 Stack
	stack1.Push("P")
	stack1.Push("F")
	stack1.Push("M")
	stack1.Push("Q")
	stack1.Push("W")
	stack1.Push("G")
	stack1.Push("R")
	stack1.Push("T")

	var stack2 Stack
	stack2.Push("H")
	stack2.Push("F")
	stack2.Push("R")

	var stack3 Stack
	stack3.Push("P")
	stack3.Push("Z")
	stack3.Push("R")
	stack3.Push("V")
	stack3.Push("G")
	stack3.Push("H")
	stack3.Push("S")
	stack3.Push("D")

	var stack4 Stack
	stack4.Push("Q")
	stack4.Push("H")
	stack4.Push("P")
	stack4.Push("B")
	stack4.Push("F")
	stack4.Push("W")
	stack4.Push("G")

	var stack5 Stack
	stack5.Push("P")
	stack5.Push("S")
	stack5.Push("M")
	stack5.Push("J")
	stack5.Push("H")

	var stack6 Stack
	stack6.Push("M")
	stack6.Push("Z")
	stack6.Push("T")
	stack6.Push("H")
	stack6.Push("S")
	stack6.Push("R")
	stack6.Push("P")
	stack6.Push("L")

	var stack7 Stack
	stack7.Push("P")
	stack7.Push("T")
	stack7.Push("H")
	stack7.Push("N")
	stack7.Push("M")
	stack7.Push("L")

	var stack8 Stack
	stack8.Push("F")
	stack8.Push("D")
	stack8.Push("Q")
	stack8.Push("R")

	var stack9 Stack
	stack9.Push("D")
	stack9.Push("S")
	stack9.Push("C")
	stack9.Push("N")
	stack9.Push("L")
	stack9.Push("P")
	stack9.Push("H")

	var stacks [10]Stack
	stacks[1] = stack1
	stacks[2] = stack2
	stacks[3] = stack3
	stacks[4] = stack4
	stacks[5] = stack5
	stacks[6] = stack6
	stacks[7] = stack7
	stacks[8] = stack8
	stacks[9] = stack9

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		noOfMoves, _ := strconv.Atoi(input[1])
		firstStack, _ := strconv.Atoi(input[3])
		secondStack, _ := strconv.Atoi(input[5])


		// Part 1
		/*for i := 0; i < noOfMoves; i++ {
			val, bol := stacks[firstStack].Pop()
			if bol {
				stacks[secondStack].Push(val)
			}
		}*/

		// Part 2
		if noOfMoves == 1 {
			val, bol := stacks[firstStack].Pop()
			if bol {
				stacks[secondStack].Push(val)
			}
		} else {
			var tempStack Stack
			for i := 0; i < noOfMoves; i++ {
				val, bol := stacks[firstStack].Pop()
				if bol {
					tempStack.Push(val)
				}
			}

			for !tempStack.IsEmpty() {
				val, bol := tempStack.Pop()
				if bol {
					stacks[secondStack].Push(val)
				}
			}
		}
	}

	for _, stack := range stacks {
		val, _ := stack.Pop()
		print(val)
	}
}

// IsEmpty check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Pop Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1 // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index] // Remove it from the stack by slicing it off.
		return element, true
	}
}
