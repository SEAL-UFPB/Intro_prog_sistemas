# Introdução ao Go(lang)
Com o Go baixado no seu sistema, vamos iniciar o nosso primeiro código

## Passo 1 (criação do projeto)

ao digitar o comando abaixo no terminal, você irá criar seu primeiro projeto usando a linguagem 

```bash

go mod init aula1

```
após você fazer isso você terá um arquivo chamado go.mod, esse arquivo é o nosso gerenciador de dependências. similar ao cargo.toml do rust e ao requirements.txt do python 

## Passo 2 (arquivo principal)

Com o projeto já iniciado , crie um arquivo chamado de main.go, que conterá sua função principal do projeto 
```bash 
touch main.go
```

```go 

package main

func main(){

}

```

## Passo 3 ( Hello world)
em go, utilizamos a biblioteca fmt para printar informações no terminal, para isso utilizamos a função Println dessa biblioteca


```go 

package main

import "fmt"

func main(){
fmt.Println("Hello world")
}

```

para rodar o programa, temos 2 formas principais 

##### compilando e depois rodando 

```bash 
go build main.go 
./main 

```
##### compilando e rodando 

```bash 
go run  main.go 

```

## Passo 4 (variáveis e tipos primitivos)



```go 

package main

import (
"fmt"
"reflect"
	"unsafe"
) // para importarmos mais do que uma biblioteca, utilizamos essa sintaxe 
func main(){
  
  var a int8 = 5 // inteiro 8 bits 
  // em go temos int8,int16,int32,int64 
  var b int64 = -10 //inteiro 64 bits 
  var c float32 = 3.5 //flutuante de precisão simples 
  var d float64 = 7.1 // flutuante de precisão dupla 
  var e uint8 =10 // inteiros positivos de 8 bits 
  // em go temos uint8,uint16,uint32 e uint64
  var f uint64 = 100 // inteiros positivos de 64 bits 
  var g int = -15 //inteiro padrão, depende da arquitetura, 32 bits para arquiteturas x86 e 64 bits para arquitetura x64 
  var h uint = 15 //inteiros positivos padrão, depende da arquitetura, 32 bits para arquiteturas x86 e 64 bits para arquitetura x64 
  var i bool = true // booleano, ou positivo ou negativo 
  var j string = "olá mundo" // sequência imutável de bytes com codificação utf-8
  var k rune = 'á' // o equivalente ao caractere em go, ocupa 32 bits e também é codificado em utf-8 
  var l byte = 'A' // em go byte é um apelido para uint8 , nele você pode representar caracteres ascii ou números positivos que cabem em 8 bits 
  m:=2.1 
  n:= 3.5 // em go podemos utilizar a sintaxe := para fazer inferência de tipos 
  var o complex128 = complex(m,n) // curiosidade, em go temos os tipos complex, complex64 que é formado por 2 float32 e complex128 que é formado por 2 float64
  
  fmt.Printf("%-5s | %-12s | %-10s | %-10s\n", "VAR", "TIPO", "TAMANHO", "VALOR")
	fmt.Println("------------------------------------------------------------")

	// Usando os verbos de formatação:
	// %T -> Exibe o Tipo da variável
	// %v -> Exibe o Valor no formato padrão
	// %d -> Exibe um número inteiro (decimal)
	
	fmt.Printf("a     | %-12T | %-2d bytes  | %v\n", a, unsafe.Sizeof(a), a)
	fmt.Printf("b     | %-12T | %-2d bytes  | %v\n", b, unsafe.Sizeof(b), b)
	fmt.Printf("c     | %-12T | %-2d bytes  | %v\n", c, unsafe.Sizeof(c), c)
	fmt.Printf("d     | %-12T | %-2d bytes  | %v\n", d, unsafe.Sizeof(d), d)
	fmt.Printf("e     | %-12T | %-2d bytes  | %v\n", e, unsafe.Sizeof(e), e)
	fmt.Printf("f     | %-12T | %-2d bytes  | %v\n", f, unsafe.Sizeof(f), f)
	fmt.Printf("g     | %-12T | %-2d bytes  | %v\n", g, unsafe.Sizeof(g), g)
	fmt.Printf("h     | %-12T | %-2d bytes  | %v\n", h, unsafe.Sizeof(h), h)
	fmt.Printf("i     | %-12T | %-2d bytes  | %v\n", i, unsafe.Sizeof(i), i)
	fmt.Printf("j     | %-12T | %-2d bytes  | %v\n", j, unsafe.Sizeof(j), j)
	fmt.Printf("k     | %-12T | %-2d bytes  | %v (Unicode: %U)\n", k, unsafe.Sizeof(k), k, k)
	fmt.Printf("l     | %-12T | %-2d bytes  | %v (ASCII: %c)\n", l, unsafe.Sizeof(l), l, l)
	fmt.Printf("m     | %-12T | %-2d bytes  | %v\n", m, unsafe.Sizeof(m), m)
	fmt.Printf("n     | %-12T | %-2d bytes  | %v\n", n, unsafe.Sizeof(n), n)
	fmt.Printf("o     | %-12T | %-2d bytes  | %v\n", o, unsafe.Sizeof(o), o)
}

```
## Passo 5 (verbos de formatação)

Já utilizamos os verbos de formatação no passo anterior, mas vou me aprofundar neles agora:

- %v (Value): O valor na sua forma padrão. O Go decide como imprimir baseando-se no tipo.
- %+v: Útil para structs, pois imprime o nome dos campos junto com os valores.
- %#v: Mostra o valor em sintaxe Go (útil para debug pesado, pois mostra como recriar aquela variável).
- %T (Type): Imprime o tipo da variável (ex: int8, string, main.Pessoa).
- %d (Decimal): Base 10 (o padrão para humanos).
- %b (Binary): Mostra os bits (0s e 1s). Essencial para ensinar operações bitwise.
- %x / %X (Hexadecimal): Base 16. Muito usado para endereços de memória e cores.
- %o (Octal): Base 8. Comum para mostrar permissões de arquivos no Linux (ex: 0755)
- %f (Float): Notação decimal padrão (ex: 3.141593).
- %.2f: Limita a quantidade de casas decimais.
- %e / %E: Notação científica (ex: 3.141593e+00)
- %s (String): Imprime o texto puro.
- %q (Quoted): Coloca aspas ao redor da string e escapa caracteres especiais (útil para ver se há um \n escondido no final).
- %c (Character): Transforma um número (byte ou rune) na sua representação visual (ex: 65 vira 'A').
- %U (Unicode): Mostra o código no formato U+0041 (para runes)
- %10d: Reserva 10 espaços e alinha à direita.
- %-10d: Reserva 10 espaços e alinha à esquerda.
- %010d: Reserva 10 espaços e preenche os vazios com zeros (ex: 0000000015).
