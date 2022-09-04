package main

import (
	"errors"
	"fmt"
)

func main () {
	var numero int16 = 100
	fmt.Println(numero)
	//var numero2 int8 = 10000000 não vai compilar número muito alto para int8
	var numero1 int = 100000000000 // utiliza a arquitetura de seu computador no meu caso 64bits
	fmt.Println(numero1)
	numero2 := -100000 //É possível utilizar números negativos sem problemas
	fmt.Println(numero2)
	var numero3 uint32 = 1000 // uint funciona similar ao int e sem declarar o número de bits irá utilizar a arquitetura do computado
	fmt.Println(numero3)
	//alias
	// INT32 = RUNE
	var numero4 rune = 12456
	fmt.Print(numero4)

	//BYTE = UINT8
	var numero5 byte = 123
	fmt.Println(numero5)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 12344545343.65
	fmt.Println(numeroReal2)

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'b' // Número da tabela asc desse caractere
	fmt.Println(char)

	var texto string
	fmt.Println(texto)
	var texto1 int
	fmt.Println(texto1)

	var booleanol bool = true
	fmt.Println(booleanol)
	var booleanol1 bool
	fmt.Println(booleanol1)

	var erro error
	fmt.Println(erro)
	var erro1 error = errors.New("Erro interno")
	fmt.Println(erro1)
}
