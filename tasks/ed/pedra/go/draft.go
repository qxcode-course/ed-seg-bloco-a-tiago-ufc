package main

Struct jogadas {

}

func main() {
	qtd := 0
	jogadas := make([]Jogada, qtd)
    for _, jog := range jogadas {
        fmt.Scan(&jog.pa, &jog.pb)
    }
    valor_melhor := 0
    fori, jog := range jogadas {
        if jog.pa < 10 || jog.pb < 10 {
            continue
        } else {
            
        }
    }

}
