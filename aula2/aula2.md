# Tipos estruturados

## Structs 
Assim como em c, no go nós temos structs e aqui elas são protagonistas. 

para definir uma struct em go podemos fazer assim: 

```go 
type Pessoa struct{
  Nome string 
  Idade uint8
}
```
para instanciarmos essa mesma struct utilizamos essa sintaxe 

```go 
func foo() Pessoa{
  P1:= Pessoa{Nome:"Mikael",Idade:19} // para receber o valor da struct 
  P2:= new(Pessoa) // para receber um ponteiro 
  P2.Nome="Mikael"
  P2.Idade=19
}
```

também podemos definir structs anônimas, muito utilizadas para lidar com respostas de api  
```go 
pessoa:= struct{
  Nome string 
  Idade uint8 
}{"Mikael",19}
```

também temos composição de structs, onde você embute uma struct na outra 

```go
type Professor struct{
  Nome string
  Cpf string
}
type Disciplina struct{
  Nome string
  Professor // podemos anonimizar o nome do campo quando utilizamos composição
}

calculo:=Disciplina{Nome:"Cálculo",Professor:Professor{Nome:"Venegas",Cpf:"000.000.000-00"}}
```

##### alinhamento de memória

Em go, temos que levar em consideração a ordem em que definimos os campos da nossa struct, uma vez que o compilador vai tender a agrupar os campos em grupos de 8 em 8 bytes, então a ordem importa (isso acontece em outras linguagens também) 

```go 

type Otimizada struct {
    a bool  // 1 byte
    b bool  // 1 byte
    c int64 // 8 bytes
} // Total: ~16 bytes (com padding)

type NaoOtimizada struct {
    a bool  // 1 byte
            // 7 bytes de padding aqui!
    c int64 // 8 bytes
    b bool  // 1 byte
} // Total: ~24 bytes
```


