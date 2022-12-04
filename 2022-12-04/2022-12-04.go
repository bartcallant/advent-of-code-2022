package main

import (
	"advent-of-code-2022/utils/files/readFileAsLinesArray"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var lines = readFileAsLinesArray.Exec("input.txt")

	var numberOfullyOverlappingPairs = 0
	var numberOfOverlappingPairs = 0

	for _, line := range lines {
		var pairs = strings.Split(line, ",")

		var pair1Sections = strings.Split(pairs[0], "-")
		var pair2Sections = strings.Split(pairs[1], "-")

		var pair1StartSection, _ = strconv.ParseInt(pair1Sections[0], 10, 0)
		var pair1EndSection, _ = strconv.ParseInt(pair1Sections[1], 10, 0)
		var pair2StartSection, _ = strconv.ParseInt(pair2Sections[0], 10, 0)
		var pair2EndSection, _ = strconv.ParseInt(pair2Sections[1], 10, 0)

		if (pair1StartSection >= pair2StartSection && pair1EndSection <= pair2EndSection) || (pair2StartSection >= pair1StartSection && pair2EndSection <= pair1EndSection) {
			numberOfullyOverlappingPairs += 1
		}

		if pair1EndSection >= pair2StartSection && pair2EndSection >= pair1StartSection {
			numberOfOverlappingPairs += 1
		}
	}

	fmt.Println("PART1:", numberOfullyOverlappingPairs)
	fmt.Println("PART2:", numberOfOverlappingPairs)
}
