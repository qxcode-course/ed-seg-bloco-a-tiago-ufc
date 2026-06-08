package main

import "fmt"

func main() {
	var N, M int
	fmt.Scan(&N)

	fila := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&fila[i])
	}

	fmt.Scan(&M)
	sairam := make(map[int]bool, M)
	for i := 0; i < M; i++ {
		var x int
		fmt.Scan(&x)
		sairam[x] = true
	}

	for _, id := range fila {
		if !sairam[id] {
			fmt.Printf("%d ", id)
		}
	}
	fmt.Println()
}
