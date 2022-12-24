package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/strangelytyped/advent-of-code-2022/utils"
)

type Position [2]int

type Blizzard struct {
	pos Position
	dir int
}

func readInput(input io.Reader) ([]Blizzard, Position, Position, int, int) {
	scanner := bufio.NewScanner(input)

	blizzMap := []Blizzard{}
	start := Position{}
	end := Position{}
	y := 0
	width := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Index(line, "##") != -1 {
			if y == 0 {
				start[0] = strings.Index(line, ".")
				width = len(line)
			} else {
				end[0] = strings.Index(line, ".")
				end[1] = y
			}
		}
		for x := 1; x < width; x++ {
			if line[x] == '^' {
				blizzMap = append(blizzMap, Blizzard{Position{x, y}, 0})
			} else if line[x] == 'v' {
				blizzMap = append(blizzMap, Blizzard{Position{x, y}, 1})
			} else if line[x] == '<' {
				blizzMap = append(blizzMap, Blizzard{Position{x, y}, 2})
			} else if line[x] == '>' {
				blizzMap = append(blizzMap, Blizzard{Position{x, y}, 3})
			}
		}
		y++
	}
	return blizzMap, start, end, width, y
}

var moveDirections = []Position{
	{0, -1},
	{0, 1},
	{-1, 0},
	{1, 0},
}

func move(pos Position, dir int, width, height int) Position {
	newPos := Position{pos[0] + moveDirections[dir][0], pos[1] + moveDirections[dir][1]}
	if newPos[0] < 1 {
		newPos[0] = width - 2
	}
	if newPos[0] >= (width-1) {
		newPos[0] = 1
	}
	if newPos[1] < 1 {
		newPos[1] = height - 2
	}
	if newPos[1] >= (height-1) {
		newPos[1] = 1
	}
	return newPos
}

func printMap(blizzMap []Blizzard, start, end Position, width, height int){
	coordMap := map[Position][]int{}
	for _, blizzard := range blizzMap {
		coordMap[blizzard.pos] = append(coordMap[blizzard.pos], blizzard.dir)
	}
	for y := 0; y < height; y++{
		for x := 0; x < width; x++ {
			coord := Position{x, y}
			if coord == start || coord == end {
				fmt.Print(".")
			}else if x == 0 || y == 0 || x == (width-1) || y == (height-1) {
				fmt.Print("#")
			} else {
				blizzes := coordMap[coord]
				if len(blizzes) == 0 {
					fmt.Print(".")
				} else if len(blizzes) == 1 {
					dir := blizzes[0]
					if dir == 0 {
						fmt.Print("^")
					} else if dir == 1 {
						fmt.Print("v")
					} else if dir == 2 {
						fmt.Print("<")
					} else if dir == 3 {
						fmt.Print(">")
					} else {
						fmt.Print("@")
					}

				} else {
					fmt.Printf("%d", len(blizzes))
				}
			}
		}
		fmt.Println()	
	}
	fmt.Println()
	fmt.Println()
}

func run(blizzMap []Blizzard, start, end Position, width, height int) (int, []Blizzard) {
	states := []Position{start}
	for round := 0; ; round++ {
		blizzMap2 := []Blizzard{}
		coordMap := map[Position]bool{}
		for _, blizzard := range blizzMap {
			newPos := move(blizzard.pos, blizzard.dir, width, height)
			coordMap[newPos] = true
			blizzMap2 = append(blizzMap2, Blizzard{newPos, blizzard.dir})
		}
		blizzMap = blizzMap2

		newStates := []Position{}
		destMap := map[Position]bool{}
		for _, state := range states {
			if coordMap[state] {
				// State occupied by blizzard, not valid, abandon
				continue
			}

			for j := -1; j < len(moveDirections); j++ {
				newPos := Position{state[0], state[1]}
				if j >= 0 {
					newPos[0] += moveDirections[j][0]
					newPos[1] += moveDirections[j][1]
				}
				if newPos[0] < 1 || (newPos[1] < 1 && newPos != start && newPos != end) || newPos[0] >= (width-1) || (newPos[1] >= (height-1) && newPos != end && newPos != start){
					continue
				}
				if destMap[newPos] {
					continue
				}
				if newPos == end {
					return round+1, blizzMap
				}
				destMap[newPos] = true
				newStates = append(newStates, newPos)
			}
		}
		if len(newStates) == 0 {
			panic("No states left")
		}
		states = newStates
		fmt.Printf("After round %d, states=%d\n", round+1, len(states))

		//printMap(blizzMap, start, end, width, height)
	}
}

func Part1(input io.Reader) int {
	blizzMap, start, end, width, height := readInput(input)
	rounds, _ := run(blizzMap, start, end, width, height)
	
	return rounds + 1
}

func Part2(input io.Reader) int {
	blizzMap, start, end, width, height := readInput(input)
	rounds1, blizzMap := run(blizzMap, start, end, width, height)
	rounds2, blizzMap := run(blizzMap, end, start, width, height)
	rounds3, blizzMap := run(blizzMap, start, end, width, height)
	return rounds1 + rounds2 + rounds3 + 1
}

func main() {
	utils.Run("day24.txt", Part1, Part2)
}
