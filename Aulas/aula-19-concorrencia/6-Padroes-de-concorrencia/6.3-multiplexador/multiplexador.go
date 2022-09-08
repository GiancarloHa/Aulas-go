<<<<<<< HEAD
package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main() {
	canal := multiplexar(escrever("Olá Mundo!"), escrever("Programando em Go!"))

	for i := 0; i < 10; i++ {
		fmt.Println(<- canal)
	}
}

func multiplexar(canalDeEntrada1, canalDeEntrada2 <- chan string) <- chan string {
	canalDeSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <- canalDeEntrada1:
				canalDeSaida <- mensagem
			case mensagem := <- canalDeEntrada2:
				canalDeSaida <- mensagem
			}
		}
	}()
	return canalDeSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
=======
package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main() {
	canal := multiplexar(escrever("Olá Mundo!"), escrever("Programando em Go!"))

	for i := 0; i < 10; i++ {
		fmt.Println(<- canal)
	}
}

func multiplexar(canalDeEntrada1, canalDeEntrada2 <- chan string) <- chan string {
	canalDeSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <- canalDeEntrada1:
				canalDeSaida <- mensagem
			case mensagem := <- canalDeEntrada2:
				canalDeSaida <- mensagem
			}
		}
	}()
	return canalDeSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
>>>>>>> 9a41db2f2f686b367023230ea880f1a970ac856c
}