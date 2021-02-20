package main

import "fmt"

func main() {
	var k, a, b int
	fmt.Scan(&k, &a, &b)
	result := "NG"

	for i := a; i <= b; i++ {
		if i%k == 0 {
			result = "OK"
		}
	}
	fmt.Println(result)
}
