package main

import "fmt"

func main() {
	var num1, num2 int

	fmt.Println("Informe o primeiro valor a somar")
	_, err := fmt.Scan(&num1)

	if err != nil {
		fmt.Println("Valor informado é inválido. Só são aceitos números inteiros")
		return
	}

	fmt.Println("Informe o segundo valor a somar")
	_, err2 := fmt.Scan(&num2)

	if err2 != nil {
		fmt.Println("Valor informado é inválido. Só são aceitos números inteiros")
		return
	}

	fmt.Println()

	fmt.Println("Valores a somar:", num1, "e", num2)
	fmt.Println("Somando os valores, nós temos o número:", SomarNumeros(num1, num2))

	fmt.Println()
	fmt.Println("Pressione qualquer tecla para encerrar a aplicação")

	var Fim string
	fmt.Scanf("%s", &Fim)
}

func SomarNumeros(primeiroNumero, segundoNumero int) int {
	return primeiroNumero + segundoNumero
}
