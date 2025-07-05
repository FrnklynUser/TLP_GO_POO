package main

import (
	"fmt"
	"strings"
	"time"
)

// ==========================================
// 1. CLOSURES BÁSICOS - CONTADORES
// ==========================================

// createCounter devuelve una función que cuenta desde 0
func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// ==========================================
// 2. CLOSURES CON ESTADO PRIVADO
// ==========================================

// NewLimiter crea un limitador de llamadas con estado encapsulado
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
			return true // Permitido
		}
		return false // Límite excedido
	}
}

// ==========================================
// 3. FUNCIONES DE ORDEN SUPERIOR
// ==========================================

// Operation es un tipo función que opera sobre dos enteros
type Operation func(int, int) int

// applyOperation aplica una operación a dos valores
func applyOperation(op Operation, a, b int) int {
	return op(a, b)
}

// applyTwice aplica una función dos veces
func applyTwice(f func(int) int, v int) int {
	return f(f(v))
}

// ==========================================
// 4. FUNCIONES QUE DEVUELVEN FUNCIONES
// ==========================================

// multiplierGenerator devuelve una función que multiplica por un factor fijo
func multiplierGenerator(factor int) func(int) int {
	return func(value int) int {
		return value * factor
	}
}

// ==========================================
// 5. EJEMPLO PRÁCTICO: SISTEMA DE DESCUENTOS
// ==========================================

type DiscountFunc func(float64) float64

func applyDiscounts(price float64, discounts ...DiscountFunc) float64 {
	for _, discount := range discounts {
		price = discount(price)
	}
	return price
}

// Descuentos comunes como funciones
func tenPercentOff() DiscountFunc {
	return func(price float64) float64 {
		return price * 0.9
	}
}

func fixedDiscount(amount float64) DiscountFunc {
	return func(price float64) float64 {
		return price - amount
	}
}

func buyOneGetOneHalfOff(items int) DiscountFunc {
	return func(price float64) float64 {
		if items >= 2 {
			return price * 0.75 // 25% extra por BOGO
		}
		return price
	}
}

// ==========================================
// FUNCIÓN PRINCIPAL
// ==========================================
func main() {
	fmt.Println("🧩 SESIÓN 04 - CIERRES Y FUNCIONES AVANZADAS")
	fmt.Println(strings.Repeat("=", 60))

	// 1. Cierre básico
	counter := createCounter()
	fmt.Println("\n🔢 Contador básico:")
	fmt.Println("Conteo:", counter()) // Output: 1
	fmt.Println("Conteo:", counter()) // Output: 2
	fmt.Println("Conteo:", counter()) // Output: 3

	// 2. Limitador de llamadas
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

	// 3. Función de orden superior
	add := func(a, b int) int { return a + b }
	result := applyOperation(add, 5, 3)
	fmt.Println("\n➕ Aplicando función de orden superior:")
	fmt.Println("Resultado de sumar 5 + 3 =", result)

	// 4. Funciones que devuelven funciones
	double := multiplierGenerator(2)
	triple := multiplierGenerator(3)
	fmt.Println("\n🔁 Funciones generadoras:")
	fmt.Println("Doble de 10:", double(10))
	fmt.Println("Triple de 10:", triple(10))

	// 5. Ejemplo práctico con descuentos
	fmt.Println("\n🛍️ Sistema de descuentos:")
	originalPrice := 100.0
	discountedPrice := applyDiscounts(originalPrice,
		tenPercentOff(),
		fixedDiscount(10),
		buyOneGetOneHalfOff(2),
	)

	fmt.Printf("Precio original: $%.2f\n", originalPrice)
	fmt.Printf("Precio con descuentos: $%.2f\n", discountedPrice)
	fmt.Printf("Descuento total: %.2f%%\n", (originalPrice-discountedPrice)/originalPrice*100)

	fmt.Println("\n🎯 Conceptos demostrados:")
	conceptos := []string{
		"✅ Uso de closures para mantener estado",
		"✅ Funciones anónimas y retornadas",
		"✅ Tipos de funciones como parámetros",
		"✅ Funciones de orden superior",
		"✅ Composición de funciones",
		"✅ Aplicación práctica de patrones funcionales",
	}
	for _, c := range conceptos {
		fmt.Println(c)
	}
}
