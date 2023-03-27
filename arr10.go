package main

import "fmt"

func main() {
	var nums = []int{8, 4, 11, 2, 15, 35, 7, 26, 18, 21}
	fmt.Println(nums)
	maior := nums[0]
	menor := nums[0]

	for i := range nums {
		if nums[i] > maior {
			maior = nums[i]
		}
		if nums[i] < menor {
			menor = nums[i]
		}
	}

	fmt.Println("O maior valor é", maior)
	fmt.Println("O menor valor é", menor)
}
