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

func loadFileToStringSlice(path string) []string {
	file, error := os.Open(path)
	if error != nil {
		log.Fatal(error)
	}
	scanner := bufio.NewScanner(file)

	var inputData []string

	for scanner.Scan() {
		line := scanner.Text()
		inputData = append(inputData, line)
	}

	file.Close()

	return inputData
}

func filterValues(list []string, position int, value byte) []string {
	var newList []string

	for _, line := range list {
		if line[position] == value {
			newList = append(newList, line)
		}
	}
	return newList
}

func getGasRating(list []string, gas string) int64 {
	output := list
	var zero, one byte = '0', '1'

	if gas == "co2" {
		zero = '1'
		one = '0'
	}

	for index := 0; index < len(list[index]); index++ {
		var zeros, ones int

		for _, value := range output {
			if value[index] == '0' {
				zeros += 1
			} else {
				ones += 1
			}
		}

		if ones >= zeros {
			// for oxygen it passes when amount of 1 is not less than amount of 0, then it leaves those with 1
			// for co2 it passes when amount of 0 is not less than amount of 1, then it leaves those with 0
			output = filterValues(output, index, one)
		} else {
			output = filterValues(output, index, zero)
		}
		if len(output) == 1 {
			break
		}
	}

	number, _ := strconv.ParseInt(output[0], 2, 64)
	return number
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
	stringDataOxygen := loadFileToStringSlice("./input.txt")
	stringDataCO2 := loadFileToStringSlice("./input.txt")

	oxygenRating := getGasRating(stringDataOxygen, "oxygen")
	co2Rating := getGasRating(stringDataCO2, "co2")

	fmt.Println("Oxygen generator rating", oxygenRating)
	fmt.Println("CO2 scrubber rating", co2Rating)
	fmt.Println("Life support rating", oxygenRating*co2Rating)
}
