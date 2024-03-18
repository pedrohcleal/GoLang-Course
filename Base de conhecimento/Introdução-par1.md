# Introdução ao Go - part1

## Tipos de dados em Golang

Golang, conhecida por sua simplicidade e poder, oferece um conjunto de tipos de dados básicos e compostos para construir programas robustos e eficientes. Essa variedade permite armazenar e manipular diferentes tipos de informações, desde números e caracteres até estruturas complexas.

**Tipos básicos:**

* **Numéricos:**
    * **Inteiros:** int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64
    * **Ponto flutuante:** float32, float64
    * **Booleano:** bool
* **Texto:** string
* **Caractere:** byte, rune

**Tipos compostos:**

* **Array:** fatia ordenada de um tipo específico
* **Slice:** fatia referenciável e mutável de um array
* **Map:** coleção não ordenada de pares chave-valor
* **Struct:** tipo definido pelo usuário com campos nomeados
* **Interface:** tipo abstrato que define um conjunto de métodos

**Inferência de tipos:**

Golang permite a inferência de tipos, onde o tipo de uma variável é automaticamente inferido pela sua atribuição. Isso simplifica o código, mas exige cuidado para evitar erros.

**Detalhes importantes:**

* Os tipos básicos são tipados estaticamente, garantindo segurança de memória e evitando erros em tempo de execução.
* Os tipos compostos oferecem flexibilidade para armazenar e organizar dados complexos.
* A biblioteca padrão fornece diversos tipos e funções para trabalhar com dados de forma eficiente.

**Exemplos:**

```go
// Inteiro
var numero int = 10

// Ponto flutuante
var pi float64 = 3.14

// Booleano
var verdadeiro bool = true

// String
var nome string = "Golang"

// Array
var array [5]int = [1, 2, 3, 4, 5]

// Slice
var slice []int = []int{1, 2, 3, 4, 5}

// Map
var mapa map[string]int = map[string]int{"a": 1, "b": 2}

// Struct
type Pessoa struct {
  nome string
  idade int
}

// Interface
var interface interface{} = 10
```

**Recursos para aprofundar:**

* Documentação oficial da linguagem Golang: [https://golang.org/pkg/](https://golang.org/pkg/)
* Artigo sobre tipos de dados em Golang: [https://medium.com/@habbema/golang-dados-22e7ff5b8131](https://medium.com/@habbema/golang-dados-22e7ff5b8131)
* Tutorial sobre tipos de dados em Golang: [https://www.digitalocean.com/community/tutorials/understanding-data-types-in-go-pt](https://www.digitalocean.com/community/tutorials/understanding-data-types-in-go-pt)

**Conclusão:**

A compreensão dos tipos de dados em Golang é fundamental para escrever programas eficientes e robustos. A variedade de tipos básicos e compostos, combinada com a inferência de tipos, torna a linguagem flexível e poderosa para lidar com diferentes tipos de informações.

### Operadores:
Go oferece uma variedade de operadores para realizar operações em variáveis e valores. Os operadores incluem:

1. **Aritméticos:** `+`, `-`, `*`, `/`, `%`.
2. **Relacionais:** `==`, `!=`, `>`, `<`, `>=`, `<=`.
3. **Lógicos:** `&&`, `||`, `!`.
4. **Bit a Bit:** `&`, `|`, `^`, `<<`, `>>`.
5. **Atribuição:** `=`, `+=`, `-=`, `*=`, `/=`, `%=`, `<<=`, `>>=`, `&=`, `|=`, `^=`.

## Lógica de controle em Golang

A lógica de controle em Golang define o fluxo de execução de um programa, permitindo tomar decisões e executar diferentes ações com base em condições e valores. As principais estruturas de controle da linguagem são:

**1. Estruturas condicionais:**

* **if:** Permite executar um bloco de código se uma condição for verdadeira.
* **else if:** Permite executar um bloco de código alternativo se a condição do `if` for falsa e outra condição for verdadeira.
* **else:** Permite executar um bloco de código alternativo se todas as condições anteriores forem falsas.
* **switch:** Permite executar um bloco de código específico com base no valor de uma variável.

**Exemplo:**

```go
if numero > 10 {
  fmt.Println("O número é maior que 10")
} else if numero == 10 {
  fmt.Println("O número é igual a 10")
} else {
  fmt.Println("O número é menor que 10")
}

switch numero {
case 1:
  fmt.Println("O número é 1")
case 2:
  fmt.Println("O número é 2")
default:
  fmt.Println("O número não é 1 nem 2")
}
```

**2. Laços de repetição:**

* **for:** Permite executar um bloco de código repetidamente enquanto uma condição for verdadeira.
* **while:** Permite executar um bloco de código repetidamente enquanto uma condição for verdadeira.
* **for range:** Permite iterar sobre um array, slice, map ou string.

**Exemplo:**

```go
for i := 0; i < 5; i++ {
  fmt.Println(i)
}

numero := 1
while numero <= 5 {
  fmt.Println(numero)
  numero++
}

for _, letra := range "Olá, mundo!" {
  fmt.Println(letra)
}
```

**3. Funções:**

* Permitem organizar o código em blocos reutilizáveis.
* Recebem parâmetros e retornam valores.
* Podem ser chamadas de outras partes do programa.

**Exemplo:**

```go
func soma(a int, b int) int {
  return a + b
}

resultado := soma(10, 20)
fmt.Println(resultado)
```

**4. GoRoutines e canais:**

* Permitem executar código concorrentemente.
* GoRoutines são unidades de execução concorrentes.
* Canais são usados para comunicação entre GoRoutines.

**Exemplo:**

```go
go func() {
  fmt.Println("Olá do mundo!")
}()

fmt.Println("Olá do outro lado!")
```

**Recursos para aprofundar:**

* Documentação oficial da linguagem Golang: [https://golang.org/pkg/](https://golang.org/pkg/)
* Curso sobre GoRoutines e canais: [https://gobyexample.com/](https://gobyexample.com/)

