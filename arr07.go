package main

import "fmt"

func main() {
	var lista = []int{2, 6, 11, 14, 20}
	var x int
	achou := false
	fmt.Print("Adicionar valor: ")
	fmt.Scan(&x)

	for i := range lista {
		if lista[i] == x {
			fmt.Println("O valor", x, "já se encontra na lista")
			achou = true
		}
	}

	if achou == false {
		lista = append(lista, x)
		fmt.Println("O valor", x, "foi adicionado")
	}

	fmt.Println(lista)
}
