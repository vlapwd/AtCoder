package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	m := map[int]bool{}
	for i := 0; i < k; i++ {
		var d int
		fmt.Scan(&d)
		for j := 0; j < d; j++ {
			var a int
			fmt.Scan(&a)
			m[a] = true
		}
	}
	fmt.Println(n - len(m))
}
