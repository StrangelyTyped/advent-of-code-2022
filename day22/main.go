package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/strangelytyped/advent-of-code-2022/utils"
)

type Position [2]int

var directions = []Position{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

type Instruction struct {
	turnDir  rune
	distance int
}

func readInput(input io.Reader) ([]string, []Instruction) {
	scanner := bufio.NewScanner(input)

	maze := []string{}

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			break
		}

		maze = append(maze, line)
	}
	scanner.Scan()
	instructions := scanner.Text()

	instructionList := []Instruction{}
	instruction := Instruction{
		turnDir: ' ',
	}

	for _, chr := range instructions {
		if chr == 'L' || chr == 'R' {
			instructionList = append(instructionList, instruction)
			instruction = Instruction{
				turnDir: chr,
			}
		} else {
			instruction.distance = instruction.distance*10 + int(chr-'0')
		}
	}
	instructionList = append(instructionList, instruction)

	return maze, instructionList
}

func Part1(input io.Reader) int {

	maze, instructions := readInput(input)

	pos := Position{strings.Index(maze[0], "."), 0}
	dir := 0

	for _, instr := range instructions {
		if instr.turnDir == 'L' {
			dir -= 1
			if dir < 0 {
				dir = len(directions) - 1
			}
		} else if instr.turnDir == 'R' {
			dir = (dir + 1) % len(directions)
		}
		for i := 0; i < instr.distance; i++ {
			nextSpace := nextSpace(maze, pos, dir)
			if maze[nextSpace[1]][nextSpace[0]] == '#' {
				break
			}
			pos = nextSpace
		}
	}

	return (1+pos[1])*1000 + (1+pos[0])*4 + dir
}

func nextSpace(maze []string, pos Position, dir int) Position {
	moved := directions[dir]
	for {
		pos = Position{pos[0] + moved[0], pos[1] + moved[1]}

		if pos[1] < 0 {
			pos[1] = len(maze) - 1
		}
		if pos[1] >= len(maze) {
			pos[1] = 0
		}

		if pos[0] < 0 {
			pos[0] = len(maze[pos[1]]) - 1
		}
		if pos[0] >= len(maze[pos[1]]) && moved[0] != 0 {
			pos[0] = 0
		}

		if pos[0] < len(maze[pos[1]]) && maze[pos[1]][pos[0]] != ' ' {
			break
		}
	}
	return pos
}

var zoneTransitionMap = map[[2]int]bool{}

func nextSpace2(maze []string, pos Position, dir int) (Position, int) {
	moved := directions[dir]
	newPos := Position{pos[0] + moved[0], pos[1] + moved[1]}

	if newPos[1] < 0 {
		if newPos[0] < 100 {
			// 1-6
			fmt.Println("1-6")
			newPos[1] = 100 + newPos[0]
			newPos[0] = 0
			dir = 0
		} else {
			// 2-6
			fmt.Println("2-6")
			newPos[1] = len(maze) - 1
			newPos[0] -= 100
		}
	} else if newPos[1] < 50 {
		if newPos[0] < 50 {
			// 1-4
			fmt.Println("1-4")
			newPos[1] = 149 - newPos[1]
			newPos[0] = 0
			dir = 0
		} else if newPos[0] >= 150 {
			// 2-5
			fmt.Println("2-5")
			newPos[1] = 149 - newPos[1]
			newPos[0] = 99
			dir = 2
		}
	} else if newPos[1] < 100 {
		if newPos[0] >= 100 {
			if dir == 0 {
				// 3-2
				fmt.Println("3-2")
				newPos[0] = newPos[1] + 50
				newPos[1] = 49
				dir = 3
			} else {
				// 2-3
				fmt.Println("2-3")
				newPos[1] = newPos[0] - 50
				newPos[0] = 99
				dir = 2
			}
		} else if newPos[0] < 50 {
			if dir == 2 {
				// 3-4
				fmt.Println("3-4")
				newPos[0] = newPos[1] - 50
				newPos[1] = 100
				dir = 1
			} else {
				// 4-3
				fmt.Println("4-3")
				newPos[1] = newPos[0] + 50
				newPos[0] = 50
				dir = 0
			}
		}
	} else if newPos[1] < 150 {
		if newPos[0] < 0 {
			// 4-1
			fmt.Println("4-1")
			newPos[1] = 149 - newPos[1]
			newPos[0] = 50
			dir += 2
		} else if newPos[0] >= 100 {
			// 5-2
			fmt.Println("5-2")
			newPos[1] = 149 - newPos[1]
			newPos[0] = 149
			dir += 2
		}
	} else if newPos[1] < 200 {
		if newPos[0] < 0 {
			// 6-1
			fmt.Println("6-1")
			newPos[0] = newPos[1] - 100
			newPos[1] = 0
			dir = 1
		} else if newPos[0] >= 50 {
			if dir == 0 {
				// 6-5
				fmt.Println("6-5")
				newPos[0] = newPos[1] - 100
				newPos[1] = 149
				dir = 3
			} else {
				// 5-6
				fmt.Println("5-6")
				newPos[1] = newPos[0] + 100
				newPos[0] = 49
				dir = 2
			}
		}
	} else {
		// 6-2
		fmt.Println("6-2")
		newPos[1] = 0
		newPos[0] += 100
	}

	return newPos, dir % len(directions)
}

func Part2(input io.Reader) int {
	maze, instructions := readInput(input)

	pos := Position{strings.Index(maze[0], "."), 0}
	dir := 0

	for _, instr := range instructions {
		if instr.turnDir == 'L' {
			dir -= 1
			if dir < 0 {
				dir = len(directions) - 1
			}
		} else if instr.turnDir == 'R' {
			dir = (dir + 1) % len(directions)
		}
		for i := 0; i < instr.distance; i++ {
			nextSpace, newDir := nextSpace2(maze, pos, dir)
			if maze[nextSpace[1]][nextSpace[0]] == '#' {
				break
			}
			pos = nextSpace
			dir = newDir
		}
		fmt.Printf("(%d, %d)\n", pos[0], pos[1])
	}

	return (1+pos[1])*1000 + (1+pos[0])*4 + dir
}

func main() {
	utils.Run("day22.txt", Part1, Part2)
}
