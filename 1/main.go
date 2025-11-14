package main

const a string = "Constante de escopo global tipo string. não precisa ser usada necessariamente e pode ser acessada em qualquer lugar do código que esteja usando o pacote main"

// B assume um valor como false ao ser declarado desta forma.
var b bool

// C assume um valor como 0 ao ser declarado desta forma.
var c int

// D assume um valor como "" ao ser declarado desta forma.
var d string

// E assume um valor como 0.0 ao ser declarado desta forma.
var e float64

/* Declaração pode ser feita desta forma
var (
	var b bool
	var c int
	var d string
)
*/

func main() {

	// Escopo não global, escopo local, obrigatório o uso dentro da função main
	var exemplo string

	// Ao usar := Ele captura o tipo e não é necessario especificar o tipo.
	exemplo2 := ""

	println(a)
	println(b)
	println(c)
	println(d)
	println(e)

	println(exemplo)
	println(exemplo2)
}
