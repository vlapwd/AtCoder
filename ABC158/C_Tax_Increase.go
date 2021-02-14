package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	result := -1
	for i := 1000; i > 0; i-- {
		e := int(float64(i) * 0.08)
		t := int(float64(i) * 0.1)
		if e == a && t == b {
			result = i
		}
	}
	fmt.Println(result)
}
