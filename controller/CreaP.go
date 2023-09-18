package controller

import (
	"cuadrangular/database"
	"cuadrangular/models"
	"encoding/json"
	"fmt"
	"net/http"
)

// r request trae datos de entrada y Response es la respuesta que da la funcion al cliente
func CreaEquipo(w http.ResponseWriter, r *http.Request) {
	var nuevoEquipo models.Equipos
	err := json.NewDecoder(r.Body).Decode(&nuevoEquipo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("equipos", &nuevoEquipo)
	err = database.InsertNewTeam(&nuevoEquipo)
	if err != nil {
		// Manejar el error, por ejemplo, respondiendo con un error HTTP
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
