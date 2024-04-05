# Coleções em Go

## Arrays

Em Go (ou Golang), os arrays são estruturas de dados fundamentais que armazenam uma coleção fixa de elementos do mesmo tipo. Aqui está uma visão geral das características dos arrays em Go:

1. **Tamanho Fixo**: Um array em Go tem um tamanho fixo definido no momento da sua declaração e esse tamanho não pode ser alterado posteriormente. Por exemplo, se você declara um array de tamanho 5, ele sempre terá exatamente 5 elementos.

2. **Tipo Homogêneo**: Todos os elementos de um array devem ser do mesmo tipo. Isso significa que você não pode ter um array que armazena, por exemplo, inteiros e strings ao mesmo tempo. Todos os elementos devem ser do mesmo tipo de dados.

3. **Índices Baseados em Zero**: Os índices de arrays em Go começam a partir de 0. Isso significa que o primeiro elemento está no índice 0, o segundo no índice 1, e assim por diante.

4. **Declaração de Array**: Você pode declarar um array em Go usando a seguinte sintaxe:
   ```go
   var nomeDoArray [tamanho]tipo
   ```

   Por exemplo:
   ```go
   var meusNumeros [5]int
   ```

5. **Inicialização**: Os arrays podem ser inicializados durante a declaração:
   ```go
   var nomes = [3]string{"Alice", "Bob", "Charlie"}
   ```

   Ou você pode deixar o compilador contar o tamanho automaticamente:
   ```go
   nomes := [...]string{"Alice", "Bob", "Charlie"}
   ```

6. **Acesso aos Elementos**: Você pode acessar elementos individuais do array usando colchetes e o índice desejado:
   ```go
   fmt.Println(meusNumeros[0]) // Acessando o primeiro elemento do array
   ```

7. **Iteração sobre Arrays**: Você pode iterar sobre os elementos de um array usando um loop `for`:
   ```go
   for i := 0; i < len(meusNumeros); i++ {
       fmt.Println(meusNumeros[i])
   }
   ```

8. **Arrays como Argumentos de Funções**: Em Go, quando você passa um array para uma função, você está passando uma cópia do array. Isso pode ser ineficiente para arrays grandes. No entanto, você também pode passar ponteiros para arrays para evitar a cópia desnecessária.

9. **Slices**: Embora os arrays sejam úteis em algumas situações, os slices são usados com mais frequência em Go. Os slices são estruturas de dados dinâmicas construídas sobre arrays que fornecem mais flexibilidade e funcionalidades extras.

Os arrays são uma parte fundamental da linguagem Go e são usados para uma variedade de finalidades, desde armazenar pequenas coleções de dados até representar estruturas de dados mais complexas.

## Métodos

Em Go, os arrays são tipos de dados simples e, portanto, não têm métodos associados a eles como em algumas outras linguagens de programação orientadas a objetos. No entanto, Go fornece funções embutidas (built-in functions) e sintaxe que permite trabalhar com arrays de maneira eficaz. Aqui estão algumas maneiras comuns de trabalhar com arrays em Go:

1. **Acesso a elementos**: Você pode acessar elementos individuais de um array usando colchetes e o índice desejado.
   ```go
   var meuArray [5]int
   meuArray[0] = 10
   primeiroElemento := meuArray[0]
   ```

2. **Iteração**: Você pode iterar sobre os elementos de um array usando um loop `for`.
   ```go
   for i := 0; i < len(meuArray); i++ {
       fmt.Println(meuArray[i])
   }
   ```

3. **Comprimento do array**: Você pode obter o comprimento de um array usando a função `len`.
   ```go
   tamanho := len(meuArray)
   ```

4. **Comparação**: Você pode comparar dois arrays usando o operador `==`. Os arrays são considerados iguais se tiverem o mesmo tipo e cada elemento correspondente for igual.
   ```go
   array1 := [3]int{1, 2, 3}
   array2 := [3]int{1, 2, 3}
   if array1 == array2 {
       fmt.Println("Os arrays são iguais")
   }
   ```

5. **Copiar**: Você pode copiar elementos de um array para outro usando a função `copy`.
   ```go
   var outroArray [5]int
   copy(outroArray, meuArray) // Copia os primeiros 5 elementos de meuArray para outroArray
   ```

Embora os arrays em Go não tenham métodos associados diretamente, a combinação de funcionalidades embutidas e sintaxe da linguagem oferece uma maneira eficiente de manipular e trabalhar com arrays em seus programas. Além disso, é importante notar que slices, uma estrutura de dados relacionada, fornecem mais funcionalidades e flexibilidade do que arrays em muitos casos de uso, especialmente para operações dinâmicas como adicionar ou remover elementos.

