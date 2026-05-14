package main

import (
	"fmt"
	"unsafe"
)

// para importarmos mais do que uma biblioteca, utilizamos essa sintaxe
func main() {
	var a int8 = 5 // inteiro 8 bits
	// em go temos int8,int16,int32,int64
	var b int64 = -10   // inteiro 64 bits
	var c float32 = 3.5 // flutuante de precisão simples
	var d float64 = 7.1 // flutuante de precisão dupla
	var e uint8 = 10    // inteiros positivos de 8 bits
	// em go temos uint8,uint16,uint32 e uint64
	var f uint64 = 100         // inteiros positivos de 64 bits
	var g int = -15            // inteiro padrão, depende da arquitetura, 32 bits para arquiteturas x86 e 64 bits para arquitetura x64
	var h uint = 15            // inteiros positivos padrão, depende da arquitetura, 32 bits para arquiteturas x86 e 64 bits para arquitetura x64
	var i bool = true          // booleano, ou positivo ou negativo
	var j string = "olá mundo" // sequência imutável de bytes com codificação utf-8
	var k rune = 'á'           // o equivalente ao caractere em go, ocupa 32 bits e também é codificado em utf-8
	var l byte = 'A'           // em go byte é um apelido para uint8 , nele você pode representar caracteres ascii ou números positivos que cabem em 8 bits
	m := 2.1
	n := 3.5                         // em go podemos utilizar a sintaxe := para fazer inferência de tipos
	var o complex128 = complex(m, n) // curiosidade, em go temos os tipos complex, complex64 que é formado por 2 float32 e complex128 que é formado por 2 float64

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
