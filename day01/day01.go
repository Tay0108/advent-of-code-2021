package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Day 01:")
	file, error := os.Open("./input.txt") 
	// fun fact - this ./ is where you are, not program
	if error != nil {
		log.Fatal(error)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	counter := 0
	prevValue := 124

	for scanner.Scan() {
		currentValue, _ := strconv.Atoi(scanner.Text())
		if(currentValue > prevValue) {
			fmt.Println(currentValue, " is higher than ", prevValue)
			counter++
		}
		prevValue = currentValue
	}

	if error := scanner.Err(); error != nil {
		log.Fatal(error)
	}

	fmt.Println(counter)
}