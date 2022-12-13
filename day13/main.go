package main

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/strangelytyped/advent-of-code-2022/utils"
)

type Token struct {
	intVal   int
	children []*Token
	isInt    bool
}

func wrap(a1 *Token) *Token {
	return &Token{children: []*Token{a1}}
}

func cmp(a1, a2 *Token) int {
	if a1.isInt && a2.isInt {
		return a1.intVal - a2.intVal
	} else if a1.isInt {
		return cmp(wrap(a1), a2)
	} else if a2.isInt {
		return cmp(a1, wrap(a2))
	} else {
		// both lists
		for i := 0; i < len(a1.children) && i < len(a2.children); i++ {
			x := cmp(a1.children[i], a2.children[i])
			if x != 0 {
				return x
			}
		}
		return len(a1.children) - len(a2.children)
	}
}

func parse(line string, token *Token) int {
	idx := 0
	if line[0] == '[' {
		// Array mode
		idx++
		for idx < len(line) {
			childToken := Token{}
			token.children = append(token.children, &childToken)
			readChrs := parse(line[idx:], &childToken)
			idx += readChrs
			if line[idx] == ']' {
				idx++
				break
			}
			idx++
		}
	} else if line[0] == ']' {
		return 0
	} else {
		// number mode
		val := 0
		for {
			if line[idx] == ',' || line[idx] == ']' {
				break
			}
			val = val*10 + int(line[idx]-'0')
			idx++
		}
		token.intVal = val
		token.isInt = true
	}
	return idx
}

func Part1(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	idx := 1
	sum := 0
	for scanner.Scan() {
		line1 := scanner.Text()
		scanner.Scan()
		line2 := scanner.Text()
		scanner.Scan()

		root1 := Token{}
		parse(line1, &root1)
		root2 := Token{}
		parse(line2, &root2)

		if cmp(&root1, &root2) <= 0 {
			fmt.Printf("Match %d\n", idx)
			sum += idx
		} else {
			fmt.Printf("Fail %d\n", idx)
		}
		idx++
	}
	return sum
}

func makeIntToken(x int) *Token {
	return &Token{intVal: x, isInt: true}
}

func Part2(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	tokenA := wrap(makeIntToken(2))
	tokenB := wrap(makeIntToken(6))

	lines := []*Token{tokenA, tokenB}

	for scanner.Scan() {
		line1 := scanner.Text()
		if len(strings.TrimSpace(line1)) == 0 {
			continue
		}

		root1 := Token{}
		parse(line1, &root1)
		lines = append(lines, &root1)
	}

	sort.Slice(lines, func(i, j int) bool {
		return cmp(lines[i], lines[j]) < 0
	})
	return (index(lines, tokenA) + 1) * (index(lines, tokenB) + 1)
}

func index(lines []*Token, tokenA *Token) int {
	for i, x := range lines {
		if tokenA == x {
			return i
		}
	}
	return -1
}

func main() {
	utils.Run("day13.txt", Part1, Part2)
}
