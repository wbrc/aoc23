package main

import (
	"bufio"
	"fmt"
	"os"
)

func euclid(a, b int) int {
	if a == 0 {
		return b
	}

	for b != 0 {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	return a
}

func lcm(nums []int) int {
	lcm := nums[0]
	for i := 1; i < len(nums); i++ {
		lcm = lcm * nums[i] / euclid(lcm, nums[i])
	}
	return lcm
}

type node struct {
	left, right string
}

func main() {
	network := map[string]node{}
	instructions := []bool{}
	aEnder := []string{}
	iterations := []int{}

	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	iStr := s.Text()
	for _, c := range iStr {
		instructions = append(instructions, c == 'L')
	}

	for s.Scan() {
		if s.Text() == "" {
			continue
		}

		var current, left, right string
		fmt.Sscanf(s.Text(), "%3s = (%3s, %3s)", &current, &left, &right)

		network[current] = node{left, right}

		if current[2] == 'A' {
			aEnder = append(aEnder, current)
		}
	}

	for i := range aEnder {
		iPtr := 0
		steps := 0
		for {
			steps++
			if instructions[iPtr] {
				aEnder[i] = network[aEnder[i]].left
			} else {
				aEnder[i] = network[aEnder[i]].right
			}
			if aEnder[i][2] == 'Z' {
				break
			}
			iPtr = (iPtr + 1) % len(instructions)
		}
		iterations = append(iterations, steps)
	}
	fmt.Println(lcm(iterations))
}
