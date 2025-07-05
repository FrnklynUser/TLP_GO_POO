package main

import (
	"fmt"
	"unicode/utf8"
)

func demoStringsRunes() {
	texto := "Hola mundo 🌍"

	fmt.Printf("Texto original: '%s'\n", texto)
	fmt.Printf("Longitud en bytes: %d\n", len(texto))
	fmt.Printf("Longitud en runes: %d\n", utf8.RuneCountInString(texto))

	runes := []rune(texto)
	fmt.Printf("Runes: %v\n", runes)

	for i, r := range texto {
		fmt.Printf("Posición %d: %c (U+%04X)\n", i, r, r)
	}

	emoji := '🚀'
	fmt.Printf("Rune: %c (%d)\n", emoji, emoji)
}
