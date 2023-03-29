package main

import "fmt"

func main() {
	var nums = [10]float64{2.5, 11, 5.1, 7.2, 3.3, 5, 9.8, 4.7, 15.9, 1.62}
	fmt.Println(nums)
	var m5 = []float64{}
	for i := range nums {
		if nums[i] > 5 {
			m5 = append(m5, nums[i])
		}
	}
	fmt.Println("Os números maiores que 5 são:", m5)
}
