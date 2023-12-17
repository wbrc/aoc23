package main

import (
	"bufio"
	"os"
)

type contraption struct {
	grid      [][]byte
	energized [][][4]bool
	energy    int
}

func isEnergized(e [4]bool) bool {
	for _, b := range e {
		if b {
			return true
		}
	}
	return false
}

type direction int

const (
	up direction = iota
	down
	left
	right
)

func (d direction) move(r, c int) (int, int) {
	switch d {
	case up:
		return r - 1, c
	case down:
		return r + 1, c
	case left:
		return r, c - 1
	case right:
		return r, c + 1
	}
	return r, c
}

func (c *contraption) beamTo(row, col int, dir direction) {
	if row < 0 || row >= len(c.grid) || col < 0 || col >= len(c.grid[row]) {
		return
	}

	if !isEnergized(c.energized[row][col]) {
		c.energy++
	}

	if c.energized[row][col][dir] {
		return
	}

	c.energized[row][col][dir] = true

	if c.grid[row][col] == '|' && (dir == left || dir == right) {
		c.beamTo(row-1, col, up)
		c.beamTo(row+1, col, down)
		return
	}

	if c.grid[row][col] == '-' && (dir == up || dir == down) {
		c.beamTo(row, col-1, left)
		c.beamTo(row, col+1, right)
		return
	}

	switch {
	case c.grid[row][col] == '/' && dir == up:
		dir = right
	case c.grid[row][col] == '/' && dir == down:
		dir = left
	case c.grid[row][col] == '/' && dir == left:
		dir = down
	case c.grid[row][col] == '/' && dir == right:
		dir = up
	case c.grid[row][col] == '\\' && dir == up:
		dir = left
	case c.grid[row][col] == '\\' && dir == down:
		dir = right
	case c.grid[row][col] == '\\' && dir == left:
		dir = up
	case c.grid[row][col] == '\\' && dir == right:
		dir = down
	}

	row, col = dir.move(row, col)
	c.beamTo(row, col, dir)
}

func (c *contraption) clear() {
	for row := range c.energized {
		for col := range c.energized[row] {
			for dir := range c.energized[row][col] {
				c.energized[row][col][dir] = false
			}
		}
	}
	c.energy = 0
}

func main() {
	contraption := &contraption{}

	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		contraption.grid = append(contraption.grid, []byte(s.Text()))
		contraption.energized = append(contraption.energized, make([][4]bool, len(s.Text())))
	}

	best := 0

	for col := range contraption.grid[0] {
		contraption.clear()
		contraption.beamTo(0, col, down)
		if contraption.energy > best {
			best = contraption.energy
		}

		contraption.clear()
		contraption.beamTo(len(contraption.grid)-1, col, up)
		if contraption.energy > best {
			best = contraption.energy
		}
	}

	for row := range contraption.grid {
		contraption.clear()
		contraption.beamTo(row, 0, right)
		if contraption.energy > best {
			best = contraption.energy
		}

		contraption.clear()
		contraption.beamTo(row, len(contraption.grid[row])-1, left)
		if contraption.energy > best {
			best = contraption.energy
		}
	}

	println(best)
}
