package day9

//Solution is copied as its from - https://topaz.github.io/paste/#XQAAAQBPCAAAAAAAAAA4GEiZzRd1JAgz+whYRQxSFI7XvmlfhtGDinguAj8sFyd/fs4fcP1q4v72r/JJnQ8R1o0pqxXdtafKwg2InUzrW7ApC/rEuOF1AbF4HN6w14kvamBsN+I67KzmaghRL3riX4YxuEJVuS0IiOjaG6D139lE+HE4RkWtGYEMNaPW2Di5h/ObvxaPZxocTCxDTTH5tTZhq5+KXUmRBwrgJL5GOJispoFd8JbQKwv+ZqmwXhLFUAsY8mV7S9RXMVKnfrk4vF62Ek+q/28rqSaxB+vM73NY8z1qSYTlS9Yc16GsPvxkPztZpAEumfzSJNgFZLc0G7Q8RcOUSYT/4OSkAkYPxpl6GFdDwL3zs9nyXjKv881QlyIBt1kJilOHWIXU0fDbWBuUcjJQn+lveH62/hudr7+ATI37Wg75Xi0VJShQyBSGFFOTK4GQeNRNLQKVfTivTwmMcWxEnG2QTjpAa31wziut/EZdBX3pLvAAHh93QWFtxFU3mCtj7AsRzGneFPmNaMGyaFVKdXdxtARCoGQrPfkQ2DwMU419/hO0J93Eqzan/snO7BmtUQY1veAIY4h7zccSU9wQNQgYOE9/A6JULkLNXHok+m0SkVDYRYU8G0ECd6Ucvw6HlmqVzmjhPb4qUlScFS/m3lFOpEvEEVFzqZ39MYPccK97Jl4z1rmrq+XVKMxF/3R8tsS/TOryvEuPEOaLEkkCTSmgrQFWRdGD1gkqHOiloM/YmlqBvAad2lHDiWNgThXSuvOFJHy9bHfK4M3OpJWksjCFNsgFOoscU0qGwe7B4q4s2/9WejgL89RKLngaH6U09OFm0UQzjHmWUFmkSnTleNZCVwOmc/gL+mtbJErDttHMWhMsY8iLDwfIpypnnMuqF6ceswaIP3KLzocuBQ/g/T3/tXwb9qn+kL9NlkImRclCaRdfb/dTTcgsrnJ6kay+wmW8YAzirh4/0M9YzWlKyp6y9R0XuicCHyr9P7UTOZC22dF0YZBGL6JNc3xevW5+jHgzDE/QuwOj4PoFxsQkzrIQWEu+VFYxmd+HthFG9GRZaXzKl4TLVKcDF9ixxud6cUeGnkuQWhfy88v/rqp0VA==
//https://www.reddit.com/r/adventofcode/comments/zgnice/2022_day_9_solutions/

import (
	"bufio"
	"fmt"
	set "github.com/deckarep/golang-set"
	"log"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	direction string
	dist      int
}

type point struct {
	x, y int
}

func Execute() {
	file, err := os.Open("day9/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var instructions []instruction
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		dist, _ := strconv.Atoi(line[1])
		instructions = append(instructions, instruction{line[0], dist})
	}

	fmt.Println("Part 1:", solvePart1(instructions))
	fmt.Println("Part 2:", solvePart2(instructions))
}

func solvePart1(instructions []instruction) int {
	seen := set.NewSet()
	seen.Add(point{0, 0})
	head := point{0, 0}
	tail := point{0, 0}

	for _, in := range instructions {
		for i := 0; i < in.dist; i++ {
			head = moveByOne(in.direction, head)
			if !adjacent(head, tail) {
				tail = chaseAfter(head, tail)
			}
			seen.Add(tail)
		}
	}

	return seen.Cardinality()
}

func moveByOne(direction string, p point) point {
	if direction == "R" {
		p.x += 1
	} else if direction == "L" {
		p.x -= 1
	} else if direction == "U" {
		p.y += 1
	} else {
		p.y -= 1
	}

	return p
}

func adjacent(head point, tail point) bool {
	for dx := -1; dx < 2; dx++ {
		for dy := -1; dy < 2; dy++ {
			nei := point{head.x + dx, head.y + dy}
			if nei == tail {
				return true
			}
		}
	}

	return false
}

func chaseAfter(head point, tail point) point {
	if tail.x != head.x {
		if tail.x < head.x {
			tail.x += 1
		} else {
			tail.x -= 1
		}
	}

	if tail.y != head.y {
		if tail.y < head.y {
			tail.y += 1
		} else {
			tail.y -= 1
		}
	}

	return tail
}

func solvePart2(instructions []instruction) int {
	const n = 10
	seen := set.NewSet()
	seen.Add(point{0,0})
	knots := [n]point{}

	for _, in := range instructions {
		for i := 0; i < in.dist; i++ {
			knots[0] = moveByOne(in.direction, knots[0])
			for j := 1; j < n; j++ {
				if !adjacent(knots[j-1], knots[j]) {
					knots[j] = chaseAfter(knots[j-1], knots[j])
				}
			}
			seen.Add(knots[n-1])
		}
	}

	return seen.Cardinality()
}
//type rope struct {
//	head int
//	tail int
//}
//
//func Execute() {
//	file, err := os.Open("day9/input.txt")
//	if err != nil {
//		fmt.Println(err)
//	}
//	defer file.Close()
//
//	//ropes := make([]rope, 0)
//	//ropes = append(ropes, rope{0,0})
//
//	board := set.NewSet()
//	board.Add(rope{0,0})
//	head := rope{0, 0}
//	tail := rope{0, 0}
//
//	scanner := bufio.NewScanner(file)
//	for scanner.Scan() {
//		input := strings.Split(scanner.Text(), " ")
//		no, _ := strconv.Atoi(input[1])
//		for i := 0; i < no; i++ {
//			if input[0] == "R" {
//				head.head += 1
//			} else if input[0] == "L" {
//				head.head -= 1
//			} else if input[0] == "U" {
//				head.tail += 1
//			} else {
//				head.tail -= 1
//			}
//
//			var moveTail = false
//			for hh := -1; hh < 2; hh++ {
//				for hy := -1; hy < 2; hy++ {
//					nei := rope{head.head + hh, head.tail + hy}
//					if nei == tail {
//						moveTail = true
//						break
//					}
//				}
//			}
//
//			if moveTail {
//				if tail.head != head.head {
//					if tail.head < head.head {
//						tail.head += 1
//					} else {
//						tail.head -= 1
//					}
//				}
//
//				if tail.tail != head.tail {
//					if tail.tail < head.tail {
//						tail.tail += 1
//					} else {
//						tail.tail -= 1
//					}
//				}
//			}
//
//			board.Add(tail)
//		}
//
//	}
//
//	println("visited count : ", board.Cardinality())
//}
