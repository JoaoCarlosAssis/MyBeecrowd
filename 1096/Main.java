public class Main {
  public static void main(String[] args) {
    int I = 1;
    int J = 7;
    for (int i = 0; i < 5; i++) {
      for (int k = 0; k < 3; k++) {
        System.out.printf("I=%d J=%d%n", I, J);
        J--;
      }
      I += 2;
      J = 7;
    }
  }
}