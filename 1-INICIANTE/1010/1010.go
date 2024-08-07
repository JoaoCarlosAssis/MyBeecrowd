package main

import "fmt"

func main() {
	var codigo1, codigo2, quantidade1, quantidade2 int
	var preco1, preco2 float64

	fmt.Scan(&codigo1, &quantidade1, &preco1)
	fmt.Scan(&codigo2, &quantidade2, &preco2)

	totalpeca1 := float64(quantidade1) * preco1
	totalpeca2 := float64(quantidade2) * preco2

	totalPagar := totalpeca1 + totalpeca2

	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", totalPagar)
}
