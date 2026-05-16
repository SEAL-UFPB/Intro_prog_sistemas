# Funções 

em go a sintaxe de funções seguem o padrão a seguir 

```go 
func Somar(a int, b int) int {
    return a + b
}
```

também podemos abreviar essa definição

```go 
func Somar(a, b int) int {
    return a + b
}
```
o que significa a mesma coisa da função anterior 

> [!IMPORTANT] Isso só funciona com parâmetros de mesmo tipo 


## Múltiplos retornos 
uma das principais características de go é a possibilidade de Múltiplos retornos ,muito utilizado para o tratamento de erros em go, assunto que abordaremos em aulas futuras 

```go 
func Dividir(a, b float64) (float64, string) {
    if b == 0 {
        return 0, "Impossível dividir por zero")
    }
    return a / b, ""
}
```

> [!IMPORTANT] Essa não é a forma que realmente utilizamos em go para lidar com erros, mas por enquanto, isso vai bastar

## cidadãs de primeira classe
As funções em go são tradadas como cidadãs de primeira classe, isso é, podemos tratar elas como qualquer outra variável.

isso implica que :
- podemos atribuir uma função à uma variável 
  - a:= foo()
- podemos passar uma função como argumento para outras funções 
  - foo(a) com a sendo uma função 
- retornar como valor,
  - foo() func(int)
- armazenar em estruturas 

um bom exemplo disso é o código a seguir 


```go 

package main

import "fmt"

// Define um tipo de função que recebe dois inteiros e retorna um inteiro
type OperacaoMatematica func(int, int) int

// Função que recebe outra função como argumento
func executarOperacao(op OperacaoMatematica, a int, b int) int {
    return op(a, b)
}

func main() {
    // 1. Atribuindo uma função anônima a uma variável
    somar := func(x int, y int) int {
        return x + y
    }

    // 2. Passando a variável de função como argumento
    resultado := executarOperacao(somar, 10, 5)
    fmt.Println("Resultado da soma:", resultado) // Saída: 15
}
```

na qual definimos um tipo OperacaoMatematica que é definido como uma função que recebe 2 inteiros e retorna um inteiro e podemos instanciar outras funções com esse tipo, como a função somar, além de que ainda podemos passar ela como um argumento da função executarOperacao


