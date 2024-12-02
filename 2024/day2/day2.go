package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func difference(x *int, y *int) *int {
	diff := *x - *y
	return &diff
}

func abs(x *int) int {
	if *x < 0 {
		return -*x
	}
	return *x
}

func parse(filename string) *[][]int {
	var levels [][]int
	dat, err := os.ReadFile(filename)
	check(err)

	lines := strings.Split(string(dat), "\n")
	//fmt.Print("lines: \n")
	//fmt.Println(lines)

	for _, line := range lines {
		if line == "" {
			continue
		}

		var currentLine []int
		stringNumbers := strings.Split(line, " ")

		for _, numStr := range stringNumbers {
			num, err := strconv.Atoi(numStr)
			check(err)
			currentLine = append(currentLine, num)
		}
		levels = append(levels, currentLine)
	}
	return &levels
}

func checkLevels(levels *[][]int) *int {
	var safeReports int
	var unsafe bool
	var direction int //1=down, 2=up ,0=not set

	for i := 0; i < len(*levels); i++ {
		//fmt.Printf("####### safeamount: %d\n", safeReports)
		//fmt.Println((*levels)[i])
		unsafe = false
		direction = 0
		for j := 0; unsafe == false && j < len((*levels)[i]); j++ {
			if j > 0 {
				diff := difference(&(*levels)[i][j], &(*levels)[i][j-1])
				//fmt.Println(*diff)
				if abs(diff) > 3 || abs(diff) < 1 {
					unsafe = true
					continue

				}
				localDirection := 0
				switch {
				case *diff < 0:
					localDirection = 1
				case *diff > 0:
					localDirection = 2
				}
				//fmt.Printf("localDirection: %d\n", localDirection)
				//fmt.Printf("direction: %d\n\n", direction)
				if direction == 0 {
					direction = localDirection
				}
				if direction != localDirection {
					unsafe = true
				} else if j == 4 {
					safeReports++
				}
			}
		}
	}

	return &safeReports
}

func main() {
	levels := parse("test.txt")
	//fmt.Println(*levels)
	safeReports := checkLevels(levels)
	fmt.Printf("task 1: %d", *safeReports)
}
