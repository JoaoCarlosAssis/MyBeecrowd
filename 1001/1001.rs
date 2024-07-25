use std::io;

fn main() {
    let mut a = String::new();
    io::stdin().read_line(&mut a).expect("Impossível ler a entrada");
    let a: i32 = a.trim().parse().expect("Digite um número válido");

    let mut b = String::new();
    io::stdin().read_line(&mut b).expect("Impossível ler entrada");
    let b: i32 = b.trim().parse().expect("Digite um numero válido");

    let x = a + b;

    println!("X = {}", x);
}
