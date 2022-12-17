package main

import (
	"bufio"
	//"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/strangelytyped/advent-of-code-2022/utils"
)

var parseRe = regexp.MustCompile(`Valve (.*) has flow rate=(\d+); tunnels? leads? to valves? (.*)`)

const debug = false

func solve(valves map[int]ValveRoom, remainingTime int, currentRoom []int, visited uint64, memo map[Memo]int, path [][]int) int {
	if remainingTime <= 0 {
		return 0
	}

	currentPos := roomListToPos(currentRoom)

	myMemo1 := Memo{currentPos, visited, remainingTime}
	if val, has := memo[myMemo1]; has {
		return val
	}

	flowOut := 0

	// for every combination of new tunnels
	newTunnels := [][]int{}
	for _, currentPersonRoom := range currentRoom {
		valveDef := valves[currentPersonRoom]
		tunnelOptions := valveDef.tunnels
		if valveDef.flow > 0 && (visited&valveDef.bitFlag) == 0 {
			tunnelOptions = append(tunnelOptions, currentPersonRoom)
		}
		newTunnels = append(newTunnels, tunnelOptions)
	}
	newTunnels = combine(newTunnels)
	for _, newRooms := range newTunnels {

		thisVisited := visited
		flowHere := 0

		for i := range currentRoom {
			if currentRoom[i] == newRooms[i] {
				valveDef := valves[currentRoom[i]]
				if valveDef.flow > 0 && remainingTime > 0 && (thisVisited&valveDef.bitFlag) == 0 {
					thisVisited |= valveDef.bitFlag
					flowHere += valveDef.flow * (remainingTime - 1)
				}
			}
		}

		newPath := path
		if debug {
			newPath = append(path, newRooms)
		}

		flowOption := solve(valves, remainingTime-1, newRooms, thisVisited, memo, newPath)
		flowOut = utils.Max(flowOut, flowHere+flowOption)

	}

	memo[myMemo1] = flowOut

	return flowOut
}

func combine(tunnelOptions [][]int) [][]int {
	out := [][]int{}

	if len(tunnelOptions) == 1 {
		for _, option := range tunnelOptions[0] {
			out = append(out, []int{option})
		}
	} else {
		extant := map[uint64]bool{}
		subCombinations := combine(tunnelOptions[0 : len(tunnelOptions)-1])
		for _, option := range tunnelOptions[len(tunnelOptions)-1] {
			for _, subCombo := range subCombinations {
				newCombo := append(subCombo, option)
				posBitflag := roomListToPos(newCombo)
				if !extant[posBitflag] {
					out = append(out, newCombo)
					extant[posBitflag] = true
				}
			}
		}
	}
	return out
}

func roomListToPos(currentRoom []int) uint64 {
	val := uint64(0)
	for _, room := range currentRoom {
		val |= 1 << room
	}
	return val
}

type ValveRoom struct {
	flow    int
	tunnels []int
	bitFlag uint64
	index   int
	name    string
}

type Memo struct {
	point         uint64
	visited       uint64
	remainingTime int
}

func Part1(input io.Reader) int {
	valves, startIdx := readInput(input)
	memo := map[Memo]int{}
	result := solve(valves, 30, []int{startIdx}, uint64(0), memo, [][]int{})
	return result
}

func readInput(input io.Reader) (map[int]ValveRoom, int) {
	scanner := bufio.NewScanner(input)
	valves := map[int]ValveRoom{}
	valveIdMap := map[string]int{}
	nextValveIdx := 0

	getValveIdx := func(valve string) int {
		id, has := valveIdMap[valve]
		if !has {
			id = nextValveIdx
			nextValveIdx++
			valveIdMap[valve] = id
		}
		return id
	}

	startIdx := 0
	for scanner.Scan() {
		line := scanner.Text()

		parts := parseRe.FindStringSubmatch(line)
		valveId := parts[1]
		flow := utils.AtoiOrPanic(parts[2])
		tunnels := strings.Split(parts[3], ", ")

		valveIdx := getValveIdx(valveId)

		if valveId == "AA" {
			startIdx = valveIdx
		}

		tunnelIds := []int{}
		for _, tunnel := range tunnels {
			tunnelIds = append(tunnelIds, getValveIdx(tunnel))
		}

		valves[valveIdx] = ValveRoom{flow, tunnelIds, uint64(1) << valveIdx, valveIdx, valveId}
	}
	return valves, startIdx
}

func Part2(input io.Reader) int {
	valves, startIdx := readInput(input)
	memo := map[Memo]int{}
	result := solve(valves, 26, []int{startIdx, startIdx}, uint64(0), memo, [][]int{})
	return result
}

func main() {
	utils.Run("day16.txt", Part1, Part2)
}
