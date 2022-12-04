package day4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type room struct {
	pair1 []int
	pair2 []int
}

func Execute() {
	file, err := os.Open("day4/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var rooms = make([]room, 0)
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var room room
		for count, pair := range strings.Split(scanner.Text(), ",") {
			startAndEnd := strings.Split(pair, "-")

			var items = make([]int, 0)

			start,_ := strconv.Atoi(startAndEnd[0])
			end,_ := strconv.Atoi(startAndEnd[1])

			for i := start; i <= end; i++ {
				items = append(items, i)
			}

			if count == 0 {
				room.pair1 = items
			} else {
				room.pair2 = items
			}

		}
		rooms = append(rooms, room)
	}

	var fullyContainCount = 0

	for _, room := range rooms {

		var allpresent = true

		if len(room.pair1) > len(room.pair2) {
			var tempMap = make(map[int]int)
			for _, item := range room.pair2 {
				for _, item1 := range room.pair1 {
					if item == item1 {
						tempMap[item] = 1
						break
					}
				}
			}

			// Part 1
			//if len(room.pair2) != len(tempMap) {
			//	allpresent = false
			//}

			// part 2
			if len(tempMap) == 0 {
				allpresent = false
			}

		} else {
			var tempMap = make(map[int]int)
			for _, item := range room.pair1 {
				for _, item1 := range room.pair2 {
					if item == item1 {
						tempMap[item] = 1
						break
					}
				}
			}

			// Part 1
			//if len(room.pair1) != len(tempMap) {
			//	allpresent = false
			//}

			// part 2
			if len(tempMap) == 0 {
				allpresent = false
			}
		}

		if allpresent {
			fullyContainCount++
		}

	}

	println("Count : ", fullyContainCount)
}
