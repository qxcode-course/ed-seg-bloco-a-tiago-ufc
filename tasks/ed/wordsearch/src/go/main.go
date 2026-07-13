package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(grid [][]rune, word string) bool {
	nRows := len(grid)
	nCols := len(grid[0])
	for r := 0; r < nRows; r++ {
		for c := 0; c < nCols; c++ {
			if backtrack(grid, r, c, word) {
				return true
			}
		}
	}
	return false
}

func backtrack(grid [][]rune, row, col int, suffix string) bool {
	if len(suffix) == 0 {
		return true
	}
	if row < 0 || row >= len(grid) ||
		col < 0 || col >= len(grid[0]) ||
		grid[row][col] != rune(suffix[0]) {
		return false
	}
	temp := grid[row][col]
	grid[row][col] = '#'
	dirs := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	for _, d := range dirs {
		if backtrack(grid, row+d[0], col+d[1], suffix[1:]) {
			return true
		}
	}
	grid[row][col] = temp
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]rune, 0)
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
