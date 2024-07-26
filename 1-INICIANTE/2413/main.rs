use std::io;

fn main(){
    
    let mut input = String::new();
    io::stdin().read_line(&mut input).expect("Falha ao ler a linha");

    let number: i32 = match input.trim().parse() {
        Ok(num) => num,
        Err(_) => {
            println!("Não é um número");
            return;
        }
    };
    let total = (number *2)*2;
    println!("{}", total);
}
