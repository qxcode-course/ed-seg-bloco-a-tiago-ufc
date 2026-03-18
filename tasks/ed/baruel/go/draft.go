package main

import "fmt"

func main() {
	qtd_album := 0
	qtd_fig := 0
	fmt.Scan(&qtd_album, &qtd_fig)
	album := make([]int, qtd_fig)
	unicos := make(map[int]bool)
	repetidos := make([]int, 0, qtd_fig)
	for i := range album {
		fmt.Scan(&album[i])
	}
	for _, fig := range album {
		if unicos[fig] {
			repetidos = append(repetidos, fig)
		} else {
			unicos[fig] = true
		}
	}
	for i, valor := range repetidos {
		if i != 0 {
			fmt.Print(" ")
		} else {
			fmt.Printf("%v", valor)
		}
	}
	fmt.Println()
	saida := ""
	for i := 1; i < qtd_album; i++ {
		if !unicos[i] {
			saida += fmt.Sprintf("%v ", i)
		}
	}
	if saida == "" {
		fmt.Println("N")
	} else {
		fmt.Println(saida[0 : len(saida)-1])
	}
	println(unicos)
	println(repetidos)
}
