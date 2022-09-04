package main

import "fmt"

func main() {
	//Arimeticos
	soma := 1 + 2
	subtracao := 2 - 1
	divisao := 4 / 2
	multipicacao := 10 * 2
	restodadivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multipicacao, restodadivisao)

	//var numero1 int16 = 10
	//var numero2 int32 = 10
	//soma1 := numero1 + numero2
	//fmt.Println(soma)
	// Não vai funcionar pois o GO não faz operação entre tipos
	var numero1 int16 = 10
	var numero2 int16 = 10
	soma1 := numero1 + numero2
	fmt.Println(soma1)

	// Atribuição
	var variavel string = "String"
	variavel2 := "String"
	fmt.Println(variavel)
	fmt.Println(variavel2)

	// Operadores relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 1)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	// Operadores lógicos
	fmt.Println("Operadores Lógicos")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	// Operadores unários
	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)

	numerod := 10
	numero--
	numero -= 2
	fmt.Println(numerod)
	numerod *= 2
	fmt.Println(numerod)
	numerod /= 1
	fmt.Println(numerod)
	numerod %= 2
	fmt.Println(numerod)

	// Operador ternario não existe em GO
	// texto := numerod > 5 ? "Maior que 5" : "Menor que 5"
	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}
	fmt.Println(texto)
}