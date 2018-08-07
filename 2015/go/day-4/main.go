package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	Part1(string(data))
	Part2(string(data))
}

// Part1 solves part 1 of day 4
func Part1(str string) {
	response := runChecker(str, "\\A00000")
	fmt.Println(response)
}

// Part2 solves part 1 of day 4
func Part2(str string) {
	response := runChecker(str, "\\A000000")
	fmt.Println(response)
}

func runChecker(str string, reg string) int {
	num := 0
	regex := regexp.MustCompile(reg)
	str = strings.TrimSuffix(str, "\n")

	for {
		num++

		hash := md5.New()
		io.WriteString(hash, fmt.Sprintf("%s%d", string(str), num))

		digest := fmt.Sprintf("%x", hash.Sum(nil))

		if regex.MatchString(digest) {
			break
		}
	}

	return num
}
