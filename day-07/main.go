package main

import (
	"aoc/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	flags := utils.GetFlags()
	input := utils.GetInput(flags.Filepath)

	switch flags.Part {
	case 1:
		part1(input)
	case 2:
		part2(input)
	default:
		fmt.Printf("%d is not a valid part number\n", flags.Part)
	}
}

type Directory struct {
	parent *Directory
	size   int
}

func (d *Directory) addToSize(size int) {
	d.size += size
	if d.parent != nil {
		d.parent.addToSize(size)
	}
}

func parseDirectoryList(input string) []*Directory {
	inputLines := strings.Split(input, "\n")
	dirList := []*Directory{}
	var currDir *Directory

	for _, inputLine := range inputLines {
		if string(inputLine[0:4]) == "$ cd" {
			if string(inputLine[5:]) == ".." {
				currDir = currDir.parent
			} else {
				dir := Directory{parent: currDir, size: 0}
				dirList = append(dirList, &dir)
				currDir = &dir
			}
		} else if string(inputLine[0]) != "$" {
			if string(inputLine[0:3]) != "dir" {
				fileInfo := strings.Split(inputLine, " ")
				fileSize, _ := strconv.Atoi(fileInfo[0])
				currDir.addToSize(fileSize)
			}
		}
	}
	return dirList
}

func part1(input string) {
	directoryList := parseDirectoryList(input)
	total := 0

	for _, directory := range directoryList {
		if directory.size < 100000 {
			total += directory.size
		}
	}

	fmt.Println("part 1 ->", total)
}

func part2(input string) {
	directoryList := parseDirectoryList(input)

	totalSize := 70000000
	updateSize := 30000000
	usedSize := directoryList[0].size
	toDeleteSize := updateSize - (totalSize - usedSize)

	candidateDirSize := math.MaxInt

	for _, directory := range directoryList {
		if directory.size > toDeleteSize && directory.size < candidateDirSize {
			candidateDirSize = directory.size
		}
	}

	fmt.Println("part 2 -> ", candidateDirSize)
}
