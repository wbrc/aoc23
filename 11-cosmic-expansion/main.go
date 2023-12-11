package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := [][]bool{}
	emptyRows := []int{}
	fullCols := map[int]bool{}

	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		line := []bool{}
		hasGalaxy := false

		for col, c := range s.Text() {
			line = append(line, c == '#')
			if c == '#' {
				hasGalaxy = true
				fullCols[col] = true
			}
		}
		input = append(input, line)

		if !hasGalaxy {
			emptyRows = append(emptyRows, len(input)-1)
		}
	}

	numCols := len(input[0])
	emptyCols := []int{}
	for col := 0; col < numCols; col++ {
		if !fullCols[col] {
			emptyCols = append(emptyCols, col)
		}
	}

	type pos struct {
		r, c int
	}

	galaxies := []pos{}

	for r := range input {
		for c, col := range input[r] {
			if col {
				galaxies = append(galaxies, pos{r, c})
			}
		}
	}

	totalDistance := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {

			colDist := abs(galaxies[i].c - galaxies[j].c)
			emptyColsBetween := emptyBetween(galaxies[i].c, galaxies[j].c, emptyCols)
			colDist -= emptyColsBetween
			colDist += emptyColsBetween * 1000000

			rowDist := abs(galaxies[i].r - galaxies[j].r)
			emptyRowsBetween := emptyBetween(galaxies[i].r, galaxies[j].r, emptyRows)
			rowDist -= emptyRowsBetween
			rowDist += emptyRowsBetween * 1000000

			totalDistance += colDist + rowDist
		}
	}

	fmt.Println(totalDistance)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func emptyBetween(lower, upper int, empties []int) int {
	if lower > upper {
		lower, upper = upper, lower
	}

	count := 0
	for _, col := range empties {
		if col > lower && col < upper {
			count++
		}
	}
	return count
}
