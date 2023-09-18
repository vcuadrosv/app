package models

type Equipos struct {
	Id_equipo     string `json:"Id_equipo"`
	Nombre_equipo string `json:"Nombre_equipo"`
}

type Marcador struct {
	Id_marcador_partido         string `json:"Id_marcador_partido"`
	Puntuacion_equipo_local     string `json:"Puntuacion_equipo_local"`
	Puntuacion_equipo_visitante string `json:"Puntuacion_equipo_visitante"`
	Tiempo_de_juego             string `json:"Tiempo_de_juego"`
	Resultado_final             string `json:"Resultado_final"`
	Fecha                       string `json:"Fecha"`
}

type Partidos struct {
	Equipo_local_id     string `json:"Equipo_local_id"`
	Equipo_visitante_id string `json:"Equipo_visitante_id"`
}
