package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sábado"
	default:
		return: "Número Inválido"
	}

}

func diaDaSemana2(numero int) string {
	var diaDaSemana string

	switch {
	case diaDaSemana == 1:
		return "Domingo"
	case diaDaSemana == 2:
		return "Segunda-Feira"
	case diaDaSemana == 3:
		return "Terça-Feira"
	case diaDaSemana == 4:
		return "Quarta-Feira"
	case diaDaSemana == 5:
		return "Quinta-Feira"
	case diaDaSemana == 6:
		return "Sexta-Feira"
	case diaDaSemana == 7:
		return "Sábado"
	default:
		return: "Número Inválido"
	}
	return diaDaSemana
}

func main () {
	fmt.Println("Switch")

	dia := diaDaSemana(3)
	fmt.Println(dia)

	dia2 := diaDaSemana2(3)
	fmt.Println(dia2)

	dia3 := diaDaSemana2(10)
	fmt.Println(dia3)
}
