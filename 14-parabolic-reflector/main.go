package main

import (
	"bufio"
	"os"
	"strings"
)

type position struct {
	row, col int
}

func (p position) north() position {
	return position{p.row - 1, p.col}
}

func (p position) south() position {
	return position{p.row + 1, p.col}
}

func (p position) east() position {
	return position{p.row, p.col + 1}
}

func (p position) west() position {
	return position{p.row, p.col - 1}
}

type cellType uint8

const (
	empty cellType = iota
	round
	cube
	side
)

type platform [][]cellType

func (p platform) at(pos position) cellType {
	if pos.row < 0 || pos.row >= len(p) {
		return side
	}
	if pos.col < 0 || pos.col >= len(p[pos.row]) {
		return side
	}
	return p[pos.row][pos.col]
}

func (p platform) set(pos position, t cellType) {
	p[pos.row][pos.col] = t
}

func (p platform) tiltNorth() {
	for r := range p {
		for c := range p[r] {
			pos := position{r, c}

			if p.at(pos) != round {
				continue
			}

			for p.at(pos.north()) == empty {
				p.set(pos, empty)
				pos = pos.north()
				p.set(pos, round)
			}
		}
	}
}

func (p platform) tiltSouth() {
	for r := len(p) - 1; r >= 0; r-- {
		for c := range p[r] {
			pos := position{r, c}

			if p.at(pos) != round {
				continue
			}

			for p.at(pos.south()) == empty {
				p.set(pos, empty)
				pos = pos.south()
				p.set(pos, round)
			}
		}
	}
}

func (p platform) tiltWest() {
	for r := range p {
		for c := range p[r] {
			pos := position{r, c}

			if p.at(pos) != round {
				continue
			}

			for p.at(pos.west()) == empty {
				p.set(pos, empty)
				pos = pos.west()
				p.set(pos, round)
			}
		}
	}
}

func (p platform) tiltEast() {
	for r := range p {
		for c := len(p[r]) - 1; c >= 0; c-- {
			pos := position{r, c}

			if p.at(pos) != round {
				continue
			}

			for p.at(pos.east()) == empty {
				p.set(pos, empty)
				pos = pos.east()
				p.set(pos, round)
			}
		}
	}
}

func (p platform) weightNorth() int {
	var w int
	for r := range p {
		for c := range p[r] {
			if p.at(position{r, c}) == round {
				w += len(p) - r
			}
		}
	}
	return w
}

func (p platform) String() string {
	sb := strings.Builder{}

	for r := range p {
		for c := range p[r] {
			switch p.at(position{r, c}) {
			case round:
				sb.WriteRune('O')
			case cube:
				sb.WriteRune('#')
			case empty:
				sb.WriteRune('.')
			}
		}
		sb.WriteRune('\n')
	}

	return sb.String()
}

func main() {
	var p platform

	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		if s.Text() == "" {
			break
		}

		row := []cellType{}
		for _, c := range s.Text() {
			switch c {
			case 'O':
				row = append(row, round)
			case '#':
				row = append(row, cube)
			case '.':
				row = append(row, empty)
			}
		}
		p = append(p, row)
	}

	configs := map[string]int{}
	ca, cb := 0, 0

	for i := 0; i < 1000000000; i++ {
		p.tiltNorth()
		p.tiltWest()
		p.tiltSouth()
		p.tiltEast()

		key := p.String()
		if _, ok := configs[key]; !ok {
			configs[key] = i
		} else {
			ca = configs[key]
			cb = i
			break
		}
	}

	for i := 1; i < (1000000000-ca)%(cb-ca); i++ {
		p.tiltNorth()
		p.tiltWest()
		p.tiltSouth()
		p.tiltEast()
	}

	println(p.weightNorth())
}
