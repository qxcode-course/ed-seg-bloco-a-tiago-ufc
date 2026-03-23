package main

import "fmt"

func main() {
	var qtdAlbum, qtdFig int
	fmt.Scan(&qtdAlbum)
	fmt.Scan(&qtdFig)

	figs := make([]int, qtdFig)
	unicos := make(map[int]int) // conta quantas vezes aparece

	repetidos := []int{}
	for i := 0; i < qtdFig; i++ {
		fmt.Scan(&figs[i])
		unicos[figs[i]]++
	}

	// Repetidos
	for fig, count := range unicos {
		if count > 1 {
			for i := 1; i < count; i++ {
				repetidos = append(repetidos, fig)
			}
		}
	}
	if len(repetidos) == 0 {
		fmt.Println("N")
	} else {
		for i, v := range repetidos {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(v)
		}
		fmt.Println()
	}

	// Faltantes
	faltando := []int{}
	for i := 1; i <= qtdAlbum; i++ {
		if unicos[i] == 0 {
			faltando = append(faltando, i)
		}
	}
	if len(faltando) == 0 {
		fmt.Println("N")
	} else {
		for i, v := range faltando {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(v)
		}
		fmt.Println()
	}
}
