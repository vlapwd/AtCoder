package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	ns := nextInts(n)
	sort.Sort(sort.IntSlice(ns))
	l, r := ns[0], ns[n-1]
	min := 1000000000
	for i := l; i <= r; i++ {
		var len int
		for _, m := range ns {
			tmp := m - i
			len += tmp * tmp
		}
		if min > len {
			min = len
		}
	}
	fmt.Println(min)
}

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	a, _ := strconv.Atoi(next())
	return a
}
func nextInts(n int) []int {
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		ret[i] = nextInt()
	}
	return ret
}
