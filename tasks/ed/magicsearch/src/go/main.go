package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// func MagicSearch(slice []int, value int) int {
// 	last := -1
// 	pos := len(slice)
// 	for i := 0; i < len(slice); i++ {
// 		if slice[i] == value {
// 			last = i
// 		}
// 		if slice[i] >= value && pos == len(slice) {
// 			pos = i
// 		}
// 	}
// 	if last != -1 {
// 		return last
// 	}
// 	return pos
// }

func MagicSearch(slice []int, value int) int {
	n := len(slice)
	if n == 0 {
		return 0
	}
	left, right := 0, n-1
	pos := n
	found := -1
	for left <= right {
		mid := (left + right) / 2
		if slice[mid] == value {
			found = mid
			left = mid + 1
		} else if slice[mid] < value {
			left = mid + 1
		} else {
			pos = mid
			right = mid - 1
		}
	}
	if found != -1 {
		return found
	}
	return pos
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	slice := make([]int, 0, 1)
	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}
	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	result := MagicSearch(slice, value)
	fmt.Println(result)
}
