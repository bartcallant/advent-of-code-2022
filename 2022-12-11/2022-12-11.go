package main

import (
	"advent-of-code-2022/utils/arrays/reduceArray"
	"advent-of-code-2022/utils/files/readFileAsLinesArray"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	id                     int
	items                  []int
	operation              string
	testValue              int
	testValueTrue          int
	testValueFalse         int
	numberOfInspectedItems int
}

func buildMonkeys(lines []string) []Monkey {
	var monkeys = []Monkey{}

	var currentMonkey = Monkey{id: -1}

	for _, line := range lines {
		if strings.HasPrefix(line, "Monkey ") {
			if currentMonkey.id != -1 {
				monkeys = append(monkeys, currentMonkey)
			}

			var monkeyIdString = strings.Fields(line)[1]
			var cleanedMonkeyIdString = strings.Replace(monkeyIdString, ":", "", 1)
			var monkeyId, _ = strconv.ParseInt(cleanedMonkeyIdString, 10, 0)

			currentMonkey = Monkey{id: int(monkeyId)}

			continue
		}

		if strings.HasPrefix(line, "  Starting items:") {
			var startingItemsString = strings.Replace(line, "  Starting items: ", "", 1)
			var startingItemsStrings = strings.Split(startingItemsString, ", ")

			var startingItems = []int{}

			for _, startingItemString := range startingItemsStrings {
				var startingItem, _ = strconv.ParseInt(startingItemString, 10, 0)

				startingItems = append(startingItems, int(startingItem))
			}

			currentMonkey.items = startingItems

			continue
		}

		if strings.HasPrefix(line, "  Operation:") {
			var operationString = strings.Replace(line, "  Operation: new = ", "", 1)

			currentMonkey.operation = operationString

			continue
		}

		if strings.HasPrefix(line, "  Test:") {
			var divisibleString = strings.Replace(line, "  Test: divisible by ", "", 1)
			var divisible, _ = strconv.ParseInt(divisibleString, 10, 0)

			currentMonkey.testValue = int(divisible)

			continue
		}

		if strings.HasPrefix(line, "    If true:") {
			var testValueTrueString = strings.Replace(line, "    If true: throw to monkey ", "", 1)
			var testValueTrue, _ = strconv.ParseInt(testValueTrueString, 10, 0)

			currentMonkey.testValueTrue = int(testValueTrue)

			continue
		}
		if strings.HasPrefix(line, "    If false:") {
			var testValueFalseString = strings.Replace(line, "    If false: throw to monkey ", "", 1)
			var testValueFalse, _ = strconv.ParseInt(testValueFalseString, 10, 0)

			currentMonkey.testValueFalse = int(testValueFalse)

			continue
		}
	}

	monkeys = append(monkeys, currentMonkey)

	return monkeys
}

func processOperation(currentValue int, operation string) int {
	var parts = strings.Fields(operation)

	var value1String = parts[0]
	var op = parts[1]
	var value2String = parts[2]

	var value1 int = 0
	var value2 int = 0

	if value1String == "old" {
		value1 = currentValue
	}

	if value2String == "old" {
		value2 = currentValue
	} else {
		var parsedValue, _ = strconv.ParseInt(value2String, 10, 0)

		value2 = int(parsedValue)
	}

	if op == "*" {
		return value1 * value2
	} else if op == "+" {
		return value1 + value2
	}

	return 0
}

func calculateMonkeyBusiness(monkeys []Monkey, numberOfRounds int, worryFactor int) int {
	var monkeyMap = make(map[int]Monkey)

	var mod = 1

	for _, monkey := range monkeys {
		monkeyMap[monkey.id] = monkey

		mod *= monkey.testValue
	}

	for round := 1; round <= numberOfRounds; round++ {
		for monkeyId := 0; monkeyId < len(monkeyMap); monkeyId++ {
			for itemId, _ := range monkeyMap[monkeyId].items {
				var item = monkeyMap[monkeyId].items[itemId]

				var worryLevel = processOperation(item, monkeyMap[monkeyId].operation)
				worryLevel = worryLevel / worryFactor % mod

				if worryLevel%monkeyMap[monkeyId].testValue == 0 {
					var monkeyToThrow, monkeyToThrowFound = monkeyMap[monkeyMap[monkeyId].testValueTrue]

					if monkeyToThrowFound {
						monkeyToThrow.items = append(monkeyMap[monkeyMap[monkeyId].testValueTrue].items, worryLevel)

						monkeyMap[monkeyToThrow.id] = monkeyToThrow
					}
				} else {
					var monkeyToThrow, monkeyToThrowFound = monkeyMap[monkeyMap[monkeyId].testValueFalse]

					if monkeyToThrowFound {
						monkeyToThrow.items = append(monkeyMap[monkeyMap[monkeyId].testValueFalse].items, worryLevel)

						monkeyMap[monkeyToThrow.id] = monkeyToThrow
					}
				}
			}

			var monkey = monkeyMap[monkeyId]
			monkey.numberOfInspectedItems += len(monkey.items)
			monkey.items = []int{}
			monkeyMap[monkeyId] = monkey
		}
	}

	var numberOfInspectedItems = []int{}

	for _, monkey := range monkeyMap {
		numberOfInspectedItems = append(numberOfInspectedItems, monkey.numberOfInspectedItems)
	}

	sort.Ints(numberOfInspectedItems)
	var twoBiggest = numberOfInspectedItems[len(numberOfInspectedItems)-2:]
	var monkeyBusiness = reduceArray.Exec(twoBiggest, func(acc int, value int) int { return acc * value }, 1)

	return monkeyBusiness
}

func main() {
	var lines = readFileAsLinesArray.Exec("input.txt")

	var monkeys = buildMonkeys(lines)

	fmt.Println("PART1:", calculateMonkeyBusiness(monkeys, 20, 3))
	fmt.Println("PART2:", calculateMonkeyBusiness(monkeys, 10000, 1))
}
