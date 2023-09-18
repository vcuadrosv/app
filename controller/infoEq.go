package controller

import (
	"cuadrangular/database"
	"encoding/json"
	"net/http"
)

// obtiene información de un equipo
func InfoEquipo(w http.ResponseWriter, r *http.Request) {
	//var equipos []models.Equipos

	equipos, err := database.SelectTeams()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Aquí puedes recuperar información de un equipo de la base de datos y responder con JSON
	//Equipos

	// Convertir los equipos a formato JSON
	jsonEquipos, err := json.Marshal(equipos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Establecer el encabezado de la respuesta como JSON
	w.Header().Set("Content-Type", "application/json")

	// Escribir la respuesta JSON en el objeto http.ResponseWriter
	w.Write(jsonEquipos)

}
