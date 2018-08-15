package main

import (
	"io/ioutil"
	"regexp"
)

var regex = regexp.MustCompile(`^(?P<Command>.*)\s(?P<Start>\d*,\d*)\sthrough\s(?P<End>\d*,\d*)$`)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	Part1(string(data), regex)
	Part2(string(data), regex)
}
