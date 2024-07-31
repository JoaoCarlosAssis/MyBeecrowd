use std::io;

fn main(){
  let mut input = String::new();
  io::stdin().read_line(&mut input).expect("Impossível ler entrada.");
  let n: usize = input.trim().parse().expect("Digite um número válido.");

  let mut numeros = vec![0, 1];

  while numeros.len() < n{
    let next_numero = numeros[numeros.len() -1] + numeros[numeros.len() - 2];
    numeros.push(next_numero);
  }

  for (i, numero) in numeros.iter().enumerate(){
    if i == numeros.len() - 1{
      println!("{}", numero);
    }else{
      print!("{} ", numero);
    }
  }
  
}