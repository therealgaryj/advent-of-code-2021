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

	input = make([]int, len(rawInput))

	for x, fish := range rawInput {
		fishAsInt, _ := strconv.Atoi(fish)

		input[x] = fishAsInt
	}

	partOne()
	partTwo()
}

func partOne() {
	fmt.Println("#### Part One ####")

	fishes := passDays(80)

	fmt.Printf("Total Fish: %d\n", len(fishes))
}

func passDays(days int) []int {

	fishes := make([]int, len(input))

	for x, fish := range input {
		fishes[x] = fish
	}

	for x := 0; x < days; x++ {
		fmt.Printf("%d, ", x)
		for y, fish := range fishes {
			if fish == 0 {
				fishes[y] = 6
				fishes = append(fishes, 8)
			} else {
				fishes[y] = fish - 1
			}
		}
	}

	return fishes
}
func partTwo() {
	fmt.Println("#### Part Two ####")

	fishes := passDays(256)

	fmt.Printf("Total Fish: %d\n", len(fishes))

}
