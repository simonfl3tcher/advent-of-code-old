package main

import "fmt"

// Part2 returns the answer to day 1 part 2 challenge
func Part2(data []byte) {
	var floor int
	for index, c := range data {
		if c == 40 {
			floor += 1
		} else if c == 41 {
			floor -= 1
		}

		if floor < 0 {
			fmt.Println(index + 1)
			break
		}
	}
}
