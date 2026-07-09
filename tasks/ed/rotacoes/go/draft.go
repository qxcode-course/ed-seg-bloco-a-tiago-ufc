package main

import "fmt"

func printFormat(v []int) {
	fmt.Print("[")
	for _, x := range v {
		fmt.Printf(" %d", x)
	}
	fmt.Println(" ]")
}

func main() {
	var t, r int
	fmt.Scan(&t)
	fmt.Scan(&r)

	vetor := make([]int, t)
	for i := 0; i < t; i++ {
		fmt.Scan(&vetor[i])
	}

	r = r % t
	rotated := append(vetor[t-r:], vetor[:t-r]...)
	printFormat(rotated)
}
