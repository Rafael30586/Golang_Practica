package main

import (
	"fmt"
	"operadores_aritmeticos/servicios"
)

func main() {
	suma := 45 + 12
	fmt.Println("La suma es:", suma)

	combinado := 2*23 + (69 - 10)
	fmt.Println("El cálculo combinado es:", combinado)

	edad := 23
	edad++
	fmt.Println("La edad es:", edad)

	fmt.Println(servicios.Sumar(56, 12))

}
