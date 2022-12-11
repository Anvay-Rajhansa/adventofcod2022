package day10

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Execute() {
	file, err := os.Open("day10/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	cycleCount := 1
	Xvalue := 1
	//sum := 0

	var spritePosition [41]string
	spritePosition = refreshSprite(spritePosition, 1)

	var CRTRows = make([][41]string, 0)
	var CurrentCRTrow [41]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		input := scanner.Text()

		if cycleCount > 40 {
			CRTRows = append(CRTRows, CurrentCRTrow)
			var newCRTrow [41]string
			CurrentCRTrow = newCRTrow
			cycleCount = 1
		}

		if strings.Contains(input, "noop") {
			if spritePosition[cycleCount] == "#" {
				CurrentCRTrow[cycleCount] = "#"
			} else {
				CurrentCRTrow[cycleCount] = "."
			}
			cycleCount++
		}

		if strings.Contains(input, "addx") {
			splittedInput := strings.Split(input, " ")

			if spritePosition[cycleCount] == "#" {
				CurrentCRTrow[cycleCount] = "#"
			} else {
				CurrentCRTrow[cycleCount] = "."
			}
			cycleCount++

			if cycleCount > 40 {
				CRTRows = append(CRTRows, CurrentCRTrow)
				var newCRTrow [41]string
				CurrentCRTrow = newCRTrow
				cycleCount = 1
			}

			if spritePosition[cycleCount] == "#" {
				CurrentCRTrow[cycleCount] = "#"
			} else {
				CurrentCRTrow[cycleCount] = "."
			}
			cycleCount++

			no, _ := strconv.Atoi(splittedInput[1])
			Xvalue = Xvalue + no
			spritePosition = refreshSprite(spritePosition, Xvalue)
		}

		//Part 1
		/*sum = sum + checkCycleAndCalculateValue(Xvalue, cycleCount)

		if strings.Contains(input, "noop") {
			cycleCount++
		}

		if strings.Contains(input, "addx") {
			splittedInput := strings.Split(input, " ")
			cycleCount++
			sum = sum + checkCycleAndCalculateValue(Xvalue, cycleCount)
			no, _ := strconv.Atoi(splittedInput[1])
			Xvalue = Xvalue + no
			cycleCount++
		}*/
	}

	CRTRows = append(CRTRows, CurrentCRTrow)

	println("cycle count : ", cycleCount)
	println("xvalue count : ", Xvalue)
	//println("sum count : ", sum)

	for _, row := range CRTRows {
		for _, char := range row {
			print(char)
		}
		println()
	}
}

func checkCycleAndCalculateValue(x int, cycleCount int) int {
	switch cycleCount {
	case 20:
		return x * 20
	case 60:
		return x * 60
	case 100:
		return x * 100
	case 140:
		return x * 140
	case 180:
		return x * 180
	case 220:
		return x * 220
	}
	return 0
}

func refreshSprite(sprite [41]string, hashBeginPosition int) [41]string {
	for i := 1; i < 41; i++ {
		sprite[i] = "."
	}

	if hashBeginPosition < 0 {
		hashBeginPosition = 1
	}

	for i := hashBeginPosition; i < hashBeginPosition+3; i++ {
		if i <= 40 {
			sprite[i] = "#"
		}
	}

	return sprite
}
