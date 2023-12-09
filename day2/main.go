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

func maxValueInSubsets(subsets string) map[string]int {
	maxValues := map[string]int{
		"r": 1,
		"g": 1,
		"b": 1,
	}
	for _, subset := range strings.Split(subsets, ";") {
		colors := strings.Split(subset, ",")
		for _, color := range colors {
			countColor := strings.Split(color, " ")
			count, err := strconv.Atoi(countColor[1])
			check(err)
			color := countColor[2]
			switch color {
			case "red":
				if count > maxValues["r"] {
					maxValues["r"] = count
				}
			case "green":
				if count > maxValues["g"] {
					maxValues["g"] = count
				}
			case "blue":
				if count > maxValues["b"] {
					maxValues["b"] = count
				}
			}
		}
	}
	return maxValues
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
	var powers []int
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), ":")
		gameId, subsets := input[0], input[1]
		id, err := strconv.Atoi(strings.Split(gameId, " ")[1])
		check(err)
		// push all gameIds and "with" the ones that fail
		gamePassed, maxValues := subsetsPassed(subsets), maxValueInSubsets(subsets)
		powers = append(powers, maxValues["r"]*maxValues["g"]*maxValues["b"])
		if gamePassed {
			passedGameIds = append(passedGameIds, id)
		}
	}
	part1, part2 := sum(passedGameIds), sum(powers)
	fmt.Printf("total part1: %d\n", part1)
	fmt.Printf("total part2: %d\n", part2)
}
