package main

import "fmt"

func main() {
	var s int
	fmt.Scan(&s)
	fmt.Println((s / 500 * 1000) + ((s % 500) / 5 * 5))
}
