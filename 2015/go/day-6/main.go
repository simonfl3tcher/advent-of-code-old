package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	Part1(string(data))
}

// Part1 solves part 1 of day 6
func Part1(str string) {
	fmt.Println("Start right here!")
}
