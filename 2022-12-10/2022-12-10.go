package main

import (
	"advent-of-code-2022/utils/arrays/doesItemExistInArray"
	"advent-of-code-2022/utils/files/readFileAsLinesArray"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var lines = readFileAsLinesArray.Exec("input.txt")

	var actions = []int{}

	for _, line := range lines {
		actions = append(actions, 0)

		if line == "noop" {
			continue
		}

		var change, _ = strconv.ParseInt(strings.Fields(line)[1], 10, 0)
		actions = append(actions, int(change))
	}

	var x = 1
	var intervals = []int{20, 60, 100, 140, 180, 220}
	var signalStrength = 0
	var displayOutput = "\n"

	for actionIndex, action := range actions {
		var actionNumber = actionIndex + 1

		if doesItemExistInArray.Exec(intervals, func(i int) bool { return i == actionNumber }) {
			signalStrength += x * actionNumber
		}

		var visiblePixels = []int{x - 1, x, x + 1}
		if doesItemExistInArray.Exec(visiblePixels, func(vp int) bool { return vp == actionIndex%40 }) {
			displayOutput += "#"
		} else {
			displayOutput += "."
		}

		if actionNumber%40 == 0 {
			displayOutput += "\n"
		}

		x += action
	}

	fmt.Println("PART1: ", signalStrength)
	fmt.Println("PART2: ", displayOutput)
}
