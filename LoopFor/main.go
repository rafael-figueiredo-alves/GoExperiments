package main

import "fmt"

func main() {

	i := 1
	for {
		fmt.Println("teste ", i)
		if i >= 5 {
			break
		}
		i += 1
	}

	a := 0
	for {
		a += 1
		if a > 10 {
			break
		}
		fmt.Println("Imprimindo ... ", a)
	}
}
