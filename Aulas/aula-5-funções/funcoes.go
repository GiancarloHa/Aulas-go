package main

import(
	"fmt"
)

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos (n1, n2 int8) (int8, int8){
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main () {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func () {
		fmt.Println("Funcao F")
	}
	f()
	var a = func (txt string) string {
		fmt.Println(txt)
		return txt
	}
	a("Texto da função 1")

	resultado := a("Texto da função a")
	fmt.Println(resultado)

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	resultadoSoma2, _ := calculosMatematicos(10, 15) // _ é utilizado para o compilador entender que o segundo valor você não deseja ele
	fmt.Println(resultadoSoma2)
}