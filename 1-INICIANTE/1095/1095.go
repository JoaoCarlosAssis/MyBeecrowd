package main

import "fmt"

func main() {
	var I int = 1
	var J int = 60

	for i := 0; i < 13; i++ {
		fmt.Printf("I=%d J=%d\n", I, J)
		J -= 5
		I += 3
	}
}
