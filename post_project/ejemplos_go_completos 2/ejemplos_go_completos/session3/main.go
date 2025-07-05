package session3

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// ============================================================
// ESTRUCTURAS (STRUCTS)
// ============================================================

// Definición básica de un struct
type Usuario struct {
	ID     int
	Nombre string
	Email  string
	Activo bool
}

// Struct con tags (útil para JSON, validaciones, etc.)
type Producto struct {
	ID     int     `json:"id" db:"product_id"`
	Nombre string  `json:"nombre" validate:"required"`
	Precio float64 `json:"precio" validate:"min=0"`
}

// Struct anidado
type Empresa struct {
	Nombre    string
	Direccion Direccion
	Empleados []Usuario
}

type Direccion struct {
	Calle  string
	Ciudad string
	CP     string
}

// Métodos con receptores
type Rectangulo struct {
	Ancho, Alto int
}

// Receptor de valor
func (r Rectangulo) Area() int {
	return r.Ancho * r.Alto
}

// Receptor de puntero
func (r *Rectangulo) Escalar(factor int) {
	r.Ancho *= factor
	r.Alto *= factor
}

// Ejemplo completo de biblioteca
type Libro struct {
	ID               int
	Titulo           string
	Autor            string
	ISBN             string
	Paginas          int
	Prestado         bool
	Disponible       bool
	Precio           float64
	FechaPublicacion time.Time
}

type Prestamo struct {
	ID              int
	LibroID         int
	UsuarioID       int
	FechaPrestamo   time.Time
	FechaDevolucion time.Time
	Devuelto        bool
}

type Biblioteca struct {
	Nombre    string
	Direccion string
	Libros    []Libro
	Usuarios  []Usuario
	Prestamos []Prestamo
	proximoID int
}

// Métodos con receptor de valor (solo lectura)
func (l Libro) ObtenerInfo() string {
	estado := "Disponible"
	if l.Prestado {
		estado = "Prestado"
	}
	return fmt.Sprintf("[%d] %s por %s - %s", l.ID, l.Titulo, l.Autor, estado)
}

func (l Libro) EsPrestable() bool {
	return !l.Prestado && l.Paginas > 0
}

// Métodos con receptor de puntero (modificación)
func (l *Libro) Prestar() error {
	if l.Prestado {
		return fmt.Errorf("el libro '%s' ya está prestado", l.Titulo)
	}
	if l.Paginas <= 0 {
		return fmt.Errorf("el libro '%s' no es válido", l.Titulo)
	}
	l.Prestado = true
	return nil
}

func (l *Libro) Devolver() error {
	if !l.Prestado {
		return fmt.Errorf("el libro '%s' no está prestado", l.Titulo)
	}
	l.Prestado = false
	return nil
}

// Constructor para Biblioteca
func NuevaBiblioteca(nombre, direccion string) *Biblioteca {
	return &Biblioteca{
		Nombre:    nombre,
		Direccion: direccion,
		Libros:    make([]Libro, 0),
		Usuarios:  make([]Usuario, 0),
		Prestamos: make([]Prestamo, 0),
		proximoID: 1,
	}
}

// Métodos de Biblioteca
func (b *Biblioteca) AgregarLibro(titulo, autor, isbn string, paginas int) (*Libro, error) {
	if titulo == "" || autor == "" {
		return nil, fmt.Errorf("titulo y autor son obligatorios")
	}

	for _, libro := range b.Libros {
		if libro.ISBN == isbn && isbn != "" {
			return nil, fmt.Errorf("ya existe un libro con ISBN: %s", isbn)
		}
	}

	libro := Libro{
		ID:               b.proximoID,
		Titulo:           titulo,
		Autor:            autor,
		ISBN:             isbn,
		Paginas:          paginas,
		Prestado:         false,
		Disponible:       true,
		Precio:           0,
		FechaPublicacion: time.Now(),
	}

	b.Libros = append(b.Libros, libro)
	b.proximoID++
	return &libro, nil
}

func (b Biblioteca) BuscarLibro(id int) *Libro {
	for i, libro := range b.Libros {
		if libro.ID == id {
			return &b.Libros[i]
		}
	}
	return nil
}

// ============================================================
// PUNTEROS
// ============================================================

