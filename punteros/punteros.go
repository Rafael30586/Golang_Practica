package main

import "fmt"

func main() {
	var a = 45
	var s *int = &a

	fmt.Println("s es igual a:", s)
	fmt.Println("&s es igual a:", &s)

	var b = 5012
	var t = &b // Esta variable t también es un puntero. El lenguaje infiere que lo es.
	fmt.Println("t es igual a:", t)
	fmt.Println("&t es igual a:", &t)
	fmt.Println(" apunta a:", *t) // En este caso usamos el asterisco para saber el valor al que apunta el puntero.

	*t = 1020 // Así cambiamos de manera indirecta el valor de b
	fmt.Println("Ahora b vale:", b)

}
