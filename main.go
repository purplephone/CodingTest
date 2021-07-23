package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

var (
	sc      *bufio.Scanner
	wr      *bufio.Writer
	treeMap map[int]*Node
)

type Node struct {
	data  int
	child []int
}

func init() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	wr = bufio.NewWriter(os.Stdout)
}

func main() {
	defer wr.Flush()
	N := scanInt()
	treeMap = make(map[int]*Node, N+1)
	treeMap[1] = &Node{}
	for i := 1; i < N; i++ {
		p, c, w := scanInt(), scanInt(), scanInt()
		if _, empty := treeMap[c]; !empty {
			treeMap[c] = &Node{data: w}
		}
		treeMap[p].child = append(treeMap[p].child, c)
	}
	var max int
	for i := 1; i < N; i++ {
		result := dfs2(i)
		if max < result {
			max = result
		}
	}
	wr.WriteString(strconv.Itoa(max))
}

func dfs(index int) int {
	if len(treeMap[index].child) == 0 {
		return treeMap[index].data
	}
	var max int
	for _, tmp := range treeMap[index].child {
		re := dfs(tmp)
		if re > max {
			max = re
		}
	}
	return max + treeMap[index].data
}

func dfs2(index int) int {
	var tmp []int
	if len(treeMap[index].child) == 0 {
		return 0
	}
	for _, i := range treeMap[index].child {
		tmp = append(tmp, dfs(i))
	}
	if len(tmp) > 1 {
		sort.Slice(tmp, func(i, j int) bool {
			return tmp[i] > tmp[j]
		})
		return tmp[0] + tmp[1]
	} else {
		return tmp[0]
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