func demostrarPunteros() {
	fmt.Println("\n=== DEMOSTRACIÓN DE PUNTEROS ===")

	// Puntero básico
	var x int = 10
	var ptr *int = &x

	fmt.Println("Valor de x:", x)
	fmt.Println("Dirección de x:", ptr)
	fmt.Println("Valor a través del puntero:", *ptr)

	// Modificar valor a través del puntero
	*ptr = 20
	fmt.Println("Nuevo valor de x:", x)

	// Puntero con new
	ptr2 := new(int)
	*ptr2 = 30
	fmt.Println("Puntero con new:", *ptr2)

	// Puntero a struct
	usuario := Usuario{ID: 1, Nombre: "Juan"}
	ptrUsuario := &usuario
	fmt.Println("Nombre a través de puntero:", ptrUsuario.Nombre) // Go hace el dereferenciado automático

	// Modificar struct a través de puntero
	ptrUsuario.Nombre = "Juan Carlos"
	fmt.Println("Nombre modificado:", usuario.Nombre)
}

// ============================================================
// INTERFACES
// ============================================================

// Interfaces básicas
type Notificador interface {
	EnviarNotificacion(destinatario, mensaje string) error
}

type ValidadorMensaje interface {
	ValidarMensaje(mensaje string) error
	ValidarDestinatario(destinatario string) error
}

// Interfaces compuestas
type NotificadorCompleto interface {
	Notificador
	ValidadorMensaje
}

// Implementación concreta: EmailNotificador
type EmailNotificador struct {
	servidor  string
	puerto    int
	usuario   string
	password  string
	registros map[string]*RegistroNotificacion
}

type RegistroNotificacion struct {
	ID           string
	Tipo         string
	Destinatario string
	Mensaje      string
	Estado       string
	Timestamp    time.Time
	Intentos     int
	Error        string
}

func NuevoEmailNotificador(servidor string, puerto int, usuario, password string) *EmailNotificador {
	return &EmailNotificador{
		servidor:  servidor,
		puerto:    puerto,
		usuario:   usuario,
		password:  password,
		registros: make(map[string]*RegistroNotificacion),
	}
}

func (e *EmailNotificador) EnviarNotificacion(destinatario, mensaje string) error {
	// Validación básica
	if !strings.Contains(destinatario, "@") {
		return fmt.Errorf("email inválido")
	}
	if len(mensaje) == 0 {
		return fmt.Errorf("mensaje no puede estar vacío")
	}

	// Registrar notificación
	id := fmt.Sprintf("email_%d", time.Now().UnixNano())
	e.registros[id] = &RegistroNotificacion{
		ID:           id,
		Tipo:         "email",
		Destinatario: destinatario,
		Mensaje:      mensaje,
		Estado:       "enviado",
		Timestamp:    time.Now(),
		Intentos:     1,
	}

	fmt.Printf("Email enviado a %s: %s\n", destinatario, mensaje)
	return nil
}

func (e *EmailNotificador) ValidarMensaje(mensaje string) error {
	if len(mensaje) == 0 {
		return errors.New("mensaje no puede estar vacío")
	}
	if len(mensaje) > 1000 {
		return errors.New("mensaje muy largo (máximo 1000 caracteres)")
	}
	return nil
}

func (e *EmailNotificador) ValidarDestinatario(destinatario string) error {
	if !strings.Contains(destinatario, "@") {
		return errors.New("email inválido: debe contener @")
	}
	if !strings.Contains(destinatario, ".") {
		return errors.New("email inválido: debe contener dominio")
	}
	return nil
}

// Implementación simple: SlackNotificador
type SlackNotificador struct {
	webhook string
	canal   string
}

func NuevoSlackNotificador(webhook, canal string) *SlackNotificador {
	return &SlackNotificador{
		webhook: webhook,
		canal:   canal,
	}
}

func (s *SlackNotificador) EnviarNotificacion(destinatario, mensaje string) error {
	fmt.Printf("Slack -> Canal: %s | Usuario: %s | Mensaje: %s\n", s.canal, destinatario, mensaje)
	return nil
}

// Servicio que usa interfaces
type ServicioNotificaciones struct {
	notificadores []Notificador
}

func NuevoServicioNotificaciones() *ServicioNotificaciones {
	return &ServicioNotificaciones{
		notificadores: make([]Notificador, 0),
	}
}

func (s *ServicioNotificaciones) AgregarNotificador(n Notificador) {
	s.notificadores = append(s.notificadores, n)
}

func (s *ServicioNotificaciones) EnviarATodos(destinatario, mensaje string) map[string]error {
	resultados := make(map[string]error)

	for _, n := range s.notificadores {
		tipo := fmt.Sprintf("%T", n)
		err := n.EnviarNotificacion(destinatario, mensaje)
		resultados[tipo] = err
	}

	return resultados
}

// ============================================================
// FUNCIONES
// ============================================================

