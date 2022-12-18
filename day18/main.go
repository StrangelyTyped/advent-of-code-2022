package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/strangelytyped/advent-of-code-2022/utils"
)

func modelDroplet(input io.Reader) [][][]int {
	scanner := bufio.NewScanner(input)

	droplet := [][][]int{}

	decrementAt := func(x, y, z int) {
		if droplet[x][y][z] == 1 {
			droplet[x][y][z] = -1
		} else {
			droplet[x][y][z]--
		}
	}

	for scanner.Scan() {
		line := scanner.Text()

		strDigits := strings.Split(line, ",")
		coords := make([]int, len(strDigits))
		for idx, strDigit := range strDigits {
			coords[idx] = utils.AtoiOrPanic(strDigit)
		}

		for len(droplet) <= coords[0]+1 {
			droplet = append(droplet, [][]int{})
		}
		for len(droplet[coords[0]]) <= coords[1]+1 {
			droplet[coords[0]] = append(droplet[coords[0]], []int{})
		}
		for len(droplet[coords[0]][coords[1]]) <= coords[2]+1 {
			droplet[coords[0]][coords[1]] = append(droplet[coords[0]][coords[1]], 0)
		}

		exposedSides := 6
		if coords[0] > 0 && len(droplet[coords[0]-1]) > coords[1] && len(droplet[coords[0]-1][coords[1]]) > coords[2] && droplet[coords[0]-1][coords[1]][coords[2]] > 0 {
			decrementAt(coords[0]-1, coords[1], coords[2])
			exposedSides--
		}
		if len(droplet[coords[0]+1]) > coords[1] && len(droplet[coords[0]+1][coords[1]]) > coords[2] && droplet[coords[0]+1][coords[1]][coords[2]] > 0 {
			decrementAt(coords[0]+1, coords[1], coords[2])
			exposedSides--
		}
		if coords[1] > 0 && len(droplet[coords[0]][coords[1]-1]) > coords[2] && droplet[coords[0]][coords[1]-1][coords[2]] > 0 {
			decrementAt(coords[0], coords[1]-1, coords[2])
			exposedSides--
		}
		if len(droplet[coords[0]][coords[1]+1]) > coords[2] && droplet[coords[0]][coords[1]+1][coords[2]] > 0 {
			decrementAt(coords[0], coords[1]+1, coords[2])
			exposedSides--
		}
		if coords[2] > 0 && droplet[coords[0]][coords[1]][coords[2]-1] > 0 {
			decrementAt(coords[0], coords[1], coords[2]-1)
			exposedSides--
		}
		if droplet[coords[0]][coords[1]][coords[2]+1] > 0 {
			decrementAt(coords[0], coords[1], coords[2]+1)
			exposedSides--
		}
		if exposedSides == 0 {
			exposedSides--
		}
		droplet[coords[0]][coords[1]][coords[2]] = exposedSides
	}
	return droplet
}

func Part1(input io.Reader) int {
	droplet := modelDroplet(input)

	sumEdges := 0
	for _, xAxis := range droplet {
		for _, yAxis := range xAxis {
			for _, zAxis := range yAxis {
				if zAxis > 0 {
					sumEdges += zAxis
				}
			}
		}
	}

	return sumEdges
}

func Part2(input io.Reader) int {
	droplet := modelDroplet(input)

	maxZ := 0
	maxY := 0

	// square off droplet and add padding around the edges of the volume so we can reach the edges
	droplet = append(append([][][]int{{}}, droplet...), [][]int{})
	for _, xAxis := range droplet {
		maxY = utils.Max(maxY, len(xAxis)+1)
		for _, yAxis := range xAxis {
			maxZ = utils.Max(maxZ, len(yAxis)+1)
		}
	}
	for x, xAxis := range droplet {
		xAxis = append([][]int{{}}, xAxis...)
		for len(xAxis) <= maxY {
			xAxis = append(xAxis, []int{})
		}
		droplet[x] = xAxis

		for y, yAxis := range xAxis {
			yAxis = append([]int{0}, yAxis...)
			for len(yAxis) <= maxZ {
				yAxis = append(yAxis, 0)
			}
			droplet[x][y] = yAxis
		}
	}

	sumEdges := 0

	openList := [][3]int{{0, 0, 0}}
	closedList := map[[3]int]bool{
		openList[0]: true,
	}

	check := func(x, y, z int) int {
		if x >= 0 && x < len(droplet) && y >= 0 && y < len(droplet[x]) && z >= 0 && z < len(droplet[x][y]) {
			if droplet[x][y][z] > 0 {
				droplet[x][y][z]--
				if droplet[x][y][z] == 0 {
					droplet[x][y][z] = -1
				}
				return 1
			} else if droplet[x][y][z] == 0 {
				key := [3]int{x, y, z}
				if !closedList[key] {
					closedList[key] = true
					openList = append(openList, key)
				}
			}
		}
		return 0
	}

	for len(openList) > 0 {
		coord := openList[len(openList)-1]
		openList = openList[0 : len(openList)-1]

		sumEdges += check(coord[0]-1, coord[1], coord[2])
		sumEdges += check(coord[0]+1, coord[1], coord[2])
		sumEdges += check(coord[0], coord[1]-1, coord[2])
		sumEdges += check(coord[0], coord[1]+1, coord[2])
		sumEdges += check(coord[0], coord[1], coord[2]-1)
		sumEdges += check(coord[0], coord[1], coord[2]+1)
	}

	return sumEdges
}

func printDroplet(droplet [][][]int) {
	for x, xAxis := range droplet {
		fmt.Printf("x=%d\n", x)
		fmt.Print("Y | Z >\nV\n")
		for y, yAxis := range xAxis {
			fmt.Printf("%2d ", y)
			for _, zAxis := range yAxis {
				if zAxis > 0 {
					fmt.Printf("%d", zAxis)
				} else if zAxis < 0 {
					fmt.Print("@")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
		fmt.Println()
		fmt.Println()
	}
}

func main() {
	utils.Run("day18.txt", Part1, Part2)
}
