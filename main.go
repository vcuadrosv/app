package main

import (
	"cuadrangular/database"
	"cuadrangular/handlers"
	"fmt"
)

func main() {

	err := database.InitDB()
	if err != nil {
		fmt.Printf("Error al iniciar la conexión a la base de datos: %v\n", err)
		return
	}
	fmt.Println("conexiòn BD OK")
	handlers.MiWebServer()

}
