package main

import (
	"bufio"
	"container/list"
	"os"
	"strconv"
)

var sc *bufio.Scanner
var wr *bufio.Writer
var N, M, V int
var visited [1001]bool
var graph [1001][1001]bool

func init() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	wr = bufio.NewWriter(os.Stdout)
}

func main() {
	defer wr.Flush()
	N = scanInt()
	M = scanInt()
	V = scanInt()

	for i := 0; i < M; i++ {
		a := scanInt()
		b := scanInt()
		graph[a][b] = true
		graph[b][a] = true
	}
	dfs(V)
	wr.WriteByte('\n')
	bfs(V)
}

func dfs(v int) {
	wr.WriteString(strconv.Itoa(v) + " ")
	visited[v] = true
	for i := 1; i < N+1; i++ {
		if visited[i] != true && graph[v][i] {
			dfs(i)
		}
	}
}

func bfs(v int) {
	Q := list.New()
	Q.PushBack(v)
	visited[v] = false
	for Q.Len() != 0 {
		tmp := Q.Front()
		Q.Remove(tmp)
		wr.WriteString(strconv.Itoa(tmp.Value.(int)) + " ")
		for i := 1; i < N+1; i++ {
			if visited[i] && graph[tmp.Value.(int)][i] {
				visited[i] = false
				Q.PushBack(i)
			}
		}
	}
}

func scanInt() (n int) {
	sc.Scan()
	n, _ = strconv.Atoi(sc.Text())
	return
}
