package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")

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
