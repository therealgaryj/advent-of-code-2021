package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var input []string

func main() {

	fmt.Println("#### Day One ####")

	inputFile, fileErr := ioutil.ReadFile("./resources/day2.txt")
	if fileErr != nil {
		panic(fileErr)
	}
	input = strings.Split(string(inputFile), "\n")

	partOne()
	partTwo()
}

func partOne() {
	fmt.Println("#### Part One ####")

	depth := 0
	distance := 0


	for i, rawInstruction := range input {
		instruction := strings.Split(string(rawInstruction), " ")

		command := instruction[0]
		value, err := strconv.Atoi(instruction[1])
		if err != nil {
			fmt.Printf("Failure on line %d due to %s\n", i, err.Error())
			return
		}

		switch command {
		case "forward":
			distance = distance + value
		case "up":
			depth = depth - value
		case "down":
			depth = depth + value

		}
	}

	fmt.Printf("Depth: %d, Distance: %d, Product: %d\n", depth, distance, depth * distance)
}

func partTwo() {
	fmt.Println("#### Part Two ####")

	depth := 0
	distance := 0
	aim := 0

	for i, rawInstruction := range input {
		instruction := strings.Split(string(rawInstruction), " ")

		command := instruction[0]
		value, err := strconv.Atoi(instruction[1])
		if err != nil {
			fmt.Printf("Failure on line %d due to %s\n", i, err.Error())
			return
		}

		switch command {
		case "forward":
			distance = distance + value
			depth = depth + (aim * value)
		case "up":
			aim = aim - value
		case "down":
			aim = aim + value

		}
	}

	fmt.Printf("Depth: %d, Distance: %d, Product: %d\n", depth, distance, depth * distance)
}