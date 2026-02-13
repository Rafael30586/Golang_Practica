package main

import "fmt"

func main() {
	var frutas = [...]string{"ananá", "naranja", "frutilla", "manzana", "banana", "ciruela"}
	fmt.Println("Las frutas son: ", frutas)
	fmt.Println("La segunda fruta es:", frutas[1])

	var personajes = []string{"Mario", "Link", "Cloud", "Master chief", "Kirby", "Kratos", "Sheppard", "Marcus"} // Esto es un slice. A diferencia de un array puede cambiar de tamaño
	fmt.Println("El tamaño del slice de personajes es:", len(personajes))
	fmt.Println("La capacidad del slice de personajes es:", cap(personajes))

	var algunasFrutas = frutas[1:4]
	fmt.Println("Algunas frutas son:", algunasFrutas) // Así se crea un slice a partir de un array (no incluye al elemento de la última posición que indicamos).

	edades := make([]int8, 5, 10)
	fmt.Println("El slice es:", edades)
	edades = append(edades, 23, 45, 10) // De esta manera se adicionan elementos a un slice
	fmt.Println("Ahora el slice es:", edades)

	var frutas2 = frutas

	fmt.Println("frutas2 es:", frutas2)
	frutas[0] = "limón"
	fmt.Println("Ahora frutas2 es:", frutas2) // frutas2 sigue inalterado, al parece el lugar en memoria es distinto a frutas

	var frutas3 = &frutas // frutas3 es un puntero que apunta a frutas

	fmt.Println("frutas3 es: ", frutas3)
	frutas[3] = "damasco"
	fmt.Println("Ahora frutas3 es igual a:", frutas3)
	frutas3[1] = "frambuesa" // cuando cambio a frutas3 también cambias frutas
	fmt.Println("Ahota frutas es ", frutas)

}
