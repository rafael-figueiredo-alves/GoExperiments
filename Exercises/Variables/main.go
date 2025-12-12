package main

import "fmt"

var globalVar = "Esta é uma variável global" //declaração de variável global

func main() {
	var num1 = 10     //declaração com inferência de tipo
	var num2 int = 20 //declaração com tipo explícito

	sum := num1 + num2 //declaração curta e atribuição
	VariavelLocal := "Esta é uma variável local"

	fmt.Println(globalVar)
	fmt.Println(VariavelLocal)

	fmt.Println("A soma de", num1, "e", num2, "é:", sum)

	var num3, num4, num5, num6 int = 1, 2, 3, 4 //declaração múltipla
	fmt.Println("Números múltiplos:", num3, num4, num5, num6)

	/* Mais exemplos de variáveis múltiplas */

	var a, b = 7, "numero sete" //declaração múltipla com inferência de tipo
	fmt.Println("Valor de a:", a)
	fmt.Println("Valor de b:", b)

	c, d := "Olá", 3.14 //declaração curta múltipla
	fmt.Println("Valor de c:", c)
	fmt.Println("Valor de d:", d)

	// Declarção de variáveis em bloco

	var (
		x int
		y int    = 15
		z string = "Variável em bloco"
	)

	x = 10

	fmt.Println("Valor de x:", x)
	fmt.Println("Valor de y:", y)
	fmt.Println("Valor de z:", z)
}
