package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Set struct {
	data     []int
	size     int
	capacity int
}

func NewSet(value int) *Set {
	return &Set{
		data:     make([]int, value),
		size:     0,
		capacity: value,
	}
}

func (s *Set) grow() {
	newCap := s.capacity + 1
	if newCap == 0 {
		newCap = 1
	}
	newData := make([]int, newCap)
	copy(newData, s.data)
	s.data = newData
	s.capacity = newCap
}

func binarySearch(slice []int, value int) int {
	n := len(slice)
	if n == 0 {
		return -1
	}
	left, right := 0, n-1
	found := -1
	for left <= right {
		mid := (left + right) / 2
		if slice[mid] == value {
			found = mid
			left = mid + 1
		} else if slice[mid] < value {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return found
}

func (s *Set) Contains(value int) bool {
	return binarySearch(s.data[:s.size], value) != -1
}

func (s *Set) Erase(value int) {
	idx := binarySearch(s.data[:s.size], value)
	if idx == -1 {
		fmt.Println("value not found")
		return
	}
	for i := idx; i < s.size-1; i++ {
		s.data[i] = s.data[i+1]
	}
	s.size--
}

func (s *Set) Insert(value int) {
	idx := binarySearch(s.data[:s.size], value)
	if idx != -1 {
		return
	}
	pos := s.size
	for i := 0; i < s.size; i++ {
		if s.data[i] > value {
			pos = i
			break
		}
	}
	if s.size == s.capacity {
		s.grow()
	}
	for i := s.size; i > pos; i-- {
		s.data[i] = s.data[i-1]
	}
	s.data[pos] = value
	s.size++
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	s := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]
		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			s = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				s.Insert(value)
			}
		case "show":
			if s.size == 0 {
				fmt.Println("[]")
			} else {
				str := "["
				for i := 0; i < s.size; i++ {
					str += fmt.Sprintf("%d", s.data[i])
					if i < s.size-1 {
						str += ", "
					}
				}
				str += "]"
				fmt.Println(str)
			}
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			s.Erase(value)
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if s.Contains(value) {
				fmt.Println(true)
			} else {
				fmt.Println(false)
			}
		// case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
