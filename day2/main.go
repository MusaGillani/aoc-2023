package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const red, green, blue = 12, 13, 14

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func sum(arr []int) int {
	total := 0
	for _, val := range arr {
		total += val
	}
	return total
}

func subsetsPassed(subsets string) bool {
	passed := true
SubsetsLoop:
	for _, subset := range strings.Split(subsets, ";") {
		colors := strings.Split(subset, ",")
		for _, color := range colors {
			countColor := strings.Split(color, " ")
			count, err := strconv.Atoi(countColor[1])
			check(err)
			color := countColor[2]
			switch color {
			case "red":
				passed = count <= red
			case "green":
				passed = count <= green
			case "blue":
				passed = count <= blue
			}
			if !passed {
				break SubsetsLoop
			}
		}
	}
	return passed
}

func main() {
	fmt.Println("Day2: ")

	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	var scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var passedGameIds []int
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), ":")
		gameId, subsets := input[0], input[1]
		id, err := strconv.Atoi(strings.Split(gameId, " ")[1])
		check(err)
		// push all gameIds and skip the ones that fail
		if subsetsPassed(subsets) {
			passedGameIds = append(passedGameIds, id)
		}
	}
	fmt.Printf("total: %d\n", sum(passedGameIds))
}
