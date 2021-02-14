package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	taka := c / b
	ao := a / d
	if c%b != 0 {
		taka++
	}
	if a%d != 0 {
		ao++
	}
	result := "No"
	if taka <= ao {
		result = "Yes"
	}
	fmt.Println(result)
}
