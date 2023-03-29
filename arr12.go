package main

import "fmt"

func main() {
	var nums = [5]int{2, 6, 8, 12, 14}
	var mult = []int{}
	for i := range nums {
		if nums[i]%3 == 0 {
			mult = append(mult, nums[i])
		}
	}
	fmt.Println("Os números múltiplos de 3 são:", mult)
}
