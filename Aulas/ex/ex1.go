// Elabore uma função que recebe um objeto com estudantes e suas notas. As notas de cada estudante estarão num array. Você deverá calcurar a média da nota de cada
// aluno e retornar um objeto com os atributos nome e media, que indica que o aluno que teve o melhor desempenho nas notas.
package main

import (
	"fmt"
)

type alunos struct {
	aluno string
	notas [4] float64
	nome string
	media float64
}

func main () {


}

func mediaDosalunos (estudante, es string) (media float64){
	totaldenotas := 4.0

	var estudante1 alunos
	estudante1.aluno = "Diego Podre"
	estudante1.notas = [4] float64 {}

	var estudante2 alunos
	estudante2.aluno = "João"
	estudante2.notas = [4] float64 {8.0, 7.6, 8.9, 6.0}

	var estudante3 alunos
	estudante3.aluno = "Carla"
	estudante3.notas = [4] float64 {7.0, 7.6, 9.0, 10.0}

	var mediaaluno1 alunos
	mediaaluno1.nome = estudante1.aluno
	mediaaluno1.media = (estudante1.notas[0] + estudante1.notas[1] + estudante1.notas[2] + estudante1.notas[3]) / (totaldenotas)

	var mediaaluno2 alunos
	mediaaluno2.nome = estudante2.aluno
	mediaaluno2.media = (estudante2.notas[0] + estudante2.notas[1] + estudante2.notas[2] + estudante2.notas[3]) / (totaldenotas)

	var mediaaluno3 alunos
	mediaaluno3.nome = estudante3.aluno
	mediaaluno3.media = (estudante3.notas[0] + estudante3.notas[1] + estudante3.notas[2] + estudante3.notas[3]) / (totaldenotas)
}