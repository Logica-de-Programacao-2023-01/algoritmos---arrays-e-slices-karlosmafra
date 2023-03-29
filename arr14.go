package main

import "fmt"

func main() {
	var nums = []int{1, 2, 3, 4, 5, 6, 7, 8}
	var i1, i2, n1 int
	fmt.Println(nums)
	fmt.Print("Digite dois índices para serem trocados (entre 0 e 7): ")
	fmt.Scan(&i1, &i2)
	n1 = nums[i1]
	nums[i1] = nums[i2]
	nums[i2] = n1
	fmt.Println(nums)
}
