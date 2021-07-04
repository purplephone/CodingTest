package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	sc            *bufio.Scanner
	wr            *bufio.Writer
	N, max, water int
	arr           [100][100]int
	visited       [100][100]bool
	dx, dy        [4]int
)

func init() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	wr = bufio.NewWriter(os.Stdout)
	dx = [4]int{0, 0, 1, -1}
	dy = [4]int{1, -1, 0, 0}
}

func main() {
	defer wr.Flush()
	N = scanInt()
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			arr[i][j] = scanInt()
			max = Max(max, arr[i][j])
		}
	}
	var result, cnt int
	for water = 2; water < max; water++ {
		cnt = 0
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				if arr[i][j] > water {
					visited[i][j] = true
					dfs(i, j)
					cnt++
				}
			}
		}
		clear()
		result = Max(result, cnt)
	}
	wr.WriteString(strconv.Itoa(result))
}

func dfs(x, y int) {
	for i := 0; i < 4; i++ {
		newX := x + dx[i]
		newY := y + dy[i]
		if newX >= 0 && newX < N && newY >= 0 && newY < N && arr[newX][newY] > water && !visited[newX][newY] {
			visited[newX][newY] = true
			dfs(newX, newY)
		}
	}
}
func clear() {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			visited[i][j] = false
		}
	}
}
func Max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func scanInt() (n int) {
	sc.Scan()
	n, _ = strconv.Atoi(sc.Text())
	return
}
