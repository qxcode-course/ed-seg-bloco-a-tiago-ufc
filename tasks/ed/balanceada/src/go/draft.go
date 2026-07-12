package main

import (
	"bufio"
	"fmt"
	"os"
)

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack[T]) Pop() T {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	value := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return value
}

func (s *Stack[T]) Push(value T) {
	s.data = append(s.data, value)
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{data: []T{}}
}

func ehBalanceado(s string) bool {
	pilha := NewStack[rune]()
	for _, ch := range s {
		switch ch {
		case '(', '[':
			pilha.Push(ch)
		case ')':
			if pilha.IsEmpty() {
				return false
			}
			if pilha.Pop() != '(' {
				return false
			}
		case ']':
			if pilha.IsEmpty() {
				return false
			}
			if pilha.Pop() != '[' {
				return false
			}
		}
	}
	return pilha.IsEmpty()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	if ehBalanceado(input) {
		fmt.Println("balanceado")
	} else {
		fmt.Println("nao balanceado")
	}
}
