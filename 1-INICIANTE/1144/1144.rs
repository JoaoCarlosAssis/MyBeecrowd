use std::io;

fn main(){
  let mut entrada = String::new();
  io::stdin().read_line(&mut entrada).expect("Impossível ler entrada.");
  let entrada: i32 = entrada.trim().parse().expect("Não é um número.");


  for i in 1..=entrada{ 
    let x = i * i;
    let y = i * x;
    println!("{} {} {}", i, x, y);
    println!("{} {} {}", i, x+1, y+1);
  }

}