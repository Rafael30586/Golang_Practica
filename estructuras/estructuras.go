package main

import "fmt"

func main() {

	var pers1 Persona
	pers1.nombre = "Miguel"
	pers1.edad = 32
	pers1.dni = 45910271
	pers1.altura = 1.86

	pers1.mascota.nombre = "Firulais"
	pers1.mascota.edad = 4
	pers1.mascota.especie = "Perro"

	mostrarPersona(pers1)
}

type Persona struct {
	nombre  string
	edad    int8
	dni     int32
	altura  float32
	mascota Mascota
}

type Mascota struct {
	nombre  string
	especie string
	edad    int8
}

func mostrarPersona(persona Persona) {
	fmt.Println(persona)
}
