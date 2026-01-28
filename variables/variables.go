package main

import (
	"fmt"
)

var dni = 32361730

// dni := 32361730 esta forma de declarar una variable no está permitida fuera de una función
var peso int                                           // Si declaramos una variable sin inicializar debemos aclarar de que tipo es.
var per1, per2, per3 string = "Mario", "Link", "Cloud" // De esta manera se pueden declarar varias variables en una sola línea de código
var coso1, coso2 = 45, true                            // De esta manera se pueden declarar variasd vairables de distintos tipos en un sola línea de código

func main() {
	var edad int = 39
	var altura float32 = 1.82
	var segundoNombre string = "Rafael"
	var primerNombre = "Fernando" // No hace falta aclarar de que tipo es la variable cuando la inicializamos al declararla. El tipo de dato es inferido.
	apellido := "Alvarez"         // Esta forma de declarar variables solo sirve dentro de funciones
	peso = 69
	recibido := true

	var ( // Así se declaran variables en un bloque
		fruta1 = "naranja"
		fruta2 = "manzana"
	)

	fmt.Println("La edad es:", edad, ". La altura es: ", altura)
	fmt.Println("El nombre es ", segundoNombre, ". El apellido es ", apellido)
	fmt.Println("El primer nombre es: ", primerNombre)
	fmt.Println("El DNI es: ", dni)
	fmt.Println("El peso es: ", peso)
	fmt.Println(recibido)
	fmt.Println("Me gusta la", fruta1, "y la", fruta2)

}
