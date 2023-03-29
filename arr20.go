package main

import "fmt"

func main() {
	var nums = []int{}
	var tam int
	var cres = true
	fmt.Print("Tamanho do slice: ")
	fmt.Scan(&tam)

	for i := 0; i < tam; i++ {
		var valor int
		fmt.Print("Valor: ")
		fmt.Scan(&valor)
		nums = append(nums, valor)
	}

	fmt.Println(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			fmt.Println("O slice não está em ordem crescente.")
			cres = false
		}
	}
	if cres == true {
		fmt.Println("O slice está em ordem crescente.")
	}
}
