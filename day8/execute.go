package day8

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Execute() {
	file, err := os.Open("day8/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var matrix [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matrixi := make([]int, 0)
		for _, char := range []rune(scanner.Text()) {
			no, _ := strconv.Atoi(string(char))
			matrixi = append(matrixi, no)
		}
		matrix = append(matrix, matrixi)
	}

	var visibalCount = (len(matrix) * 4) - 4
	for i := 1; i < len(matrix) - 1; i++ {
		for j := 1; j < len(matrix)-1; j++ {
			val := matrix[i][j]
			var topVisible = true
			for top := 0; top < i; top++ {
				if matrix[top][j] >= val {
					topVisible = false
					break
				}
			}
			if topVisible {
				visibalCount++
				continue
			}

			var rightVisible = true
			for right := j+1; right < len(matrix); right++ {
				if matrix[i][right] >= val {
					rightVisible = false
					break
				}
			}
			if rightVisible {
				visibalCount++
				continue
			}

			var bottomVisible = true
			for bottom := i+1; bottom < len(matrix); bottom++ {
				if matrix[bottom][j] >= val {
					bottomVisible = false
					break
				}
			}
			if bottomVisible {
				visibalCount++
				continue
			}

			var leftVisible = true
			for left := 0; left < j; left++ {
				if matrix[i][left] >= val {
					leftVisible = false
					break
				}
			}
			if leftVisible {
				visibalCount++
				continue
			}
		}
	}

	println("visible count : ", visibalCount)
}
