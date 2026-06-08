package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getMen(vet []int) []int {
	list := []int{}
	for _, val := range vet {
		if val > 0 {
			list = append(list, val)
		}
	}
	return list
}

func getCalmWomen(vet []int) []int {
	list := []int{}
	for _, val := range vet {
		if val < 0 && val > -10 {
			list = append(list, val)
		}
	}
	return list
}

func sortVet(vet []int) []int {
	sort.Ints(vet)
	return vet
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sortStress(vet []int) []int {
	sort.Slice(vet, func(i, j int) bool {
		return abs(vet[i]) < abs(vet[j])
	})
	return vet
}

func reverse(vet []int) []int {
	rev := make([]int, len(vet))
	copy(rev, vet)
	for i, j := 0, len(rev)-1; i < j; i, j = i+1, j-1 {
		rev[i], rev[j] = rev[j], rev[i]
	}
	return rev
}

func unique(vet []int) []int {
	visto := make(map[int]bool)
	uni := make([]int, 0, len(vet))

	for _, val := range vet {
		if !visto[val] {
			visto[val] = true
			uni = append(uni, val)
		}
	}
	return uni
}

func repeated(vet []int) []int {
    visto := make(map[int]bool)
    rep := []int{}

    for _, val := range vet {
        if visto[val] {
            rep = append(rep, val)
        } else {
            visto[val] = true
        }
    }
    return rep
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}
