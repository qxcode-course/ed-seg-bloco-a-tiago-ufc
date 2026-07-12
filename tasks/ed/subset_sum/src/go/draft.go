package main

import (
	"fmt"
)

func somaSub(a []int, k int) bool {
	var bcktr func(i, soma int) bool
	bcktr = func(i, soma int) bool {
		if soma == k {
			return true
		}
		if i == len(a) {
			return false
		}
		if bcktr(i+1, soma+a[i]) {
			return true
		}
		if bcktr(i+1, soma) {
			return true
		}
		return false
	}
	return bcktr(0, 0)
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	if somaSub(a, k) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
