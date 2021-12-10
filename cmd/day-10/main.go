package main

import (
	"fmt"
	"github.com/therealgaryj/advent-of-code-2021/internal/utils"
	"io/ioutil"
	"sort"
	"strings"
)

var scoresPart1 = map[rune]int{')':3,']':57,'}':1197,'>':25137}
var scoresPart2 = map[rune]int{')':1,']':2,'}':3,'>':4}
var symbols = map[rune]rune{'(':')','[':']','{':'}','<':'>'}
var input []string

func main() {

	fmt.Println("#### Day Ten ####")

	inputFile, fileErr := ioutil.ReadFile("./resources/day10.txt")
	if fileErr != nil {
		panic(fileErr)
	}
	input = strings.Split(string(inputFile), "\n")

	partOne()
	partTwo()
}

func partOne() {
	fmt.Println("#### Part One ####")

	totalScore := 0
	for _, noteLine := range input {
		stack := make([]rune, 0)
		for _, symbol := range noteLine {
			if _, ok := symbols[symbol]; ok {
				utils.Push(&stack, symbol)
			} else {
				previous, ok := utils.Pop(&stack)

				if ok {
					if !isOpposite(previous, symbol) {
						totalScore = totalScore + scoresPart1[symbol]
						break
					}
				}
			}
		}
	}

	fmt.Printf("Total: %d\n", totalScore)

}

func partTwo() {
	fmt.Println("#### Part Two ####")

	scores := make([]int, 0)

	for _, noteLine := range input {
		stack := make([]rune, 0)
		score := 0
		for x, symbol := range noteLine {
			if _, ok := symbols[symbol]; ok {
				utils.Push(&stack, symbol)
			} else {
				previous, ok := utils.Pop(&stack)

				if ok {
					if !isOpposite(previous, symbol) {
						break
					}
				}
			}

			if x == (len(noteLine) - 1) {
				// pop the array
				for x := len(stack); x > 0; x-- {
					popped, _ := utils.Pop(&stack)
					opposite := symbols[popped]

					score = (score * 5) + scoresPart2[opposite]
				}

				if score > 0 {
					scores = append(scores, score)
				}
			}
		}
	}

	sort.Ints(scores)

	fmt.Printf("%d", scores[int(len(scores) / 2)])
}

func isOpposite(previous rune, symbol rune) bool {
	for opening, closing := range symbols {
		if closing == symbol && opening == previous {
			return true
		}
	}

	return false
}