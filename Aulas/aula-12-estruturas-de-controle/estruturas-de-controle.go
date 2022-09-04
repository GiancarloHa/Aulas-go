package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if outronumero := numero; outronumero > 0 {
		fmt.Println("Número é maior que zero")
	} else if numero < - 10 {
		fmt.Println("Número é menor que -10")
	} else {
		fmt.Println("Entre 0 e -10")
	}

	// fmt.Println(outronumero) variavel é criada no escopo do if, fora do if a variavel não existe

}