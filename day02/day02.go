package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, error := os.Open("./input.txt")
	if error != nil {
		log.Fatal(error)
	}
	scanner := bufio.NewScanner(file)

	forwardPosition := 0
	depth := 0
	depth2 := 0
	aim := 0

	for scanner.Scan() {
		command := scanner.Text()
		commandAndValue := strings.Fields(command)
		value, _ := strconv.Atoi(commandAndValue[1])

		switch commandAndValue[0] {
		case "up":
			depth -= value
			aim -= value
		case "down":
			depth += value
			aim += value
		case "forward":
			forwardPosition += value
			depth2 += aim * value
		}
	}

	file.Close()

	// part1:
	fmt.Println("part1:")
	fmt.Println("forwardPosition:", forwardPosition)
	fmt.Println("depth:", depth)
	fmt.Println("forwardPositon * depth:", forwardPosition*depth)

	// part2:
	fmt.Println("part2:")
	fmt.Println("forwardPosition:", forwardPosition)
	fmt.Println("depth2:", depth2)
	fmt.Println("forwardPositon * depth2:", forwardPosition*depth2)
}
