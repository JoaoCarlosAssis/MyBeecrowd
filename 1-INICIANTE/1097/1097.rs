fn main() {
  let mut j = 7;
  for i in (1..10).step_by(2) {
      for _ in 0..3 {
          println!("I={} J={}", i, j);
          j -= 1;
      }
      j += 5;
  }
}