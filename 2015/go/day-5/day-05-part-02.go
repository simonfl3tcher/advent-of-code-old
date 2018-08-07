package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func Part2(str string) {
	var wg sync.WaitGroup
	var counter uint64

	s := strSlices(str)
	wg.Add(len(s))

	for _, v := range s {
		go func(str string) {
			defer wg.Done()

			if containsPairOfLetters(str) && containsTwoLettersWithOneLetterInbetween(str) {
				atomic.AddUint64(&counter, 1) // Add one if it is valid
			}
		}(v)
	}

	wg.Wait()
	fmt.Println(counter)
}
