package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Day2: ")

	const r, g, b = 12, 13, 14
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	var scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	counter := 0
	for scanner.Scan() {
		if counter < 5 {
			input := strings.Split(scanner.Text(), ":")
			gameId, subsets := input[0], input[1]
			// for each game sum the values of r,g,b in each subset
			// for gameIds having sums equal or less then provided rgb values, store those gameIds
			fmt.Print(gameId)
			for _, subset := range strings.Split(subsets, ";") {
				// subsets provided in each game
				fmt.Printf("%s ", subset)
			}
			fmt.Println()
			counter++
		}
	}
}
