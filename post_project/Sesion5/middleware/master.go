package master

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapper := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}
		next(wrapper, r)
		duration := time.Since(start)
		log.Printf(
			"%s %s %d %s %s",
			r.Method, r.URL.Path, wrapper.statusCode, duration, r.RemoteAddr,
		)
	}
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func CORSMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next(w, r)
	}
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || len(authHeader) < 7 || authHeader[:7] != "Bearer " {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, `{"error": "Token de autorización no proporcionado o inválido"}`)
			return
		}
		token := authHeader[7:]
		if token != "secret123" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, `{"error": "Token de autorización inválido"}`)
			return
		}
		next(w, r)
	}
}

func RateLimitMiddleware(next http.HandlerFunc) http.HandlerFunc {
	requests := make(map[string]int)

	return func(w http.ResponseWriter, r *http.Request) {
		clientIP := r.RemoteAddr
		requests[clientIP]++
		if requests[clientIP] > 10 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusTooManyRequests)
			fmt.Fprintf(w, `{"error": "Demasiadas solicitudes desde esta IP"}`)
			return
		}
		next(w, r)
	}
}

func ChainMiddleware(handler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}
func main() {
	initSampleData()

	http.HandleFunc("/", ChainMiddleware(homeHandler, LoggingMiddleware, CORSMiddleware))
	http.HandleFunc("/health", ChainMiddleware(healthHandler, LoggingMiddleware, CORSMiddleware))

	http.HandleFunc("/tasks", ChainMiddleware(tasksHandler, LoggingMiddleware, CORSMiddleware))

	http.HandleFunc("/tasks/create", ChainMiddleware(
		createTaskHandler, LoggingMiddleware, CORSMiddleware, AuthMiddleware,
	))
	http.HandleFunc("/tasks/", ChainMiddleware(
		taskByIDHandler, LoggingMiddleware, CORSMiddleware, AuthMiddleware,
	))

	fmt.Println("Servidor escuchando en :8080")
	fmt.Println("Rutas protegidas: Bearer token secret123")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	handleCreateTask(w, r)
}

func initSampleData() {
	tasks = []Task{
		{
			ID:          1,
			Title:       "Tarea 1",
			Description: "Descripción de la tarea 1",
			Completed:   false,
		},
		{
			ID:          2,
			Title:       "Tarea 2",
			Description: "Descripción de la tarea 2",
			Completed:   false,
		},
	}
	nextID = 3
}
