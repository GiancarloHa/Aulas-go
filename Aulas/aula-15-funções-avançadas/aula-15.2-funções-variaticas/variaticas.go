package main

import "fmt"

func soma(numeros ...int){
	total := 0
	for_, numero := range numeros {
		total += numero
	}

	return total
}

func escrever(texto, string, numeros ...int) {
	for _, numero := range numero {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalDaSoma := soma(1, 2, 3, 0)
	fmt.Println(totalDaSoma)

	escrever("Olá mundo", 10, 30, 15, 6)
}