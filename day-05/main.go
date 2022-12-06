package main

import (
	"aoc/utils"
	"fmt"
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

func parseStacks(input string) []utils.Stack[string] {
	textSections := strings.Split(input, "\n\n")
	stackLayers := strings.Split(textSections[0], "\n")
	stackCount := len(strings.Fields(stackLayers[len(stackLayers)-1]))
	stackList := make([]utils.Stack[string], stackCount)

	for i := len(stackLayers) - 2; i >= 0; i-- {
		stackListIndex := 0

		for j := 1; j < len(stackLayers[i]); j += 4 {
			letter := string(stackLayers[i][j])

			if letter != " " {
				stackList[stackListIndex].Push(letter)

			}

			stackListIndex++
		}
	}

	return stackList
}

type Move struct {
	Count int
	From  int
	To    int
}

func parseMoves(input string) []Move {
	sections := strings.Split(input, "\n\n")
	lines := strings.Split(sections[1], "\n")
	moveList := make([]Move, len(lines))

	for i, line := range lines {
		fields := strings.Fields(line)
		count, _ := strconv.Atoi(fields[1])
		from, _ := strconv.Atoi(fields[3])
		to, _ := strconv.Atoi(fields[5])
		moveList[i].Count = count
		moveList[i].From = from - 1
		moveList[i].To = to - 1
	}

	return moveList
}

func part1(input string) {
	stacks := parseStacks(input)
	moves := parseMoves(input)
	output := ""

	for _, move := range moves {
		for i := 0; i < move.Count; i++ {
			stacks[move.To].Push(stacks[move.From].Pop())
		}
	}

	for _, stack := range stacks {
		output += stack.Peek()
	}

	fmt.Println("part 1 ->", output)
}

func part2(input string) {
	stacks := parseStacks(input)
	moves := parseMoves(input)
	output := ""

	for _, move := range moves {
		stacks[move.To].PushN(stacks[move.From].PopN(move.Count))

	}

	for _, stack := range stacks {
		output += stack.Peek()
	}

	fmt.Println("part 2 -> ", output)
}
