package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

var regex = regexp.MustCompile(`^(\w+)\sto\s(\w+)\s=\s(\d+)$`)

func strSlices(str string) []string {
	strSlices := strings.Split(str, "\n")
	strSlices = strSlices[:len(strSlices)-1] // Remove the last line as it is always empty
	return strSlices
}

func example() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	graph := &Graph{
		structure: make(map[string]nodeEdge),
		nodes:     []string{},
	}

	for _, line := range strSlices(string(data)) {
		fmt.Println(line)
		// match := regex.FindStringSubmatch(strings.TrimSuffix(line, "\n"))
		// start := match[1]
		// end := match[2]
		// length, _ := strconv.Atoi(match[3])
	}

	fmt.Println(graph.structure)

	var count float64
	for _, m := range graph.structure {
		lowest := 0.00
		for _, v := range m {
			if v > lowest {
				lowest = v
			}
		}
		count += lowest
	}

	fmt.Println(count)
}

func heapPermutations(a []int, size int, n int) {
	if size == 1 {
		for i := 0; i < n; i++ {
			fmt.Println(a[i])
		}
		fmt.Println("--")
		return
	}

	for i := 0; i < n; i++ {
		heapPermutations(a, size-1, n)

		if size%2 == 1 {
			temp := a[0]
			a[0] = a[size-1]
			a[size-1] = temp
		} else {
			temp := a[i]
			a[i] = a[size-1]
			a[size-1] = temp
		}
	}
}

func main() {
	a := []int{1, 2, 3}
	heapPermutations(a, len(a), len(a))
}

func lowestDistance(distances []float64) float64 {
	var n float64
	n = distances[0]
	for _, v := range distances {
		if v < n {
			n = v
		}
	}
	return n
}
