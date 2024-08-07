package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	var C int = 0
	var R int = 0
	var S int = 0

	for N > 0 {
		var quant int
		var tipo string
		fmt.Scan(&quant, &tipo)

		switch tipo {
		case "C":
			C += quant
		case "R":
			R += quant
		case "S":
			S += quant
		}
		N--
	}
	total_cobaias := C + R + S
	porC := float64(C) / float64(total_cobaias) * 100
	porR := float64(R) / float64(total_cobaias) * 100
	porS := float64(S) / float64(total_cobaias) * 100

	fmt.Printf("Total: %d cobaias\n", total_cobaias)
	fmt.Printf("Total de coelhos: %d\n", C)
	fmt.Printf("Total de ratos: %d\n", R)
	fmt.Printf("Total de sapos: %d\n", S)
	fmt.Printf("Percentual de coelhos: %.2f %%\n", porC)
	fmt.Printf("Percentual de ratos: %.2f %%\n", porR)
	fmt.Printf("Percentual de sapos: %.2f %%\n", porS)

}
