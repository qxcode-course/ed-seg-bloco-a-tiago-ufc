package main

import "fmt"

func prox_vivo(elementos []int, pos int) int {
    size := len(elementos)
    for {
        pos = (pos + 1) % size
        if elementos[pos] != 0 {
            return pos
        }
    }
}

func imprimir(elementos []int, pos int) {
    fmt.Printf("[ ")
    for i := 0; i < len(elementos); i++ {
        if elementos[i] != 0 {
            fmt.Printf("%d", elementos[i])
            if i == pos {
                fmt.Printf(">")
            }
            fmt.Printf(" ")
        }
    }
    fmt.Println("]")
}

func executar(elementos []int, pos int) int {
    vivos := len(elementos)
    for vivos > 1 {
        imprimir(elementos, pos)
        alvo := prox_vivo(elementos, pos)
        elementos[alvo] = 0
        vivos--
        pos = prox_vivo(elementos, alvo)
    }
    imprimir(elementos, pos)
    return pos
}

func main() {
    var size, pos int
    fmt.Scanf("%d %d", &size, &pos)
    pos-- // ajustar para índice 0
    elementos := make([]int, size)
    for i := 0; i < size; i++ {
        elementos[i] = i + 1
    }
    executar(elementos, pos)    
}
