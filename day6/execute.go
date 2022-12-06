package day6

import (
	"bufio"
	"fmt"
	"os"
)

func Execute()  {
	file, err := os.Open("day6/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var input string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = scanner.Text()
	}

	var charMap = make(map[string]int, 0)
	intputArray := []rune(input)

	for i, char := range intputArray {
		if i < 14 {
			charMap[string(char)] = charMap[string(char)] + 1
			continue
		} else {
			removeChar := intputArray[i-14]
			charMap[string(removeChar)] = charMap[string(removeChar)] - 1
			if charMap[string(removeChar)] == 0 {
				delete(charMap, string(removeChar))
			}
			charMap[string(char)] = charMap[string(char)] + 1
		}


		if len(charMap) == 14 {
			var find = true
			for _, count := range charMap {
				if count != 1 {
					find = false
				}
			}

			if find {
				println(i + 1)
				break
			}
		}
	}
}
