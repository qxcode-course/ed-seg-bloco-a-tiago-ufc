package main

import (
	"bufio"
	"fmt"
	"os"
)

func valido(seq []rune, pos int, dig int, L int) bool {
	for i := pos - L; i < pos; i++ {
		if i >= 0 && seq[i] == rune('0'+dig) {
			return false
		}
	}
	for i := pos + 1; i <= pos+L && i < len(seq); i++ {
		if seq[i] == rune('0'+dig) {
			return false
		}
	}
	return true
}

func solve(seq []rune, L int, pos int) bool {
	if pos == len(seq) {
		return true
	}
	if seq[pos] != '.' {
		return solve(seq, L, pos+1)
	}
	for d := 0; d <= L; d++ {
		if valido(seq, pos, d, L) {
			seq[pos] = rune('0' + d)
			if solve(seq, L, pos+1) {
				return true
			}
			seq[pos] = '.'
		}
	}
	return false
}

func main() {
	var L int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	seq := []rune(scanner.Text())
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &L)
	solve(seq, L, 0)
	fmt.Println(string(seq))
}
