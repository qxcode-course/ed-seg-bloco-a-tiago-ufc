package main

import (
	"fmt"
	"math"
)

type Jogada struct {
	A int
	B int
}

func main() {
	var N int
	fmt.Scan(&N)
	jogadas := make([]Jogada, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&jogadas[i].A, &jogadas[i].B)
	}
	melhorIndice := -1
	melhorValor := math.MaxInt32
	for i, jog := range jogadas {
		if jog.A < 10 || jog.B < 10 {
			continue
		}
		diff := int(math.Abs(float64(jog.A - jog.B)))
		if diff < melhorValor {
			melhorValor = diff
			melhorIndice = i
		}
	}
	if melhorIndice == -1 {
		fmt.Println("sem ganhador")
	} else {
		fmt.Println(melhorIndice)
	}
}
