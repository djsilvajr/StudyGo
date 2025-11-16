package main

import "errors"

//Criação de tipagens

// Essas criações de tipagens geralmente valem criar em outro package, por exemplo package email,
// la poderemos ter uma função de validação de email + a criação do tipo, que vai validar o tipo da variavel

type CPF string
type RG string
type EMAIL string
type SENHA string

var (
	// poderia ser senha passwords.SENHA onde eu pegaria o tipo senha do package password validando os campos obrigatórios.
	senha            SENHA
	confirmacaoSenha SENHA
	email            EMAIL
	emailRecuperacao EMAIL
	rg               RG
	cpf              CPF
)

//esta função teria como responsabilidade validar o tipo dentro do package cpf
func newCpf(str string) (CPF, error) {
	// Aqui teriamos as funções de validação verificando str
	if str == "" {
		return "", errors.New("CPF invalido")
	}
	cpf := CPF(str)
	return cpf, nil
}

func main() {

	//Exemplo de como utilizar a criação de tipos. Chamaria a função de verificação direto do package cpf
	str, err := newCpf("teste")
	if err != nil {
		println("erro")
		return
	}

	println(str)
}
