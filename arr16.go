package main

import "fmt"

func main() {
	var nums = [10]int{1, 4, 7, 8, 11, 15, 18, 20, 21, 26}
	fmt.Println(nums)
	var par = []int{}
	for i := range nums {
		if nums[i]%2 == 0 {
			par = append(par, nums[i])
		}
	}
	fmt.Println("Os números pares são:", par)
}
