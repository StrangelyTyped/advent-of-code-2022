package main

import (
	"bufio"
	"io"
	"strings"

	"github.com/strangelytyped/advent-of-code-2022/utils"
)

type MonkeyOp func(a, b int) int

var add MonkeyOp = func(a, b int) int { return a + b }
var sub MonkeyOp = func(a, b int) int { return a - b }
var sub2 MonkeyOp = func(a, b int) int { return b - a }
var mul MonkeyOp = func(a, b int) int { return a * b }
var div MonkeyOp = func(a, b int) int { return a / b }
var div2 MonkeyOp = func(a, b int) int { return b / a }

type Monkey struct {
	name string
	dep1 string
	dep2 string
	op MonkeyOp
	reverseOp1 MonkeyOp
	reverseOp2 MonkeyOp
	literalValue int
	isLiteral bool
}

func readInput(input io.Reader) map[string]Monkey {
	scanner := bufio.NewScanner(input)

	monkeys := map[string]Monkey{}
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ": ")
		monkey := Monkey{name: parts[0]}

		if strings.Index(parts[1], "+") != -1 {
			parts2 := strings.Split(parts[1], " + ")
			monkey.op = add
			monkey.reverseOp1 = sub
			monkey.reverseOp2 = sub
			monkey.dep1 = parts2[0]
			monkey.dep2 = parts2[1]
		} else if strings.Index(parts[1], "-") != -1 {
			parts2 := strings.Split(parts[1], " - ")
			monkey.op = sub
			monkey.reverseOp1 = add
			monkey.reverseOp2 = sub2
			monkey.dep1 = parts2[0]
			monkey.dep2 = parts2[1]
		} else if strings.Index(parts[1], "/") != -1 {
			parts2 := strings.Split(parts[1], " / ")
			monkey.op = div
			monkey.reverseOp1 = mul
			monkey.reverseOp2 = div2
			monkey.dep1 = parts2[0]
			monkey.dep2 = parts2[1]
		} else if strings.Index(parts[1], "*") != -1 {
			parts2 := strings.Split(parts[1], " * ")
			monkey.op = mul
			monkey.reverseOp1 = div
			monkey.reverseOp2 = div
			monkey.dep1 = parts2[0]
			monkey.dep2 = parts2[1]
		} else { 
			monkey.literalValue = utils.AtoiOrPanic(parts[1])
			monkey.isLiteral = true
		}
		monkeys[monkey.name] = monkey
	}
	return monkeys
}

func Part1(input io.Reader) int {
	monkeys := readInput(input)
	return getValue(monkeys, "root")
}

func getValue(monkeys map[string]Monkey, name string) int {
	monkey := monkeys[name]
	if !monkey.isLiteral {
		monkey.literalValue = monkey.op(getValue(monkeys, monkey.dep1), getValue(monkeys, monkey.dep2))
		monkey.isLiteral = true
	}
	return monkey.literalValue
}

func Part2(input io.Reader) int {
	monkeys := readInput(input)

	humanPath := searchForHuman(monkeys, "root")
	rootMonkey := monkeys["root"]

	otherMonkey := rootMonkey.dep1
	if rootMonkey.dep1 == humanPath[0] {
		otherMonkey = rootMonkey.dep2
	}
	target := getValue(monkeys, otherMonkey)

	return unweave(monkeys, humanPath[0], humanPath[1:], target)
}

func unweave(monkeys map[string]Monkey, name string, targetRoute []string, valueIn int) int {
	if name == "humn" {
		return valueIn
	}
	monkey := monkeys[name]
	valueOut := 0
	if monkey.dep1 == targetRoute[0] {
		otherValue := getValue(monkeys, monkey.dep2)
		valueOut = monkey.reverseOp1(valueIn, otherValue)
	} else {
		otherValue := getValue(monkeys, monkey.dep1)
		valueOut = monkey.reverseOp2(valueIn, otherValue)
	}
	return unweave(monkeys, targetRoute[0], targetRoute[1:], valueOut)
}

func searchForHuman(monkeys map[string]Monkey, currentNode string) []string {
	monkey := monkeys[currentNode]
	if monkey.isLiteral {
		return nil
	}
	if monkey.dep1 == "humn" {
		return []string{monkey.dep1}
	}
	if monkey.dep2 == "humn" {
		return []string{monkey.dep2}
	}
	deep1 := searchForHuman(monkeys, monkey.dep1)
	if len(deep1) != 0{
		return append([]string{monkey.dep1}, deep1...)
	}
	deep2 := searchForHuman(monkeys, monkey.dep2)
	if len(deep2) != 0{
		return append([]string{monkey.dep2}, deep2...)
	}
	return nil
}



func main() {
	utils.Run("day21.txt", Part1, Part2)
}
