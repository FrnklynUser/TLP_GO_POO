package main

import (
	"fmt"
	"time"
)

func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func NewLimiter(maxCalls int, duration time.Duration) func() bool {
	var count int
	resetTime := time.Now().Add(duration)

	return func() bool {
		if time.Now().After(resetTime) {
			count = 0
			resetTime = time.Now().Add(duration)
		}

		if count < maxCalls {
			count++
			return true
		}
		return false
	}
}

func applyOperation(op func(int, int) int, a, b int) int {
	return op(a, b)
}

func applyTwice(f func(int) int, v int) int {
	return f(f(v))
}

func multiplierGenerator(factor int) func(int) int {
	return func(value int) int {
		return value * factor
	}
}

func main() {
	counter := createCounter()
	fmt.Println("\n🔢 Contador básico:")
	fmt.Println("Conteo:", counter())
	fmt.Println("Conteo:", counter())
	fmt.Println("Conteo:", counter())

	limiter := NewLimiter(3, 5*time.Second)
	fmt.Println("\n⏱ Limitador de llamadas (máximo 3 en 5 segundos):")
	for i := 1; i <= 5; i++ {
		if limiter() {
			fmt.Printf("✅ Llamada #%d permitida\n", i)
		} else {
			fmt.Printf("❌ Llamada #%d rechazada - límite excedido\n", i)
		}
		time.Sleep(1 * time.Second)
	}

	add := func(a, b int) int { return a + b }
	result := applyOperation(add, 5, 3)
	fmt.Println("\n➕ Aplicando función de orden superior:")
	fmt.Println("Resultado de sumar 5 + 3 =", result)

	double := multiplierGenerator(2)
	triple := multiplierGenerator(3)
	fmt.Println("\n🔁 Funciones generadoras:")
	fmt.Println("Doble de 10:", double(10))
	fmt.Println("Triple de 10:", triple(10))

	fmt.Println("\n🎯 Conceptos demostrados:")
	fmt.Println("✅ Uso de closures para mantener estado")
	fmt.Println("✅ Funciones anónimas y retornadas")
	fmt.Println("✅ Tipos de funciones como parámetros")
	fmt.Println("✅ Funciones de orden superior")
	fmt.Println("✅ Composición de funciones")
	fmt.Println("✅ Aplicación práctica de patrones funcionales")
}
