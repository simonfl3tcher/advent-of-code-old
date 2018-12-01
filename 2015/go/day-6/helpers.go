package main

import (
	"strconv"
	"strings"
)

type fn = func(lights *[][1000]int, i, j int)

func strSlices(str string) []string {
	strSlices := strings.Split(str, "\n")
	strSlices = strSlices[:len(strSlices)-1] // Remove the last line as it is always empty
	return strSlices
}

func stringToMap(str string) map[string]string {
	match := regex.FindStringSubmatch(str)
	result := make(map[string]string)
	for i, name := range regex.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}
	return result
}

func runInstructions(str string, lights *[][1000]int, turnOn fn, turnOff fn, toggle fn) {
	for _, instruction := range strSlices(str) {
		instruction := stringToMap(instruction)
		start := strings.Split(instruction["Start"], ",")
		end := strings.Split(instruction["End"], ",")
		runInstructionLine(lights, instruction["Command"], start, end, turnOn, turnOff, toggle)
	}
}

func runInstructionLine(lights *[][1000]int, command string, start, end []string, turnOn fn, turnOff fn, toggle fn) {
	rowStart, _ := strconv.Atoi(start[1])
	rowEnd, _ := strconv.Atoi(end[1])
	for i := rowStart; i <= rowEnd; i++ {
		columnStart, _ := strconv.Atoi(start[0])
		columnEnd, _ := strconv.Atoi(end[0])
		for j := columnStart; j <= columnEnd; j++ {
			switch command {
			case "turn on":
				turnOn(lights, i, j)
			case "toggle":
				toggle(lights, i, j)
			case "turn off":
				turnOff(lights, i, j)
			}
		}
	}
}
