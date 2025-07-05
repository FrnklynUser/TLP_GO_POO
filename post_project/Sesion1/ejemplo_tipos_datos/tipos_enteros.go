package main

import (
	"fmt"
	"unsafe"
)

func demoEnteros() {
	// Enteros con signo
	var a int8 = 127
	var b int16 = 32767
	var c int32 = 2147483647
	var d int64 = 9223372036854775807

	fmt.Printf("Tamaños en bytes:\n")
	fmt.Printf("int8: %d, int16: %d, int32: %d, int64: %d\n",
		unsafe.Sizeof(a), unsafe.Sizeof(b), unsafe.Sizeof(c), unsafe.Sizeof(d))

	// Enteros sin signo
	var ua uint8 = 255
	var ub uint16 = 65535
	fmt.Printf("uint8: %d, uint16: %d\n", unsafe.Sizeof(ua), unsafe.Sizeof(ub))

	// Tipos dependientes de la arquitectura
	var e int = 42
	var f uint = 42
	fmt.Printf("int: %d, uint: %d\n", unsafe.Sizeof(e), unsafe.Sizeof(f))

	// Tipos especiales
	var g byte = 255 // alias para uint8
	var h rune = 'A' // alias para int32
	var i uintptr = 0x12345678
	fmt.Printf("byte: %d, rune: %d, uintptr: %d\n", unsafe.Sizeof(g), unsafe.Sizeof(h), unsafe.Sizeof(i))
}
