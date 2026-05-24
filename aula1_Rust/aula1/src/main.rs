fn main() 
{
// Mutabilidade
println!("=== Mutabilidade ===\n");

let mut uvas = 100;

println!("Antes de ir ao mercado, há {} uvas em casa.", uvas);

uvas = 8001;

println!("Depois de ir ao mercado, há {} uvas em casa.", uvas);

print!("\n");

println!("=== Sombreamento ===\n");

let caixas_de_chocolate = 10;

println!("Antes de chegar o natal, há {} caixas de chocolate.", caixas_de_chocolate);

let caixas_de_chocolate = 0;

println!("Depois de chegar o natal, há {} caixas de chocolate.", caixas_de_chocolate);

print!("\n");

println!("=== Tipos escalares ===\n");

let vida : i32 = 100; // Inteiro com sinal
let bonus: u8 = 0; // Inteiro sem sinal
let multiplicador: f64 = 1.5; // Padrão float
let ativo: bool = true;     // Booleano
let classe: char = '⚔';    // Char no Rust aceita emoji (4 bytes Unicode)

println!("Jogador ativo: {ativo} | Classe: {classe} | Vida: {vida} | Bonus: {bonus} | Multplicador {multiplicador}\n");

println!("=== Tipos compostos ===\n");

// Tupla: agrupa tipos diferentes (Coordenadas (x,y) e nome da região)
let localizacao: (f64, f64, &str) = (45.2, 108.7, "Main Base");
println!("Local: {} nas coordenadas ({}, {})", localizacao.2, localizacao.0, localizacao.1);

// Array: mesmo tipo, tamanho fixo na Stack (histórico dos últimos 3 danos causados)
let historico_danos: [i32; 3] = [45, 12, 88];
println!("Último dano causado: {} pontos.\n", historico_danos[2]);

// Teste de segurança (Descomente a linha abaixo para ver o compilador/programa dar 'panic')
// let erro = historico_danos[5];

}
