package main

import (
	"fmt"
	"sync"
)

// Part1 returns solution for part 1 of day 2
func Part1(str string) {
	var wg sync.WaitGroup
	var counter uint64

	ss := strSlices(str)
	wg.Add(len(ss))

	for _, c := range ss {
		go calculate(c, &wg, &counter)
	}

	wg.Wait()
	fmt.Println(counter)
}
