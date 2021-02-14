package main

import (
	"fmt"
	"math"
)

func main() {
	var n, k int64
	fmt.Scan(&n, &k)

	ret := n % k
	for {
		temp := int64(math.Abs(float64(ret - k)))
		if ret > temp {
			ret = temp
		}
		if ret-k < 0 {
			break
		}
	}
	fmt.Println(ret)
}
