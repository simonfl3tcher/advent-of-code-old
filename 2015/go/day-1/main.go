package main

import "io/ioutil"

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	Part1(data)
	Part2(data)
}
