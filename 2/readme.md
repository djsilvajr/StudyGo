# README

Projeto de exemplo em Go — definição de tipos fortes (type aliases) e validação básica

Este repositório contém um exemplo simples (package `main`) que demonstra a criação de tipos específicos para valores como CPF, RG, EMAIL e SENHA, e uma função de construção/validação (`newCpf`) que ilustra onde a lógica de validação deve ficar.

> Observação importante: o código de exemplo contém uma condição de validação invertida em `newCpf` (veja a seção "Correção sugerida").

## Conteúdo do arquivo

- `main.go` (exemplo):
  - Define tipos personalizados:
    - `type CPF string`
    - `type RG string`
    - `type EMAIL string`
    - `type SENHA string`
  - Variáveis de pacote para exemplificar uso.
  - Função `newCpf(str string) (CPF, error)` com propósito de validar e retornar um CPF tipado.
  - `main()` demonstra a chamada de `newCpf`.

## Objetivo

- Demonstrar como tipar valores primitivos (como `string`) para aumentar segurança de tipos e clareza no código.
- Sugerir a separação das responsabilidades: validação e criação de tipos devem ficar em packages separados (ex.: `cpf`, `email`, `password`), enquanto `main` apenas consome essas APIs.

## Problema no exemplo

A implementação original de `newCpf` contém um erro lógico:

```go
func newCpf(str string) (CPF, error) {
	// Aqui teriamos as funções de validação verificando str
	if str != "" {
		return "", errors.New("CPF invalido")
	}
	cpf := CPF(str)
	return cpf, nil
}
```

Nessa versão, qualquer string não vazia é considerada inválida (o teste `if str != ""` retornará erro). O comportamento esperado normalmente é rejeitar strings vazias ou que não batem com a validação de CPF.

## Correção sugerida (exemplo mínimo)

```go
import "errors"

// Validação mínima: rejeita string vazia.
// Em um projeto real, substitua por validação formal de CPF (tamanho, dígitos verificadores, formatação).
func newCpf(str string) (CPF, error) {
    if str == "" {
        return "", errors.New("CPF inválido: string vazia")
    }
    // TODO: adicionar validação real (regex + cálculo de dígitos verificadores)
    return CPF(str), nil
}
```

## Boas práticas recomendadas

1. Separar em packages:
   - `package cpf` com `type CPF string` e `func New(value string) (CPF, error)` que encapsula todas as validações.
   - `package email` com validação de e-mail e tipo `EMAIL`.
   - `package password` com tipo `SENHA` e possíveis regras (mínimo de caracteres, complexidade).
   Isso deixa o `main` focado apenas em orquestrar e usar as APIs.

2. Escrever testes unitários para cada package:
   - Casos válidos e inválidos.
   - Testar bordas (ex.: CPFs com caracteres extras, espaços, formato com pontos/traço).

3. Validadores robustos:
   - Para CPF, calcule os dígitos verificadores (algoritmo conhecido) em vez de confiar só em regex.
   - Para e-mail, utilize regex confiável ou bibliotecas de validação.

4. Documentação:
   - Documente no package o formato aceito (com ou sem máscara, se aceita espaços, etc.).

## Como executar

Assumindo que o arquivo de exemplo está em `main.go`:

- Executar diretamente:
  - `go run main.go`

- Compilar:
  - `go build -o exemplo .`
  - `./exemplo`

- Rodar testes (após mover a lógica para packages e criar testes):
  - `go test ./...`

## Exemplo de estrutura de projeto recomendada

- /cmd/app (ou /)
  - main.go
- /internal/cpf
  - cpf.go
  - cpf_test.go
- /internal/email
  - email.go
  - email_test.go
- /internal/password
  - password.go
  - password_test.go

## Exemplo de API do package `cpf`

```go
package cpf

import "errors"

type CPF string

func New(v string) (CPF, error) {
    if v == "" {
        return "", errors.New("cpf: valor vazio")
    }
    // Remover formatação, validar dígitos, etc.
    if !validFormat(v) {
        return "", errors.New("cpf: formato inválido")
    }
    return CPF(v), nil
}
```

## Próximos passos sugeridos

- Corrigir a condição em `newCpf` no código atual.
- Implementar a validação real de CPF.
- Extrair validação para packages específicos e escrever testes.
- Adicionar exemplos de uso no README (ex.: como criar um CPF válido, como tratar erros).

---

Se quiser, eu posso:
- Gerar uma versão refatorada separando `cpf`, `email` e `password` em packages com testes.
- Escrever testes unitários de exemplo para `newCpf`.
- Corrigir o arquivo `main.go` e criar um pequeno projeto com a estrutura recomendada.