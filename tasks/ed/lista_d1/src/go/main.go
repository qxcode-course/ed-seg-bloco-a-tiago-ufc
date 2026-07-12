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
	next  *Node
	prev  *Node
}

type LList struct {
	root Node
	Size int
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
	n := &Node{value: value}
	last := l.root.prev
	last.next = n
	n.prev = last
	n.next = &l.root
	l.root.prev = n
	l.Size++
}

func (l *LList) PushFront(value int) {
	n := &Node{value: value}
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
	for n := l.root.next; n != &l.root; n = n.next {
		n.prev = nil
	}
	l.root.next = &l.root
	l.root.prev = &l.root
}

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
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
