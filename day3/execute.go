package day3

import (
	"bufio"
	"fmt"
	"os"
)

type backpack struct {
	comp1 string
	comp2 string
}

type backpackGroup struct {
	backpack1 string
	backpack2 string
	backpack3 string
}

func Execute() {
	//Part 1
	/*file, err := os.Open("day3/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var backpacks = make([]backpack, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		backpacks = append(backpacks, backpack{
			comp1: input[:len(input)/2], comp2: input[len(input)/2:],
		})
	}

	var priorityItems = make([]int, 0)
	for _, backpackItem := range backpacks {
		var commonMap = make(map[string]int)
		for _, item := range []rune(backpackItem.comp1) {
			for _, item2 := range []rune(backpackItem.comp2) {
				if string(item) == string(item2) {
					if commonMap[string(item)] == 0 {
						commonMap[string(item)] = valueMap[string(item)]
					}

					break
				}
			}
		}

		for _, commonChar := range commonMap {
			priorityItems = append(priorityItems, commonChar)
		}
	}

	var total = 0
	for _, items := range priorityItems {
		total = total + items
	}

	println("total : ", total)*/

	// Part 2
	file, err := os.Open("day3/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var backPacksGroup = make([]backpackGroup, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var backPackGroupVar = backpackGroup{}
		backPackGroupVar.backpack1 = scanner.Text()

		if scanner.Scan() {
			backPackGroupVar.backpack2 = scanner.Text()
		}

		if scanner.Scan() {
			backPackGroupVar.backpack3 = scanner.Text()
		}

		backPacksGroup = append(backPacksGroup, backPackGroupVar)
	}

	var priorityItems = make([]int, 0)

	for _, backpackItem := range backPacksGroup {
		var tempMap = make(map[string]int)
		var commonMap = make(map[string]int)

		for _, item := range []rune(backpackItem.backpack1) {
			for _, item2 := range []rune(backpackItem.backpack2) {
				if string(item) == string(item2) {
					if tempMap[string(item)] == 0 {
						tempMap[string(item)] = valueMap[string(item)]
					}
					break

				}
			}

			for _, item2 := range []rune(backpackItem.backpack3) {
				if string(item) == string(item2) {
					if tempMap[string(item)] != 0 {
						commonMap[string(item)] = valueMap[string(item)]
					}

					break
				}
			}

		}

		for _, commonChar := range commonMap {
			priorityItems = append(priorityItems, commonChar)
		}
	}

	var total = 0
	for _, items := range priorityItems {
		total = total + items
	}

	println("total : ", total)
}

var valueMap = map[string]int{
	"A": 27,
	"B": 28,
	"C": 29,
	"D": 30,
	"E": 31,
	"F": 32,
	"G": 33,
	"H": 34,
	"I": 35,
	"J": 36,
	"K": 37,
	"L": 38,
	"M": 39,
	"N": 40,
	"O": 41,
	"P": 42,
	"Q": 43,
	"R": 44,
	"S": 45,
	"T": 46,
	"U": 47,
	"V": 48,
	"W": 49,
	"X": 50,
	"Y": 51,
	"Z": 52,

	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
	"i": 9,
	"j": 10,
	"k": 11,
	"l": 12,
	"m": 13,
	"n": 14,
	"o": 15,
	"p": 16,
	"q": 17,
	"r": 18,
	"s": 19,
	"t": 20,
	"u": 21,
	"v": 22,
	"w": 23,
	"x": 24,
	"y": 25,
	"z": 26,
}
