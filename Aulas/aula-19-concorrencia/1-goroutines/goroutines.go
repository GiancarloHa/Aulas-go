<<<<<<< HEAD
package main

import (
	"fmt"
	"time"
)

func main() {
	// Concorrencia != Paralelismo
	go escrever("Olá Mundo!") // goroutine
	escrever("Programando em Go!")
}

func escrever(texto string) {
	for{
	fmt.Println(texto)
	time.Sleep(time.Second)
	}
=======
package main

import (
	"fmt"
	"time"
)

func main() {
	// Concorrencia != Paralelismo
	go escrever("Olá Mundo!") // goroutine
	escrever("Programando em Go!")
}

func escrever(texto string) {
	for{
	fmt.Println(texto)
	time.Sleep(time.Second)
	}
>>>>>>> 9a41db2f2f686b367023230ea880f1a970ac856c
}