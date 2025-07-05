package main

import "fmt"

func causaPanico() {
	fmt.Println("Antes del pánico 🚨")
	panic("¡Algo salió mal!")
	// Esta línea no se ejecuta porque el panic detiene la función
	// fmt.Println("Después del pánico ❌")
}

func main() {
	causaPanico()
	// Esta línea tampoco se ejecuta porque el panic termina el programa
	fmt.Println("Esto nunca se ejecuta ❌")
}
