package main

import (
	"fmt"
	"regexp"
)

// Part2 solves part 2 of day 6
func Part2(str string, regex *regexp.Regexp) {
	lights := make([][1000]int, 1000)
	runInstructions(str, &lights, turnOnP2, turnOffP2, toggleP2)
	var counter int
	for _, row := range lights {
		for _, light := range row {
			counter += light
		}
	}
	fmt.Println(counter)
}

func turnOnP2(lights *[][1000]int, i, j int) {
	(*lights)[i][j]++
}

func turnOffP2(lights *[][1000]int, i, j int) {
	(*lights)[i][j]--
	if (*lights)[i][j] < 0 {
		(*lights)[i][j] = 0
	}
}

func toggleP2(lights *[][1000]int, i, j int) {
	(*lights)[i][j] += 2
}
