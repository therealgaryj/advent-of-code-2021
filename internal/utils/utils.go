package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
)

func GetCharacterAtPosition(input string, index int) (s string) {
	for i, rune := range input {
		if i == index {
			return string(rune)
		}
	}

	return ""
}

func LoadInput(filename string, target *interface{}) {
	// Open the file
	rawInputFile, err := os.Open("./resources/" + filename)
	if err != nil {
		panic(err)
	}
	defer rawInputFile.Close()

	// Parse the file
	rawJson, readErr := ioutil.ReadAll(rawInputFile)
	if readErr != nil {
		panic(readErr)
	}

	jsonErr := json.Unmarshal(rawJson, target)
	if jsonErr != nil {
		panic(jsonErr)
	}
}

func ConvertInputToInts(input []string) []int {
	newInput := make([]int, len(input))
	for x := range input {
		number, err := strconv.Atoi(input[x])
		if err != nil {
			panic(err)
		}
		newInput[x] = number
	}

	return newInput
}

func IntArrayContains(haystack []int, needle int) bool {
	for _, candidate := range haystack {
		if candidate == needle {
			return true
		}
	}

	return false
}

func Diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func Push(stack *[]rune, element rune) {
	*stack = append(*stack, element)
}

func Pop(stack *[]rune) (rune, bool) {

	if len(*stack) == 0 {
		return 0, false
	}

	lastIndex := len(*stack) - 1
	popped := (*stack)[lastIndex]
	*stack = (*stack)[:lastIndex]

	return popped, true
}