package main

import "fmt"

func main() {
	var nums = []int{}
	var tam int
	var soma int
	fmt.Print("Tamanho do slice: ")
	fmt.Scan(&tam)

	for i := 0; i < tam; i++ {
		var valor int
		fmt.Print("Valor: ")
		fmt.Scan(&valor)
		nums = append(nums, valor)
		soma += valor
	}
	fmt.Println(nums)
	fmt.Println("A soma dos valores é igual a", soma)
}
