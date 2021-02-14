package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	all := strings.Count(s, "R") * strings.Count(s, "G") * strings.Count(s, "B")
	//fmt.Println("all :", all)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			k := j*2 - i
			//fmt.Println("i:", i, "j:", j, "k:", k)
			if s[i] == s[j] {
				continue
			}
			if k >= n || s[k] == s[i] || s[k] == s[j] {
				continue
			}
			//fmt.Println("exclude i:", i, "j:", j, "k:", k)
			all--
		}
	}
	fmt.Println(all)
}
