package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}

	fmt.Println("Pressione qualquer tecla para encerrar aplicação")
	fmt.Scanln()
}
