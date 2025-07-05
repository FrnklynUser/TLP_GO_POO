package main

import (
	"fmt"
)

func demoConstantes() {
	// Constantes básicas
	//const companyName = "TechCorp"
	//const version = "1.0.0"

	// Enumeración básica con iota
	const (
		Lunes = iota
		Martes
		Miercoles
		Jueves
		Viernes
		Sabado
		Domingo
	)

	fmt.Printf("Días de la semana:\n")
	fmt.Printf("Lunes: %d, Martes: %d, Miércoles: %d\n", Lunes, Martes, Miercoles)

	// Enumeración con valores específicos
	const (
		StatusInactivo = iota + 1
		StatusActivo
		StatusSuspendido
		StatusBloqueado
	)

	fmt.Printf("Estados:\n")
	fmt.Printf("Inactivo: %d, Activo: %d, Suspendido: %d, Bloqueado: %d\n",
		StatusInactivo, StatusActivo, StatusSuspendido, StatusBloqueado)

	// Enumeración con potencias de 2 (flags)
	const (
		Read    = 1 << iota // 1<<0 = 1
		Write               // 1<<1 = 2
		Execute             // 1<<2 = 4
	)

	fmt.Printf("Permisos (flags):\n")
	fmt.Printf("Read: %d, Write: %d, Execute: %d\n", Read, Write, Execute)

	// Combinación de permisos
	const ReadWrite = Read | Write            // 3
	const FullAccess = Read | Write | Execute // 7

	fmt.Printf("Permisos combinados:\n")
	fmt.Printf("ReadWrite: %d, FullAccess: %d\n", ReadWrite, FullAccess)
}
