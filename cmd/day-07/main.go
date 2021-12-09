package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/therealgaryj/advent-of-code-2021/internal/utils"
)

var input []int
var crabPositions []int

func main() {

	fmt.Println("#### Day One ####")

	inputFile, fileErr := ioutil.ReadFile("./resources/day7.txt")
	if fileErr != nil {
		panic(fileErr)
	}
	rawInput := strings.Split(string(inputFile), ",")

	input = make([]int, len(rawInput))

	largestPos := 0

	for x, subPos := range rawInput {
		subPosAsInt, _ := strconv.Atoi(subPos)

		input[x] = subPosAsInt

		if largestPos < subPosAsInt {
			largestPos = subPosAsInt
		}
	}

	crabPositions = make([]int, largestPos + 1)

	for _, position := range input {
		crabPositions[position] = crabPositions[position] + 1
	}

	partOne()
	partTwo()
}

func partOne() {
	fmt.Println("#### Part One ####")

	smallestTargetIndex := 0
	smallestTargetFuel := 0

	for x := 0; x < len(crabPositions); x++ {
		totalFuelToMove := 0

		for y, numberOfSubs := range crabPositions {
			totalFuelToMove = totalFuelToMove + (utils.Diff(x, y) * numberOfSubs)
		}

		if smallestTargetFuel == 0 || totalFuelToMove < smallestTargetFuel {
			smallestTargetFuel = totalFuelToMove
			smallestTargetIndex = x
		}
	}

	fmt.Printf("Fuel to move them all to position %d: %d\n", smallestTargetIndex, smallestTargetFuel)

}

func partTwo() {
	fmt.Println("#### Part Two ####")

	smallestTargetIndex := 0
	smallestTargetFuel := 0

	for x := 0; x < len(crabPositions); x++ {
		totalFuelToMove := 0

		for y, numberOfSubs := range crabPositions {
			totalFuelToMove = totalFuelToMove + (sumMoves(utils.Diff(x, y)) * numberOfSubs)
		}

		if smallestTargetFuel == 0 || totalFuelToMove < smallestTargetFuel {
			smallestTargetFuel = totalFuelToMove
			smallestTargetIndex = x
		}
	}
	fmt.Printf("Fuel to move them all to position %d: %d\n", smallestTargetIndex, smallestTargetFuel)

}

func sumMoves(moves int) int {

	if moves == 0 {
		return 0
	} else {
		return moves + sumMoves(moves-1)
	}
}