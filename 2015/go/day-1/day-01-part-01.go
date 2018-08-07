package main

import "fmt"

// Part1 returns the answer to day 1 part 1 challenge
func Part1(data []byte) {
	var floor int
	for _, c := range data {
		if c == 40 {
			floor++
		} else if c == 41 {
			floor--
		}
	}

	fmt.Println(floor)
}
