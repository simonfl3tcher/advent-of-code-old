package main

import (
	"fmt"
	"strings"
)

type counter struct {
	x int
	y int
}

// Part1 returns part 1 answer for day 3
func Part1(str string) {
	var c counter

	var store []counter
	store = append(store, counter{})

	for _, a := range strings.Split(str, "") {
		switcher(a, &c, &store)
	}
	fmt.Println(len(store))
}
