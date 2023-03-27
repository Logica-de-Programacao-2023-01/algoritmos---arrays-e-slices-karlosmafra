package main

import "fmt"

func main() {
	var matriz = [3][2]int{}

	for linha := range matriz {
		for coluna := range matriz[linha] {
			var valor int
			fmt.Print("Valor: ")
			fmt.Scan(&valor)
			matriz[linha][coluna] = valor
		}
	}

	fmt.Println(matriz)
}
