package main

import "fmt"

func main() {
	var H, P, F, D int
	fmt.Scanf("%d %d %d %d", &H, &P, &F, &D)
	pos := F
	for {
		pos = (pos + D + 16) % 16
		if pos == H {
			fmt.Println("S")
			return
		}
		if pos == P {
			fmt.Println("N")
			return
		}
	}
}
