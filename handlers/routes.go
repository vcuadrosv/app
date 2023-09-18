package handlers

/* Sí, la carpeta handlers en una aplicación web en Go generalmente contiene
funciones de controlador (handlers) que se encargan de procesar las solicitudes
HTTP entrantes que provienen de los clientes. */

import (
	"cuadrangular/controller"
	"net/http"
	/* que es un paquete estándar en Go que proporciona funcionalidad
	   para crear aplicaciones web y manejar solicitudes HTTP. */)

// HandleRoot es un controlador para la ruta "/" r es lo que entra y w es lo que sale

func MiWebServer() {
	http.HandleFunc("/", home)
	http.HandleFunc("/crear", controller.CreaEquipo)
	http.HandleFunc("/consultarequipos", controller.InfoEquipo)
	http.ListenAndServe(":3000", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/index.html")
}
