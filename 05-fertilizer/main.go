package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type rangeMap struct {
	sourceStart, destStart, length int
}

type mapping []rangeMap

func (m *mapping) add(destStart, sourceStart, length int) {
	*m = append(*m, rangeMap{sourceStart, destStart, length})
	sort.Slice(*m, func(i, j int) bool {
		return (*m)[i].sourceStart < (*m)[j].sourceStart
	})
}

func (m mapping) lookup(source int) int {
	for _, r := range m {
		if source >= r.sourceStart && source < r.sourceStart+r.length {
			return r.destStart + (source - r.sourceStart)
		}
	}
	return source
}

func main() {

	currentMapping := ""
	mappings := make(map[string]*mapping)
	var seeds []int

	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		if s.Text() == "" {
			continue
		}

		if strings.HasPrefix(s.Text(), "seeds:") {
			for _, numStr := range strings.Split(strings.TrimSpace(strings.Split(s.Text(), ": ")[1]), " ") {
				num, _ := strconv.Atoi(numStr)
				seeds = append(seeds, num)
			}
			continue
		}

		if s.Text()[0] < '0' || s.Text()[0] > '9' {
			currentMapping = strings.Split(s.Text(), " ")[0]
			mappings[strings.Split(s.Text(), " ")[0]] = &mapping{}
			continue
		}

		var dest, source, length int
		fmt.Sscanf(s.Text(), "%d %d %d", &dest, &source, &length)

		mappings[currentMapping].add(dest, source, length)
	}

	lookup := func(id int) int {
		id = mappings["seed-to-soil"].lookup(id)
		id = mappings["soil-to-fertilizer"].lookup(id)
		id = mappings["fertilizer-to-water"].lookup(id)
		id = mappings["water-to-light"].lookup(id)
		id = mappings["light-to-temperature"].lookup(id)
		id = mappings["temperature-to-humidity"].lookup(id)
		return mappings["humidity-to-location"].lookup(id)
	}

	lowest := math.MaxInt64
	for i := 0; i < len(seeds); i += 2 {
		fmt.Println(seeds[i], seeds[i+1])
		for j := seeds[i]; j <= seeds[i]+seeds[i+1]; j++ {
			if t := lookup(j); t < lowest {
				lowest = t
			}
		}
	}
	fmt.Println(lowest)
}
