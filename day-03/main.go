package main

import (
	"aoc/utils"
	"fmt"
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

func getRepeatedRuneInRucksack(rucksack string) rune {
	mid := len(rucksack) / 2
	left := rucksack[0:mid]
	right := rucksack[mid:]

	leftRuneMap := make(map[rune]bool)

	for _, r := range left {
		leftRuneMap[r] = true
	}

	for _, r := range right {
		if leftRuneMap[r] {
			return r
		}
	}

	return 0
}

func findCommonRuneInRucksacks(rucksacks []string) rune {
	set := make(map[rune]bool)
	subset := make(map[rune]bool)

	for _, rune := range rucksacks[0] {
		set[rune] = true
	}

	for _, rune := range rucksacks[1] {
		if set[rune] {
			subset[rune] = true
		}
	}

	for _, rune := range rucksacks[2] {
		if subset[rune] {
			return rune
		}
	}

	return 0
}

func mapRuneToPriority(r rune) int {
	if r >= 97 {
		return int(r) - 96
	} else {
		return int(r) - 38
	}
}

func part1(input string) {
	rucksacks := strings.Split(input, "\n")
	var total int

	for _, rucksack := range rucksacks {
		repeatedRune := getRepeatedRuneInRucksack(rucksack)
		total += mapRuneToPriority(repeatedRune)
	}

	fmt.Println("part 1 ->", total)
}

func part2(input string) {
	rucksacks := strings.Split(input, "\n")
	var total int

	for i := 0; i < len(rucksacks); i = i + 3 {
		rucksackGroup := rucksacks[i : i+3]
		commonRune := findCommonRuneInRucksacks(rucksackGroup)
		total += mapRuneToPriority(commonRune)
	}

	fmt.Println("part 2 ->", total)
}
