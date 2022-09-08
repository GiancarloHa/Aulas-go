<<<<<<< HEAD
package main

import "fmt"

func main() {
	canal := make(chan string, 2)
	canal <- "Olá mundo"
	canal <- "Programando em Go!"

	mensagem := <- canal
	mensagem2 := <- canal
	fmt.Println(mensagem)
	fmt.Println(mensagem2)
=======
package main

import "fmt"

func main() {
	canal := make(chan string, 2)
	canal <- "Olá mundo"
	canal <- "Programando em Go!"

	mensagem := <- canal
	mensagem2 := <- canal
	fmt.Println(mensagem)
	fmt.Println(mensagem2)
>>>>>>> 9a41db2f2f686b367023230ea880f1a970ac856c
}