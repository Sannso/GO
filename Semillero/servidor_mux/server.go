package main

import (
	"go_example/servidor_mux/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const PORT = ":8000"

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/usuarios", handlers.ObtenerUsuarios).Methods("GET", "OPTIONS")
	log.Println("Escuchando en el puerto 8000")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
