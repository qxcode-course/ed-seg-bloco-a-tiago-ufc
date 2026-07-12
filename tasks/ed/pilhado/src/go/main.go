package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func inside(mat [][]rune, p Pos) bool {
	return p.l >= 0 && p.l < len(mat) && p.c >= 0 && p.c < len(mat[0])
}

func vizinhos(p Pos) []Pos {
	return []Pos{
		{p.l, p.c - 1},
		{p.l - 1, p.c},
		{p.l, p.c + 1},
		{p.l + 1, p.c},
	}
}

func main() {
	var nl, nc int
	var ini, fim Pos
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d %d", &nl, &nc)
	mat := make([][]rune, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		linha := []rune(scanner.Text())
		mat[i] = linha
		for j, ch := range linha {
			if ch == 'I' {
				ini = Pos{i, j}
			}
			if ch == 'F' {
				fim = Pos{i, j}
			}
		}
	}
	caminho := NewStack[Pos]()
	becos := NewStack[Pos]()
	visitado := make(map[Pos]bool)
	caminho.Push(ini)

	for !caminho.IsEmpty() {
		atual := caminho.Top()
		visitado[atual] = true
		if atual == fim {
			break
		}
		validos := []Pos{}
		for _, v := range vizinhos(atual) {
			if inside(mat, v) && mat[v.l][v.c] != '#' && !visitado[v] {
				validos = append(validos, v)
			}
		}
		if len(validos) > 0 {
			caminho.Push(validos[0])
		} else {
			becos.Push(atual)
			caminho.Pop()
		}
	}
	for !becos.IsEmpty() {
		b := becos.Pop()
		if mat[b.l][b.c] == '.' {
			mat[b.l][b.c] = ' '
		}
	}
	for _, p := range caminho.data {
		if mat[p.l][p.c] == ' ' || mat[p.l][p.c] == 'I' || mat[p.l][p.c] == 'F' {
			mat[p.l][p.c] = '.'
		}
	}
	for _, linha := range mat {
		fmt.Println(string(linha))
	}
}
