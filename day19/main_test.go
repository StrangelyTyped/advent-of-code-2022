package main

import (
	"strings"
	"testing"

	"github.com/strangelytyped/advent-of-code-2022/utils"
)

const blueprint1 string = "Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian."
const blueprint2 string = "Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian."

const maxGeodesBlueprint1Part1 = 9
const maxGeodesBlueprint2Part1 = 12

const maxGeodesBlueprint1Part2 = 56
const maxGeodesBlueprint2Part2 = 62

const testInput = blueprint1 + "\n" + blueprint2
const testResultPart1 = 33

func TestBlueprint1Part1(t *testing.T) {
	result := maxForBlueprint(parseBlueprint(blueprint1), 24)
	if result != maxGeodesBlueprint1Part1 {
		t.Errorf("expected %v, got %v", maxGeodesBlueprint1Part1, result)
	}
}

func TestBlueprint2Part1(t *testing.T) {
	result := maxForBlueprint(parseBlueprint(blueprint2), 24)
	if result != maxGeodesBlueprint2Part1 {
		t.Errorf("expected %v, got %v", maxGeodesBlueprint2Part1, result)
	}
}

func TestBlueprint1Part2(t *testing.T) {
	result := maxForBlueprint(parseBlueprint(blueprint1), 32)
	if result != maxGeodesBlueprint1Part2 {
		t.Errorf("expected %v, got %v", maxGeodesBlueprint1Part2, result)
	}
}

func TestBlueprint2Part2(t *testing.T) {
	result := maxForBlueprint(parseBlueprint(blueprint2), 32)
	if result != maxGeodesBlueprint2Part2 {
		t.Errorf("expected %v, got %v", maxGeodesBlueprint2Part2, result)
	}
}

func TestPart1(t *testing.T) {
	result := Part1(strings.NewReader(utils.CleanInput(testInput)))
	if result != testResultPart1 {
		t.Errorf("expected %v, got %v", testResultPart1, result)
	}
}
