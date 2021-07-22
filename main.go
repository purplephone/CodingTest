package main

import (
	"bufio"
	"os"
)

var (
	sc      *bufio.Scanner
	wr      *bufio.Writer
	visited []bool
	result  []int
	treeMap map[int]*Node
)

type Node struct {
	data  int
	child *Node
}

func init() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	wr = bufio.NewWriter(os.Stdout)
}

func main() {
	defer wr.Flush()
	N := scanInt()
	for i := 1; i < N; i++ {
		p, c, d := scanInt(), scanInt(), scanInt()
		if _, empty := treeMap[p]; !empty {
			treeMap[p] = &Node{}
		}
		treeMap[p].child = append()
	}
}

func dfs(index int) {
	visited[index] = true
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
