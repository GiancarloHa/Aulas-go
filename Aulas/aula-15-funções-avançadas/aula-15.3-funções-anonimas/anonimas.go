package main

func main () {

	func () {
		fmt.Println("Olá mundo")
	}()

	func (texto string) {
		fmt.Println(texto)
	}("Passando Parâmetro")

	retorno := func(texto1 string) string {
		return fmt.Sprintf("Recebid -> %s", texto1)
	}("Passando Parâmetro")

	fmt.Println(retorno)
}