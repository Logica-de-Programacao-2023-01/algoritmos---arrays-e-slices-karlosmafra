package main

import "fmt"

func main() {
	var nums = [4]float64{1.1, 2.5, 5, 10.5}
	var prod float64
	prod = 1
	for i := 0; i < len(nums); i++ {
		prod *= nums[i]
	}
	fmt.Println("O produto do array é", prod)
}
