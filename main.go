package main

import (
	"bufio"
	"os"
	"strconv"
)

var sc *bufio.Scanner
var wr *bufio.Writer
var N, M int
var visited [1001][1001]bool
var arr [1001][1001]int
var sum int

func init() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	wr = bufio.NewWriter(os.Stdout)
}

func main() {
	defer wr.Flush()
	N = scanInt()
	M = scanInt()
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			arr[i][j] = scanInt()
		}
	}
	dfs(1)
	wr.WriteString(strconv.Itoa(sum))
}

func dfs(index int) {

}

func scanInt() (n int) {
	sc.Scan()
	n, _ = strconv.Atoi(sc.Text())
	return
}
