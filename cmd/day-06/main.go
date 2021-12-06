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

	inputFile, fileErr := ioutil.ReadFile("./resources/day6.txt")
	if fileErr != nil {
		panic(fileErr)
	}
	rawInput := strings.Split(string(inputFile), ",")

	input = make([]int, 9)

	for _, fish := range rawInput {
		fishAsInt, _ := strconv.Atoi(fish)

		input[fishAsInt] += 1
	}

	partOne()
	partTwo()
}

func partOne() {
	fmt.Println("#### Part One ####")

	fishes := passDays(80)

	fmt.Printf("Total Fish: %d\n", countFish(fishes))
}

func partTwo() {
	fmt.Println("#### Part Two ####")

	fishes := passDays(256)

	fmt.Printf("Total Fish: %d\n", countFish(fishes))

}

func countFish(fishes []int) interface{} {
	count := 0
	for _, fish := range fishes {
		count = count + fish
	}

	return count
}

func passDays(daysToPass int) []int {

	fishes := make([]int, 9)

	for days, count := range input {
		fishes[days] = count
	}

	for x := 0; x < daysToPass; x++ {
		//fmt.Printf("%d, ", x)

		toGiveBirth := fishes[0]
		for y := 1; y < len(fishes); y++ {
			fishes[y-1] = fishes[y]
		}
		fishes[8] = toGiveBirth
		fishes[6] = fishes[6] + toGiveBirth
	}

	return fishes
}