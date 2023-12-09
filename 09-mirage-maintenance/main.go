package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func extrapolate(seq []int) int {
	var subSeq []int
	numZ := 0
	for i := 0; i < len(seq)-1; i++ {
		diff := seq[i+1] - seq[i]
		subSeq = append(subSeq, diff)
		if diff == 0 {
			numZ++
		}
	}

	if numZ == len(subSeq) {
		return seq[len(seq)-1]
	}

	return seq[len(seq)-1] + extrapolate(subSeq)
}

func main() {
	sum := 0
	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		var seq []int
		for _, nStr := range strings.Split(s.Text(), " ") {
			n, _ := strconv.Atoi(nStr)
			seq = append(seq, n)
		}

		for i, j := 0, len(seq)-1; i < j; i, j = i+1, j-1 {
			seq[i], seq[j] = seq[j], seq[i]
		}

		sum += extrapolate(seq)
	}
	fmt.Println(sum)
}
