package main

import (
	"fmt"
	"strconv"
)

func demoConversiones() {
	// Conversiones numéricas
	var i int = 42
	var f float64 = 3.14
	var result1 = float64(i) + f
	var result2 = i + int(f)
	fmt.Printf("int(%d) + float64(%.2f) = %.2f\n", i, f, result1)
	fmt.Printf("int(%d) + int(%.2f) = %d\n", i, f, result2)

	// Conversión de string a número
	strNumero := "123"
	numero, err := strconv.Atoi(strNumero)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("String '%s' -> int: %d\n", strNumero, numero)
	}

	// Conversión de número a string
	entero := 456
	strDesdeInt := strconv.Itoa(entero)
	fmt.Printf("int(%d) -> string: '%s'\n", entero, strDesdeInt)
}
