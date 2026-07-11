package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	data     []int
	size     int
	capacity int
}

func NewVector(cap int) *Vector { // fugiu do escopo, é dinamico na heap
	vector := Vector{
		data:     make([]int, cap), // nunca use len(data) ou cap(data) ou qq método do go de manipulação de array
		size:     0,
		capacity: cap,
	}
	return &vector
}

func (v *Vector) Reserve(newCapacity int) {
	if newCapacity > v.capacity {
		v.capacity = newCapacity
	}
}

func (v *Vector) Status() string {
	return fmt.Sprintf("size:%d capacity:%d", v.size, v.capacity)
}

func (v *Vector) PushBack(valor int) {
	if v.size == v.capacity {
		newCapacity := v.capacity * 2
		if newCapacity == 0 {
			newCapacity = 2
		}
		newData := make([]int, newCapacity)
		copy(newData, v.data)
		v.data = newData
		v.capacity = newCapacity
	}
	v.data[v.size] = valor
	v.size = v.size + 1
}

func (v *Vector) PopBack() (int, error) {
	if v.size == 0 {
		return 0, fmt.Errorf("vector is empty")
	}
	popValue := v.data[v.size-1]
	v.data[v.size-1] = 0
	v.size = v.size - 1
	return popValue, nil
}

func (v *Vector) Erase(index int) error {
	if index < 0 || index >= v.size {
		return fmt.Errorf("index out of range")
	}
	for i := index; i < v.size-1; i++ {
		v.data[i] = v.data[i+1]
	}
	v.data[v.size-1] = 0
	v.size--
	return nil
}

func (v *Vector) Insert(index int, value int) error {
	if index < 0 || index >= v.size {
		return fmt.Errorf("index out of range")
	}
	if v.size == v.capacity {
		newCapacity := v.capacity * 2
		if newCapacity == 0 {
			newCapacity = 2
		}
		newData := make([]int, newCapacity)
		copy(newData, v.data)
		v.data = newData
		v.capacity = newCapacity
	}
	for i := v.size; i > index; i-- {
		v.data[i] = v.data[i-1]
	}
	v.data[index] = value
	v.size++
	return nil
}

func (v *Vector) Clear() {
	for range v.capacity {
		v.PopBack()
	}
}

func (v *Vector) String() string {
	return "[" + Join(v.data[0:v.size], ", ") + "]"
}

func (v *Vector) At(index int) (int, error) {
	if index < 0 || index >= v.size {
		return 0, fmt.Errorf("index out of range")
	}
	return v.data[index], nil
}

func (v *Vector) Set(index int, value int) error {
	if index < 0 || index >= v.size {
		return fmt.Errorf("index out of range")
	}
	v.data[index] = value
	return nil
}

func (v *Vector) IndexOf(value int) int {
	for i := 0; i < v.size; i++ {
		if value == v.data[i] {
			return i
		}
	}
	return -1
}

func (v *Vector) Contains(value int) bool {
	for i := 0; i < v.size; i++ {
		if value == v.data[i] {
			return true
		}
	}
	return false
}

func (v *Vector) Slice(start int, end int) *Vector {
	if v.size == 0 {
		return v
	}

	start = (start%v.size + v.size) % v.size
	end = (end%v.size + v.size) % v.size

	sub := v.data[start:end]

	return &Vector{
		data:     sub,
		size:     len(sub),
		capacity: cap(sub),
	}
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	var result strings.Builder
	fmt.Fprintf(&result, "%d", slice[0])
	for _, value := range slice[1:] {
		fmt.Fprintf(&result, "%s%d", sep, value)
	}
	return result.String()
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewVector(0)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
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
			v = NewVector(value)
		case "push":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.PushBack(value)
			}
		case "show":
			fmt.Println(v)
		case "status":
			fmt.Println(v.Status())
		case "pop":
			_, err := v.PopBack()
			if err != nil {
				fmt.Println(err)
			}
		case "insert":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Insert(index, value)
			if err != nil {
				fmt.Println(err)
			}
		case "erase":
			index, _ := strconv.Atoi(parts[1])
			err := v.Erase(index)
			if err != nil {
				fmt.Println(err)
			}
		case "indexOf":
			value, _ := strconv.Atoi(parts[1])
			index := v.IndexOf(value)
			fmt.Println(index)
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if v.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "clear":
			v.Clear()
		case "capacity":
			// fmt.Println(v.Capacity())
		case "get":
			index, _ := strconv.Atoi(parts[1])
			value, err := v.At(index)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(value)
			}
		case "set":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Set(index, value)
			if err != nil {
				fmt.Println(err)
			}

		case "reserve":
			newCapacity, _ := strconv.Atoi(parts[1])
			v.Reserve(newCapacity)
		case "slice":
			start, _ := strconv.Atoi(parts[1])
			end, _ := strconv.Atoi(parts[2])
			slice := v.Slice(start, end)
			fmt.Println(slice)
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
