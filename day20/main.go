package main

import (
	"bufio"
	"io"

	"github.com/strangelytyped/advent-of-code-2022/utils"
)

type Node struct {
	value int
	currentIndex int
}

func mix(numbers, nMutable []*Node){
	for _, node := range numbers {
		srcIdx := node.currentIndex
		destIdx := (srcIdx + node.value) % (len(numbers)-1)
		if destIdx <= 0 {
			destIdx += (len(numbers)-1)
		}

		if destIdx != srcIdx {
			direction := utils.Sign(destIdx - srcIdx)
			for i := srcIdx; i != destIdx; i += direction {
				nMutable[i].currentIndex, nMutable[i + direction].currentIndex = nMutable[ i+ direction].currentIndex, nMutable[i].currentIndex
				nMutable[i], nMutable[i + direction] = nMutable[ i + direction], nMutable[i]
			}
		}
	}
}

func readInput(input io.Reader) ([]*Node, *Node) {
	scanner := bufio.NewScanner(input)
	numbers := []*Node{}
	var zeroVal *Node
	for scanner.Scan() {
		number := utils.AtoiOrPanic(scanner.Text())
		numbers = append(numbers, &Node{number, len(numbers)})
		if number == 0 {
			zeroVal = numbers[len(numbers)-1]
		}
	}
	return numbers, zeroVal
}

func getAnswer(numbers2 []*Node, zeroVal *Node) int {
	return numbers2[(zeroVal.currentIndex + 1000) % len(numbers2)].value + numbers2[(zeroVal.currentIndex + 2000) % len(numbers2)].value + numbers2[(zeroVal.currentIndex + 3000) % len(numbers2)].value
}

func Part1(input io.Reader) int {
	numbers, zeroVal := readInput(input)

	numbers2 := append([]*Node{}, numbers...)

	mix(numbers, numbers2)

	return getAnswer(numbers2, zeroVal)
}

func Part2(input io.Reader) int {
	key := 811589153
	numbers, zeroVal := readInput(input)
	for _, node := range numbers {
		node.value *= key
	}

	numbers2 := append([]*Node{}, numbers...)

	for i := 0; i < 10; i++ {
		mix(numbers, numbers2)
	}

	return getAnswer(numbers2, zeroVal)
}

func main() {
	utils.Run("day20.txt", Part1, Part2)
}
