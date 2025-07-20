package main

import "fmt"

func main() {
	const (
		Lunes = iota
		Martes
		Miercoles
	)
	fmt.Println("Lunes:", Lunes, "Martes:", Martes, "Miércoles:", Miercoles)
}
