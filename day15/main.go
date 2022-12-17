package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"regexp"

	"github.com/strangelytyped/advent-of-code-2022/utils"
)

const BEACON byte = 'B'
const EMPTY byte = '#'
const SENSOR byte = 'S'

type Position [2]int

type SensorResult struct {
	beaconPos Position
	delta     Position
	distance  int
}

var parseRe = regexp.MustCompile(`-?\d+`)

func distanceFromDelta(delta Position) int {
	return delta[0] + delta[1]
}

func delta(d1, d2 Position) (Position, int) {
	delta := Position{utils.Abs(d1[0] - d2[0]), utils.Abs(d1[1] - d2[1])}
	return delta, distanceFromDelta(delta)
}

func Part1(input io.Reader, targetY int) int {
	sensors, minX, maxX := parseInput(input)

	cave := mapAtY(sensors, targetY)

	emptyLocs := 0

	for x := minX; x <= maxX; x++ {
		spot := cave[x]

		if spot == EMPTY {
			emptyLocs++
		}
	}

	return emptyLocs
}

func mapAtY(sensors map[Position]SensorResult, y int) map[int]byte {
	cave := map[int]byte{}
	for sensor, beaconData := range sensors {
		if beaconData.beaconPos[1] == y {
			cave[beaconData.beaconPos[0]] = BEACON
		}
		if sensor[1] == y {
			cave[sensor[0]] = SENSOR
		}
		distanceLimit := beaconData.distance - (utils.Abs(y - sensor[1]))
		if distanceLimit < 0 {
			continue
		}

		for x := -distanceLimit; x <= distanceLimit; x++ {
			if _, isMarked := cave[sensor[0] + x]; !isMarked {
				cave[sensor[0] + x] = EMPTY
			}
		}			
	}
	return cave
}


func parseInput(input io.Reader) (map[Position]SensorResult, int, int) {
	scanner := bufio.NewScanner(input)

	sensors := map[Position]SensorResult{}
	minX := math.MaxInt32
	maxX := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := parseRe.FindAllString(line, -1)
		sensor := Position{utils.AtoiOrPanic(parts[0]), utils.AtoiOrPanic(parts[1])}
		beacon := Position{utils.AtoiOrPanic(parts[2]), utils.AtoiOrPanic(parts[3])}
		delta, distance := delta(sensor, beacon)
		sensors[sensor] = SensorResult{beacon, delta, distance}

		minX = utils.Min(sensor[0]-distance, minX)
		maxX = utils.Max(sensor[0]+distance, maxX)
	}
	return sensors, minX, maxX
}

func check(sensors map[Position]SensorResult, x, y, coordLimit int) int {
	if x < 0 || y < 0 || x > coordLimit || y > coordLimit {
		return 0
	}
	for sensor, result := range sensors {

		_, distance := delta(sensor, Position{x, y})
		
		if distance <= result.distance {
			return 0
		}
	}
	fmt.Printf("Found %d/%d\n", x, y)
	return x * 4000000 + y
}

func Part2(input io.Reader, coordLimit int) int {

	sensors, _, _ := parseInput(input)

	for sensor, beaconData := range sensors { 
		distance := beaconData.distance + 1
		for x := -distance; x <= distance; x++ {
			y1 := distance - x
			y2 := -y1

			if r := check(sensors, sensor[0] + x, sensor[1] + y1, coordLimit); r != 0 {
				return r
			}
			if r := check(sensors, sensor[0] + x, sensor[1] + y2, coordLimit); r != 0 {
				return r
			}
		}
	}

	return -1
}

func main() {
	utils.Run("day15.txt", func(r io.Reader) int { return Part1(r, 2000000) }, func(r io.Reader) int {return Part2(r, 4000000)})
}
