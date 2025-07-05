package main

import "fmt"

func protegido() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("🧯 Se recuperó del pánico:", r)
		}
	}()

	fmt.Println("Ejecutando función protegida")
	panic("🔥 Error inesperado")
}

func main() {
	protegido()
	fmt.Println("✅ El programa continúa después del recover")
}
