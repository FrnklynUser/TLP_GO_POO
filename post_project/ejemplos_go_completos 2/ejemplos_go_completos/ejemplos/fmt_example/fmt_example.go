package main

import "fmt"

func main() {
	if true {
		fmt.Println("Mal formateado")
	}
	for i := 0; i < 5; i++ {
		fmt.Printf("Número: %d\n", i)
	}
}
