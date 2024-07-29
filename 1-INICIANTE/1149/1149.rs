use std::io;

fn main(){
  let mut input = String::new();
  io::stdin().read_line(&mut input).expect("Impossivel ler entrada.");

  let a: Vec<i32> = input.trim().split_whitespace().map(|x| x.parse().expect("Falha ao converter para inteiro")).collect();

  let mut total = 0;

  for x in 0..a[a.len() -1] {
    total += a[0] + x;
  }

  println!("{}", total);
}