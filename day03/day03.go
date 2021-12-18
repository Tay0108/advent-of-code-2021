package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func loadFileToSlice(path string) []int64 {
	file, error := os.Open(path)
	if error != nil {
		log.Fatal(error)
	}
	scanner := bufio.NewScanner(file)

	var inputData []int64

	for scanner.Scan() {
		num, _ := strconv.ParseInt(scanner.Text(), 2, 64)
		inputData = append(inputData, num)
	}

	file.Close()

	return inputData
}

func bitHasNotLessOnesThanZeroes(position int, bitCount [12]int, numberOfLines int) bool {
	return numberOfLines/2-bitCount[position] <= 0
}

func main() {
	bitCount := [12]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	numSlice := loadFileToSlice("./input.txt")
	lineCounter := 0

	for j := 0; j < len(numSlice); j++ {
		num32 := int(numSlice[j])

		for i := 1; i <= len(bitCount); i++ {
			reminder := num32 % 2
			bitCount[len(bitCount)-i] += reminder
			num32 = num32 / 2
		}
		lineCounter++
	}

	gamma := 0
	epsilon := 0

	for i := range bitCount {
		if lineCounter/2-bitCount[i] > 0 {
			gamma = gamma<<1 + 1
			epsilon = epsilon << 1
		} else {
			gamma = gamma << 1
			epsilon = epsilon<<1 + 1
		}
	}

	fmt.Println("gamma:", gamma)
	fmt.Println("epsilon:", epsilon)
	fmt.Println("Power comsumption:", gamma*epsilon)

	// part2:

}
