package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n, _ = strconv.Atoi(sc.Text())

	set := make(map[string]struct{})
	for i := 0; i < n; i++ {
		sc.Scan()
		set[sc.Text()] = struct{}{}
	}

	fmt.Println(len(set))
}
