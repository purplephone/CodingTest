package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	sc        *bufio.Scanner
	wr        *bufio.Writer
	N, M, max int
	visited   []bool
	dx, dy    [4]int
	v         []virus
)

type virus struct {
	y, x int
}

func init() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	wr = bufio.NewWriter(os.Stdout)
	dx = [4]int{0, 0, 1, -1}
	dy = [4]int{1, -1, 0, 0}
}

func main() {
	defer wr.Flush()

}

func scanInt() (n int) {
	sc.Scan()
	n, _ = strconv.Atoi(sc.Text())
	return
}
