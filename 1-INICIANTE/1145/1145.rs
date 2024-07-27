use std::io;

fn main(){
  let mut values = String::new();
  io::stdin().read_line(&mut values).expect("Impossivel ler o valor.");
  let values = values.trim();

  //entendendo as diferenças entre split(" ") e split_whitespace()
  let parts: Vec<&str> = values.split_whitespace().collect();

  let x: i32 = parts[0].parse().expect("Impossivel ler o numero: ");
  let y: i32 = parts[1].parse().expect("Impossivel ler o numero: ");

  for i in 1..=y{
    if i % x == 0{
      println!("{}", i);
    }else{
      print!("{} ", i);
    };
  }
}