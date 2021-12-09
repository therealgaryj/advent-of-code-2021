package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var displayRules = [][]int{{0,1,2,4,5,6},{2,5},{0,2,3,4,6},{0,2,3,5,6},{1,2,3,5},{0,1,3,5,6},{0,1,3,4,5,6},{0,2,5},{0,1,2,3,4,5,6},{0,1,2,3,5,6}}
var notes []map[string][][]string

func main() {

	fmt.Println("#### Day One ####")

	inputFile, fileErr := ioutil.ReadFile("./resources/day8.txt")
	if fileErr != nil {
		panic(fileErr)
	}
	input := strings.Split(string(inputFile), "\n")

	notes = make([]map[string][][]string, 0)
	for _, noteLine := range input {
		noteParts := strings.Split(noteLine, " | ")
		inputNote := strings.Split(noteParts[0], " ")
		outputNote := strings.Split(noteParts[1], " ")

		sortedInput := make([][]string, 8)
		for _, word := range inputNote {
			sortedInput[len(word)] = append(sortedInput[len(word)], word)
		}

		sortedOutput := make([][]string, 8)
		for _, word := range outputNote {
			sortedOutput[len(word)] = append(sortedOutput[len(word)], word)
		}
		notes = append(notes, map[string][][]string{"input":sortedInput,"output":sortedOutput})
	}

	partOne()
	partTwo()
}

func partOne() {
	fmt.Println("#### Part One ####")

	numberOfUniqueOutputs := 0
	for _, inputLine := range notes {
		for _, outputWords := range inputLine["output"] {
			if len(outputWords) == 2  || len(outputWords) == 3  || len(outputWords) == 4  || len(outputWords) == 7 {
				numberOfUniqueOutputs = numberOfUniqueOutputs + 1
			}
		}
	}

	fmt.Printf("1, 4, 7, 8 count is: %d\n", numberOfUniqueOutputs)

}

func partTwo() {
	fmt.Println("#### Part Two ####")

	// total output
	for _, line := range notes {

		results := make(map[string]int)

		// Compare len(2) and len(3)

		lettersAtPos2And5 := make([]string, 0)

		for _, letterFrom7 := range line["input"][3][0] {
			foundLetterFrom7 := false
			for _, letterFrom1 := range line["input"][2][0] {
				if letterFrom7 == letterFrom1 {
					foundLetterFrom7 = true
					break
				}
			}
			if foundLetterFrom7 == false {
				results[string(letterFrom7)] = 0
			} else {
				lettersAtPos2And5 = append(lettersAtPos2And5, string(letterFrom7))
			}
		}

		for _, lengthOf6 := range line["input"][6] {
			for _, lettersInLength6 := range lengthOf6 {
				if !strings.Contains(string(lettersInLength6), lettersAtPos2And5[0]) && strings.Contains(string(lettersInLength6), lettersAtPos2And5[1]) {
					results[lettersAtPos2And5[0]] = 5
					results[lettersAtPos2And5[1]] = 3
				} else if strings.Contains(string(lettersInLength6), lettersAtPos2And5[0]) && !strings.Contains(string(lettersInLength6), lettersAtPos2And5[1]) {
					results[lettersAtPos2And5[0]] = 3
					results[lettersAtPos2And5[1]] = 5
				}
			}
		}

	}
}