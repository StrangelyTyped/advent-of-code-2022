package main

import (
	"bufio"
	"io"

	"github.com/strangelytyped/advent-of-code-2022/utils"
)

func readInput(input io.Reader) (trees [][]int) {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		treeRow := []int{}

		for _, chr := range line {
			treeRow = append(treeRow, int(chr - '0'))
		}

		trees = append(trees, treeRow)
	}
	return trees
}

func checkVisibility(trees [][]int, startX, startY, deltaX, deltaY int) bool {
	height := trees[startY][startX]
	startX += deltaX
	startY += deltaY
	for startX >= 0 && startY >= 0 && startY < len(trees) && startX < len(trees[startY]) {
		if trees[startY][startX] >= height {
			return false
		}
		
		startX += deltaX
		startY += deltaY
	}
	return true
}

func checkScore(trees [][]int, startX, startY, deltaX, deltaY int) int {
	height := trees[startY][startX]
	startX += deltaX
	startY += deltaY
	treeCount := 0
	for startX >= 0 && startY >= 0 && startY < len(trees) && startX < len(trees[startY]) {
		treeCount++
		if trees[startY][startX] >= height {
			break
		}
		
		startX += deltaX
		startY += deltaY
	}
	return treeCount
}

func Part1(input io.Reader) int {
	trees := readInput(input)
	visibleCount := 0

	for y := 0; y < len(trees); y++ {
		treeRow := trees[y]
		for x := 0; x < len(treeRow); x++ {
			// Short circuit the perimiter
			if x == 0 || y == 0 || x == (len(treeRow) - 1) || y == (len(trees) - 1) {
				visibleCount++
				continue
			}

			if checkVisibility(trees, x, y, 1, 0) ||
				checkVisibility(trees, x, y, -1, 0) || 
				checkVisibility(trees, x, y, 0, 1) ||
				checkVisibility(trees, x, y, 0, -1) {
				visibleCount++
			}
		}
	}
	return visibleCount
}

func Part2(input io.Reader) int {
	trees := readInput(input)
	maxScore := 0

	for y := 0; y < len(trees); y++ {
		treeRow := trees[y]
		for x := 0; x < len(treeRow); x++ {
			score := checkScore(trees, x, y, 1, 0) *
				checkScore(trees, x, y, -1, 0) *
				checkScore(trees, x, y, 0, 1) *
				checkScore(trees, x, y, 0, -1) 
				
			if score > maxScore {
				maxScore = score
			}
		}
	}
	return maxScore
}

func main() {
	utils.Run("day08.txt", Part1, Part2)
}