// Función básica con múltiples retornos
func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("división por cero no permitida")
	}
	return a / b, nil
}

// Función con retornos nombrados
func analizarTexto(texto string) (palabras int, caracteres int, lineas int) {
	caracteres = len(texto)
	lineas = strings.Count(texto, "\n") + 1
	palabras = len(strings.Fields(texto))
	return // return implícito
}

// Función variádica
func calcularPromedio(numeros ...float64) float64 {
	if len(numeros) == 0 {
		return 0
	}
	var suma float64
	for _, num := range numeros {
		suma += num
	}
	return suma / float64(len(numeros))
}

// Función que acepta otra función como parámetro
type FiltroLibro func(Libro) bool

func FiltrarLibros(libros []Libro, filtro FiltroLibro) []Libro {
	var resultado []Libro
	for _, libro := range libros {
		if filtro(libro) {
			resultado = append(resultado, libro)
		}
	}
	return resultado
}

// Función que retorna una función (closure)
func CrearFiltroAutor(autorBuscado string) FiltroLibro {
	return func(libro Libro) bool {
		return strings.Contains(strings.ToLower(libro.Autor), strings.ToLower(autorBuscado))
	}
}

// ============================================================
// ESTRUCTURA DE PROYECTO (EJEMPLO SIMPLIFICADO)
// ============================================================

// Ejemplo de estructura de paquetes (simulado en este archivo único)

// pkg/logger/logger.go (simulado)
func LoggerInfo(mensaje string) {
	fmt.Printf("[INFO] %s\n", mensaje)
}

// models/usuario.go (simulado)
type ModelsUsuario struct {
	ID     int
	Nombre string
	Email  string
}

// internal/servicio/servicio.go (simulado)
func ServicioProcesar() {
	fmt.Println("Procesando servicio...")
}

// ============================================================
// FUNCIÓN PRINCIPAL
// ============================================================

func main() {
	fmt.Println("=== DEMOSTRACIÓN COMPLETA DE LA SESIÓN 3 ===")

	// 1. Structs y métodos
	fmt.Println("\n--- Structs y Métodos ---")
	rect := Rectangulo{Ancho: 10, Alto: 5}
	fmt.Println("Área inicial:", rect.Area())
	rect.Escalar(2)
	fmt.Println("Área después de escalar:", rect.Area())

	// 2. Biblioteca completa
	fmt.Println("\n--- Sistema de Biblioteca ---")
	biblio := NuevaBiblioteca("Central", "Av. Principal 123")
	libro, _ := biblio.AgregarLibro("El Quijote", "Cervantes", "123-456", 863)
	fmt.Println("Libro agregado:", libro.ObtenerInfo())

	// 3. Punteros
	demostrarPunteros()

	// 4. Interfaces y polimorfismo
	fmt.Println("\n--- Interfaces y Polimorfismo ---")
	servicio := NuevoServicioNotificaciones()
	servicio.AgregarNotificador(NuevoEmailNotificador("smtp.example.com", 587, "user", "pass"))
	servicio.AgregarNotificador(NuevoSlackNotificador("https://hooks.slack.com", "#general"))

	resultados := servicio.EnviarATodos("user@example.com", "Hola desde Go!")
	for tipo, err := range resultados {
		if err != nil {
			fmt.Printf("Error en %s: %v\n", tipo, err)
		} else {
			fmt.Printf("%s: Éxito\n", tipo)
		}
	}

	// 5. Funciones avanzadas
	fmt.Println("\n--- Funciones Avanzadas ---")
	prom := calcularPromedio(8.5, 9.0, 7.5, 8.0)
	fmt.Printf("Promedio: %.2f\n", prom)

	libros := []Libro{
		{ID: 1, Titulo: "El Quijote", Autor: "Cervantes", Disponible: true},
		{ID: 2, Titulo: "Cien Años de Soledad", Autor: "García Márquez", Disponible: false},
	}

	filtroCervantes := CrearFiltroAutor("Cervantes")
	librosCervantes := FiltrarLibros(libros, filtroCervantes)
	fmt.Println("Libros de Cervantes:", librosCervantes)

	// 6. Estructura de proyecto (simulada)
	fmt.Println("\n--- Estructura de Proyecto ---")
	LoggerInfo("Este es un mensaje de log")
	ServicioProcesar()
	usuario := ModelsUsuario{ID: 1, Nombre: "Juan", Email: "juan@example.com"}
	fmt.Println("Usuario:", usuario)

	fmt.Println("\n¡Demostración completada!")
}
