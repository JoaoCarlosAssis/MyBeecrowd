public class Main {
  public static void main(String[] args){
    int I = 1;
    int J = 60;
    for (int i = 0; i < 13; i++) {
      System.out.printf("I=%d J=%d%n", I, J);
      J -=5;
      I +=3;
    }
  }
}