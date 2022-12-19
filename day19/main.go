package main

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/strangelytyped/advent-of-code-2022/utils"
)
func maxForBlueprint(blueprint Blueprint, tMax int) int {
	initialState := PuzzleState{
		1, 0, 0, 0, 
		0, 0, 0, 0,
	}

	maxOre := utils.Max(utils.Max(utils.Max(blueprint.oreCost, blueprint.clayCost), blueprint.obsOreCost), blueprint.geoOreCost)
	maxClay := blueprint.obsClayCost
	maxObs := blueprint.geoObsCost

	pendingStates := []PuzzleState{initialState}
	for t := 0; t < tMax; t++ {		
		newStates := []PuzzleState{}
		for len(pendingStates) > 0 {
			state := pendingStates[len(pendingStates)-1]
			pendingStates = pendingStates[0:len(pendingStates)-1]

			canBuyOre := state.ore >= blueprint.oreCost
			canBuyClay := state.ore >= blueprint.clayCost
			canBuyObs := state.ore >= blueprint.obsOreCost && state.clay >= blueprint.obsClayCost
			canBuyGeo := state.ore >= blueprint.geoOreCost && state.obs >= blueprint.geoObsCost

			// Can only build 1 thing per round, no point in overproducing
			shouldBuyOre := state.oreBots < maxOre
			shouldBuyClay := state.clayBots < maxClay
			shouldBuyObs := state.obsBots < maxObs

			state.ore += state.oreBots
			state.clay += state.clayBots
			state.obs += state.obsBots
			state.geo += state.geoBots

			// Only consider not-buying-something if there's something worth saving for
			if !(canBuyOre && canBuyClay && canBuyObs && canBuyGeo){
				newStates = checkAndAppend(newStates, state)
			}
			if canBuyGeo {
				newState := PuzzleState{
					state.oreBots,
					state.clayBots,
					state.obsBots,
					state.geoBots + 1,
					state.ore - blueprint.geoOreCost,
					state.clay,
					state.obs - blueprint.geoObsCost,
					state.geo,
				}
				newStates = checkAndAppend(newStates, newState)
			}
			if canBuyObs && shouldBuyObs {
				newState := PuzzleState{
					state.oreBots,
					state.clayBots,
					state.obsBots + 1,
					state.geoBots,
					state.ore - blueprint.obsOreCost,
					state.clay - blueprint.obsClayCost,
					state.obs,
					state.geo,
				}
				newStates = checkAndAppend(newStates, newState)
			}
			if canBuyClay && shouldBuyClay {
				newState := PuzzleState{
					state.oreBots,
					state.clayBots + 1,
					state.obsBots,
					state.geoBots,
					state.ore - blueprint.clayCost,
					state.clay,
					state.obs,
					state.geo,
				}
				newStates = checkAndAppend(newStates, newState)
			}
			if canBuyOre && shouldBuyOre {
				newState := PuzzleState{
					state.oreBots + 1,
					state.clayBots,
					state.obsBots,
					state.geoBots,
					state.ore - blueprint.oreCost,
					state.clay,
					state.obs,
					state.geo,
				}
				newStates = checkAndAppend(newStates, newState)
			}
		}
		pendingStates = newStates
		fmt.Printf("At t=%d, state# = %d\n", t, len(pendingStates))
	}

	maxGeodes := 0
	for _, state := range pendingStates {
		maxGeodes = utils.Max(maxGeodes, state.geo)
	}
	return maxGeodes
}

func checkAndAppend(newStates []PuzzleState, state PuzzleState) []PuzzleState {
	for i := range newStates {
		otherState := newStates[i]

		if otherState.ore >= state.ore && otherState.clay >= state.clay && otherState.obs >= state.obs && otherState.geo >= state.geo && 
			otherState.oreBots >= state.oreBots && otherState.clayBots >= state.clayBots && otherState.obsBots >= state.obsBots && otherState.geoBots >= state.geoBots {
			// This other state is quantifiably better or the same
			return newStates
		}
	}
	return append(newStates, state)
}

type PuzzleState struct {
	oreBots int
	clayBots int
	obsBots int
	geoBots int

	ore int
	clay int
	obs int
	geo int
}

type Blueprint struct {
	id int
	oreCost int
	clayCost int
	obsOreCost int
	obsClayCost int
	geoOreCost int
	geoObsCost int
}

func Part1(input io.Reader) int {
	sum := 0
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		blueprint := parseBlueprint(line)
		fmt.Printf("Solving for blueprint %d\n", blueprint.id)
		sum += blueprint.id * maxForBlueprint(blueprint, 24)
	}
	return sum
}

var re = regexp.MustCompile(`(\d+) (?:ore|clay|obsidian)`)

func parseBlueprint(line string) Blueprint {
	id := utils.AtoiOrPanic(line[strings.Index(line, " ")+1:strings.Index(line,":")])
	line = line[strings.Index(line, ":"):]
	costTokens := strings.Split(line, ". ")

	oreCost := re.FindStringSubmatch(costTokens[0])
	clayCost := re.FindStringSubmatch(costTokens[1])
	obsidianCost := re.FindAllStringSubmatch(costTokens[2], -1)
	geodeCost := re.FindAllStringSubmatch(costTokens[3], -1)
	return Blueprint{
		id,
		utils.AtoiOrPanic(oreCost[1]),
		utils.AtoiOrPanic(clayCost[1]),
		utils.AtoiOrPanic(obsidianCost[0][1]),
		utils.AtoiOrPanic(obsidianCost[1][1]), 
		utils.AtoiOrPanic(geodeCost[0][1]),
		utils.AtoiOrPanic(geodeCost[1][1]),
	}
}

func Part2(input io.Reader) int {
	sum := 1
	scanner := bufio.NewScanner(input)
	for i := 0; i < 3; i++ {
		scanner.Scan()
		line := scanner.Text()
		blueprint := parseBlueprint(line)
		fmt.Printf("Solving for blueprint %d\n", blueprint.id)
		geodeCount := maxForBlueprint(blueprint, 32)
		fmt.Printf("Count for blueprint %d is %d\n", blueprint.id, geodeCount)
		sum *= geodeCount
	}
	return sum
}

func main() {
	utils.Run("day19.txt", Part1, Part2)
}
