package main

import (
	"advent-of-code-2022/utils/files/readFileAsLinesArray"
	"fmt"
	"strconv"
)

type Tree struct {
	row    int
	column int
	height int
}

func buildPosition(row int, column int) string {
	return fmt.Sprintf("%d-%d", row, column)
}

func buildTreeMap(lines []string) map[string]Tree {
	var treeMap = make(map[string]Tree)

	for rowIndex, line := range lines {
		for columnIndex, heightAsRune := range line {
			var heightAsString = fmt.Sprintf("%c", heightAsRune)
			var heightAsInt, _ = strconv.ParseInt(heightAsString, 10, 0)

			var position = buildPosition(rowIndex, columnIndex)

			var tree = Tree{row: rowIndex, column: columnIndex, height: int(heightAsInt)}

			treeMap[position] = tree
		}
	}

	return treeMap
}

func isVisible(tree Tree, treeMap map[string]Tree, row int, column int) bool {
	var positionToCheck = buildPosition(row, column)

	var positionTree, positionTreeExists = treeMap[positionToCheck]

	if !positionTreeExists {
		return true
	}

	if positionTree.height >= tree.height {
		return false
	}

	return true
}

func calculateTreesVisibleFromOutsideGrid(treeMap map[string]Tree, rows int, columns int) int {
	var treesVisibleFromOUtsideGridCount = 0

	for _, tree := range treeMap {
		var visibleFromLeft = true
		for left := 0; left < tree.column; left++ {
			visibleFromLeft = isVisible(tree, treeMap, tree.row, left)

			if !visibleFromLeft {
				break
			}
		}

		if visibleFromLeft {
			treesVisibleFromOUtsideGridCount += 1
			continue
		}

		var visibleFromRight = true
		for right := columns - 1; right > tree.column; right-- {
			visibleFromRight = isVisible(tree, treeMap, tree.row, right)

			if !visibleFromRight {
				break
			}
		}

		if visibleFromRight {
			treesVisibleFromOUtsideGridCount += 1
			continue
		}

		var visibleFromTop = true
		for top := 0; top < tree.row; top++ {
			visibleFromTop = isVisible(tree, treeMap, top, tree.column)

			if !visibleFromTop {
				break
			}
		}

		if visibleFromTop {
			treesVisibleFromOUtsideGridCount += 1
			continue
		}

		var visibleFromBottom = true
		for bottom := rows - 1; bottom > tree.row; bottom-- {
			visibleFromBottom = isVisible(tree, treeMap, bottom, tree.column)

			if !visibleFromBottom {
				break
			}
		}

		if visibleFromBottom {
			treesVisibleFromOUtsideGridCount += 1
			continue
		}
	}

	return treesVisibleFromOUtsideGridCount
}

func calculateHighestScenicScore(treeMap map[string]Tree, rows int, columns int) int {
	var highestScenicScore = 0

	for _, tree := range treeMap {
		var leftVisibility = 0
		for left := tree.column - 1; left >= 0; left-- {
			var positionTree = treeMap[buildPosition(tree.row, left)]

			leftVisibility += 1

			if positionTree.height < tree.height {
				continue
			}

			break
		}

		var rightVisibility = 0
		for right := tree.column + 1; right < columns; right++ {
			var positionTree = treeMap[buildPosition(tree.row, right)]

			rightVisibility += 1

			if positionTree.height < tree.height {
				continue
			}

			break
		}

		var topVisibility = 0
		for top := tree.row - 1; top >= 0; top-- {
			var positionTree = treeMap[buildPosition(top, tree.column)]

			topVisibility += 1

			if positionTree.height < tree.height {
				continue
			}

			break
		}

		var bottomVisibility = 0
		for bottom := tree.row + 1; bottom < rows; bottom++ {
			var positionTree = treeMap[buildPosition(bottom, tree.column)]

			bottomVisibility += 1

			if positionTree.height < tree.height {
				continue
			}

			break
		}

		var treeScenicScore = leftVisibility * rightVisibility * topVisibility * bottomVisibility
		if treeScenicScore > highestScenicScore {
			highestScenicScore = treeScenicScore
		}
	}

	return highestScenicScore
}

func main() {
	var lines = readFileAsLinesArray.Exec("input.txt")

	var rows = len(lines)
	var columns = len(lines[0])

	var treeMap = buildTreeMap(lines)

	fmt.Println("PART1: ", calculateTreesVisibleFromOutsideGrid(treeMap, rows, columns))
	fmt.Println("PART2: ", calculateHighestScenicScore(treeMap, rows, columns))
}
