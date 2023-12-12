package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type field struct {
	state uint32
	qpos  []int
}

func (f *field) set(pos int) {
	f.state |= 1 << uint(pos)
}

func (f *field) setQpos(qpos int) {
	f.qpos = append(f.qpos, qpos)
}

func (f *field) next() {
	for _, q := range f.qpos {
		f.state ^= 1 << uint(q)

		if f.state&(1<<uint(q)) != 0 {
			break
		}
	}
}

func (f *field) matchGroups(groupArr [32]int, numGroups int) bool {
	groups := groupArr[:numGroups]

	currentGroup := 0
	prev := 0
	for i := 0; i < 32; i++ {
		if f.state&(1<<uint(i)) == 0 && prev == 0 {
			continue
		}

		if f.state&(1<<uint(i)) == 0 {
			currentGroup++
			prev = 0
			continue
		}

		if currentGroup >= len(groups) {
			return false
		}

		groups[currentGroup]--

		if groups[currentGroup] < 0 {
			return false
		}

		prev = 1
	}

	for _, g := range groups {
		if g != 0 {
			return false
		}
	}

	return true
}

func (f *field) numMatches(groupArr [32]int, numGroups int) int {
	matches := 0
	numIt := 1 << uint(len(f.qpos))
	for i := 0; i < numIt; i++ {
		if f.matchGroups(groupArr, numGroups) {
			matches++
		}
		f.next()
	}
	return matches
}

func main() {

	sum := 0
	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		mask := strings.Split(s.Text(), " ")[0]
		groups := strings.Split(s.Text(), " ")[1]

		field := &field{}
		for i, r := range mask {
			switch r {
			case '?':
				field.setQpos(i)
			case '#':
				field.set(i)
			}
		}

		groupArr := [32]int{}
		numGroups := 0
		for _, g := range strings.Split(groups, ",") {
			number, _ := strconv.Atoi(g)
			groupArr[numGroups] = number
			numGroups++
		}

		sum += field.numMatches(groupArr, numGroups)
	}

	fmt.Println(sum)
}
