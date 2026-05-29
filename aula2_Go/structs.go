package main

import (
	"fmt"
	"unsafe"
)

type Pessoa struct {
	Nome  string
	Idade uint8
}

type Professor struct {
	Nome string
	Cpf  string
}

type Disciplina struct {
	Nome string
	Professor
}
type Otimizada struct {
	a bool  // 1 byte
	b bool  // 1 byte
	c int64 // 8 bytes
} // Total: ~16 bytes (com padding)

type NaoOtimizada struct {
	a bool // 1 byte
	// 7 bytes de padding aqui!
	c int64 // 8 bytes
	b bool  // 1 byte
} // Total: ~24 bytes

func main() {
	P1 := Pessoa{Nome: "Mikael", Idade: 19}
	fmt.Println("")

	fmt.Printf("P1     | %-12T | %-2d bytes  | %+v\n", P1, unsafe.Sizeof(P1), P1)
	P2 := new(Pessoa)
	P2.Nome = "Kaio"
	P2.Idade = 19

	fmt.Printf("P2     | %-12T | %-2d bytes  | %+v\n", P2, unsafe.Sizeof(P2), P2)

	calculo := Disciplina{Nome: "Cálculo", Professor: Professor{Nome: "Venegas", Cpf: "000.000.000-00"}}

	fmt.Printf("calculo    | %-12T | %-2d bytes  | %+v\n", calculo, unsafe.Sizeof(calculo), calculo)
	otimo := Otimizada{a: true, b: true, c: 42}
	notimo := NaoOtimizada{a: true, b: true, c: 42}

	fmt.Printf("Otimizada    | %-12T | %-2d bytes  | %+v\n", otimo, unsafe.Sizeof(otimo), otimo)
	fmt.Printf("Não otimizada     | %-12T | %-2d bytes  | %+v\n", notimo, unsafe.Sizeof(notimo), notimo)
}
