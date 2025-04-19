package main

import "fmt"

func main() {
	var num1 int
	var num2 int

	fmt.Print("Informe os dois números que deseja somar:")
	fmt.Scanln(&num1, &num2)

	fmt.Println("O resultado da soma entre", num1, " e ", num2, " é ", num1+num2)
}
