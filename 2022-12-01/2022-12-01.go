package main

import (
	"fmt"
	"log"
	"sort"

	"advent-of-code-2022/utils/arrays/reduceArray"
	"advent-of-code-2022/utils/arrays/reverseArray"
	"advent-of-code-2022/utils/files/readFileAsLinesArray"
)

func calculateTotals(lines []string) []int {
	result := []int{}

	total := 0

	for _, line := range lines {
		if line == "" {
			result = append(result, total)

			total = 0

			continue
		}

		intValue := 0

		_, err := fmt.Sscan(line, &intValue)

		if err != nil {
			log.Fatal(err)
		}

		total += intValue
	}

	result = append(result, total)

	return result
}

func main() {
	fileAsLineArray := readFileAsLinesArray.Exec("input.txt")
	totals := calculateTotals(fileAsLineArray)

	sort.Ints(totals)
	descSortedTotals := reverseArray.Exec(totals)

	first := descSortedTotals[0]
	fmt.Println("Highest total of calories", first)

	topThreeSum := reduceArray.Exec(descSortedTotals[:3], func(acc, current int) int {
		return acc + current
	}, 0)
	fmt.Println("Sum of te top three combined calories", topThreeSum)
}
