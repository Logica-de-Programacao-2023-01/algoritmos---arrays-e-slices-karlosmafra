package main

import "fmt"

func main() {
	var lista = []string{"T", "E", "R", "M", "I", "N", "A", "L"}
	var x string
	tam := len(lista)
	fmt.Print("Digite uma letra (maiúscula): ")
	fmt.Scan(&x)

	for i := 0; i < tam; i++ {
		if lista[i] == x {
			lista = append(lista[:i], lista[i+1:]...)
			tam--
		}
	}

	fmt.Println("Letra removida.")
	fmt.Println(lista)
}
