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
    sc.nextLine();

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
    double pRatos = (double) totalRatos / totalCobaias * 100;
    double pSapos = (double) totalSapos / totalCobaias * 100;
    double pCoelhos = (double) totalCoelhos / totalCobaias * 100;

    System.out.printf("Total: %d cobaias%n", totalCobaias);
    System.out.printf("Total de coelhos: %d%n", totalCoelhos);
    System.out.printf("Total de ratos: %d%n", totalRatos);
    System.out.printf("Total de sapos: %d%n", totalSapos);
    System.out.printf("Percentual de coelhos: %.2f %%%n", pCoelhos);
    System.out.printf("Percentual de ratos: %.2f %%%n", pRatos);
    System.out.printf("Percentual de sapos: %.2f %%%n", pSapos);

    sc.close();
  }
}
