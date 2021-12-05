package main

import (
	"fmt"
	"github.com/therealgaryj/advent-of-code-2021/internal/utils"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var input []string

type coord struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

var coords = make([]coord, 0)
var vents [][]int

func main() {

	fmt.Println("#### Day One ####")

	inputFile, fileErr := ioutil.ReadFile("./resources/day5.txt")
	if fileErr != nil {
		panic(fileErr)
	}
	input = strings.Split(string(inputFile), "\n")

	r := regexp.MustCompile(`(?P<X1>\d+),(?P<Y1>\d+) -> (?P<X2>\d+),(?P<Y2>\d+)`)

	maxX := 0
	maxY := 0

	for _, instruction := range input {
		res := r.FindStringSubmatch(instruction)
		x1, _ := strconv.Atoi(res[r.SubexpIndex("X1")])
		y1, _ := strconv.Atoi(res[r.SubexpIndex("Y1")])
		x2, _ := strconv.Atoi(res[r.SubexpIndex("X2")])
		y2, _ := strconv.Atoi(res[r.SubexpIndex("Y2")])

		if x1 > maxX {
			maxX = x1
		}
		if x2 > maxX {
			maxX = x2
		}
		if y1 > maxY {
			maxY = y1
		}
		if y2 > maxY {
			maxY = y2
		}

		coords = append(coords, coord{
			x1: x1,
			y1: y1,
			x2: x2,
			y2: y2,
		})
	}

	vents = make([][]int, maxY + 1)
	for y := 0; y < len(vents); y++ {
		vents[y] = make([]int, maxX + 1)
	}

	partOne()
	partTwo()
}

func partOne() {
	fmt.Println("#### Part One ####")

	for _, coord := range coords {
		if coord.x1 == coord.x2 {
			if coord.y1 > coord.y2 {
				for y := coord.y2; y <= coord.y1; y++ {
					incrementPoint(coord.x1, y)
				}
			} else {
				for y := coord.y1; y <= coord.y2; y++ {
					incrementPoint(coord.x1, y)
				}
			}
		} else if coord.y1 == coord.y2 {
			if coord.x1 > coord.x2 {
				for x := coord.x2; x <= coord.x1; x++ {
					incrementPoint(x, coord.y1)
				}
			} else {
				for x := coord.x1; x <= coord.x2; x++ {
					incrementPoint(x, coord.y1)
				}
			}
		}
	}

	total := countMoreThanOneVent()

	fmt.Printf("Total: %d\n", total)
}

func countMoreThanOneVent() int {
	total := 0
	for _, rows := range vents {
		for _, column := range rows {
			if column > 1 {
				total = total + 1
			}
		}
	}

	return total
}

func partTwo() {
	fmt.Println("#### Part Two ####")

	resetVents()

	for _, coord := range coords {

		if coord.x1 > coord.x2 {
			for j := 0; j <= utils.Diff(coord.x1, coord.x2); j++ {
				if coord.y1 == coord.y2 {
					incrementPoint(coord.x1 - j, coord.y1)
				} else if coord.y1 > coord.y2 {
					incrementPoint(coord.x1 - j, coord.y1 - j)
				} else {
					incrementPoint(coord.x1 - j, coord.y1 + j)
				}
			}
		} else if coord.x2 > coord.x1 {
			for j := 0; j <= utils.Diff(coord.x1, coord.x2); j++ {
				if coord.y1 == coord.y2 {
					incrementPoint(coord.x1 + j, coord.y1)
				} else  if coord.y1 > coord.y2 {
					incrementPoint(coord.x1 + j, coord.y1 - j)
				} else {
					incrementPoint(coord.x1 + j, coord.y1 + j)
				}
			}
		} else { //x is the same, do y
			if coord.y1 > coord.y2 {
				for j := 0; j <= utils.Diff(coord.y1, coord.y2); j++ {
					incrementPoint(coord.x1, coord.y1 - j)
				}
			} else if coord.y2 > coord.y1 {
				for j := 0; j <= utils.Diff(coord.y1, coord.y2); j++ {
					incrementPoint(coord.x1, coord.y1 + j)
				}
			}
		}
	}

	total := countMoreThanOneVent()

	fmt.Printf("Total: %d\n", total)
}

func incrementPoint(x int, y int) {
	vents[y][x] = vents[y][x] + 1
}

func resetVents() {
	for x := 0; x < len(vents); x++ {
		for y := 0; y < len(vents[x]); y++ {
			vents[x][y] = 0
		}
	}
}