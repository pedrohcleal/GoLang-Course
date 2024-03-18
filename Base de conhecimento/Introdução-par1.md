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

### Lógica de Controle:
Go suporta estruturas de controle de fluxo padrão, incluindo:

1. **Condicional If-Else:**
   ```go
   if condição {
       // código a ser executado se a condição for verdadeira
   } else {
       // código a ser executado se a condição for falsa
   }
   ```

2. **Switch:**
   ```go
   switch expressao {
   case valor1:
       // código a ser executado se a expressão for igual a valor1
   case valor2:
       // código a ser executado se a expressão for igual a valor2
   default:
       // código a ser executado se nenhuma das condições for satisfeita
   }
   ```

3. **Loop For:**
   ```go
   for inicialização; condição; pós-iteração {
       // código a ser repetido enquanto a condição for verdadeira
   }
   ```

4. **Loop Range:**
   ```go
   for índice, valor := range coleção {
       // código a ser repetido para cada elemento da coleção
   }
   ```

5. **Loop Infinito:**
   ```go
   for {
       // código a ser executado repetidamente em um loop infinito
   }
   ```

Esses são alguns dos conceitos básicos de tipos de dados, declaração de variáveis e constantes, operadores e lógica de controle em Go. Essas construções são fundamentais para escrever programas eficazes e eficientes em Go.
