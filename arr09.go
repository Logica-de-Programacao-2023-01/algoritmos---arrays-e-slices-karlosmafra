package main

import "fmt"

func main() {
	var nums = [6]float64{0.4, 2.55, 4.7, 5.5, 10.2, 21.6}
	var adc float64
	fmt.Print("Valor para adicionar: ")
	fmt.Scan(&adc)

	for i := range nums {
		nums[i] += adc
	}

	fmt.Println("Array resultante:")
	fmt.Println(nums)
}
