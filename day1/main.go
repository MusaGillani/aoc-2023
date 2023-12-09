package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func sum(arr []int) int {
	total := 0
	for _, val := range arr {
		total += val
	}
	return total
}

func main() {
	fmt.Println("Day1 ")

	file, err := os.Open("./input.txt")
	check(err)
	var scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lineNumbers []int
	for scanner.Scan() {
		input := scanner.Text()
		digitCounter, firstNumber, lastNumber := 0, 0, 0
		for _, char := range input {
			if unicode.IsDigit(char) {
				digit := int(char - '0')
				if digit >= 0 || digit <= 9 {
					digitCounter++
					if digitCounter == 1 {
						firstNumber = digit
					} else {
						// set digit to last number each time until the loop ends
						// last value set will be the last digit
						lastNumber = digit
					}
				}
			}
		}
		var lineNumber int
		if digitCounter == 1 {
			lineNumber = (firstNumber * 10) + firstNumber
		} else {
			lineNumber = (firstNumber * 10) + lastNumber
		}
		lineNumbers = append(lineNumbers, lineNumber)
	}
	file.Close()
	fmt.Printf("total: %d\n", sum(lineNumbers))
}
