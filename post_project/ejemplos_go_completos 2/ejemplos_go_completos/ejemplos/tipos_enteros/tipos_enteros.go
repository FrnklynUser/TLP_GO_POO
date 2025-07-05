package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a int8 = 127
	var b int16 = 32767
	var ua uint8 = 255

	fmt.Printf("Tamaños: int8=%d, int16=%d, uint8=%d\n", unsafe.Sizeof(a), unsafe.Sizeof(b), unsafe.Sizeof(ua))
}
