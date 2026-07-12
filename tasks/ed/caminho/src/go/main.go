package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func (p Pos) getNeig() []Pos {
	return nil
}

func inside(grid [][]rune, pos Pos) bool {
	nrows := len(grid)
	ncols := len(grid[0])
	return pos.l >= 0 && pos.l < nrows && pos.c >= 0 && pos.c < ncols
}

func match(grid [][]rune, pos Pos, char rune) bool {
	return inside(grid, pos) && grid[pos.l][pos.c] == char
}

func vizinhos(p Pos) []Pos {
	return []Pos{
		{p.l, p.c - 1},
		{p.l - 1, p.c},
		{p.l, p.c + 1},
		{p.l + 1, p.c},
	}
}

func search(grid [][]rune, startPos Pos, endPos Pos) {
	fila := NewQueue[Pos]()
	fila.Enqueue(startPos)
	visitado := make(map[Pos]bool)
	pai := make(map[Pos]Pos)
	achou := false
	for !fila.IsEmpty() {
		atual, _ := fila.Dequeue()
		if atual == endPos {
			achou = true
			break
		}
		for _, v := range vizinhos(atual) {
			if inside(grid, v) && grid[v.l][v.c] != '#' && !visitado[v] {
				visitado[v] = true
				pai[v] = atual
				fila.Enqueue(v)
			}
		}
	}
	if achou {
		voltar(grid, pai, startPos, endPos)
	}
}

func voltar(grid [][]rune, pai map[Pos]Pos, ini Pos, fim Pos) {
	atual := fim
	for atual != ini {
		grid[atual.l][atual.c] = '.'
		atual = pai[atual]
	}
	grid[ini.l][ini.c] = '.'
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var nl, nc int
	scanner.Scan()
	line := scanner.Text()
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	mat := make([][]rune, nl) // Inicializa a matriz de runes

	// Carregando matriz
	for i := 0; i < nl; i++ {
		scanner.Scan()
		line := scanner.Text()
		mat[i] = []rune(line)
	}

	var inicio, fim Pos

	// Procurando inicio e fim e colocando ' ' nas posições iniciais
	for l := range nl {
		for c := 0; c < nc; c++ {
			if mat[l][c] == 'I' {
				mat[l][c] = ' '
				inicio = Pos{l, c}
			}
			if mat[l][c] == 'F' {
				mat[l][c] = ' '
				fim = Pos{l, c}
			}
		}
	}

	search(mat, inicio, fim)

	for _, line := range mat {
		fmt.Println(string(line)) // Converte o slice de runes de volta para string para imprimir
	}
}
