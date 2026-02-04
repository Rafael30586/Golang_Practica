package main

import "fmt"

func main() {
	var edad = 27
	var nombre = "Rafael"
	var clase = "A"

	if edad >= 18 {
		fmt.Println("Es mayor de edad")
	} else {
		fmt.Println("Es menor de edad")
	}

	if len(nombre) <= 4 {
		fmt.Println("Tiene pocas letras")
	} else if len(nombre) <= 10 && len(nombre) > 4 {
		fmt.Println("Tiene varias letras")
	} else {
		fmt.Println("Tiene muchas letras")
	}

	switch clase {
	case "A":
		fmt.Println("Es de primera")
	case "B":
		fmt.Printf("Es de segunda")
	case "C":
		fmt.Println("Es de tercera")
	default:
		fmt.Println("¡Es de lo peor!")
	}
}
