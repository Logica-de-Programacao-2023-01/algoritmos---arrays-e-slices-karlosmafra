package main

import "fmt"

func main() {
	var primos = [25]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53,
		59, 61, 67, 71, 73, 79, 83, 89, 97}
	var n int
	fmt.Print("Digite um número inteiro positivo até 25: ")
	fmt.Scan(&n)

	fmt.Println(n, "primeiros números primos:")
	for i := 0; i < n; i++ {
		fmt.Print(primos[i], " ")
	}
}
