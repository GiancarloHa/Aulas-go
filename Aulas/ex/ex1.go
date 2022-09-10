// Elabore uma função que recebe um objeto com estudantes e suas notas. As notas de cada estudante estarão num array. Você deverá calcurar a média da nota de cada
// aluno e retornar um objeto com os atributos nome e media, que indica que o aluno que teve o melhor desempenho nas notas.
package main

import (
	"fmt"
)

type alunos struct {
	aluno string
	notas [4] float64
	media float64
}

func main () {
	var estudante1 alunos
	estudante1.aluno = "Diego"
	estudante1.notas = [4] float64 {6.0, 6.0, 6.0, 6.0}

	var estudante2 alunos
	estudante2.aluno = "João"
	estudante2.notas = [4] float64 {8.0, 7.6, 8.9, 6.0}

	var estudante3 alunos
	estudante3.aluno = "Carla"
	estudante3.notas = [4] float64 {7.0, 7.6, 9.0, 10.0}

	var mediaaluno1 alunos
	mediaaluno1.media = calculodemedia(estudante1.notas[0], estudante1.notas[1], estudante1.notas[2], estudante1.notas[3])

	var mediaaluno2 alunos
	mediaaluno2.media = calculodemedia(estudante2.notas[0], estudante2.notas[1], estudante2.notas[2], estudante2.notas[3])

	var mediaaluno3 alunos
	mediaaluno3.media = calculodemedia(estudante3.notas[0], estudante3.notas[1], estudante3.notas[2], estudante3.notas[3])

	if mediaaluno1.media > mediaaluno2.media && mediaaluno1.media > mediaaluno3.media {
		fmt.Println(estudante1.aluno, mediaaluno1.media)
	} else if mediaaluno2.media > mediaaluno1.media && mediaaluno2.media > mediaaluno3.media {
		fmt.Println(estudante2.aluno, mediaaluno2.media)
	} else {
		fmt.Println(estudante3.aluno, mediaaluno3.media)
	} 
}

func calculodemedia (n1, n2, n3, n4 float64) (media float64) {
	media = (n1 + n2 + n3 + n4) / 4
	return
}