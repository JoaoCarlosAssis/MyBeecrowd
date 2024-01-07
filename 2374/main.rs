using std::io;

fn main(){
    let mut input = String::new();
    io::stdin().read_line(&mut, input).except("Falha ao ler linha");

    let number: i32 = match input.trim().parse(){
        Ok(num)=> num,
        Err(_) => {
            println!("Não é um numero");
            return;
        }
    };
    
    
}
