fn main() {
  let numbers = vec![1, 2, 3, 4, 5];
  let double_numbers: Vec<i32> = numbers.iter().map(|&x| x*2).collect();

  println!("{:?}", double_numbers);
}