## SLices

Em Go, slices são estruturas de dados dinâmicas construídas sobre arrays que fornecem uma maneira flexível e poderosa de trabalhar com sequências de elementos. Enquanto os arrays têm um tamanho fixo definido no momento da sua declaração, slices permitem uma abordagem mais dinâmica, onde o tamanho do slice pode mudar durante a execução do programa. Aqui está uma descrição detalhada sobre slices em Go:

1. **Declaração de Slices**: Em vez de especificar um tamanho fixo, a declaração de um slice em Go é feita fornecendo uma parte de um array existente, ou usando a função embutida `make` para criar um novo slice.
   ```go
   // Usando parte de um array existente
   var slice1 []int = meuArray[1:4] // Contém elementos do índice 1 até o índice 3 (exclusivo) de meuArray

   // Criando um novo slice usando make
   slice2 := make([]int, 5) // Cria um slice com capacidade para 5 elementos, inicializados com zero
   ```

2. **Índices de Slices**: Um slice consiste em três componentes: um ponteiro para o array subjacente, um comprimento e uma capacidade. O comprimento é o número de elementos no slice, e a capacidade é o número de elementos no array subjacente a partir do primeiro elemento no slice.

3. **Modificação de Slices**: Slices podem ser facilmente modificados adicionando ou removendo elementos.
   ```go
   // Adicionando um elemento a um slice
   slice = append(slice, novoElemento)

   // Removendo um elemento de um slice
   slice = append(slice[:indice], slice[indice+1:]...)
   ```

4. **Acesso a Elementos**: Elementos de um slice são acessados da mesma forma que em um array, usando índices.
   ```go
   primeiroElemento := slice[0]
   ```

5. **Iteração sobre Slices**: Você pode iterar sobre os elementos de um slice usando um loop `for` ou a função `range`.
   ```go
   for indice, valor := range slice {
       fmt.Println(indice, valor)
   }
   ```

6. **Fatia Vazia**: Um slice vazio tem comprimento e capacidade zero e é representado por `nil`.
   ```go
   var sliceVazio []int
   if sliceVazio == nil {
       fmt.Println("O slice está vazio")
   }
   ```

7. **Cópia de Slices**: A função `copy` pode ser usada para copiar elementos de um slice para outro.
   ```go
   sliceDestino := make([]int, len(sliceOrigem))
   copy(sliceDestino, sliceOrigem)
   ```

Slices são uma parte fundamental da linguagem Go e são amplamente utilizados em muitas situações, especialmente quando se trata de manipulação de coleções de dados dinâmicas. Eles fornecem uma maneira eficiente e conveniente de trabalhar com sequências de elementos de tamanho variável.

## Métodos com slices

Em Go, os slices são estruturas de dados embutidas que não possuem métodos diretamente associados a eles. No entanto, Go fornece uma série de funções embutidas (built-in functions) que são comumente usadas para manipular e trabalhar com slices de maneira eficaz. Aqui estão algumas das funções mais úteis para trabalhar com slices:

1. **append**: Adiciona elementos a um slice. Se necessário, o slice será redimensionado automaticamente para acomodar os novos elementos.
   ```go
   slice = append(slice, elemento1, elemento2, ...)
   ```

2. **copy**: Copia elementos de um slice para outro. Retorna o número de elementos copiados, que é o mínimo entre o comprimento dos dois slices.
   ```go
   quantidadeCopiada := copy(destino, origem)
   ```

3. **len**: Retorna o comprimento de um slice, ou seja, o número de elementos no slice.
   ```go
   tamanho := len(slice)
   ```

4. **cap**: Retorna a capacidade de um slice, ou seja, o número de elementos que o slice pode conter antes de precisar de realocação de memória.
   ```go
   capacidade := cap(slice)
   ```

5. **make**: Cria um novo slice com um comprimento e capacidade específicos.
   ```go
   novoSlice := make([]T, comprimento, capacidade)
   ```

6. **append**: Adiciona elementos a um slice.
   ```go
   slice = append(slice, elemento1, elemento2, ...)
   ```

7. **delete**: Não existe um método nativo para deletar um elemento de um slice. Em vez disso, você pode usar uma combinação de slicing para remover um elemento de um slice.
   ```go
   slice = append(slice[:indice], slice[indice+1:]...)
   ```

8. **range**: Usado para iterar sobre os elementos de um slice. Retorna o índice e o valor de cada elemento em cada iteração.
   ```go
   for indice, valor := range slice {
       // Faça algo com o índice e o valor
   }
   ```

