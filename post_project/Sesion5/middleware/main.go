package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/home", home)
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("¡Bienvenido a la API de tareas! Aquí puedes gestionar tus tareas."))
}
