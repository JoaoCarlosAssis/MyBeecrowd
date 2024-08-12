package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bufio_teste() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite a quantidade de casos de teste: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	N, _ := strconv.Atoi(input)

	for i := 0; i < N; i++ {
		fmt.Print("Digite os valores de X e Y: ")
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)
		values := strings.Split(input, " ")

		X, _ := strconv.Atoi(values[0])
		Y, _ := strconv.Atoi(values[1])

		if X%2 == 0 {
			X++ // Se X for par, comece do próximo ímpar
		}

		sum := 0
		for j := 0; j < Y; j++ {
			sum += X
			X += 2 // Pular para o próximo ímpar
		}

		fmt.Printf("A soma é: %d\n", sum)
	}
}
