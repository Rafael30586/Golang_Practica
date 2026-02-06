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

	for indice, valor := range numeros { // Así se hacen los bucles for each en go
		fmt.Println("El índice es", indice, "y el valor es", valor)
	}

	fmt.Println("-----------------------------------")

	var iterador = 0

	for iterador < 5 { // Esto es lo más parecido a un bucle while en Go
		fmt.Println("El iterador vale", iterador)
		iterador++
	}

}
