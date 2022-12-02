package main

import (
	"aoc/utils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	flags := utils.GetFlags()
	input := utils.GetInput(flags.Filepath)

	if flags.Part == 1 {
		part1(string(input))
	} else if flags.Part == 2 {
		part2(string(input))
	}
}

func part1(data string) {
	elfCollections := strings.Split(string(data), "\n\n")

	elfLoadRecucer := func(acc int, elfCollection string) int {
		items := strings.Split(elfCollection, "\n")
		var elfLoad int

		for _, item := range items {
			parsedItem, _ := strconv.Atoi(item)
			elfLoad += parsedItem
		}

		if elfLoad > acc {
			return elfLoad
		} else {
			return acc
		}
	}

	heaviestElfLoad := utils.Reduce(elfCollections, elfLoadRecucer, math.MinInt)
	fmt.Println(heaviestElfLoad)
}

func part2(data string) {
	var elfLoadTotal int
	elfLoadList := []int{}
	elfCollections := strings.Split(string(data), "\n\n")

	for _, elfCollection := range elfCollections {
		var elfLoad int
		items := strings.Split(elfCollection, "\n")

		for _, item := range items {
			parsedItem, _ := strconv.Atoi(item)
			elfLoad += parsedItem
		}

		elfLoadList = append(elfLoadList, elfLoad)
	}

	sort.Ints(elfLoadList)

	for i := len(elfLoadList) - 3; i < len(elfLoadList); i++ {
		elfLoadTotal += elfLoadList[i]
	}

	fmt.Println(elfLoadTotal)
}
