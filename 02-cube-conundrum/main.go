package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type set struct {
	r, g, b int
}

func parseSet(line string) set {
	var s set
	tokens := strings.Split(line, ",")
	for _, t := range tokens {
		t = strings.TrimSpace(t)
		var count int
		var color string

		fmt.Sscanf(t, "%d %s", &count, &color)
		switch color {
		case "red":
			s.r = count
		case "green":
			s.g = count
		case "blue":
			s.b = count
		}
	}

	return s
}

func parseGame(line string) (int, []set) {
	idStr := strings.Split(line, ":")[0]
	var id int
	fmt.Sscanf(idStr, "Game %d", &id)

	var parsedSets []set
	sets := strings.TrimSpace(strings.Split(line, ":")[1])
	for _, setStr := range strings.Split(sets, "; ") {
		parsedSets = append(parsedSets, parseSet(setStr))
	}

	return id, parsedSets
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func max(a, b set) set {
	return set{
		r: maxInt(a.r, b.r),
		g: maxInt(a.g, b.g),
		b: maxInt(a.b, b.b),
	}
}

func main() {
	sum := 0
	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		line := s.Text()
		_, sets := parseGame(line)
		var min set
		for _, s := range sets {
			min = max(min, s)
		}
		power := min.r * min.g * min.b
		sum += power
	}
	fmt.Println(sum)
}
