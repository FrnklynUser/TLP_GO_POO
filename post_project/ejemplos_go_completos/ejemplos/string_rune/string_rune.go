package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	texto := "Hola 世界 🌍"
	fmt.Printf("Longitud en bytes: %d\n", len(texto))
	fmt.Printf("Longitud en runes: %d\n", utf8.RuneCountInString(texto))
}
