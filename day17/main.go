package main

import (
	"bufio"
	"fmt"
	"io"

	"github.com/strangelytyped/advent-of-code-2022/utils"
)

var shapes = [][]int{
	{
		0b0011110,
	},
	{
		0b0001000,
		0b0011100,
		0b0001000,
	},
	{
		0b0011100,
		0b0000100,
		0b0000100,
	},
	{
		0b0010000,
		0b0010000,
		0b0010000,
		0b0010000,
	},
	{
		0b0011000,
		0b0011000,
	},
}

func canMoveLeft(shape []int, tower []int, y int) bool {
	for idx, row := range shape {
		if row&0b1000000 != 0 {
			return false
		}
		if len(tower) > (y+idx) && tower[y+idx]&(row<<1) != 0 {
			return false
		}
	}
	return true
}

func canMoveRight(shape []int, tower []int, y int) bool {
	for idx, row := range shape {
		if row&0b0000001 != 0 {
			return false
		}
		if len(tower) > (y+idx) && tower[y+idx]&(row>>1) != 0 {
			return false
		}
	}
	return true
}

func drop(tower []int, moves string, settledRocks, inputPosition int) ([]int, int) {
	shape := append([]int{}, shapes[settledRocks%len(shapes)]...)

	y := len(tower) + 3
	for y >= 0 {
		shift := moves[inputPosition]
		inputPosition = (inputPosition + 1) % len(moves)
		if shift == '<' && canMoveLeft(shape, tower, y) {
			for idx := range shape {
				shape[idx] <<= 1
			}
		} else if shift == '>' && canMoveRight(shape, tower, y) {
			for idx := range shape {
				shape[idx] >>= 1
			}
		}

		if y == 0 {
			break
		}

		clear := true
		for rowIdx, row := range shape {
			if rowIdx+(y-1) >= len(tower) {
				continue
			}
			if row&tower[rowIdx+(y-1)] != 0 {
				clear = false
				break
			}
		}
		if clear {
			y--
		} else {
			break
		}
	}

	for len(tower) < len(shape)+y {
		tower = append(tower, 0)
	}
	for idx := range shape {
		tower[y+idx] |= shape[idx]
	}
	return tower, inputPosition
}

func Part1(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Scan()
	moves := scanner.Text()

	settledRocks := 0
	inputPosition := 0
	tower := []int{}

	for ; settledRocks < 2022; settledRocks++ {
		tower, inputPosition = drop(tower, moves, settledRocks, inputPosition)
	}

	return len(tower)
}

type TowerState struct {
	topRows     uint64
	rockOffset  int
	inputOffset int
}

func Part2(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Scan()
	moves := scanner.Text()

	settledRocks := 0
	inputPosition := 0
	tower := []int{}

	history := map[TowerState][2]int{}
	targetCycles := 1000000000000
	foundLoop := false
	cyclesToTarget := 0
	cycleTowerHeightIncrease := 0

	for settledRocks < targetCycles {
		tower, inputPosition = drop(tower, moves, settledRocks, inputPosition)
		settledRocks++

		if !foundLoop {
			topRows := uint64(0)
			for i := 0; i < 9 && i < len(tower); i++ {
				topRows |= uint64(tower[len(tower)-(i+1)]) << (7 * i)
			}

			historyEntry := TowerState{
				topRows,
				settledRocks % len(shapes),
				inputPosition,
			}
			if oldEntry, has := history[historyEntry]; has {
				foundLoop = true

				towerHeightAtStartOfCycle := oldEntry[0]
				settledRocksAtStartOfCycle := oldEntry[1]

				cycleLength := settledRocks - settledRocksAtStartOfCycle
				cycleTowerHeightIncrease = len(tower) - towerHeightAtStartOfCycle
				fmt.Printf("Found cycle starting at rocks=%d / tower=%d, assuming tower will repeat every %d rocks / %d height\n", settledRocksAtStartOfCycle, towerHeightAtStartOfCycle, cycleLength, cycleTowerHeightIncrease)
				cyclesToTarget = ((targetCycles - settledRocksAtStartOfCycle) / cycleLength)
				settledRocks = (cyclesToTarget * cycleLength) + settledRocksAtStartOfCycle
				fmt.Printf("Skipping to rocks=%d\n", settledRocks)

			} else {
				history[historyEntry] = [2]int{len(tower), settledRocks}
			}
		}
	}

	return ((cyclesToTarget - 1) * cycleTowerHeightIncrease) + len(tower)
}

func main() {
	utils.Run("day17.txt", Part1, Part2)
}
