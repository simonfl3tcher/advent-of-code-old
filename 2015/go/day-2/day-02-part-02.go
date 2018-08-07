package main

import (
	"fmt"
	"sync"
)

// Part2 returns solution for part 1 of day 2
func Part2(str string) {
	var wg sync.WaitGroup
	var counter uint64

	ss := strSlices(str)
	wg.Add(len(ss))

	for _, c := range ss {
		go calculateRibon(c, &wg, &counter)
	}

	wg.Wait()
	fmt.Println(counter)
}
