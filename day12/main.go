package main

import (
	"bufio"
	"container/heap"
	"io"
	"math"

	"github.com/strangelytyped/advent-of-code-2022/utils"
)

type Position [2]int

var validDirections = []Position{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

type OpenEntry struct {
	pos    Position
	length int
}

type OpenEntryList []*OpenEntry

func (pq OpenEntryList) Len() int { return len(pq) }

func (pq OpenEntryList) Less(i, j int) bool {
	return pq[i].length < pq[j].length
}

func (pq OpenEntryList) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *OpenEntryList) Push(x any) {
	item := x.(*OpenEntry)
	*pq = append(*pq, item)
}

func (pq *OpenEntryList) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func search(terrain [][]int, goal Position, openList OpenEntryList) int {
	closedList := map[Position]bool{
		openList[0].pos: true,
	}
	for len(openList) > 0 {
		entry := heap.Pop(&openList).(*OpenEntry)

		for _, dir := range validDirections {
			p := Position{entry.pos[0] + dir[0], entry.pos[1] + dir[1]}

			if p[0] < 0 || p[1] < 0 || p[1] >= len(terrain) || p[0] >= len(terrain[p[1]]) || 
				(terrain[p[1]][p[0]]-terrain[entry.pos[1]][entry.pos[0]] > 1) {
				continue
			}

			if goal == p {
				return entry.length
			}

			_, found := closedList[p]
			if found {
				continue
			}

			newEntry := OpenEntry{
				p,
				entry.length + 1,
			}
			heap.Push(&openList, &newEntry)
			closedList[newEntry.pos] = true
		}
	}
	return -1
}

func readTerrain(input io.Reader) ([][]int, Position, Position) {
	terrain := [][]int{}
	var start, end Position
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		terrain = append(terrain, make([]int, len(line)))
		row := &terrain[len(terrain)-1]
		for pos, chr := range line {
			if chr == 'S' {
				start = Position{pos, len(terrain) - 1}
			} else if chr == 'E' {
				end = Position{pos, len(terrain) - 1}
				(*row)[pos] = 'z' - 'a'
			} else {
				(*row)[pos] = int(chr - 'a')
			}
		}
	}
	return terrain, start, end
}

func Part1(input io.Reader) int {
	terrain, start, end := readTerrain(input)

	startEntry := OpenEntry{start, 1}

	return search(terrain, end, OpenEntryList{&startEntry})
}

func Part2(input io.Reader) int {
	min := math.MaxInt

	terrain, _, end := readTerrain(input)

	for y := 0; y < len(terrain); y++ {
		for x := 0; x < len(terrain[y]); x++ {
			if terrain[y][x] == 0 {
				startEntry := OpenEntry{Position{x, y}, 1}
				result := search(terrain, end, OpenEntryList{&startEntry})
				if result != -1 && result < min {
					min = result
				}
			}
		}
	}

	return min
}

func main() {
	utils.Run("day12.txt", Part1, Part2)
}
