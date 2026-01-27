package main // Se le suele poner "main" por convención
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hola mundo")
	bufio.NewReader(os.Stdin).ReadBytes('\n') // Esto sirve para que la terminal no se cierre luego de ejecutar el archivo .exe. Si queremos que se cierre la ventana debemos presionar Enter.
}
