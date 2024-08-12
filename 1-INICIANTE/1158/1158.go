package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	results := []int{}

	for i := 0; i < N; i++ {
		var X, Y int
		fmt.Scan(&X, &Y)

		if X%2 == 0 {
			X++
		}

		sum := 0
		for j := 0; j < Y; j++ {
			sum += X
			X += 2
		}
		results = append(results, sum)
	}

	for _, i := range results {
		fmt.Println(i)
	}
}
