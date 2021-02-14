package main

import "fmt"

func main() {
	var n, w int
	fmt.Scan(&n, &w)
	result := "unsafe"
	if n > w {
		result = "safe"
	}
	fmt.Println(result)
}
