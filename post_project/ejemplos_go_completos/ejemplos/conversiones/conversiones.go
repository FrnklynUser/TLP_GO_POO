package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "123"
	n, _ := strconv.Atoi(str)
	fmt.Printf("Convertido: %d\n", n)
}
