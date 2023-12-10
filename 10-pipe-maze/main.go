package main

import (
	"bufio"
	"fmt"
	"os"
)

type pipeType uint8

const (
	pipeNone pipeType = iota
	pipeVertical
	pipeHorizontal
	pipeLeftUp
	pipeLeftDown
	pipeRightUp
	pipeRightDown
	pipeStart
)

func pipeTypeFromRune(r rune) pipeType {
	switch r {
	case '|':
		return pipeVertical
	case '-':
		return pipeHorizontal
	case 'L':
		return pipeRightUp
	case 'J':
		return pipeLeftUp
	case '7':
		return pipeLeftDown
	case 'F':
		return pipeRightDown
	case 'S':
		return pipeStart
	default:
		return pipeNone
	}
}

type cameFrom uint8

const (
	cameFromUp cameFrom = iota
	cameFromDown
	cameFromLeft
	cameFromRight
)

type pipe struct {
	pipeType
	cameFrom
	distance int
}

type pos struct {
	row, col int
}

func (p pos) up() pos {
	return pos{p.row - 1, p.col}
}

func (p pos) down() pos {
	return pos{p.row + 1, p.col}
}

func (p pos) left() pos {
	return pos{p.row, p.col - 1}
}

func (p pos) right() pos {
	return pos{p.row, p.col + 1}
}

func (p pos) direction(other pos) cameFrom {
	if p.up() == other {
		return cameFromUp
	} else if p.down() == other {
		return cameFromDown
	} else if p.left() == other {
		return cameFromLeft
	} else if p.right() == other {
		return cameFromRight
	}
	panic("invalid direction")
}

type pipeMaze [][]pipe

func (m pipeMaze) at(p pos) pipe {
	if p.row < 0 || p.row >= len(m) || p.col < 0 || p.col >= len(m[p.row]) {
		return pipe{pipeNone, 0, 0}
	}
	return m[p.row][p.col]
}

func (m pipeMaze) setFrom(p pos, from cameFrom) {
	m[p.row][p.col].cameFrom = from
}

func (m pipeMaze) pre(p pos) pos {
	switch m.at(p).cameFrom {
	case cameFromUp:
		return p.down()
	case cameFromDown:
		return p.up()
	case cameFromLeft:
		return p.right()
	case cameFromRight:
		return p.left()
	default:
		return p
	}
}

func (m pipeMaze) next(p pos) pos {
	switch m.at(p).pipeType {
	case pipeVertical:
		if m.at(p).cameFrom == cameFromUp {
			return p.down()
		}
		return p.up()
	case pipeHorizontal:
		if m.at(p).cameFrom == cameFromLeft {
			return p.right()
		}
		return p.left()
	case pipeLeftUp:
		if m.at(p).cameFrom == cameFromUp {
			return p.left()
		}
		return p.up()
	case pipeLeftDown:
		if m.at(p).cameFrom == cameFromDown {
			return p.left()
		}
		return p.down()
	case pipeRightUp:
		if m.at(p).cameFrom == cameFromUp {
			return p.right()
		}
		return p.up()
	case pipeRightDown:
		if m.at(p).cameFrom == cameFromDown {
			return p.right()
		}
		return p.down()
	default:
		panic("invalid pipe type")
	}
}

func (m pipeMaze) setDistance(p pos, distance int) int {
	if m[p.row][p.col].distance == 0 {
		m[p.row][p.col].distance = distance
		return distance
	}
	if m[p.row][p.col].distance < distance {
		return m[p.row][p.col].distance
	}

	m[p.row][p.col].distance = distance
	return distance
}

func main() {
	maze := pipeMaze{}
	var startPos pos

	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		line := []pipe{}
		for _, r := range s.Text() {
			if r == 'S' {
				startPos = pos{len(maze), len(line)}
			}
			line = append(line, pipe{pipeTypeFromRune(r), 0, 0})
		}
		maze = append(maze, line)
	}

	var startPosition pos

	switch maze.at(startPos.up()).pipeType {
	case pipeVertical, pipeLeftDown, pipeRightDown:
		startPosition = startPos.up()
	}

	switch maze.at(startPos.down()).pipeType {
	case pipeVertical, pipeLeftUp, pipeRightUp:
		startPosition = startPos.down()
	}

	switch maze.at(startPos.left()).pipeType {
	case pipeHorizontal, pipeRightUp, pipeRightDown:
		startPosition = startPos.left()
	}

	switch maze.at(startPos.right()).pipeType {
	case pipeHorizontal, pipeLeftUp, pipeLeftDown:
		startPosition = startPos.right()
	}

	p := startPosition
	maze.setFrom(p, p.direction(startPos))
	dist := 2
	for {
		maze.setDistance(p, dist)

		nextPos := maze.next(p)

		if maze.at(nextPos).pipeType == pipeStart {
			break
		}

		var from cameFrom
		if nextPos.down() == p {
			from = cameFromDown
		} else if nextPos.up() == p {
			from = cameFromUp
		} else if nextPos.left() == p {
			from = cameFromLeft
		} else if nextPos.right() == p {
			from = cameFromRight
		}
		maze.setFrom(nextPos, from)

		p = nextPos

		dist++
	}
	fmt.Println(dist / 2)
}
