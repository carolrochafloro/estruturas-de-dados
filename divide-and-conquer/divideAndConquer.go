/* DC: encontrar o caso base (mais simples possível). Nesse caso, realizar a soma de todos os
itens de um array utilizando uma função recursiva. */

package main

import "fmt"

var arrayToSum = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func sumDivideAndConquer(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	return arr[0] + sumDivideAndConquer(arr[1:])
}



func main() {
	result := sumDivideAndConquer(arrayToSum)

	fmt.Printf("A soma dos elementos do array é: %d\n", result)
}