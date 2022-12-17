package main

import (
	"bufio"
	"fmt"
	"sort"

	"io"
	"regexp"
	"strings"

	"github.com/strangelytyped/advent-of-code-2022/utils"
)

var parseRe = regexp.MustCompile(`Valve (.*) has flow rate=(\d+); tunnels? leads? to valves? (.*)`)


func solve(valves map[string]ValveRoom, remainingTime int, currentRoom string, visited uint64, flowIn int, pathPlan []string) int {
	if remainingTime <= 0 {
		return flowIn
	}

	if len(pathPlan) > 0 {
		fmt.Printf("Moving to %s\n", pathPlan[0])
		return solve(valves, remainingTime - 1, pathPlan[0], visited, flowIn, pathPlan[1:])
	}

	currentValve := valves[currentRoom]
	if len(pathPlan) == 0 && currentValve.flow > 0 && visited & currentValve.bitFlag == 0 {
		thisRoom := currentValve.flow * (remainingTime - 1)
		fmt.Printf("Opening %s for %d (cumulative %d)\n", currentRoom, thisRoom, flowIn + thisRoom)
		return solve(valves, remainingTime - 2, currentRoom, visited | currentValve.bitFlag, flowIn + thisRoom, nil)
	}

	pathPlan = calculatePathPlan(valves, currentRoom, visited, remainingTime)
	if len(pathPlan) > 0 {
		return solve(valves, remainingTime, currentRoom, visited, flowIn, pathPlan)
	}
	return flowIn
}

type ValveCandidate struct {
	valve ValveRoom
	path []string
	stops []string
	reward int
}

var pathCache = map[string][]string{}

func calculatePathPlan(valves map[string]ValveRoom, currentRoom string, visited uint64, remainingTime int) []string {

	candidateValves := map[string]*ValveCandidate{}
	candidateList := []*ValveCandidate{}
	for valveId, valve := range valves {
		if visited & valve.bitFlag == 0 && valve.flow > 0 {
			candidate := &ValveCandidate{
				valve: valve,
			}
			candidateValves[valveId] = candidate
			candidateList = append(candidateList, candidate)
		}
	}

	for valveId, candidate := range candidateValves {
		path := calculatePathBetween(valves, currentRoom, valveId)
		candidate.path = path
		candidate.reward = (remainingTime - (len(path) + 1)) * valves[valveId].flow
		candidate.stops = []string{valveId}
	}

	// TODO: check stops for viability

	sort.Slice(candidateList, func(i, j int) bool {return candidateList[i].reward < candidateList[j].reward})
	return candidateList[0].path
}

func calculatePathBetween(valves map[string]ValveRoom, from, to string) []string {
	key1 := from + "/" + to
	//key2 := to + "/" + from
	if path, has := pathCache[key1]; has {
		return path
	}

	path := findPath(valves, from, to, []string{from})



	pathCache[key1] = path
	//reversePath := 
	// ugh. 
	//pathCache[key2] = reversePath
	return path
}

func findPath(valves map[string]ValveRoom, from, to string, visited []string) []string {
	thisRoom := valves[from]
	var bestCandidate []string
	for _, candidate := range thisRoom.tunnels {
		if candidate == to {
			return []string{to}
		}
		
		
		alreadyVisited := false
		for _, visitedEntry := range visited {
			if visitedEntry == candidate {
				alreadyVisited = true
				break
			}
		}
		if alreadyVisited {
			continue
		}

		subRoute := findPath(valves, candidate, to, append(visited, candidate))
		if subRoute != nil {
			subRoute = append([]string{candidate}, subRoute...)
			if bestCandidate == nil || len(subRoute) < len(bestCandidate) {
				bestCandidate = subRoute
			}
		}
	}

	return bestCandidate
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
	tunnels []string
	bitFlag uint64
	index   int
	name    string
}

func Part1(input io.Reader) int {
	valves, startIdx := readInput(input)
	result := solve(valves, 30, startIdx, uint64(0), 0, nil)
	return result
}

func readInput(input io.Reader) (map[string]ValveRoom, string) {
	scanner := bufio.NewScanner(input)
	valves := map[string]ValveRoom{}
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

	startId := ""
	for scanner.Scan() {
		line := scanner.Text()

		parts := parseRe.FindStringSubmatch(line)
		valveId := parts[1]
		flow := utils.AtoiOrPanic(parts[2])
		tunnels := strings.Split(parts[3], ", ")

		valveIdx := getValveIdx(valveId)

		if valveId == "AA" {
			startId = valveId
		}

		valves[valveId] = ValveRoom{flow, tunnels, uint64(1) << valveIdx, valveIdx, valveId}
	}
	return valves, startId
}

func Part2(input io.Reader) int {
	return 0
}

func main() {
	utils.Run("day16.txt", Part1, Part2)
}
