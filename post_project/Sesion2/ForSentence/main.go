package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("=== ESTRUCTURAS FOR ===")

	// FOR CLÁSICO (C-style)
	fmt.Println("--- For clásico ---")
	for i := 0; i < 5; i++ {
		fmt.Printf("Iteración %d\n", i)
	}

	// FOR COMO WHILE
	fmt.Println("\n--- For como while ---")
	contador := 0
	for contador < 3 {
		fmt.Printf("Contador: %d\n", contador)
		contador++
	}

	// FOR INFINITO
	fmt.Println("\n--- For infinito con break ---")
	i := 0
	for {
		if i >= 3 {
			break
		}
		fmt.Printf("Bucle infinito - iteración: %d\n", i)
		i++
	}

	// FOR CON MÚLTIPLES VARIABLES
	fmt.Println("\n--- For con múltiples variables ---")
	for i, j := 0, 10; i < j; i, j = i+1, j-1 {
		fmt.Printf("i=%d, j=%d, suma=%d\n", i, j, i+j)
	}

	// FOR CON CONDICIONES COMPLEJAS
	fmt.Println("\n--- For con condiciones complejas ---")
	x, y := 1, 1
	for x < 100 && y < 100 {
		fmt.Printf("Fibonacci: x=%d, y=%d\n", x, y)
		x, y = y, x+y
	}

	// RANGE CON SLICES
	fmt.Println("\n--- Range con slices ---")
	frutas := []string{"manzana", "banana", "naranja", "uva"}

	fmt.Println("Índice y valor:")
	for indice, fruta := range frutas {
		fmt.Printf("%d: %s\n", indice, fruta)
	}

	fmt.Println("Solo valores:")
	for _, fruta := range frutas {
		fmt.Printf("- %s\n", fruta)
	}

	fmt.Println("Solo índices:")
	for indice := range frutas {
		fmt.Printf("Índice: %d\n", indice)
	}

	// RANGE CON MAPS
	fmt.Println("\n--- Range con maps ---")
	edades := map[string]int{
		"Ana":   25,
		"Luis":  30,
		"María": 28,
	}
	for nombre, edad := range edades {
		fmt.Printf("%s tiene %d años\n", nombre, edad)
	}

	// RANGE CON STRINGS
	fmt.Println("\n--- Range con strings ---")
	texto := "Hola 世界"
	for i, caracter := range texto {
		fmt.Printf("Posición %d: %c (U+%04X)\n", i, caracter, caracter)
	}

	// CASOS PRÁCTICOS
	demonstrarCasosPracticosFor()
}

func demonstrarCasosPracticosFor() {
	fmt.Println("\n--- Casos prácticos con for ---")

	// 1. Procesamiento en lotes
	fmt.Println("1. Procesamiento en lotes:")
	datos := make([]int, 100)
	for i := range datos {
		datos[i] = i + 1
	}

	tamañoLote := 10
	for i := 0; i < len(datos); i += tamañoLote {
		fin := i + tamañoLote
		if fin > len(datos) {
			fin = len(datos)
		}
		lote := datos[i:fin]
		fmt.Printf("  Procesando lote %d: %d elementos\n", i/tamañoLote+1, len(lote))
		time.Sleep(50 * time.Millisecond)
	}

	// 2. Búsqueda con múltiples criterios
	fmt.Println("\n2. Búsqueda de usuarios:")
	usuarios := []struct {
		ID     int
		Nombre string
		Edad   int
		Activo bool
		Ciudad string
	}{
		{1, "Ana García", 25, true, "Lima"},
		{2, "Luis Martín", 30, false, "Cusco"},
		{3, "María López", 28, true, "Lima"},
		{4, "Carlos Ruiz", 35, true, "Arequipa"},
		{5, "Elena Torres", 29, true, "Lima"},
	}

	fmt.Println("Usuarios activos de Lima > 25 años:")
	for _, usuario := range usuarios {
		if usuario.Activo && usuario.Ciudad == "Lima" && usuario.Edad > 25 {
			fmt.Printf("  - %s (%d años)\n", usuario.Nombre, usuario.Edad)
		}
	}

	// 3. Validación de datos con acumuladores
	fmt.Println("\n3. Validación de formulario:")
	campos := map[string]string{
		"nombre":   "Juan Pérez",
		"email":    "juan@email.com",
		"telefono": "123456789",
		"edad":     "25",
		"ciudad":   "",
	}

	errores := make([]string, 0)
	camposValidos := 0
	for campo, valor := range campos {
		if valor == "" {
			errores = append(errores, fmt.Sprintf("Campo '%s' es requerido", campo))
		} else {
			camposValidos++
			fmt.Printf("  ✅ %s: %s\n", campo, valor)
		}
	}

	if len(errores) > 0 {
		fmt.Println("  Errores encontrados:")
		for _, err := range errores {
			fmt.Printf("  ❌ %s\n", err)
		}
	}
	fmt.Printf("  Campos válidos: %d/%d\n", camposValidos, len(campos))

	// 4. Reporte de ventas por región
	fmt.Println("\n4. Reporte de ventas por región:")
	ventas := []struct {
		Producto string
		Region   string
		Monto    float64
	}{
		{"Laptop", "Norte", 2500.00},
		{"Mouse", "Norte", 45.50},
		{"Laptop", "Sur", 2500.00},
		{"Teclado", "Centro", 120.00},
		{"Mouse", "Sur", 45.50},
		{"Laptop", "Centro", 2500.00},
	}

	ventasPorRegion := make(map[string]float64)
	contadorPorRegion := make(map[string]int)
	for _, venta := range ventas {
		ventasPorRegion[venta.Region] += venta.Monto
		contadorPorRegion[venta.Region]++
	}

	for region, total := range ventasPorRegion {
		promedio := total / float64(contadorPorRegion[region])
		fmt.Printf("  %s: $%.2f total (%d ventas, promedio: $%.2f)\n",
			region, total, contadorPorRegion[region], promedio)
	}

	// 5. Retry con backoff
	fmt.Println("\n5. Simulación de retry con backoff:")
	maxIntentos := 5
	for intento := 1; intento <= maxIntentos; intento++ {
		fmt.Printf("  Intento %d/%d", intento, maxIntentos)
		if rand.Float32() < 0.7 {
			fmt.Println(" - ❌ Falló")
			if intento < maxIntentos {
				delay := time.Duration(intento*intento) * 100 * time.Millisecond
				fmt.Printf("    Esperando %v antes del siguiente intento...\n", delay)
				time.Sleep(delay)
			}
		} else {
			fmt.Println(" - ✅ Éxito")
			break
		}
	}

	// 6. Ordenamiento burbuja
	fmt.Println("\n 6. Ordenamiento burbuja:")
	numeros := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("  Array original: %v\n", numeros)
	n := len(numeros)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if numeros[j] > numeros[j+1] {
				numeros[j], numeros[j+1] = numeros[j+1], numeros[j]
			}
		}
	}
	fmt.Printf("  Array ordenado: %v\n", numeros)
}
