package main

import (
	"fmt"
	"math"
)

func demoFlotantes() {
	var precio32 float32 = 19.99
	var precio64 float64 = 19.99
	var pi = 3.14159265359 // inferido como float64

	fmt.Printf("Precisión float32: %.10f\n", precio32)
	fmt.Printf("Precisión float64: %.10f\n", precio64)

	area := pi * math.Pow(5.0, 2)
	fmt.Printf("Área del círculo: %.2f\n", area)

	// Problemas comunes con punto flotante
	a := 0.1
	b := 0.2
	c := 0.3
	fmt.Printf("0.1 + 0.2 == 0.3: %t\n", a+b == c)

	// Comparación correcta con tolerancia
	const epsilon = 1e-9
	fmt.Printf("Comparación con tolerancia: %t\n", math.Abs((a+b)-c) < epsilon)
}
