package main

import (
	"aoc/utils"
	"fmt"
	"strings"
)

func main() {
	flags := utils.GetFlags()
	input := utils.GetInput(flags.Filepath)

	if flags.Part == 1 {
		part1(input)
	} else if flags.Part == 2 {
		part2(input)
	}
}

func splitPair(inputPair string) (string, string) {
	pairArray := strings.Split(inputPair, " ")
	return pairArray[0], pairArray[1]
}

func part1(input string) {
	shapeScoreMap := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	isRightWinner := func(left, right string) bool {
		if left == "A" && right == "Y" {
			return true
		}
		if left == "B" && right == "Z" {
			return true
		}
		if left == "C" && right == "X" {
			return true
		}
		return false
	}

	isTie := func(left, right string) bool {
		if left == "A" && right == "X" {
			return true
		}
		if left == "B" && right == "Y" {
			return true
		}
		if left == "C" && right == "Z" {
			return true
		}
		return false
	}

	inputPairs := strings.Split(input, "\n")
	var total int

	for _, inputPair := range inputPairs {
		left, right := splitPair(inputPair)
		total += shapeScoreMap[right]

		if isRightWinner(left, right) {
			total += 6
		} else if isTie(left, right) {
			total += 3
		}
	}

	fmt.Println(total)
}

func part2(input string) {
	inputPairs := strings.Split(input, "\n")
	var total int

	resultScoreMap := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}

	losingShapeScoreMap := map[string]int{
		"A": 3, // rock -> scissors
		"B": 1, // paper -> rock
		"C": 2, // scissors -> paper
	}

	winningShapeScoreMap := map[string]int{
		"A": 2, // rock -> paper
		"B": 3, // paper -> scissors
		"C": 1, // scissors -> rock
	}

	tieShapeScoreMap := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}

	for _, inputPair := range inputPairs {
		left, right := splitPair(inputPair)
		total += resultScoreMap[right]

		if right == "X" {
			total += losingShapeScoreMap[left]
		} else if right == "Y" {
			total += tieShapeScoreMap[left]
		} else if right == "Z" {
			total += winningShapeScoreMap[left]
		}
	}

	fmt.Println(total)
}
