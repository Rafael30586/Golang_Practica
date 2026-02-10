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
	Sumar_Numeros()
	fmt.Println(MultiplicarPorDos(45, 90))
	var a, b = MultiplicarPorDos(60, 15) // Esta función que retorna dos números les está asignando valores a las dos variables.
	fmt.Print("a y b son", a, b)
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

func Sumar_Numeros() {
	var n1 int
	var n2 int
	fmt.Println("Ingrese el primer número")
	fmt.Scanln(&n1)
	fmt.Println("Ingrese el segundo número")
	fmt.Scanln(&n2)
	fmt.Println("El resultado es:", (n1 + n2))
}

func MultiplicarPorDos(n1 int, n2 int) (int, int) { // Las funcionaes pueden retornar más de un valor, por ejemplo, esta función retorna dos valores...
	return 2 * n1, 2 * n2 //... de tipo entero
}
