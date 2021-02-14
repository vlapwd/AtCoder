package main

import (
	"fmt"
)

func main() {
	var n, a int
	fmt.Scan(&n)
	ret := make([]int, n)
	for i := 0; i < n-1; i++ {
		fmt.Scan(&a)
		ret[i] = a
	}
	result := make([]int, n+1)
	for _, i := range ret {
		result[i]++
	}
	for i := 1; i < len(result); i++ {
		fmt.Println(result[i])
	}
}
