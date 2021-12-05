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

	betterInput := make([][]rune, len(input))
	for x, line := range input {
		betterInput[x] = []rune(line)
	}

	o2 := filterInput(betterInput, 0, true)
	co2 := filterInput(betterInput, 0, false)

	decimalO2, _ := strconv.ParseInt(o2, 2, 32)
	decimalCo2, _ := strconv.ParseInt(co2, 2, 32)
	fmt.Printf("O2: %s / %d, CO2: %s / %d, Product: %d\n", o2, decimalO2, co2, decimalCo2, decimalO2 * decimalCo2)
}

func filterInput(inputs [][]rune, iteration int, isO2 bool) string {

	zeroMostSig := make([][]rune, 0)
	oneMostSig := make([][]rune, 0)

	for _, diagnosticLine := range inputs {
		if diagnosticLine[iteration] == '0' {
			zeroMostSig = append(zeroMostSig, diagnosticLine)
		} else {
			oneMostSig = append(oneMostSig, diagnosticLine)
		}
	}

	nextIteration := iteration + 1

	if len(zeroMostSig) > len(oneMostSig) {
		if len(zeroMostSig) == 1 {
			return string(zeroMostSig[0])
		} else {
			if isO2 {
				return filterInput(zeroMostSig, nextIteration, isO2)
			} else {
				return filterInput(oneMostSig, nextIteration, isO2)
			}
		}
	} else if len(zeroMostSig) < len(oneMostSig) {
		if len(oneMostSig) == 1 {
			return string(oneMostSig[0])
		} else {
			if isO2 {
				return filterInput(oneMostSig, nextIteration, isO2)
			} else {
				return filterInput(zeroMostSig, nextIteration, isO2)
			}
		}
	} else { //equal
		if isO2 {
			if len(oneMostSig) == 1 {
				return string(oneMostSig[0])
			} else {
				return filterInput(oneMostSig, nextIteration, isO2)
			}
		} else {
			if len(zeroMostSig) == 1 {
				return string(zeroMostSig[0])
			} else {
				return filterInput(zeroMostSig, nextIteration, isO2)
			}
		}
	}
}