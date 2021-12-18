package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// fun fact - this ./ is where you are, not program
func loadFileToSlice(path string) []int {
	file, error := os.Open(path)
	if error != nil {
		log.Fatal(error)
	}
	scanner := bufio.NewScanner(file)

	var inputData []int

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		inputData = append(inputData, num)
	}

	file.Close()

	return inputData
}

func main() {
	// part1:
	input := loadFileToSlice("./input.txt")

	counter := 0

	for i := range input {
		if i == 0 {
			continue
		}
		if input[i] > input[i-1] {
			counter++
		}
	}

	fmt.Println("counter", counter)

	// part2:
	sumCounter := 0
	prevSum := 0

	for i := range input {
		if i == 0 {
			prevSum = input[i] + input[i+1] + input[i+2]
			continue
		}

		currentSum := input[i]

		if i+1 < len(input) {
			currentSum += input[i+1]

			if i+2 < len(input) {
				currentSum += input[i+2]
			}
		}

		if currentSum > prevSum {
			sumCounter++
		}

		prevSum = currentSum
	}

	fmt.Println("Sum counter", sumCounter)

}
