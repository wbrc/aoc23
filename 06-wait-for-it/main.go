package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var times, distances int
	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		if strings.HasPrefix(s.Text(), "Time:") {
			ts := ""
			for _, numStr := range strings.Split(strings.Split(s.Text(), ":")[1], " ") {
				if numStr == "" {
					continue
				}
				ts += numStr
			}
			times, _ = strconv.Atoi(ts)
		} else if strings.HasPrefix(s.Text(), "Distance:") {
			ns := ""
			for _, numStr := range strings.Split(strings.Split(s.Text(), ":")[1], " ") {
				if numStr == "" {
					continue
				}
				ns += numStr
			}
			distances, _ = strconv.Atoi(ns)
		}
	}

	fmt.Println(numWaysToWin(times, distances))
}

func numWaysToWin(timeAvail, minDist int) int {
	lower := -1
	upper := -1

	for i := 1; i <= timeAvail-1; i++ {
		if calcDistance(i, timeAvail) > minDist {
			lower = i
			break
		}
	}

	for i := timeAvail - 1; i >= 1; i-- {
		if calcDistance(i, timeAvail) > minDist {
			upper = i
			break
		}
	}

	if lower == -1 || upper == -1 {
		return 0
	}

	return upper - lower + 1
}

func calcDistance(timePressed, totalTime int) int {
	return timePressed * (totalTime - timePressed)
}
