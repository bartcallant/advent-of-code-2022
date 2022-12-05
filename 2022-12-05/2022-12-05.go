package main

import (
	"advent-of-code-2022/utils/arrays/chunkArrayByChunkSize"
	"advent-of-code-2022/utils/arrays/reverseArray"
	"advent-of-code-2022/utils/files/readFileAsLinesArray"
	"fmt"
	"strings"
)

func buildCrateStacks(lines []string) (map[string][]string, []string) {
	var stackDefs = []string{}
	var actions = []string{}

	for _, line := range lines {
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "move") {
			actions = append(actions, line)

			continue
		}

		stackDefs = append(stackDefs, line)
	}

	var reversedStackDefs = reverseArray.Exec(stackDefs)

	var buildStacks = map[string][]string{}

	for _, stackDef := range reversedStackDefs {
		if !strings.Contains(stackDef, "[") {
			// Process header row aka number of stacks
			var stackIds = strings.Fields(stackDef)

			for _, id := range stackIds {
				buildStacks[id] = []string{}
			}

			continue
		}

		var stackCrateDefs = chunkArrayByChunkSize.Exec([]rune(stackDef), 4)

		for index, crate := range stackCrateDefs {
			for _, c := range crate {
				if c == 32 || c == 91 || c == 93 {
					continue
				}

				buildStacks[fmt.Sprintf("%d", index+1)] = append(buildStacks[fmt.Sprintf("%d", index+1)], fmt.Sprintf("%c", c))
			}
		}
	}

	return buildStacks, actions
}

func processPart1(stacks map[string][]string, actions []string) {
	for _, action := range actions {
		var parsedAction = strings.Fields(action)

		var count = 0
		fmt.Sscan(parsedAction[1], &count)

		var from = parsedAction[3]
		var to = parsedAction[5]

		for i := 1; i <= count; i++ {
			var sourceList = stacks[from]
			var destinationList = stacks[to]

			var last = sourceList[len(sourceList)-1]

			if len(sourceList)-1 < 0 {
				continue
			}

			stacks[to] = append(destinationList, last)

			if len(sourceList)-1 >= 0 {
				stacks[from] = sourceList[0 : len(sourceList)-1]
			}
		}
	}
}

func processPart2(stacks map[string][]string, actions []string) {
	for _, action := range actions {
		var parsedAction = strings.Fields(action)

		var count = 0
		fmt.Sscan(parsedAction[1], &count)

		var from = parsedAction[3]
		var to = parsedAction[5]

		var sourceList = stacks[from]
		var destinationList = stacks[to]

		var last = sourceList[len(sourceList)-count:]

		stacks[to] = append(destinationList, last...)

		if len(sourceList)-count >= 0 {
			stacks[from] = sourceList[0 : len(sourceList)-count]
		}
	}
}

func buildResult(stacks map[string][]string) string {
	var resultList = []string{}

	for i := 1; i <= len(stacks); i++ {
		var list = stacks[fmt.Sprintf("%d", i)]
		var last = list[len(list)-1]

		resultList = append(resultList, last)
	}

	return strings.Join(resultList, "")
}

func main() {
	var lines = readFileAsLinesArray.Exec("input.txt")

	var part1Stacks, part1Actions = buildCrateStacks(lines)
	processPart1(part1Stacks, part1Actions)
	var part2Stacks, part2Actions = buildCrateStacks(lines)
	processPart2(part2Stacks, part2Actions)

	fmt.Println("PART1: ", buildResult(part1Stacks))
	fmt.Println("PART2: ", buildResult(part2Stacks))
}
