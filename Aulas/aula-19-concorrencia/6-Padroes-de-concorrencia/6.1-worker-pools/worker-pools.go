<<<<<<< HEAD
package main

import (
	"fmt"
)

func main() {
	tarefas := make(chan int, 45)
	resultado := make(chan int, 45)

	go worker(tarefas, resultado)
	go worker(tarefas, resultado)
	go worker(tarefas, resultado)
	go worker(tarefas, resultado)

	for i:=0; i < 45; i++{
		tarefas <- i
	}
	close(tarefas)
	for i:=0; i < 45; i++{
		resultado := <- resultado
		fmt.Println(resultado)
	}
}

func worker (tarefas <- chan int, resultado chan <- int) {
	for numero := range tarefas {
		resultado <- fibonacci(numero)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
=======
package main

import (
	"fmt"
)

func main() {
	tarefas := make(chan int, 45)
	resultado := make(chan int, 45)

	go worker(tarefas, resultado)
	go worker(tarefas, resultado)
	go worker(tarefas, resultado)
	go worker(tarefas, resultado)

	for i:=0; i < 45; i++{
		tarefas <- i
	}
	close(tarefas)
	for i:=0; i < 45; i++{
		resultado := <- resultado
		fmt.Println(resultado)
	}
}

func worker (tarefas <- chan int, resultado chan <- int) {
	for numero := range tarefas {
		resultado <- fibonacci(numero)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
>>>>>>> 9a41db2f2f686b367023230ea880f1a970ac856c
}