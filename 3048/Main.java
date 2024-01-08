import java.util.Scanner;

public class Main {
  static public void main(String[] args) {
    Scanner sc = new Scanner(System.in);
    int N = sc.nextInt();

    int[] numbers = new int[N];

    for (int i = 0; i < N; i++) {
      numbers[i] = sc.nextInt();
    }

    int count = 0;
    int numberAt = 0;
    int i = 0;
    while (i < numbers.length -1) {
      numberAt = numbers[i];
      if (numberAt != numbers[i+1]) {
        count++;
      }
      i++;
    }
    System.out.println(count+1);
    sc.close();
  }
}