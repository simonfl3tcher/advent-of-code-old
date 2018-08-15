package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "1321131112"
	Part1(input)
	Part2(input)
}

func Part1(str string) {
	result := recursiveFunc(str, 1, 40)
	fmt.Println(len(result))
}

func Part2(str string) {
	result := recursiveFunc(str, 1, 50)
	fmt.Println(len(result))
}

func recursiveFunc(str string, c int, r int) string {
	if c == r+1 {
		return str
	}

	strSlice := strings.Split(str, "")
	current := strSlice[0]
	counter := 0
	var newNum []string

	for i, v := range strSlice {
		if (i+1) != len(strSlice) && current == strSlice[i+1] {
			counter++
		} else {
			counter++
			val, _ := strconv.Atoi(v)
			t := fmt.Sprintf("%d%d", counter, val)
			newNum = append(newNum, t)

			var key int
			if (i + 1) != len(strSlice) {
				key = i + 1
			} else {
				key = i
			}
			current = strSlice[key]
			counter = 0
		}
	}

	return recursiveFunc(strings.Join(newNum, ""), c+1, r)
}
