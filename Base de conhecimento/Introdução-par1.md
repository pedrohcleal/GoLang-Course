# Introdução ao Go - part1

## Comandos

Em Go, existem vários comandos úteis para compilar, executar, testar e gerenciar projetos. Aqui estão alguns dos comandos mais comuns:

1. **go run**: Este comando é usado para compilar e executar um programa Go em um único passo. Por exemplo, se você tiver um arquivo chamado `meuprograma.go`, você pode executá-lo usando:

   ```
   go run meuprograma.go
   ```

   Isso compilará e executará o arquivo `meuprograma.go`.

2. **go build**: O comando `go build` é usado para compilar um programa Go em um executável. Por exemplo, se você quiser compilar o arquivo `meuprograma.go` em um executável, você pode fazer isso:

   ```
   go build meuprograma.go
   ```

   Isso criará um executável chamado `meuprograma` que você pode executar.

3. **go test**: Este comando é usado para executar testes automatizados em um pacote Go. Ele procura por arquivos com nomes que começam com `*_test.go` no diretório atual e nos subdiretórios e executa os testes definidos nesses arquivos. Por exemplo:

   ```
   go test
   ```

   Isso executará todos os testes no diretório atual.

4. **go mod**: O comando `go mod` é usado para gerenciar módulos Go. Módulos são uma forma de organizar e versionar o código em Go, introduzidos oficialmente a partir do Go 1.11. Alguns subcomandos úteis de `go mod` incluem:

   - `go mod init`: Inicializa um novo módulo Go no diretório atual.
   - `go mod tidy`: Garante que as dependências do módulo estejam consistentes com o código fonte.
   - `go mod vendor`: Copia as dependências do módulo para o diretório `vendor`.

5. **go get**: Este comando é usado para baixar e instalar pacotes Go e suas dependências. Por exemplo:

   ```
   go get github.com/pacote/exemplo
   ```

   Isso baixará e instalará o pacote `exemplo` do GitHub.

6. **go fmt**: O comando `go fmt` é usado para formatar automaticamente o código Go de acordo com as convenções de formatação do Go. Por exemplo:

   ```
   go fmt meuprograma.go
   ```

   Isso formatará o arquivo `meuprograma.go` de acordo com as convenções de formatação do Go.

Esses são apenas alguns dos comandos mais comuns em Go.

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

## Funções

Em Go, uma função é um bloco de código que executa uma tarefa específica. As funções em Go são definidas usando a palavra-chave `func`, seguida pelo nome da função, uma lista de parâmetros entre parênteses e, opcionalmente, um tipo de retorno. Aqui está um exemplo básico de uma função em Go:

```go
package main

import "fmt"

// Função que retorna a soma de dois números inteiros
func soma(a, b int) int {
    return a + b
}

func main() {
    // Chamando a função soma e imprimindo o resultado
    resultado := soma(3, 5)
    fmt.Println("A soma é:", resultado)
}
```

Aqui está uma explicação passo a passo do código acima:

1. `package main`: Define o pacote ao qual este arquivo pertence. O pacote `main` é especial em Go e é usado para criar executáveis.

2. `import "fmt"`: Importa o pacote `fmt`, que é usado para formatar a saída.

3. `func soma(a, b int) int { ... }`: Define uma função chamada `soma` que aceita dois parâmetros inteiros (`a` e `b`) e retorna um resultado do tipo `int`. Dentro do corpo da função, `return a + b` retorna a soma dos dois parâmetros.

4. `func main() { ... }`: Define a função `main`, que é o ponto de entrada para o programa. Dentro desta função, chamamos a função `soma` com os argumentos `3` e `5` e armazenamos o resultado em `resultado`. Em seguida, imprimimos o resultado usando `fmt.Println`.

Este é apenas um exemplo simples de uma função em Go. As funções em Go podem aceitar vários parâmetros, retornar múltiplos valores e até mesmo serem utilizadas como tipos de dados em Go. Elas são uma parte fundamental da linguagem e são usadas extensivamente para organizar e reutilizar código.

## Structs

## Structs em Go: Um guia completo

Structs em Go são tipos de dados compostos que permitem agrupar dados relacionados em uma única estrutura. Elas são semelhantes às classes em outras linguagens de programação.

**Definição de structs:**

```go
type nomeDaStruct struct {
  campo1 tipoDeDado
  campo2 tipoDeDado
  ...
}
```

**Elementos-chave:**

* **Nome da struct:** Identificador único que referencia a estrutura.
* **Campos:** Lista de variáveis com seus tipos de dados que compõem a estrutura.

**Exemplos:**

1. Struct simples para representar uma pessoa:

