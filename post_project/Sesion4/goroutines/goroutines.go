package main

import (
	"fmt"
	"sync"
	"time"
)

func downloadURL(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("⏳ Descargando %s...\n", url)
	time.Sleep(time.Second)
	fmt.Printf("✅ %s descargado\n", url)
}

func main() {
	urls := []string{
		"https://github.com/tiaguinho/gosoap ",
		"https://github.com/FrnklynUser/go_facturacion ",
		"https://example.com/3 ",
	}

	var wg sync.WaitGroup

	fmt.Println("=== DESCARGA SECUENCIAL ===")
	start := time.Now()

	for _, url := range urls {
		// Llamada secuencial sin usar WaitGroup
		downloadURL(url, &sync.WaitGroup{})
	}

	sequentialTime := time.Since(start)
	fmt.Printf("⏱ Tiempo secuencial: %v\n\n", sequentialTime)

	fmt.Println("=== DESCARGA CONCURRENTE ===")
	start = time.Now()

	// Usamos el WaitGroup real para concurrencia
	wg.Add(len(urls))
	for _, url := range urls {
		go downloadURL(url, &wg)
	}

	wg.Wait()
	concurrentTime := time.Since(start)
	fmt.Printf("⏱ Tiempo concurrente: %v\n", concurrentTime)
	fmt.Printf("⚡ Mejora: %.2fx más rápido\n", float64(sequentialTime)/float64(concurrentTime))
}
