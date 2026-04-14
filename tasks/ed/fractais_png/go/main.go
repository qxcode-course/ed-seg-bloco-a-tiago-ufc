package main

import (
	"fmt"
	"math/rand"
)

func randInt(min, max int) float64 {
	return float64(min + rand.Intn(max-min+1))
}

func arvere(pen *Pen, dist float64) {
	ang := 30.0
	fator := 0.75
	ang_dir := ang + randInt(5, 5)
	if dist < 10 {
		return
	}
	pen.Walk(dist)
	pen.Right(ang)
	arvere(pen, dist-fator)
	pen.Left(2 * ang)
	arvere(pen, dist-fator)
}

func embuah(pen *Pen, dist float64) {
	if dist < 10 {
		return
	}
	pen.Up()
	pen.SetLineWidth(dist / 30)
	pen.Walk(dist)
	pen.Right(90)
	embuah(pen, dist-20)
	pen.Down()
	pen.Left(90)
	pen.Walk(-dist)
}

func main() {
	pen := NewPen(1000, 1000) // cria um canvas de 500 de largura por 500 de altura
	pen.SetRGB(255, 0, 0)     // muda a cor do pincel para vermelho
	pen.SetPosition(250, 400) // move o pincel para x 250, y 500
	pen.SetHeading(90)
	// embuah(pen, 70)
	arvere(pen, 70)

	// pen.SetHeading(90)        // coloca o pincel apontando para cima
	// pen.Walk(100)             // anda 100 pixels
	// pen.Left(30)              // dobra 30 graus para esquerda
	// pen.Walk(100)             // anda 100 pixels
	// pen.DrawCircle(50)        // desenha um círculo de raio 50
	// pen.Right(60)             // gira para direita 60 graus
	// pen.Walk(150)
	// for range 10 {
	// 	pen.Up()
	// 	pen.Walk(30) // anda sem riscar
	// 	pen.Down()
	// 	pen.DrawCircle(10) //desenha um circulo pequeno
	// 	pen.Up()
	// 	pen.Walk(-30) // volta sem riscar
	// 	pen.Down()
	// 	pen.Left(36) // gira
	// }
	pen.SavePNG("tree.png")
	fmt.Println("PNG file created successfully.")
}
