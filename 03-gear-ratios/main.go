package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"sort"
	"strings"
)

type motor [][]rune

type coord struct {
	x, y int
}

func (m motor) numberBegin(c coord) coord {
	for c.x > 0 && isDigit(m[c.y][c.x-1]) {
		c.x--
	}
	return c
}

func (m motor) numberAt(c coord) int {
	var num int
	for _, r := range m[c.y][c.x:] {
		if !isDigit(r) {
			break
		}
		num *= 10
		num += int(r - '0')
	}
	return num
}

func (m motor) adj(x, y int) []coord {
	maxX := len(m[0])
	maxY := len(m)

	var adj []coord
	if x > 0 {
		adj = append(adj, coord{x - 1, y})
	}
	if x < maxX-1 {
		adj = append(adj, coord{x + 1, y})
	}
	if y > 0 {
		adj = append(adj, coord{x, y - 1})
	}
	if y < maxY-1 {
		adj = append(adj, coord{x, y + 1})
	}
	if x > 0 && y > 0 {
		adj = append(adj, coord{x - 1, y - 1})
	}
	if x < maxX-1 && y > 0 {
		adj = append(adj, coord{x + 1, y - 1})
	}
	if x > 0 && y < maxY-1 {
		adj = append(adj, coord{x - 1, y + 1})
	}
	if x < maxX-1 && y < maxY-1 {
		adj = append(adj, coord{x + 1, y + 1})
	}

	return adj
}

func (m motor) adjNumbers(c coord) []coord {
	var numbers []coord
	for _, a := range m.adj(c.x, c.y) {
		if isDigit(m[a.y][a.x]) {
			numbers = append(numbers, m.numberBegin(a))
		}
	}

	sort.Slice(numbers, func(i, j int) bool {
		if numbers[i].y == numbers[j].y {
			return numbers[i].x < numbers[j].x
		}
		return numbers[i].y < numbers[j].y
	})

	var adjNums []coord
	for i, num := range numbers {
		if i == 0 || num != numbers[i-1] {
			adjNums = append(adjNums, num)
		}
	}

	return adjNums
}

func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

func main() {
	var grid motor

	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		line := strings.TrimSpace(s.Text())
		var row []rune
		for _, r := range line {
			row = append(row, r)
		}
		grid = append(grid, row)
	}

	var sum int
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] != '*' {
				continue
			}

			adjNums := grid.adjNumbers(coord{x, y})
			if len(adjNums) != 2 {
				continue
			}

			num1 := grid.numberAt(adjNums[0])
			num2 := grid.numberAt(adjNums[1])
			sum += num1 * num2
		}
	}
	fmt.Println(sum)
}
