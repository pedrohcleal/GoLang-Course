# Introdução ao Go - part1

### Tipos de Dados:
Em Go, os tipos de dados incluem tipos básicos, compostos e definidos pelo usuário. Alguns dos tipos básicos incluem:

1. **Inteiros:** `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`.
2. **Ponto flutuante:** `float32`, `float64`.
3. **Booleano:** `bool`.
4. **String:** `string`.
5. **Caractere:** `rune`.
6. **Complexos:** `complex64`, `complex128`.

Além desses, Go também possui tipos compostos, como arrays, slices, maps, structs e tipos de interface.

### Declaração de Variáveis e Constantes:
Em Go, a declaração de variáveis segue a seguinte sintaxe:

```go
var nomeDaVariavel tipo // Declaração de variável
const nomeDaConstante tipo = valor // Declaração de constante
```

Exemplo:

```go
var x int // Declaração de uma variável inteira
const PI float64 = 3.14 // Declaração de uma constante
```

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
