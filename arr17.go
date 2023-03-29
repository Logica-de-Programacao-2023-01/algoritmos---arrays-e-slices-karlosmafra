package main

import "fmt"

func main() {
	var nums = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(nums)
	var soma int
	for i := range nums {
		if i%2 == 0 {
			soma += nums[i]
		}
	}
	fmt.Println("A soma dos números em posições pares é", soma)
}
