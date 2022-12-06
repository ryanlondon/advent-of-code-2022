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

func parseInput(input string) [][][]int {
	pairAssignments := strings.Split(input, "\n")

	return utils.Map(pairAssignments, func(pairAssignment string) [][]int {
		output := [][]int{[]int{}, []int{}}
		individualAssignments := strings.Split(pairAssignment, ",")
		for i, assignment := range individualAssignments {
			zones := strings.Split(assignment, "-")
			parsedZones := utils.Map(zones, func(zone string) int {
				parsedZone, _ := strconv.Atoi(zone)
				return parsedZone
			})
			output[i] = parsedZones
		}
		return output
	})
}

func part1(input string) {
	parsedInput := parseInput(input)
	var total int

	for _, pairAssignment := range parsedInput {
		min := int(math.Min(float64(pairAssignment[0][0]), float64(pairAssignment[1][0])))
		max := int(math.Max(float64(pairAssignment[0][1]), float64(pairAssignment[1][1])))

		for _, individualAssignment := range pairAssignment {
			if individualAssignment[0] == min && individualAssignment[1] == max {
				total += 1
				break
			}
		}
	}

	fmt.Println("part 1 ->", total)
}

func part2(input string) {
	parsedInput := parseInput(input)
	var total int

	for _, pairAssignment := range parsedInput {
		low := 0
		high := 1

		if pairAssignment[0][0] > pairAssignment[1][0] {
			low = 1
			high = 0
		}

		if pairAssignment[low][1] >= pairAssignment[high][0] {
			total += 1
		}
	}

	fmt.Println("part 2 ->", total)
}
