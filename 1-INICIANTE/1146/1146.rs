use std::io;

fn main() {
    let mut x:i32 = 1;
    
    while x != 0 {
        let mut input = String::new();
        io::stdin()
            .read_line(&mut input)
            .expect("Impossível ler entrada.");

        x = input.trim().parse().expect("Impossivel ler o numero");
    
        for i in 1..=x{
          if i == x{
            println!("{}", i);
          }else{
            print!("{} ", i);
          }
        }
    }
}
