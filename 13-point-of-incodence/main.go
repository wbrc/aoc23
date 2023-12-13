package main

import (
	"bufio"
	"fmt"
	"os"
)

func calcRC(grid [][]bool) ([]uint64, []uint64) {
	var rows, cols []uint64

	for _, row := range grid {
		var r uint64
		for _, b := range row {
			n := 0
			if b {
				n = 1
			}
			r = r<<1 | uint64(n)
		}
		rows = append(rows, r)
	}

	for i := 0; i < len(grid[0]); i++ {
		var c uint64
		for j := 0; j < len(grid); j++ {
			n := 0
			if grid[j][i] {
				n = 1
			}
			c = c<<1 | uint64(n)
		}
		cols = append(cols, c)
	}

	return rows, cols
}

func testMirrorPos(seq []uint64, pos int) (int, bool) {
	length := 0
	for {
		if pos < 0 || pos+1+2*length >= len(seq) {
			return length, true
		}

		if seq[pos] == seq[pos+1+2*length] {
			pos--
			length++
			continue
		}

		return length, false
	}
}

func leftOf(seq []uint64) (int, bool) {
	hasperfect := false
	max := 0
	maxPos := 0
	for i := 0; i < len(seq)-1; i++ {
		l, perf := testMirrorPos(seq, i)
		if perf && l > max {
			hasperfect = true
			max = l
			maxPos = i
		}
	}

	return maxPos + 1, hasperfect
}

func main() {
	var grids [][][]bool
	var grid [][]bool

	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		if s.Text() == "" {
			grids = append(grids, grid)
			grid = nil
			continue
		}

		var row []bool
		for _, r := range s.Text() {
			row = append(row, r == '#')
		}
		grid = append(grid, row)
	}
	if len(grid) > 0 {
		grids = append(grids, grid)
	}

	sum := 0
	for _, g := range grids {
		rows, cols := calcRC(g)
		l, hasperfect := leftOf(rows)
		if hasperfect {
			sum += l * 100
			continue
		}

		l, hasperfect = leftOf(cols)
		if hasperfect {
			sum += l
			continue
		}

		panic("no perfect mirror")
	}

	fmt.Println(sum)
}
