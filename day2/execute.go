package day2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type GameInput struct {
	PlayerInput1 string
	PlayerInput2 string
}

func Execute() {
	file, err := os.Open("day2/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var game = make([]GameInput, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), "")
		game = append(game, GameInput{
			PlayerInput1: input[0], PlayerInput2: input[2],
		})
	}

	valueMap := map[string]int{
		"A": 1, // Rock
		"B": 2, // Paper
		"C": 3, // Scissors
		"X": 1, // Rock
		"Y": 2, // Paper
		"Z": 3, // Scissors
	}

	var myScore = make([]int, 0)
	for _, value := range game {
		//Round 2
		if value.PlayerInput2 == "X" {
			if value.PlayerInput1 == "A" {
				myScore = append(myScore, valueMap["Z"] + 0)
				continue
			}
			if value.PlayerInput1 == "B" {
				myScore = append(myScore, valueMap["X"] + 0)
				continue
			}
			if value.PlayerInput1 == "C" {
				myScore = append(myScore, valueMap["Y"] + 0)
				continue
			}
		}

		if value.PlayerInput2 == "Y" {
			if value.PlayerInput1 == "A" {
				myScore = append(myScore, valueMap["X"] + 3)
				continue
			}
			if value.PlayerInput1 == "B" {
				myScore = append(myScore, valueMap["Y"] + 3)
				continue
			}
			if value.PlayerInput1 == "C" {
				myScore = append(myScore, valueMap["Z"] + 3)
				continue
			}
		}

		if value.PlayerInput2 == "Z" {
			if value.PlayerInput1 == "A" {
				myScore = append(myScore, valueMap["Y"] + 6)
				continue
			}
			if value.PlayerInput1 == "B" {
				myScore = append(myScore, valueMap["Z"] + 6)
				continue
			}
			if value.PlayerInput1 == "C" {
				myScore = append(myScore, valueMap["X"] + 6)
				continue
			}
		}



		//Round 1
		/*if value.PlayerInput1 == "A" && value.PlayerInput2 == "X" {
			myScore = append(myScore, valueMap["X"] + 3)
			continue
		}
		if value.PlayerInput1 == "B" && value.PlayerInput2 == "Y" {
			myScore = append(myScore, valueMap["Y"] + 3)
			continue
		}
		if value.PlayerInput1 == "C" && value.PlayerInput2 == "Z" {
			myScore = append(myScore, valueMap["Z"] + 3)
			continue
		}

		if value.PlayerInput1 == "A" && value.PlayerInput2 == "Y" {
			myScore = append(myScore, valueMap["Y"] + 6)
			continue
		}
		if value.PlayerInput1 == "A" && value.PlayerInput2 == "Z" {
			myScore = append(myScore, valueMap["Z"] + 0)
			continue
		}

		if value.PlayerInput1 == "B" && value.PlayerInput2 == "X" {
			myScore = append(myScore, valueMap["X"] + 0)
			continue
		}
		if value.PlayerInput1 == "B" && value.PlayerInput2 == "Z" {
			myScore = append(myScore, valueMap["Z"] + 6)
			continue
		}

		if value.PlayerInput1 == "C" && value.PlayerInput2 == "X" {
			myScore = append(myScore, valueMap["X"] + 6)
			continue
		}
		if value.PlayerInput1 == "C" && value.PlayerInput2 == "Y" {
			myScore = append(myScore, valueMap["Y"] + 0)
			continue
		}*/
	}

	var total = 0
	for _, i := range myScore {
		total = total + i
	}

	println("my score : ", total)
}
