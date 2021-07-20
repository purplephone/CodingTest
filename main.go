package main

import (
	"bufio"
	"os"
)

var (
	sc *bufio.Scanner
	wr *bufio.Writer
)

type Node struct {
	data []int
}

func init() {
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
}

func main() {
	defer wr.Flush()
	N := scanInt()
	treeMap := make(map[int]*Node)
	for i := 0; i < N; i++ {
		s := scanInt()
		if _, empty := treeMap[s]; !empty {

		}
	}
}

func scanInt() (n int) {
	sc.Scan()
	n = 0
	for _, i := range sc.Bytes() {
		n *= 10
		n += int(i - '0')
	}
	return
}
