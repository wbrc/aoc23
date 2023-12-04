package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sum := 0
	nextID := 0
	open := make(map[int]int)
	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		currentID := nextID
		nextID++

		open[currentID]++

		nums := strings.TrimSpace(strings.Split(s.Text(), ":")[1])
		winStr := strings.TrimSpace(strings.Split(nums, "|")[0])
		betStr := strings.TrimSpace(strings.Split(nums, "|")[1])

		wins := make(map[int]struct{})
		for _, v := range strings.Split(winStr, " ") {
			n, _ := strconv.Atoi(v)
			wins[n] = struct{}{}
		}

		score := 0
		have := make(map[int]struct{})
		for _, v := range strings.Split(betStr, " ") {
			if v == "" {
				continue
			}
			n, _ := strconv.Atoi(v)

			if _, ok := have[n]; ok {
				continue
			}

			have[n] = struct{}{}
			if _, ok := wins[n]; ok {
				score++
			}
		}

		for i := 0; i < score; i++ {
			for j := 0; j < open[currentID]; j++ {
				open[currentID+1+i]++
			}
		}
	}

	for _, v := range open {
		sum += v
	}

	fmt.Println(sum)
}
