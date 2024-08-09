package main

import "fmt"

func main() {
	j := 7

	for i := 1; i < 10; i += 2 {
		for k := 0; k < 3; k++ {
			fmt.Printf("I=%d J=%d\n", i, j)
			j--
		}
		j += 5
	}

}
