package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Execute() {
	file, err := os.Open("day1/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var elfCal = make([]int, 0)
	scanner := bufio.NewScanner(file)
	var total = 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			elfCal = append(elfCal, total)
			total = 0
			continue
		}

		cal, _ := strconv.Atoi(scanner.Text())
		total = total + cal
	}

	var firstMaximum = 0
	var secondMaximum = 0
	var thirdMaximum = 0
	for index, element := range elfCal {
		if element > firstMaximum {
			thirdMaximum = secondMaximum
			secondMaximum = firstMaximum
			firstMaximum = element
		} else if element > secondMaximum {
			thirdMaximum = secondMaximum
			secondMaximum = element
		} else if element > thirdMaximum {
			thirdMaximum = element
		}
		fmt.Println(fmt.Sprintf("Elf %v is having total cal %v : 1 max is  -> %v : 2 max is  -> %v : 3 max is -> %v",
			index, element, firstMaximum, secondMaximum, thirdMaximum))
	}

	fmt.Println(fmt.Sprintf("Total top 3 : %v", firstMaximum + secondMaximum + thirdMaximum))
}
