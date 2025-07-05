package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== DEFER ===")

	// DEFER BÁSICO
	demostrarDeferBasico()

	// MÚLTIPLES DEFERS
	demostrarMultiplesDefers()

	// DEFER CON VALORES
	demostrarDeferConValores()

	// CASOS PRÁCTICOS
	demostrarCasosPracticosDefer()
}

// 1. Defer básico
func demostrarDeferBasico() {
	fmt.Println("\n--- Defer básico ---")
	fmt.Println("1. Inicio de función")
	defer fmt.Println("4. Este mensaje se ejecuta al final (defer)")
	fmt.Println("2. En medio de función")
	fmt.Println("3. Antes del return")
}

// 2. Múltiples defers (orden LIFO)
func demostrarMultiplesDefers() {
	fmt.Println("\n--- Múltiples defers (LIFO) ---")
	defer fmt.Println("🥉 Tercer defer (se ejecuta primero)")
	defer fmt.Println("🥈 Segundo defer (se ejecuta segundo)")
	defer fmt.Println("🥇 Primer defer (se ejecuta último)")
	fmt.Println("Código normal ejecutándose...")
}

// 3. Defer con valores capturados
func demostrarDeferConValores() {
	fmt.Println("\n--- Defer con valores capturados ---")
	x := 10
	defer fmt.Printf("Defer clásico: x = %d (valor capturado)\n", x)

	// Captura por referencia (closure)
	defer func() {
		fmt.Printf("Defer closure: x = %d (valor al ejecutarse)\n", x)
	}()

	x = 30
	fmt.Printf("Valor final de x: %d\n", x)
}

// 4. Casos prácticos de uso de defer
func demostrarCasosPracticosDefer() {
	fmt.Println("\n--- Casos prácticos con defer ---")

	// 1. Manejo de archivos
	fmt.Println("1. Manejo de archivos:")
	manejarArchivo()

	// 2. Medición de tiempo
	fmt.Println("\n2. Medición de tiempo:")
	medirTiempoEjecucion()

	// 3. Limpieza de recursos
	fmt.Println("\n3. Limpieza de recursos:")
	simularConexionDB()

	// 4. Logging de entrada y salida
	fmt.Println("\n4. Logging:")
	funcionConLogging("parametro_importante")

	// 5. Mutex unlocking
	fmt.Println("\n5. Manejo de mutex:")
	simularMutex()
}

// Simular apertura y cierre de archivo
func manejarArchivo() {
	fmt.Println("📂 Abriendo archivo...")
	defer fmt.Println("📁 Cerrando archivo (defer)")
	fmt.Println("📝 Escribiendo datos...")
	fmt.Println("📖 Leyendo datos...")
}

// Medir duración de ejecución
func medirTiempoEjecucion() {
	inicio := time.Now()
	defer func() {
		duracion := time.Since(inicio)
		fmt.Printf("⏱ Tiempo de ejecución: %v\n", duracion)
	}()

	fmt.Println("⌛ Simulando operación costosa...")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("✅ Operación completada")
}

// Simular conexión a base de datos
func simularConexionDB() {
	fmt.Println("🔌 Conectando a base de datos...")
	defer fmt.Println("🔌 Desconectando de base de datos (defer)")
	fmt.Println("📡 Ejecutando query 1...")
	fmt.Println("📡 Ejecutando query 2...")
	fmt.Println("📡 Ejecutando query 3...")
}

// Logging simple
func funcionConLogging(parametro string) {
	fmt.Printf("🚀 ENTRADA: funcionConLogging(%s)\n", parametro)
	defer fmt.Println("📤 SALIDA: funcionConLogging")

	if parametro == "error" {
		fmt.Println("❌ Error simulado")
		return
	}

	fmt.Println("⚙ Procesando lógica de negocio...")
	fmt.Println("✅ Procesamiento exitoso")
}

// Simular uso de mutex con defer
func simularMutex() {
	var mu sync.Mutex
	mu.Lock()
	fmt.Println("🔒 Lock adquirido")
	defer mu.Unlock()
	defer fmt.Println("🔓 Lock liberado (defer)")

	fmt.Println("⚙ Trabajando en sección crítica...")
	time.Sleep(50 * time.Millisecond)
}
