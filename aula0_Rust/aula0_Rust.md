# Introdução ao Rust

Baseado no capítulo 1 do livro oficial do Rust, atualizado para práticas modernas de 2026.

---

# O que é Rust?

Rust é uma linguagem de programação de sistemas focada em:

* segurança de memória
* performance
* concorrência segura
* previsibilidade
* ferramentas modernas

Ela foi criada para resolver problemas comuns de linguagens como C e C++, principalmente:

* segmentation faults
* data races
* gerenciamento manual de memória
* comportamento indefinido

Hoje Rust é utilizada em:

* sistemas operacionais
* back-end de alta performance
* motores de jogos
* ferramentas CLI
* WebAssembly
* infraestrutura cloud
* sistemas embarcados
* blockchain
* IA e inferência local

Grandes empresas como:

- Microsoft
- Google
- Amazon
- Cloudflare
- Discord

utilizam Rust em produção.

---

# Instalando o Rust

A forma oficial e recomendada atualmente é utilizando o `rustup`.

## Linux / macOS

```bash
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
```

## Windows

Baixe o instalador em:

```txt
https://rustup.rs
```

---

# Verificando a instalação

Após instalar, reinicie o terminal e execute:

```bash
rustc --version
cargo --version
```

Você verá algo parecido com:

```bash
rustc 1.xx.x
cargo 1.xx.x
```

---

# O que é o Cargo?

O Cargo é o gerenciador oficial do ecossistema Rust.

Ele funciona como:

* sistema de build
* gerenciador de dependências
* executor de testes
* executor de projetos
* gerador de documentação

O Cargo é uma das partes mais fortes do ecossistema Rust.

---

# Criando o primeiro projeto

Para criar um projeto:

```bash
cargo new hello_rust
```

Entre na pasta:

```bash
cd hello_rust
```

Estrutura criada:

```txt
hello_rust/
├── Cargo.toml
└── src/
    └── main.rs
```

---

# Cargo.toml

O arquivo `Cargo.toml` define:

* nome do projeto
* versão
* edição do Rust
* dependências
* configurações de build

Exemplo moderno:

```toml
[package]
name = "hello_rust"
version = "0.1.0"
edition = "2024"

[dependencies]
```

> Atualmente a edição 2024 é a mais moderna e recomendada.

---

# Primeiro programa

Abra o arquivo `src/main.rs`.

Código:

```rust
fn main() {
    println!("Hello, world!");
}
```

---

# Entendendo o código

## fn

Define uma função.

```rust
fn
```

---

## main

É a função principal do programa.

Todo programa executável em Rust começa nela.

---

## println!

Macro usada para imprimir texto no terminal.

```rust
println!("Olá");
```

O `!` indica que `println!` é uma macro, não uma função comum.

---

# Compilando e executando

## Executar diretamente

```bash
cargo run
```

O Cargo:

1. compila o projeto
2. executa automaticamente

---

## Apenas compilar

```bash
cargo build
```

---

## Build otimizada (release)

```bash
cargo build --release
```

---

# Atualizando o Rust

Para atualizar o compilador:

```bash
rustup update
```

---

# Documentação oficial

O ecossistema Rust possui uma das melhores documentações da indústria.

Principais recursos:

## Rust Book

Livro oficial:

```txt
https://doc.rust-lang.org/book/
```

---

## Rust By Example

Aprendizado prático:

```txt
https://doc.rust-lang.org/rust-by-example/
```

---

## Docs.rs

Documentação das bibliotecas do ecossistema:

```txt
https://docs.rs
```

---

# Ferramentas modernas recomendadas

## rust-analyzer

Atualmente é a principal ferramenta de suporte para IDEs.

Fornece:

* autocomplete
* análise estática
* navegação de código
* diagnósticos
* refatoração

Funciona muito bem com:

* VS Code
* Zed
* Neovim
* IntelliJ

---

# Formatador oficial

Rust possui um formatador padrão:

```bash
cargo fmt
```

Isso ajuda a manter o código padronizado.

---

# Linter oficial

Rust também possui o Clippy:

```bash
cargo clippy
```

Ele detecta:

* código ruim
* padrões ineficientes
* bugs comuns
* melhorias de estilo

---

# Segurança de memória

O principal diferencial do Rust é o sistema de ownership.

Ele permite:

* evitar vazamentos de memória
* evitar ponteiros inválidos
* evitar data races
* evitar uso após free

Tudo isso sem garbage collector.

Esse assunto será aprofundado posteriormente.

---

# Compilação rigorosa

O compilador Rust é conhecido por ser extremamente rigoroso.

Isso pode parecer difícil no começo, mas reduz muitos bugs em produção.

Uma frase comum da comunidade:

> “Se compila, normalmente funciona.”

---

# Comunidade

Rust possui uma comunidade muito ativa.

Locais importantes:

* GitHub
* Reddit
* Discord
* users.rust-lang.org
* crates.io

---

# crates.io

O `crates.io` é o repositório oficial de bibliotecas Rust.

Equivalente ao:

* npm do JavaScript
* PyPI do Python
* Maven do Java

---

# Primeira dependência

Exemplo adicionando a biblioteca `rand`:

```toml
[dependencies]
rand = "0.9"
```

Depois:

```bash
cargo build
```

O Cargo irá:

* baixar dependências
* resolver versões
* compilar automaticamente

---

# Exemplo simples com biblioteca externa

```rust
use rand::Rng;

fn main() {
    let numero = rand::thread_rng().gen_range(1..=10);

    println!("Número: {}", numero);
}
```

---

# Filosofia do Rust

Rust valoriza:

* segurança explícita
* controle do programador
* abstrações de custo zero
* código previsível
* concorrência segura

A linguagem evita “mágicas escondidas”.

---

# Próximos assuntos importantes

Depois dessa introdução, normalmente os próximos tópicos estudados são:

1. variáveis
2. mutabilidade
3. tipos primitivos
4. funções
5. ownership
6. borrowing
7. structs
8. enums
9. pattern matching
10. traits
11. lifetimes
12. async

---

# Conclusão

Rust se tornou uma das linguagens mais relevantes da computação moderna.

Ela combina:

* performance próxima de C/C++
* segurança forte
* ferramentas modernas
* excelente ecossistema
* ótima documentação

Apesar da curva de aprendizado inicial ser maior que linguagens como Python ou JavaScript, o ganho em confiabilidade costuma compensar em projetos grandes e críticos.
