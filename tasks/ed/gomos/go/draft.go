package main

import "fmt"

func main() {
	var Q int
	var D byte

	fmt.Scanf("%d %c", &Q, &D)

	x := make([]int, Q)
	y := make([]int, Q)

	for i := 0; i < Q; i++ {
		fmt.Scan(&x[i], &y[i])
	}

	pX, pY := x[0], y[0]

	switch D {
	case 'L':
		x[0]--
	case 'R':
		x[0]++
	case 'U':
		y[0]--
	case 'D':
		y[0]++
	}

	for i := 1; i < Q; i++ {
		tempX, tempY := x[i], y[i]
		x[i], y[i] = pX, pY
		pX, pY = tempX, tempY
	}

	for i := 0; i < Q; i++ {
		fmt.Println(x[i], y[i])
	}
}
