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

	inputFile, fileErr := ioutil.ReadFile("./resources/day3.txt")
	if fileErr != nil {
		panic(fileErr)
	}
	input = strings.Split(string(inputFile), "\n")

	partOne()
	partTwo()
}

func partOne() {
	fmt.Println("#### Part One ####")

	mostSigBits := make([]rune, len(input[0]))

	for x := 0; x < len(input[0]); x++ {

		countOfZero := 0

		for y := 0; y < len(input); y++ {
			bit := string(([]rune(input[y]))[x])

			if bit == "0" {
				countOfZero = countOfZero + 1
			}
		}

		if countOfZero > (len(input) / 2) {
			mostSigBits[x] = '1'
		} else {
			mostSigBits[x] = '0'
		}
	}

	gammaString := string(mostSigBits)
	epsilonString := ""

	for _, bit := range mostSigBits {
		if bit == '0' {
			epsilonString = epsilonString + "1"
		} else {
			epsilonString = epsilonString + "0"
		}
	}

	gamma, err1 := strconv.ParseInt(gammaString, 2, 64)
	epsilon, err2 := strconv.ParseInt(epsilonString, 2, 64)

	if err1 != nil || err2 != nil {
		fmt.Printf("Failed to parse binary: %s or %s\n", err1, err2)
		return
	}

	fmt.Printf("Gamma: %s, Epsilon: %s, Product: %d\n", gammaString, epsilonString, gamma * epsilon)
}

func partTwo() {
	fmt.Println("#### Part Two ####")

	mostSigBits := make([]rune, len(input[0]))

	for x := 0; x < len(input[0]); x++ {

		countOfZero := 0

		for y := 0; y < len(input); y++ {
			bit := string(([]rune(input[y]))[x])

			if bit == "0" {
				countOfZero = countOfZero + 1
			}
		}

		if countOfZero > (len(input) / 2) {
			mostSigBits[x] = '1'
		} else {
			mostSigBits[x] = '0'
		}
	}

	gammaString := string(mostSigBits)
	epsilonString := ""

	for _, bit := range mostSigBits {
		if bit == '0' {
			epsilonString = epsilonString + "1"
		} else {
			epsilonString = epsilonString + "0"
		}
	}
}

func findMostSig([][]rune inputs, iteration int) string {

	filtered := make([][]rune)

	for _,
}