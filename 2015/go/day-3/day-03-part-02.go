package main

import (
	"fmt"
	"strings"
)

// Part2 solves part two of day 3
func Part2(str string) {
	var santa counter
	var robosanta counter

	var store []counter
	store = append(store, counter{})

	for index, a := range strings.Split(str, "") {
		if index%2 == 0 {
			switcher(a, &santa, &store)
		} else {
			switcher(a, &robosanta, &store)
		}
	}
	fmt.Println(len(store))
}