Embora não haja métodos nativos associados diretamente a slices em Go, essas funções embutidas fornecem uma maneira poderosa de manipular e trabalhar com slices em seus programas. Elas são fundamentais para a eficiência e conveniência ao lidar com coleções de dados dinâmicas em Go.

## make()

Em Go, o comando `make()` é usado principalmente para criar instâncias de tipos de dados que requerem inicialização especial, como slices, maps e channels. A sintaxe básica para `make()` é:

```go
make(tipo, tamanho, capacidade)
```

Onde:
- `tipo`: o tipo de dado que você deseja criar. Pode ser `slice`, `map` ou `channel`.
- `tamanho`: para slices e arrays, é o comprimento inicial do slice (ou array). Para maps, é uma estimativa do número de elementos que o map pode conter inicialmente (embora o map possa crescer automaticamente conforme necessário).
- `capacidade`: opcional, é usado apenas para slices e especifica a capacidade inicial do slice. Isso geralmente é usado para pré-alocar memória para o slice, o que pode melhorar o desempenho se você souber aproximadamente quantos elementos o slice conterá.

Exemplos de uso do `make()`:

1. Criando um slice:
   ```go
   novoSlice := make([]int, 5, 10) // Cria um slice de inteiros com comprimento inicial de 5 e capacidade inicial de 10
   ```

2. Criando um map:
   ```go
   novoMap := make(map[string]int) // Cria um map com chaves do tipo string e valores do tipo int
   ```

3. Criando um channel:
   ```go
   novoChannel := make(chan int) // Cria um channel para comunicação entre goroutines com valores inteiros
   ```

Em resumo, `make()` é uma função útil em Go para criar e inicializar tipos de dados especiais, como slices, maps e channels, com valores iniciais específicos. Isso é útil especialmente quando você precisa de estruturas de dados dinâmicas que devem ser inicializadas de uma maneira específica.

## map

Em Go, maps são uma estrutura de dados que fornecem uma coleção não ordenada de pares chave-valor, onde cada chave é única e associada a um único valor. Maps são implementados como uma tabela de dispersão (hash table), o que permite uma busca eficiente de valores com base em suas chaves. Aqui estão algumas características importantes sobre maps em Go:

1. **Declaração de um Map**: Você pode declarar um map em Go usando a seguinte sintaxe:
   ```go
   var nomeDoMap map[TipoDaChave]TipoDoValor
   ```

   Por exemplo:
   ```go
   var meuMap map[string]int
   ```

2. **Inicialização**: Maps são tipos de dados de referência em Go, o que significa que eles são inicializados com `nil` se não forem explicitamente inicializados. Você pode inicializar um map usando a função embutida `make`:
   ```go
   meuMap := make(map[string]int)
   ```

   Você também pode inicializar um map com valores iniciais durante a declaração:
   ```go
   outroMap := map[string]string{"a": "apple", "b": "banana", "c": "cherry"}
   ```

3. **Acesso a Elementos**: Você pode acessar elementos de um map usando sua chave entre colchetes:
   ```go
   valor := meuMap[chave]
   ```

4. **Inserção e Atualização de Elementos**: Para inserir ou atualizar um elemento em um map, basta atribuir um valor a uma chave:
   ```go
   meuMap[chave] = valor
   ```

5. **Remoção de Elementos**: Você pode remover um elemento de um map usando a palavra-chave `delete`:
   ```go
   delete(meuMap, chave)
   ```

6. **Verificação de Existência**: Você pode verificar se uma chave existe em um map usando uma atribuição múltipla:
   ```go
   valor, existe := meuMap[chave]
   ```

   O segundo valor retornado (`existe`) será `true` se a chave existir no map e `false` caso contrário.

7. **Iteração sobre Maps**: Você pode iterar sobre todos os pares chave-valor em um map usando o loop `for` ou a função `range`:
   ```go
   for chave, valor := range meuMap {
       fmt.Println(chave, valor)
   }
   ```

8. **Comprimento do Map**: Você pode obter o número de pares chave-valor em um map usando a função embutida `len`:
   ```go
   tamanho := len(meuMap)
   ```

Maps são uma estrutura de dados muito útil em Go e são amplamente utilizados para mapear chaves para valores em uma variedade de situações, como armazenamento de dados, contagem de ocorrências, cacheamento e muito mais. Eles oferecem uma maneira eficiente de associar informações e realizar pesquisas rápidas com base em chaves únicas.

## métodos em map


