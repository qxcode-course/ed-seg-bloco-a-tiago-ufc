package main

import "fmt"

func prox_vivo(elementos []int, pos int) int {
	size := len(elementos)
	for {
		pos = (pos + 1) % size
		if elementos[pos] != 0 {
			return pos
		}
	}
}

func imprimir(elementos []int, pos int, size int) {
	fmt.Printf("[ ")
	for i := 0; size > i ; i++ {
		fmt.Printf("%d", elementos[i])
		if pos == i {
			fmt.Printf(">")
		}
		if pos != 0 || pos != size {
			fmt.Printf(" ")
		}
	}
	fmt.Println("]")
}

func executar(elementos []int, pos int) int {
	vivos := len(elementos)	
	for vivos > 1 {
		imprimir(elementos, pos-1, vivos)
		alvo := prox_vivo(elementos, pos)
		elementos[alvo] = 0
		vivos--
		pos = prox_vivo(elementos, alvo)
	}
	return pos
}

func main() {
	var size, pos int
	fmt.Scanf("%d", &size)
	fmt.Scanf("%d", &pos)
	elementos := make([]int, size)
	for i := 0; i < size; i++ {
		elementos[i] = i + 1
	}
	executar(elementos, pos)
	//fmt.Println(sobrevivente)
}
