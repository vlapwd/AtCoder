package main

import (
	"fmt"
	"strings"
)

func main() {
	var n string
	fmt.Scan(&n)
	result := "No"
	if strings.Contains(n, "7") {
		result = "Yes"
	}
	fmt.Println(result)
}
