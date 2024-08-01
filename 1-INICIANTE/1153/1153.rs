use std::io;

fn calcular_fatorial(n: i32)-> i32{
  let mut fatorial = 1;
  for i in 1..=n{
    fatorial *= i;
  }
  fatorial
}


fn main (){
  let mut input = String::new();
  io::stdin().read_line(&mut input).expect("Impossivel ler a linha!");
  let n:i32 = input.trim().parse().expect("Não é um número!");

  let resultado = calcular_fatorial(n);

  println!("{}", resultado); 
}