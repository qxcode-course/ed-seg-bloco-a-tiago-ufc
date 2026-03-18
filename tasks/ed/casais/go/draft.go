package main

import "fmt"

func main() {
	solteiros := make(map[int]int)
	qtd := 0
	pares := 0
	fmt.Scan(&qtd)
	for range qtd {
		animal := 0
		fmt.Scan(&animal)
		qtd, existe := solteiros[-animal]
		if existe && qtd > 0 {
			solteiros[-animal] = qtd - 1
			pares += 1
		} else {
			solteiros[animal]++
		}
	}
	fmt.Println(pares)
}
