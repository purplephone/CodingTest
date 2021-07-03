package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var sc *bufio.Scanner
var wr *bufio.Writer
var N, min, start int
var buf []int
var visited [14]bool

func init() {
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
}

func main() {
	defer wr.Flush()
	for {
		nextInts()
		if buf[0] == 0 {
			break
		}
		dfs(0, 0)
		wr.WriteByte('\n')
	}
}

func dfs(depth, index int) {
	if depth == 6 {
		for i := 1; i < len(buf); i++ {
			if visited[i] == true {
				wr.WriteString(strconv.Itoa(buf[i]) + " ")
			}
		}
		wr.WriteByte('\n')
		return
	}
	for i := index + 1; i < len(buf); i++ {
		if visited[i] != true {
			visited[i] = true
			dfs(depth+1, i)
			visited[i] = false
		}
	}
}

func nextInts() {
	sc.Scan()
	strs := strings.Fields(sc.Text())
	buf = make([]int, len(strs))
	for i, str := range strs {
		buf[i], _ = strconv.Atoi(str)
	}
}

func scanInt() (n int) {
	sc.Scan()
	n, _ = strconv.Atoi(sc.Text())
	return
}
