# Exemplo: Uso de if em Go (arquivo único)

Este repositório contém um único arquivo Go que demonstra diferentes formas de usar a instrução `if` em Go: checagem simples de boolean, `if` com declaração curta, `if` com múltiplas variáveis e tratamento de erro com a forma curta de declaração (`v, err := fn()`).

Arquivo apresentado: `main.go`

## Conteúdo do `main.go`

```go
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
```

## Requisitos

- Go (versão 1.20+ recomendada, mas funciona em versões modernas do Go)

## Como executar

1. Salve o código acima em um arquivo chamado `main.go`.
2. No terminal, execute:

- Para executar diretamente:
  - go run main.go

- Para compilar e executar:
  - go build -o exemplo && ./exemplo

## Saída esperada

Quando executado, o programa imprime algo similar a:

The boolean variable is true
tamanho de string caso maior que 0 caracteres: 17
5 é menor que 10
Erro retornado pela função: 'Hello, Go!' Este é um erro
Resposta da função (err == nil): 2

Observações:
- O primeiro `println` usa a função embutida `println` (útil para debug simples). As demais impressões usam `fmt.Println`, que é a prática mais comum em programas Go.
- As declarações curtas em `if`, como `if n := len(s); n > 0 { ... }`, limitam o escopo da variável `n` ao bloco do `if`.
- O padrão `if v, err := fn(); err != nil { ... }` é o estilo idiomático em Go para tratamento de erros; no código aqui foi usado `err == nil`/`else` apenas para ilustrar outra forma válida de verificar o erro.

## Melhoria sugerida

- Tornar o tratamento de erro mais idiomático (verificar `err != nil` primeiro).
- Usar consistentemente `fmt.Println` em vez de `println` para saída formatada.
- Adicionar testes unitários para as funções `retornarErro` e `retornarAcerto` caso o programa cresça.

---
Feito: gerei este README.md explicando o código e mostrando como rodá-lo e o resultado esperado. Se quiser, posso também criar o arquivo `main.go` no repositório ou ajustar o README com instruções de build/CI específicas.