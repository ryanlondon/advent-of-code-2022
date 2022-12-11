package main

import (
	"aoc/utils"
	"fmt"
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

func part1(input string) {
	start, end, currLen := 0, 0, 0
	cache := make(map[byte]int)

	for end < len(input) {
		lastIdx, exists := cache[input[end]]
		if exists && lastIdx >= start {
			start = lastIdx + 1

		}
		cache[input[end]] = end
		currLen = end - start

		if currLen == 4 {
			break
		}

		end++
	}

	fmt.Println("part 1 -> ", end)
}

func part2(input string) {
	start, end, currLen := 0, 0, 0
	cache := make(map[byte]int)

	for end < len(input) {
		lastIdx, exists := cache[input[end]]
		if exists && lastIdx >= start {
			start = lastIdx + 1

		}
		cache[input[end]] = end
		currLen = end - start

		if currLen == 14 {
			break
		}

		end++
	}

	fmt.Println("part 2 -> ", end)
}
