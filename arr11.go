package main

import "fmt"

func main() {
	var matriz = [2][3]int{{2, 4, 6}, {3, 5, 7}}
	var l int
	var c int
	fmt.Print("Linha (entre 0 e 1): ")
	fmt.Scan(&l)
	fmt.Print("Coluna (entre 0 e 2): ")
	fmt.Scan(&c)
	fmt.Println(matriz[l][c])
}
