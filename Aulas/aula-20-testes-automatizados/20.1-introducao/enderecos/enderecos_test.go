package enderecos_test

import (
	. "introducao-testes/enderecos"
	"testing"
)
type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado string
}

func TestTipoDeEndereco(t *testing.T) {
	t.Parallel()

	cenarioDeTeste := []cenarioDeTeste{
		{ "Rua ABC", "Rua" },
		{ "Avenida Paulista", "Avenida" },
		{ "Rodovia dos Imigrantes", "Rodovia" },
		{ "Praça das Rosas", "Tipo Inválido" },
		{ "Estrada Qualquer", "Estrada" },
		{ "RUA DOS BOBOS", "Rua" },
		{ "AVENIDA REBOUÇAS", "Avenida" },
		{ "", "Tipo Inválido" },
	}

	for _, cenario := range cenarioDeTeste {
		TipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if TipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				TipoDeEnderecoRecebido,
				cenario.retornoEsperado,
			)
		}
	}
}

func TestQualquer(t *testing.T) {
	t.Parallel()

	if 1 > 2 {
		t.Error("Teste quebrou!")
	}
}