package main

import (
	"io"
	"log"
	"net/http"
)

// Endpoints - Servicio en especifico
// Handlers

// Lo que se retorna segun lo que pase en el servicio (Estandares)
// 200 OK
// 300 Redirect
// 400 Bad Request - Client Error
// 500 Internal Server Error - Server Error
// API REST - RESTFUL - Cuendo se utilizan los estandares

func main() {

	// Crear handler
	http.HandleFunc("/usuarios", func(w http.ResponseWriter, r *http.Request) {

		// Para saber desde que metodo se esta accediendo
		if r.Method == "GET" {
			// Para generar un objeto de respuesta
			io.WriteString(w, "Servicio de usuarios")
		} else {
			http.NotFound(w, r)
		}

	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		println("Reenviando a Usuarios")
		http.Redirect(w, r, "/usuarios", 301)
	})

	log.Println("Escuchando en el puerto 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
	//Levanta el servidor en el puerto 8000 y si hay error
	// devuelve un mensaje

}
