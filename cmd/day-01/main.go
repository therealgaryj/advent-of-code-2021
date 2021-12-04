package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var input []int

func main() {

	fmt.Println("#### Day One ####")

	inputFile, fileErr := ioutil.ReadFile("./resources/day1.txt")
	if fileErr != nil {
		panic(fileErr)
	}
	rawInput := strings.Split(string(inputFile), "\n")

	for i, x := range rawInput {
		depth, err :=  strconv.Atoi(x)
		if err != nil {
			fmt.Println("Input format issue for value %s on line %n", x, i)
			return
		}

		input = append(input, depth)
	}
	partOne()
	partTwo()
}

func partOne() {
	fmt.Println("#### Part One ####")

	previousDepth := 0
	countOfIncreases := 0

	for i, depth := range input {
		if i > 0 && previousDepth < depth {
			countOfIncreases++
		}

		previousDepth = depth
	}

	fmt.Printf("Number of increased depths: %d\n", countOfIncreases)
}

func partTwo() {
	fmt.Println("#### Part Two ####")

	sumOfPreviousWindow := 0
	countOfIncreases := 0

	for i := 0; i <= len(input) - 3; i++ {

		sumOfCurrentWindow := 0
		for j := 0; j < 3; j++ {
			sumOfCurrentWindow = sumOfCurrentWindow + input[j + i]
		}

		if i > 0 && sumOfCurrentWindow > sumOfPreviousWindow {
			countOfIncreases++
		}

		sumOfPreviousWindow = sumOfCurrentWindow
	}

	fmt.Printf("Number of increased depths: %d\n", countOfIncreases)
}