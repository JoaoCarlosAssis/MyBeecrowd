fn main() {
  let mut s: f64 = 0.0;

  for i in 1..=100 {
      s += 1.0 / i as f64;
  }

  println!("{:.2}", s);
}
