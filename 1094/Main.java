import java.util.Locale;
import java.util.Scanner;

public class Main {
  public static void main(String[] args) {
    Locale.setDefault(Locale.US);
    Scanner sc = new Scanner(System.in);
    int totalCobaias = 0;
    int totalCoelhos = 0;
    int totalRatos = 0;
    int totalSapos = 0;

    int N = sc.nextInt();

    for (int i = 0; i < N; i++) {
      String entrada = sc.nextLine().replaceAll("\\s", "");
      int quantidade = Integer.parseInt(entrada.substring(0, entrada.length() - 1));
      char tipo = entrada.charAt(entrada.length() - 1);
      totalCobaias += quantidade;

      switch (tipo) {
        case 'C':
          totalCoelhos += quantidade;
          break;
        case 'S':
          totalSapos += quantidade;
          break;
        case 'R':
          totalRatos += quantidade;
          break;
      }
    }

    System.out.printf("Total: %d cobaias%n", totalCobaias);
    System.out.printf("Total de coelhos: %d%n", totalCoelhos);
    System.out.printf("Total de ratos: %d cobaias%n", totalRatos);
    System.out.printf("Total de sapos: %d cobaias%n", totalSapos);

    

    sc.close();
  }
}
