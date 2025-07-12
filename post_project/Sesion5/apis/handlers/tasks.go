package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Completed   bool      `json:"completed"`
}

var tasks []Task
var nextID = 1

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		handleGetTasks(w, r)
	case "POST":
		handleCreateTasks(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Método no permitido",
		})
	}
}

func handleGetTasks(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	completed := queryParams.Get("completed")
	limit := queryParams.Get("limit")

	filteredTasks := tasks
	if completed != "" {
		isCompleted, err := strconv.ParseBool(completed)
		if err == nil {
			var filtered []Task
			for _, task := range tasks {
				if task.Completed == isCompleted {
					filtered = append(filtered, task)
				}
			}
			filteredTasks = filtered
		}
	}
	if limit != "" {
		if limitNum, err := strconv.Atoi(limit); err == nil && limitNum > 0 {
			if limitNum < len(filteredTasks) {
				filteredTasks = filteredTasks[:limitNum]
			}
		}
	}
	userAgent := r.Header.Get("User-Agent")
	log.Printf("Cliente conectado: %s", userAgent)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"tasks": filteredTasks,
		"total": len(filteredTasks),
	})
}

func handleCreateTasks(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Tipo de contenido no soportado deben ser JSON",
		})
		return
	}

	var newTask Task
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&newTask); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Solicitud inválida: " + err.Error(),
		})
		return
	}

	if strings.TrimSpace(newTask.Title) == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "El título de la tarea no puede estar vacío",
		})
		return
	}
	newTask.ID = nextID
	nextID++
	newTask.CreatedAt = time.Now()
	newTask.UpdatedAt = newTask.CreatedAt
	newTask.Completed = false
	tasks = append(tasks, newTask)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Tarea creada exitosamente",
		"task":    newTask,
	})
}

func taskByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	path := strings.TrimPrefix(r.URL.Path, "/tasks/")
	if path == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "ID de tarea no proporcionado",
		})
		return
	}

	id, err := strconv.Atoi(path)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "ID de tarea inválido",
		})
		return
	}

	var foundTask *Task
	for i := range tasks {
		if tasks[i].ID == id {
			foundTask = &tasks[i]
			break
		}
	}

	if foundTask == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Tarea no encontrada",
		})
		return
	}

	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(foundTask)
	case "PUT":
		handleUpdateTask(w, r, foundTask)
	case "DELETE":
		handleDeleteTask(w, r, id)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Método no permitido",
		})
	}
}

func handleUpdateTask(w http.ResponseWriter, r *http.Request, task *Task) {
	var updates Task
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "JSON inválida",
		})
		return
	}

	if updates.Title != "" {
		task.Title = updates.Title
	}
	if updates.Description != "" {
		task.Description = updates.Description
	}
	task.Completed = updates.Completed
	task.UpdatedAt = time.Now()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Tarea actualizada exitosamente",
		"task":    task,
	})
}

func handleDeleteTask(w http.ResponseWriter, r *http.Request, id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Tarea eliminada exitosamente",
			})
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"error": "Tarea no encontrada",
	})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bienvenido a la API de tareas!"))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	tasks = []Task{
		{
			ID:          1,
			Title:       "Tarea de ejemplo",
			Description: "Esta es una tarea de ejemplo",
			CreatedAt:   time.Now().Add(24 * time.Hour),
			UpdatedAt:   time.Now().Add(24 * time.Hour),
			Completed:   false,
		},
		{
			ID:          2,
			Title:       "Tarea de ejemplo 2",
			Description: "Esta es otra tarea de ejemplo",
			CreatedAt:   time.Now().Add(-12 * time.Hour),
			UpdatedAt:   time.Now().Add(-12 * time.Hour),
			Completed:   true,
		},
	}
	nextID = 3

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/tasks", tasksHandler)
	http.HandleFunc("/tasks/", taskByIDHandler)
	http.HandleFunc("/health", healthHandler)

	fmt.Println("Servidor escuchando en :8080")
	fmt.Println("Rutas disponibles:")
	fmt.Println(" - GET /tasks - Obtener todas las tareas")
	fmt.Println(" - POST /tasks - Crear una nueva tarea")
	fmt.Println(" - GET /tasks/{id} - Obtener una tarea por ID")
	fmt.Println(" - PUT /tasks/{id} - Actualizar una tarea por ID")
	fmt.Println(" - DELETE /tasks/{id} - Eliminar una tarea por ID")
	fmt.Println(" - GET /health - Verificar estado del servidor")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
