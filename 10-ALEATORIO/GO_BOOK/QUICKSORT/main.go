//Este exemplo é uma implementação simples do famoso algoritimo de ordenação quicksort. A ideia básica deste algoritmo é:
//Eleger um elemento da lista como pivô e removê-lo da lista.
//Particionar a lista em duas listas distintas: uma contendo elementos menores que o pivô, e a outra sontendo elementos menores.
//Ordenar duas listas recursivamente.
//Retornar a combinação da lista ordenada de elementos menores, o próprio pivô, e a lista ordenada de elementos maiores.

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	entrada := os.Args[1:]
	numeros := make([]int, len(entrada))

	for i, n := range entrada {
		numero, err := strconv.Atoi(n)
		if err != nil {
			fmt.Printf("%s não é um número válido!\n", n)
			os.Exit(1)
		}
		numeros[i] = numero
	}
	fmt.Println(quicksort(numeros))
}

func quicksort(numeros []int) []int {
	//condição de parada para função recursiva
	//previne que a função seja executada eternamente ou até a sobrecarga da pilha de execução// STACK OVERFLOW.
	if len(numeros) <= 1 {
		return numeros
	}

	n := make([]int, len(numeros))
	copy(n, numeros)

	indicePivo := len(n) / 2
	pivo := n[indicePivo]

	n = append(n[:indicePivo], n[indicePivo+1:]...) //... informa que todos os elementos do slice precisam ser adicionados

	menores, maiores := particionar(n, pivo)

	return append(append(quicksort(menores), pivo), quicksort(maiores)...)
}

func particionar(numeros []int, pivo int) (menores []int, maiores []int) {
	for _, n := range numeros {
		if n <= pivo {
			menores = append(menores, n)
		} else {
			maiores = append(maiores, n)
		}
	}
	return menores, maiores
}
