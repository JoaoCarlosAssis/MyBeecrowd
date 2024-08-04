fn main(){
  let mut s: f64 = 1.0;
  let mut x: f64 = 3.0;
  let mut y: f64 = 2.0;

  while x <=39 as f64{
    s += x/y;
    x += 2.0;
    y *= 2.0;
  }
  println!("{:.2}", s);

}