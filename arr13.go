package main

import "fmt"

func main() {
	var nums = [7]int{1, 2, 3, 4, 5, 6, 7}
	var adc int
	fmt.Println(nums)
	fmt.Print("Valor para adicionar: ")
	fmt.Scan(&adc)
	nums[0] += adc
	nums[6] += adc
	fmt.Println(nums)
}
