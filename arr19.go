package main

import "fmt"

func main() {
	var n1 = [8]int{2, 5, 8, 10, 12, 15, 18, 21}
	var n2 = [8]int{4, 8, 5, 11, 3, 13, 20, 16}
	fmt.Println(n1)
	fmt.Println(n2)
	for i := range n1 {
		n1[i] += n2[i]
	}
	fmt.Println("Soma:")
	fmt.Println(n1)
}
