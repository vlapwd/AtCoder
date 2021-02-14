package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	np := 0
	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			np++
		}
	}
	mp := 0
	for i := m - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			mp++
		}
	}
	fmt.Println(np + mp)
}
