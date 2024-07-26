use std::io;

fn main() {
    let mut a = String::new();
    io::stdin().read_line(&mut a).expect("Impossível ler a entrada.");
    let a: f64 = a.trim().parse().expect("Digite um número válido");
    let mut b = String::new();
    io::stdin().read_line(&mut b).expect("Impossível ler a entrada.");
    let b: f64 = b.trim().parse().expect("Digite um número válido.");

    let a = a * 3.5;
    let b = b * 7.5;

    let media:f64 = (a + b) / 11.0;

    println!("MEDIA = {:.5}", media);
}