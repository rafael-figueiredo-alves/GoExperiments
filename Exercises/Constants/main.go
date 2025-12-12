package main

import "fmt"

func main() {
	const pi = 3.14                       // Declaração de constante simples
	const greeting string = "Olá, Mundo!" // Declaração de constante com tipo explícito
	const (
		euler    = 2.71828
		language = "Go"
		version  = 1.18
	) // Declaração de múltiplas constantes em bloco

	fmt.Println("Valor de pi:", pi)
	fmt.Println(greeting)
	fmt.Println("Constante Euler:", euler)
	fmt.Println("Linguagem:", language)
	fmt.Println("Versão:", version)
	// Tentativa de modificar uma constante (isso causará um erro se descomentado)
	// pi = 3.14159 // Erro: cannot assign to pi
}
