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

func main() {
	fmt.Println("Day1 ")

	file, err := os.Open("./input.txt")
	check(err)
	// fmt.Printf("%s", data)
	var scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	// var textLines []stringln
	counter := 0
	for scanner.Scan() {
		// textLines = append(textLines, scanner.Text())
		if counter < 5 {
			input := scanner.Text()
			// inputLen := len(input)
			fmt.Printf("word %s: ", input)
			digitCounter, firstNumber, lastNumber := 0, 0, 0
			for _, char := range input {
				if unicode.IsDigit(char) {
					digit := int(char - '0')
					if digit >= 0 || digit <= 9 {
						digitCounter++
						if digitCounter == 1 {
							// find the last number in the string in this case
							firstNumber = digit
						} else {
							lastNumber = digit
						}
						// fmt.Print(digit)
					}
				}
			}
			fmt.Printf("first %d last %d ", firstNumber, lastNumber)
			fmt.Println()
		}
		counter++
	}
	file.Close()

}
