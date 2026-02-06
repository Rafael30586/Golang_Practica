package main

import (
	"fmt"
	"strconv"
)

func main() {
	Saludar()
	Sumar(16, 9)
	fmt.Println(Presentarse("Gustavo", 27))
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
