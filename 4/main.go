package main

import (
	"fmt"
)

func main() {
	var boolVar bool = true
	// Sintaxe do if : Poderia ser booleanVar == true
	if boolVar {
		println("The boolean variable is true")
	} else {
		println("The boolean variable is false")
	}

	var s string = "Exemplo de string"
	// Sintaxe do if com declaração curta
	// Aqui, n é declarado e inicializado com o comprimento de s, caso s for maior que 0 ele executará a função fmt.Println
	if n := len(s); n > 0 {
		fmt.Println("tamanho de string caso maior que 0 caracteres:", n)
	}

	// if com multiplas variáveis declaradas: Aqui eu declaro a como 5 e b como 10, e verifico se a é menor que b
	if a, b := 5, 10; a < b {
		fmt.Println(a, "é menor que", b)
	}

	// utilização com erro
	// Exemplo 1: chamada que retorna erro
	if example1, err := retornarErro(); err == nil {
		fmt.Println("Resposta da função (err == nil):", example1)
	} else {
		fmt.Println("Erro retornado pela função:", err)
	}
	// Exemplo 2: chamada que não retorna erro
	if example2, err := retornarAcerto(); err == nil {
		fmt.Println("Resposta da função (err == nil):", example2)
	} else {
		fmt.Println("Erro retornado pela função:", err)
	}
}

// função para retornar um erro
func retornarErro() (int, error) {
	str := "Hello, Go!"
	err := fmt.Errorf("'%s' Este é um erro", str)
	return len(str), err
}

// função para retornar um acerto
func retornarAcerto() (int, error) {
	str := "Go"
	return len(str), nil
}