```go
type Pessoa struct {
  Nome string
  Idade int
}
```

2. Struct com campos aninhados:

```go
type Endereco struct {
  Rua string
  Numero int
  Cidade string
}

type Pessoa struct {
  Nome string
  Idade int
  Endereco Endereco
}
```

**Características importantes:**

* **Tipadas:** Structs em Go são tipadas, o que significa que você precisa especificar o tipo de dado de cada campo.
* **Valores por valor:** Structs são passadas por valor para funções e métodos.
* **Comparação:** Structs podem ser comparadas usando o operador `==`.
* **Literais:** Structs podem ser inicializadas usando literais.
* **Métodos:** É possível definir métodos para structs, que são como funções que operam em um valor específico da struct.

**Recursos para aprofundar:**

* Aprenda Golang: Structs: [https://aprendagolang.com.br/2021/10/15/o-que-sao-structs-e-como-usa-las/](https://aprendagolang.com.br/2021/10/15/o-que-sao-structs-e-como-usa-las/)
* Go por Exemplo: Structs: [URL inválido removido]
* DigitalOcean: Definindo Structs em Go: [https://www.digitalocean.com/community/tutorials/defining-structs-in-go-pt](https://www.digitalocean.com/community/tutorials/defining-structs-in-go-pt)
* Blog da TreinaWeb: Orientação a objetos na linguagem Go: [https://www.treinaweb.com.br/blog/orientacao-a-objetos-na-linguagem-go](https://www.treinaweb.com.br/blog/orientacao-a-objetos-na-linguagem-go)

**Dicas:**

* Use structs para organizar seus dados em unidades menores e mais gerenciáveis.
* Dê nomes descritivos aos seus campos para facilitar a compreensão do código.
* Documente suas structs com comentários para explicar o que cada campo representa.
* Use métodos para encapsular a lógica de manipulação de dados da struct.

Dominar o uso de structs é essencial para escrever código Go eficiente e legível.

**Tópicos avançados:**

* **Ponteiros para structs:** Permitem modificar o valor original da struct em uma função ou compartilhá-la eficientemente entre diferentes partes do código.
* **Composição de structs:** Permite criar structs mais complexas combinando structs menores.
* **Tags em structs:** Permitem adicionar informações extras às structs, como metadados para serialização ou validação.

Explorar esses tópicos avançados pode levar seu conhecimento de structs em Go para o próximo nível.

## Heranças

Em Go, não há conceito direto de herança como em algumas outras linguagens de programação orientada a objetos, como Java ou Python. Em vez disso, Go utiliza um modelo de composição para reutilização de código e compartilhamento de funcionalidades entre tipos.

No entanto, é possível alcançar um comportamento semelhante ao da herança por meio da composição e da incorporação de tipos. Vou explicar esses conceitos:

1. **Composição**: Em vez de herdar comportamentos e campos de uma classe pai, em Go, você pode criar um novo tipo incorporando ou aninhando outros tipos. Isso é chamado de composição. Por exemplo:

```go
package main

import "fmt"

// Definindo um tipo Pessoa
type Pessoa struct {
    Nome string
    Idade int
}

// Definindo um tipo Funcionario que incorpora o tipo Pessoa
type Funcionario struct {
    Pessoa
    Salario float64
}

func main() {
    // Criando uma instância de Funcionario
    funcionario := Funcionario{
        Pessoa: Pessoa{
            Nome:  "João",
            Idade: 30,
        },
        Salario: 3000.00,
    }

    // Acessando campos de Pessoa e Funcionario
    fmt.Println("Nome:", funcionario.Nome)
    fmt.Println("Idade:", funcionario.Idade)
    fmt.Println("Salário:", funcionario.Salario)
}
```

No exemplo acima, `Funcionario` incorpora `Pessoa`, o que significa que todas as propriedades e métodos de `Pessoa` estão disponíveis para `Funcionario`.

2. **Incorporação de Tipos**: Em Go, uma struct pode incorporar outros tipos diretamente. No exemplo anterior, `Funcionario` incorpora `Pessoa`, permitindo que `Funcionario` acesse os campos e métodos de `Pessoa`. Essa incorporação é uma técnica poderosa que permite a reutilização de código sem introduzir complexidade desnecessária.

Embora Go não tenha herança no sentido clássico, o modelo de composição e incorporação de tipos oferece uma maneira eficaz de compartilhar comportamentos e dados entre tipos de uma maneira flexível e eficiente. Isso promove o princípio de design de Go de simplicidade e clareza, evitando a complexidade excessiva muitas vezes associada ao uso extensivo de herança em outras linguagens.
