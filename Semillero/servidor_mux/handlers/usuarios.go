package handlers

import (
	"encoding/json"
	"net/http"
)

func ObtenerUsuarios(w http.ResponseWriter, r *http.Request) {
	usuarios := make([]string, 0)
	usuarios = append(usuarios, "Juan", "Pedro", "Maria")
	json.NewEncoder(w).Encode(usuarios)
}
