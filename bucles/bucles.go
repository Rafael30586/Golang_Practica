package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println("I es igual a:", i)
	}

	fmt.Println("-----------------------------------")

	for i := 4; i < 29; i += 3 {
		if (i % 2) == 0 {
			continue
		}
		if i == 19 {
			break
		}
		fmt.Println("i es igual a :", i)
	}

	fmt.Println("-----------------------------------")

	var numeros = [...]int16{34, 90, 120, 82, 50, 12, 94}

	for indice, valor := range numeros {
		fmt.Println("El índice es", indice, "y el valor es", valor)
	}

}
