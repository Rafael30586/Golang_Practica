package main

import (
	"fmt"
	"reflect"
)

func main() {

	var pers1 Persona
	pers1.nombre = "Miguel"
	pers1.edad = 32
	pers1.dni = 45910271
	pers1.altura = 1.86

	pers1.mascota.nombre = "Firulais"
	pers1.mascota.edad = 4
	pers1.mascota.especie = "Perro"
	pers1.mascota.ruido = "guau guau"
	pers1.mascota.hacer_ruido = func(ruido string) string { // Con una función anónima podemos inicializar la función de la estructura
		return ruido
	}

	mostrarPersona(pers1)

	var juego1 VideoJuego
	juego1.nombre = "Final fantasy XII"
	juego1.desarrollador = "Square enix"

	var juego2 VideoJuego
	juego2.nombre = "Final fantasy XII"
	juego2.desarrollador = "Square enix"

	fmt.Println("¿Los dos video juegos son iguales?", reflect.DeepEqual(juego1, juego2)) // Esta es una forma de comparar dos estructuras
	fmt.Println("¿Los dos video juegos son iguales?", (juego1 == juego2))                // Esta forma tambiéen funciona

	var consola1 = struct { // Así se hace una estructura anónima, es una estructura que se hace para usar en el momento.
		nombre           string
		anio_lanzamiento int16
	}{"Nintendo Switch", 2017}

	fmt.Println("La consola es:", consola1)

}

type Persona struct {
	nombre  string
	edad    int8
	dni     int32
	altura  float32
	mascota Mascota
}

type Mascota struct {
	nombre      string
	especie     string
	edad        int8
	ruido       string
	hacer_ruido func(ruido string) string // De esta manera podemos colocar funciones dentro de una estructura
}

type VideoJuego struct {
	nombre        string
	desarrollador string
}

func mostrarPersona(persona Persona) {
	fmt.Println(persona)
}
