package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")

	heaviestElfLoad := math.MinInt
	elfCollections := strings.Split(string(data), "\n\n")

	for _, elfCollection := range elfCollections {
		var elfLoad int
		items := strings.Split(elfCollection, "\n")

		for _, item := range items {
			parsedItem, _ := strconv.Atoi(item)
			elfLoad += parsedItem
		}

		if elfLoad > heaviestElfLoad {
			heaviestElfLoad = elfLoad
		}
	}

	fmt.Println(heaviestElfLoad)
}
