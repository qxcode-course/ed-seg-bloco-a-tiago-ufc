package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	value int
	root  *Node
	prev  *Node
	next  *Node
}

type LList struct {
	root Node
	Size int
}

func (n *Node) Next() *Node {
	if n.next == n.root {
		return nil
	}
	return n.next
}

func (n *Node) Prev() *Node {
	if n.prev == n.root {
		return nil
	}
	return n.prev
}

func (l *LList) String() string {
	if l.Size == 0 {
		return "[]"
	}
	str := "["
	for n := l.root.next; n != &l.root; n = n.next {
		str += fmt.Sprintf("%d", n.value)
		if n.next != &l.root {
			str += ", "
		}
	}
	str += "]"
	return str
}

func NewLList() *LList {
	l := &LList{}
	l.root.next = &l.root
	l.root.prev = &l.root
	return l
}

func (l *LList) PushBack(value int) {
	n := l.newNode(value)
	last := l.root.prev
	last.next = n
	n.prev = last
	n.next = &l.root
	l.root.prev = n
	l.Size++
}

func (l *LList) PushFront(value int) {
	n := l.newNode(value)
	first := l.root.next
	n.next = first
	n.prev = &l.root
	first.prev = n
	l.root.next = n
	l.Size++
}

func (l *LList) PopBack() {
	if l.Size == 0 {
		return
	}
	last := l.root.prev
	last_prev := last.prev
	last_prev.next = &l.root
	l.root.prev = last_prev
	l.Size--
	last.prev = nil
	last.next = nil
}

func (l *LList) PopFront() {
	if l.Size == 0 {
		return
	}
	first := l.root.next
	first_next := first.next
	first_next.prev = &l.root
	l.root.next = first_next
	l.Size--
	first.prev = nil
	first.next = nil
}

func (l *LList) Clear() {
	if l.Size == 0 {
		return
	}
	for n := l.root.next; n != &l.root; {
		next := n.next
		n.prev = nil
		n.next = nil
		n = next
	}
	l.root.next = &l.root
	l.root.prev = &l.root
	l.Size = 0
}

func (l *LList) Front() *Node {
	if l.Size == 0 {
		return nil
	}
	return l.root.next
}

func (l *LList) Back() *Node {
	if l.Size == 0 {
		return nil
	}
	return l.root.prev
}

func (l *LList) newNode(value int) *Node {
	return &Node{
		value: value,
		root:  &l.root,
	}
}

func (l *LList) Search(value int) *Node {
	for n := l.Front(); n != nil; n = n.Next() {
		if n.value == value {
			return n
		}
	}
	return nil
}

func (l *LList) Insert(node *Node, value int) {
	n := l.newNode(value)
	prev := node.prev
	prev.next = n
	n.prev = prev
	n.next = node
	node.prev = n
	l.Size++
}

func (l *LList) Remove(node *Node) {
	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev
	node.next = nil
	node.prev = nil
	l.Size--
}

// class LList {
//     root * Node                   // Nó sentinela que marca o começo e o fim da lista
//     size int                      // tamanho da lista
//     Size()                        // retorna o tamanho da lista
//     Clear()                       // apaga todos os nós da lista
//     PushFront(value int)          // adiciona um novo nó com esse valor no início da lista
//     PushBack(value int)           // adiciona um novo nó com esse valor no fim da lista
//     PopFront()                    // remove o primeiro valor da lista se existir
//     PopBack()                     // remove o último valor da lista se existir
//     Front() *Node                 // retorna o primeiro nó válido da lista ou nulo
//     Back() *Node                  // retorna o último nó válido da lista ou nulo
//     Search(value int) *Node       // retorna o nó que contém a primeira ocorrência desse valor ou nulo
//     Insert(node *Node, value int) // insere um novo nó antes do nó passado por referência
//     Remove(node *Node) *Node      // remove o nó passado por referência retornando o nó que ficou no lugar dele
// }

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			fmt.Println(ll.String())
		case "size":
			fmt.Println(ll.Size)
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "walk":
			fmt.Print("[ ")
			for node := ll.Front(); node != nil; node = node.Next() {
				fmt.Printf("%v ", node.value)
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != nil; node = node.Prev() {
				fmt.Printf("%v ", node.value)
			}
			fmt.Println("]")
		case "replace":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				node.value = newvalue
			} else {
				fmt.Println("fail: not found")
			}
		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Insert(node, newvalue)
			} else {
				fmt.Println("fail: not found")
			}
		case "remove":
			oldvalue, _ := strconv.Atoi(args[1])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Remove(node)
			} else {
				fmt.Println("fail: not found")
			}
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
