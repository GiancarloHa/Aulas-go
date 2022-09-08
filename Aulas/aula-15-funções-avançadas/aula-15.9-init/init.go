package main

import "fmt"

var n int

func init () {
	fmt.Println("Executantando a função init")
	n = 10
}

func main () {
	fmt.Println("Função main sendo executada")
	fmt.Println(n)
}