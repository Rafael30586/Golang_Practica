package main

import (
	"fmt"
	"strconv"
)

func main() {
	Saludar()
	Sumar(16, 9)
	fmt.Println(Presentarse("Gustavo", 27))
	NombrarPersonajes("Master Chief", "Kratos", "Kirby", "Sonic", "Link")
	var Gato = func() { // Esta es una función anónima siendo asignada a una variable
		fmt.Println("¡miau! ¡miau!")
	}
	Gato()
}

func Saludar() {
	fmt.Println("Hola")
}

func Sumar(a int, b int) {
	fmt.Println("La suma es:", (a + b))
}

func Presentarse(nombre string, edad int) string {
	presentacion := "Me llamo " + nombre + " y tengo " + strconv.Itoa(edad) + " años de edad."
	return presentacion
}

func NombrarPersonajes(personajes ...string) { //Esta es una función variádica. Se les llama así a las funciones que tienen estos parámetros
	longitud := len(personajes)

	for i := 0; i < longitud; i++ {
		fmt.Println(personajes[i])
	}

}

func init() { // Esta es la función init, no tiene parámetros ni retorna nada. Está presente en cada paquete y es llamada cuando el paquete es...
	fmt.Println("Ya viene...") //... inicializado. Se ejecuta antes de la función main y su propósito es el de inicializar variables globales.
} // Cada paquete puede tener más de una función init.
