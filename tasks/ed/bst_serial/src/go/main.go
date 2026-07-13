package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func insert(root *Node, value int) *Node {
	if root == nil {
		return &Node{Value: value}
	}
	if value < root.Value {
		root.Left = insert(root.Left, value)
	} else if value > root.Value {
		root.Right = insert(root.Right, value)
	}
	return root
}

func BstInsert(values []int) *Node {
	var root *Node
	for _, v := range values {
		root = insert(root, v)
	}
	return root
}

// Dica: crie um vetor compartilhado e vá preenchendo conforme anda na recursão
// Depois use o strings.Join para gerar o serial
func Serialize(root *Node) string {
	out := []string{}
	var dfs func(n *Node)
	dfs = func(n *Node) {
		if n == nil {
			out = append(out, "#")
			return
		}
		out = append(out, fmt.Sprintf("%d", n.Value))
		dfs(n.Left)
		dfs(n.Right)
	}
	dfs(root)
	return strings.Join(out, " ")
}

// -----------------------------------------------------------------------------------
func BShow(node *Node, history string) {
	if node != nil && (node.Left != nil || node.Right != nil) {
		BShow(node.Left, history+"l")
	}
	for i := 0; i < len(history)-1; i++ {
		if history[i] != history[i+1] {
			fmt.Print("│   ")
		} else {
			fmt.Print("    ")
		}
	}
	if history != "" {
		if history[len(history)-1] == 'l' {
			fmt.Print("╭───")
		} else {
			fmt.Print("╰───")
		}
	}
	if node == nil {
		fmt.Println("#")
		return
	}
	fmt.Println(node.Value)
	if node.Left != nil || node.Right != nil {
		BShow(node.Right, history+"r")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	values := make([]int, 0, len(parts))
	for _, elem := range parts {
		v, err := strconv.Atoi(elem)
		if err == nil {
			values = append(values, v)
		}
	}
	root := BstInsert(values)
	BShow(root, "") // Chama a função de impressão formatada
	fmt.Println(Serialize((root)))
}
