package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type play struct {
	hand     [5]byte
	strength int
	bid      int
}

func str(hand [5]byte) int {
	var m [13]int
	numPairs := 0
	numTrips := 0
	numJoker := 0
	numVar := 0

	for _, c := range hand {
		if c == 0 {
			numJoker++
			continue
		}

		if m[c] == 0 {
			numVar++
		}

		m[c]++

		if m[c] == 2 {
			numPairs++
		} else if m[c] == 3 {
			numPairs--
			numTrips++
		} else if m[c] == 4 {
			numTrips--
		}

	}

	switch numVar {
	case 5:
		return 0
	case 4:
		return 1
	case 3:
		if numJoker == 0 && numPairs == 2 {
			return 2
		}
		return 3
	case 2:
		if numJoker == 0 && numTrips == 1 {
			return 4
		}
		if numJoker == 0 {
			return 5
		}
		if numJoker == 1 && numPairs == 2 {
			return 4
		}
		if numJoker == 1 && numTrips == 1 {
			return 5
		}
		if numJoker >= 2 {
			return 5
		}
	}
	return 6
}

func less(a, b [5]byte) bool {
	for i := 0; i < 5; i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return false
}

func main() {
	var hands []play
	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		h := strings.Split(s.Text(), " ")[0]
		bid, _ := strconv.Atoi(strings.Split(s.Text(), " ")[1])

		var hand [5]byte
		for i, c := range h {
			switch c {
			case 'J':
				hand[i] = 0
			case 'A':
				hand[i] = 12
			case 'K':
				hand[i] = 11
			case 'Q':
				hand[i] = 10
			case 'T':
				hand[i] = 9
			default:
				hand[i] = byte(c) - '1'
			}
		}

		hands = append(hands, play{hand, str(hand), bid})
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].strength != hands[j].strength {
			return hands[i].strength < hands[j].strength
		}
		return less(hands[i].hand, hands[j].hand)
	})

	sum := 0
	for i, h := range hands {
		sum += h.bid * (i + 1)
	}

	fmt.Println(sum)
}
