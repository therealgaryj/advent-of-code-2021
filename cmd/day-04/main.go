	package main

import (
	"fmt"
	"github.com/therealgaryj/advent-of-code-2021/internal/utils"
	"io/ioutil"
	"strconv"
	"strings"
)

var numbers = []int64{79,9,13,43,53,51,40,47,56,27,0,14,33,60,61,36,72,48,83,42,10,86,41,75,16,80,15,93,95,45,68,96,84,11,85,63,18,31,35,74,71,91,39,88,55,6,21,12,58,29,69,37,44,98,89,78,17,64,59,76,54,30,65,82,28,50,32,77,66,24,1,70,92,23,8,49,38,73,94,26,22,34,97,25,87,19,57,7,2,3,46,67,90,62,20,5,52,99,81,4}
var input []string
var boards [][][]int64
var results [][][]bool

func main() {

	fmt.Println("#### Day One ####")

	inputFile, fileErr := ioutil.ReadFile("./resources/day4.txt")
	if fileErr != nil {
		panic(fileErr)
	}
	input = strings.Split(string(inputFile), "\n")

	bingoBoard := make([][]int64, 0)
	resultBoard := make([][]bool, 0)

	for _, boardRow := range input {
		if len(boardRow) > 0 { // skip board separator
			potentialColumns := strings.Split(boardRow, " ")
			columns := make([]int64, 0)
			resultColumns := make([]bool, 0)

			for _, val := range potentialColumns {
				number, err := strconv.ParseInt(val, 10, 64)
				if err == nil {
					columns = append(columns, number)
					resultColumns = append(resultColumns, false)
				}
			}

			bingoBoard = append(bingoBoard, columns)
			resultBoard = append(resultBoard, resultColumns)
		} else {
			boards = append(boards, bingoBoard)
			results = append(results, resultBoard)

			bingoBoard = make([][]int64, 0)
			resultBoard = make([][]bool, 0)
		}
	}

	partOne()
	partTwo()
}

func partOne() {
	fmt.Println("#### Part One ####")

	for x, number := range numbers {
		for y, board := range boards {
			hasMatch, row, column := checkBoard(board, number)

			if hasMatch {
				results[y][row][column] = true
				isBingo, misses := isBingo(board, results[y])

				if isBingo {
					fmt.Printf("BINGO! Board: %d, Iteration: %d, Misses Total: %d, Product: %d\n", y+1, x, misses, number * misses)
					return
				}
			}
		}
	}
}

func isBingo(board [][]int64, boardStatus [][]bool) (bool, int64) {

	isBingo := false
	var sumOfHits int64 = 0
	var sumOfMisses int64 = 0

	for y, row := range boardStatus {
		isRowHit := true
		for x, column := range row {
			if column == false {
				isRowHit = false
				sumOfMisses = sumOfMisses + board[y][x]
			} else if isBingo == false {
				sumOfHits = sumOfHits + board[y][x]
			}
		}

		if isRowHit {
			isBingo = true
		} else {
			isRowHit = true
			sumOfHits = 0
		}
	}

	if isBingo == false {

		for x := 0; x < len(boardStatus[0]); x++ {
			isColumnHit := true

			for y := 0; y < len(boardStatus); y++ {
				if boardStatus[y][x] == false {
					isColumnHit = false
					sumOfMisses = sumOfMisses + board[y][x]
				} else if isBingo == false {
					sumOfHits = sumOfHits + board[y][x]
				}
			}

			if isColumnHit && isBingo == false {
				isBingo = true
			} else {
				isColumnHit = true
				sumOfHits = 0
			}
		}
	}

	return isBingo, sumOfMisses
}

func checkBoard(board [][]int64, number int64) (bool, int, int) {
	for y, row := range board {
		for x, column := range row {
			if column == number {
				return true, y, x
			}
		}
	}

	return false, 0, 0
}


func partTwo() {
	fmt.Println("#### Part Two ####")

	boardsThatHitBingo := make([]int, 0)

	for x, number := range numbers {
		for y, board := range boards {

			if !utils.IntArrayContains(boardsThatHitBingo, y) {
				hasMatch, row, column := checkBoard(board, number)

				if hasMatch {
					results[y][row][column] = true
					isBingo, misses := isBingo(board, results[y])

					if isBingo {
						fmt.Printf("BINGO! Board: %d, Iteration: %d, Misses Total: %d, Product: %d\n", y+1, x, misses, number*misses)

						boardsThatHitBingo = append(boardsThatHitBingo, y)
					}
				}
			}
		}
	}
}
