package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	Part1(string(data))
	Part2(string(data))
}

func Part1(str string) {
	strSlices := strSlices(str)

	var codeChars int
	var memoryChars int

	for _, line := range strSlices {
		codeChars += len(line)
		unescaped, _ := strconv.Unquote(line)
		memoryChars += len(unescaped)
	}

	fmt.Println(codeChars - memoryChars)
}

func Part2(str string) {
	strSlices := strSlices(str)

	var codeChars int
	var newChars int

	for _, line := range strSlices {
		codeChars += len(line)
		escaped := strconv.Quote(line)
		newChars += len(escaped)
	}

	fmt.Println(newChars - codeChars)
}

func strSlices(str string) []string {
	strSlices := strings.Split(str, "\n")
	strSlices = strSlices[:len(strSlices)-1] // Remove the last line as it is always empty
	return strSlices
}
