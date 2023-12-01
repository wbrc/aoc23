package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	sum := 0
	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		line := s.Text()

		loopnum := 0

		numIdx, numStr := firstNumString(line)
		charIdx, charStr := firstNumChar(line)

		if numIdx < charIdx {
			loopnum = numStr
		} else {
			loopnum, _ = strconv.Atoi(charStr)
		}

		numIdx, numStr = lastNumString(line)
		charIdx, charStr = lastNumChar(line)

		if numIdx > charIdx {
			loopnum *= 10
			loopnum += numStr
		} else {
			loopnum *= 10
			tmp, _ := strconv.Atoi(charStr)
			loopnum += tmp
		}

		fmt.Println(loopnum)

		sum += loopnum
	}
	fmt.Println(sum)
}

var numStrings = []string{
	"zero", "one", "two", "three", "four",
	"five", "six", "seven", "eight", "nine",
}

func firstNumString(s string) (int, int) {
	min := math.MaxInt
	val := 0
	for pos, ns := range numStrings {
		i := strings.Index(s, ns)
		if i >= 0 && i < min {
			min = i
			val = pos
		}
	}

	return min, val
}

func lastNumString(s string) (int, int) {
	max := -1
	val := 0
	for pos, ns := range numStrings {
		i := strings.LastIndex(s, ns)
		if i >= 0 && i > max {
			max = i
			val = pos
		}
	}

	if max < 0 {
		return -1, 0
	}

	return max + len(numStrings[val]), val
}

func firstNumChar(s string) (int, string) {
	for i, c := range []byte(s) {
		if c >= '0' && c <= '9' {
			return i, string(c)
		}
	}
	return -1, ""
}

func lastNumChar(s string) (int, string) {
	for i := len(s) - 1; i >= 0; i-- {
		c := s[i]
		if c >= '0' && c <= '9' {
			return i, string(c)
		}
	}
	return -1, ""
}
