package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	len := 0
	for n > 0 {
		n /= k
		len++
	}
	fmt.Println(len)
}
