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

	var maxVisibleScore = 0
	for i := 1; i < len(matrix) - 1; i++ {
		for j := 1; j < len(matrix)-1; j++ {
			val := matrix[i][j]
			var topVisibleScore = 0
			for top := i-1; top >= 0; top-- {
				if matrix[top][j] >= val {
					topVisibleScore++
					break
				}
				topVisibleScore++
			}

			var rightVisibleScore = 0
			for right := j+1; right < len(matrix); right++ {
				if matrix[i][right] >= val {
					rightVisibleScore++
					break
				}
				rightVisibleScore++
			}

			var bottomVisibleScore = 0
			for bottom := i+1; bottom < len(matrix); bottom++ {
				if matrix[bottom][j] >= val {
					bottomVisibleScore++
					break
				}
				bottomVisibleScore++
			}

			var leftVisibleScore = 0
			for left := j-1; left >= 0; left-- {
				if matrix[i][left] >= val {
					leftVisibleScore++
					break
				}
				leftVisibleScore++
			}

			if (topVisibleScore * rightVisibleScore * bottomVisibleScore * leftVisibleScore) > maxVisibleScore {
				maxVisibleScore = topVisibleScore * rightVisibleScore * bottomVisibleScore * leftVisibleScore
			}
		}
	}

	println("visible count : ", maxVisibleScore)
}
