package main

import (
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

func strSlices(str string) []string {
	strSlices := strings.Split(str, "\n")
	strSlices = strSlices[:len(strSlices)-1] // Remove the last line as it is always empty
	return strSlices
}

func calculateRibon(str string, wg *sync.WaitGroup, counter *uint64) {
	defer wg.Done()
	s := splitAndConv(str)

	// Sort the slice
	sort.Ints(s)

	// Get the two lowest values
	lowestVals := s[:2]

	// Calculate Ribon size
	total := lowestVals[0] + lowestVals[0] + lowestVals[1] + lowestVals[1]
	total += s[0] * s[1] * s[2]

	// Add value
	atomic.AddUint64(counter, uint64(total))
}

func calculate(str string, wg *sync.WaitGroup, counter *uint64) {
	defer wg.Done()
	s := splitAndConv(str)
	a := runCalculations(s)

	var sum int
	var smallest = a[0]
	for _, f := range a {
		sum += 2 * f
		if f < smallest {
			smallest = f
		}
	}

	atomic.AddUint64(counter, uint64(sum+smallest))
}

func splitAndConv(str string) []int {
	f := strings.Split(str, "x")
	var intSlice []int
	for _, c := range f {
		val, _ := strconv.Atoi(c)
		intSlice = append(intSlice, val)
	}
	return intSlice
}

func runCalculations(sliceOfInts []int) []int {
	var newSlice []int
	newSlice = append(newSlice, doSum(sliceOfInts[0], sliceOfInts[1]))
	newSlice = append(newSlice, doSum(sliceOfInts[1], sliceOfInts[2]))
	newSlice = append(newSlice, doSum(sliceOfInts[2], sliceOfInts[0]))
	return newSlice
}

func doSum(a, b int) int {
	return a * b
}
