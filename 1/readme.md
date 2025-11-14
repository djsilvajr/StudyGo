# README — Estudos de Declaração de Variáveis em Go

Este projeto contém exemplos práticos de como funcionam **declarações de variáveis**, **constantes**, **escopos** e **tipagem** na linguagem Go (Golang). O código demonstra tanto declarações globais quanto locais, além do uso do operador `:=` para inferência de tipos.

---

## 🔤 Declaração de Constantes

```go
const a string = "Constante de escopo global tipo string. não precisa ser usada necessariamente e pode ser acessada em qualquer lugar do código que esteja usando o pacote main"
```

* Constantes possuem **valor imutável**.
* Podem ter **escopo global**, como no exemplo.
* Podem ser usadas em qualquer parte do pacote onde forem declaradas.

---

## 🧩 Declaração de Variáveis Globais

```go
var b bool   // assume false
var c int    // assume 0
var d string // assume ""
var e float64 // assume 0.0
```

Em Go, quando declaramos uma variável sem atribuir valor, ela recebe o **valor zero** (
*zero value*):

* `bool` → `false`
* `int` → `0`
* `string` → `""`
* `float64` → `0.0`

Também é possível declarar várias variáveis em bloco:

```go
var (
    b bool
    c int
    d string
)
```

Essas variáveis globais podem ser usadas em qualquer função do pacote.

---

## 📍 Variáveis Locais e Escopo

Dentro da função `main`, temos variáveis de **escopo local**:

```go
var exemplo string // precisa ser usada dentro do escopo
```

* Variáveis locais **devem ser utilizadas** dentro da função onde são declaradas.
* Também recebem valor zero caso não seja atribuído valor inicial.

---

## ⚡ Inferência de Tipo com :=

```go
exemplo2 := ""
```

* O operador `:=` infere automaticamente o tipo da variável.
* Só pode ser usado **dentro de funções**.
* É equivalente a:

```go
var exemplo2 string = ""
```

---

## ▶️ Execução e Impressão

O código imprime todas as variáveis declaradas:

```go
println(a)
println(b)
println(c)
println(d)
println(e)

println(exemplo)
println(exemplo2)
```

Isso demonstra na prática os valores padrão das variáveis e a acessibilidade de cada escopo.
