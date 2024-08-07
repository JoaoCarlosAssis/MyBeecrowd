package main

import "fmt"

func main() {
	var A, B float64

	fmt.Scan(&A)
	fmt.Scan(&B)

	A *= 3.5
	B *= 7.5

	media := (A + B) / 11

	fmt.Printf("MEDIA = %.5f\n", media)
}
