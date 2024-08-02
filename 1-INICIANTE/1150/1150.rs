use std::io;

fn main() {
    let mut input = String::new();
    io::stdin()
        .read_line(&mut input)
        .expect("Impossivel ler o numero");
    let x: i32 = input.trim().parse().expect("Impossivel converter numero");

    let mut input = String::new();
    io::stdin()
        .read_line(&mut input)
        .expect("Impossivel ler o numero");
    let mut z:i32 = input.trim().parse().expect("Impossivel converter numero");

    while z <= x {
      input.clear();
      io::stdin()
          .read_line(&mut input)
          .expect("Impossivel ler o numero");
      z = input.trim().parse().expect("Impossivel converter numero");
    }

    let mut count = 0;
    let mut soma = x;

    while soma <= z{
      soma += x+1;
      count +=1;
    }    

    println!("{}", count);
}
