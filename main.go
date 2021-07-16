package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	sc           *bufio.Scanner
	wr           *bufio.Writer
	N, M, min    int
	arr, visited [1000][1000]bool
	dx, dy       [4]int
)

func init() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	wr = bufio.NewWriter(os.Stdout)
	dx = [4]int{0, 1, 0, -1}
	dy = [4]int{1, 0, -1, 0}
}

func main() {
	defer wr.Flush()
	N = scanInt()
	M = scanInt()
	min = N * M
	for i := 0; i < N; i++ {
		sc.Scan()
		for j, x := range sc.Bytes() {
			if int(x-'0') == 0 {
				arr[i][j] = false
			} else {
				arr[i][j] = true
			}
		}
	}
	dfs(0, 0, 1, false)
	if min == N*M {
		wr.WriteString("-1")
	} else {
		wr.WriteString(strconv.Itoa(min))
	}
}

func dfs(y, x, cnt int, brk bool) {
	if y == N-1 && x == M-1 {
		if cnt < min {
			min = cnt
		}
		return
	}
	for i := 0; i < 4; i++ {
		newY := y + dy[i]
		newX := x + dx[i]
		if check(newY, newX) && !visited[newY][newX] {
			if !arr[newY][newX] {
				visited[newY][newX] = true
				dfs(newY, newX, cnt+1, brk)
				visited[newY][newX] = false
			}
			if arr[newY][newX] && !brk {
				dfs(newY, newX, cnt+1, true)
			}
		}
	}
}

func check(i, j int) bool {
	if i >= 0 && i < N && j >= 0 && j < M {
		return true
	} else {
		return false
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
