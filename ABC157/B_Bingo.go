package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	ret := make([]int, 9)
	for i := 0; i < 9; i++ {
		ret[i] = nextInt()
	}
	n := nextInt()
	game := make([]int, n)
	for i := 0; i < n; i++ {
		game[i] = nextInt()
	}

	res := make([]bool, 9)
	for i, r := range ret {
		for _, g := range game {
			if r == g {
				res[i] = true
			}
		}
	}
  fmt.Println(bingo(res))
}

func next() string {
	sc.Scan()
	return sc.Text()
}
func nextInt() int {
	a, _ := strconv.Atoi(next())
	return a
}

func bingo(res []bool) string {
	b := "No"
	if res[0] && res[1] && res[2] ||
	res[3] && res[4] && res[5] ||
	res[6] && res[7] && res[8] ||
	res[0] && res[3] && res[6] ||
	res[1] && res[4] && res[7] ||
	res[2] && res[5] && res[8] ||
	res[0] && res[4] && res[8] ||
	res[2] && res[4] && res[6] {
		b = "Yes"
	}
	return b
}
