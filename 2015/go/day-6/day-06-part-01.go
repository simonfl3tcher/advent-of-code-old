package main

import (
	"fmt"
	"regexp"
)

// Part1 solves part 2 of day 6
func Part1(str string, regex *regexp.Regexp) {
	lights := make([][1000]int, 1000)
	runInstructions(str, &lights, turnOnP1, turnOffP1, toggleP1)
	var counter int
	for _, row := range lights {
		for _, light := range row {
			counter += light
		}
	}
	fmt.Println(counter)
}

func turnOnP1(lights *[][1000]int, i, j int) {
	(*lights)[i][j] = 1
}

func turnOffP1(lights *[][1000]int, i, j int) {
	(*lights)[i][j] = 0
}

func toggleP1(lights *[][1000]int, i, j int) {
	if (*lights)[i][j] == 0 {
		(*lights)[i][j] = 1
	} else {
		(*lights)[i][j] = 0
	}
}
