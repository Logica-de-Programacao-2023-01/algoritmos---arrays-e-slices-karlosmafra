package main

import "fmt"

func main() {
	var lista = [10]int{0, 1, 4, 5, 7, 8, 11, 15, 21, 25}
	var x int
	var achou int
	fmt.Print("Procurar valor: ")
	fmt.Scan(&x)

	for i := 0; i < 10; i++ {
		if lista[i] == x {
			fmt.Println("O valor", x, "foi encontrado na lista.")
			achou++
		}
	}
	if achou < 1 {
		fmt.Println("O valor", x, "não foi encontrado na lista.")
	}
}
