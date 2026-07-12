package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	queue := NewQueue[string]()
	for char := 'A'; char <= 'P'; char++ {
		queue.Enqueue(string(char))
	}
	for i := 0; i < 15; i++ {
		scanner.Scan()
		var j1, j2 int
		fmt.Sscan(scanner.Text(), &j1, &j2)
		t1 := queue.Dequeue()
		t2 := queue.Dequeue()
		if j2 > j1 {
			queue.Enqueue(t2)
		} else {
			queue.Enqueue(t1)
		}
	}
	fmt.Println(queue.Dequeue())
}
