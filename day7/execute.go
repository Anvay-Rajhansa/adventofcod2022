package day7

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type advfile struct {
	filename string
	size int
}

type directory struct {
	directoryName string
	file []advfile
	directories []*directory
	parentDirectory *directory
	totalSize int
}

func Execute() {
	file, err := os.Open("day7/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var currentDirectory = &directory{
		directoryName: "root",
		parentDirectory: nil,
		directories: nil,
		file: nil,
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()

		if strings.Contains(input, "$ cd /") || strings.Contains(input, "$ ls") {
			continue
		}

		if strings.Contains(input, "$ cd ..") {
			currentDirectory = currentDirectory.parentDirectory
			continue
		}

		if strings.Contains(input, "dir") {
			addDirectory(currentDirectory, strings.Split(input, " ")[1])
			continue
		}

		if strings.Contains(input, "$ cd") {
			travelDirectory := findDirectory(currentDirectory, strings.Split(input, " ")[2])
			currentDirectory = travelDirectory
			continue
		}


		splittedInput := strings.Split(input, " ")
		fileSize, _ := strconv.Atoi(splittedInput[0])
		var newFile = advfile{
			filename: splittedInput[1],
			size: fileSize,
		}

		addFile(currentDirectory, newFile)

	}

	for currentDirectory.directoryName != "root" {
		currentDirectory = currentDirectory.parentDirectory
	}
	rootDirectory := currentDirectory

	calculateTotalSizeForDirectories(rootDirectory)
	println("parent total", rootDirectory.totalSize)

	var total = new(int)
	*total = 0
	println("deletion total count", deletionCandidateDirectories(rootDirectory, total))

	var selected = new(int)
	*selected = rootDirectory.totalSize
	println("selected total count", smallestDeleteRequired(rootDirectory, selected))
}

func addFile(current *directory, file advfile) {
	current.totalSize = current.totalSize + file.size
	current.file = append(current.file, file)
}

func addDirectory(current *directory, newDirectoryName string) {
	var newDirectory = directory{
		directoryName: newDirectoryName,
		parentDirectory: current,
		directories: nil,
		file: nil,
	}

	current.directories = append(current.directories, &newDirectory)
}

func findDirectory(current *directory, findaName string) *directory  {
	for _, directory1 := range current.directories {
		if directory1.directoryName == findaName {
			return directory1
		}
	}

	return &directory{}
}

func calculateTotalSizeForDirectories(directory2 *directory) {
	for _, d := range directory2.directories {
		calculateTotalSizeForDirectories(d)
		d.parentDirectory.totalSize = d.parentDirectory.totalSize + d.totalSize
	}
}

func deletionCandidateDirectories(directory2 *directory, totalCount *int) int {
	for _, d := range directory2.directories {
		deletionCandidateDirectories(d, totalCount)
		if d.totalSize < 100000 {
			*totalCount = *totalCount + d.totalSize
		}
	}

	return *totalCount
}

// Part 2
// 70000000 - 44125990 = 2,58,74,010
// 30000000 - 25874010 = 41,25,990
func smallestDeleteRequired(directory2 *directory, selected *int) int {
	for _, d := range directory2.directories {
		smallestDeleteRequired(d, selected)
		if d.totalSize < *selected && d.totalSize > 4125990 {
			*selected = d.totalSize
		}
	}

	return *selected
}