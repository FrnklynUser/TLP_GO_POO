package main

import "fmt"

func demoBooleanos() {
	// Declaración explícita
	var activo bool = true
	var disponible bool = false

	// Inferencia de tipo
	var validado = true // Go infiere bool

	// Zero value
	var configurado bool // false por defecto

	fmt.Printf("activo: %t, disponible: %t, validado: %t, configurado: %t\n",
		activo, disponible, validado, configurado)

	// Operaciones lógicas
	resultado := activo && disponible
	negacion := !activo

	fmt.Printf("resultado: %t, negación: %t\n", resultado, negacion)
}
