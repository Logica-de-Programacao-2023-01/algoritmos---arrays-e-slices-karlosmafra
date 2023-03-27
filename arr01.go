package main

import "fmt"

func main() {
	var nums = [3]int{2, 5, 8}
	var soma int
	for i := 0; i < len(nums); i++ {
		soma += nums[i]
	}
	fmt.Println("A soma dos valores no array é igual a", soma)
}
