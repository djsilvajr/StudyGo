package main

import (
	"errors"
	"fmt"
)

func main() {
	// ========================================================
	// Uso básico de Print, Println e Printf
	// ========================================================

	var a string = "string"
	var b int = 10
	var c bool = true

	// Imprime na mesma linha (sem quebra automática)
	fmt.Print(a, b, c, "\n")

	// Imprime com espaço automático + quebra de linha
	fmt.Println(a, b, c)

	// ========================================================
	// Principais verbos do fmt
	// ========================================================
	// Strings --------
	// %s = string
	// %q = string com aspas

	// Inteiros -------
	// %d = decimal
	// %b = binário
	// %o = octal
	// %x = hexadecimal
	// %X = hexadecimal maiúsculo
	// %c = caractere Unicode representado pelo número

	// Floats ---------
	// %f = float
	// %.2f = float com 2 casas
	// %e = notação científica

	// Boolean --------
	// %t = true/false

	// Outros ---------
	// %v = valor conforme a tipagem
	// %+v = struct com nomes dos campos
	// %T = tipo da variável
	// %p = endereço na memória

	fmt.Printf("Teste de print: String %s, Int %d, Bool %t\n", a, b, c)
	fmt.Printf("Teste de print com %%v: %v, %v, %v\n", a, b, c)

	// ========================================================
	// Exemplo com Errorf
	// ========================================================
	err := impressaoExemploResposta(404)
	fmt.Println("Erro retornado:", err)

	err = impressaoExemploResposta(500)
	fmt.Println("Erro retornado:", err)

	err = impressaoExemploResposta(123)
	fmt.Println("Erro desconhecido:", err)
}

// ========================================================
// Sistema de respostas de erro
// ========================================================

type RespostasErro struct {
	Retorno   func() error
	Descricao string
}

var respostas = map[int]RespostasErro{
	404: {
		Retorno:   erro404,
		Descricao: "exemplo de retorno erro 404",
	},
	500: {
		Retorno:   erro500,
		Descricao: "exemplo de retorno erro 500",
	},
}

func impressaoExemploResposta(statusCode int) error {
	r, ok := respostas[statusCode]
	if !ok {
		return fmt.Errorf("status code %d não encontrado", statusCode)
	}

	if r.Retorno == nil {
		return errors.New("função de erro não definida")
	}

	return r.Retorno()
}

func erro404() error {
	return errors.New("Valor não encontrado")
}

func erro500() error {
	return errors.New("Erro interno inesperado no servidor")
}
