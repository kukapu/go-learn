package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Estructura para recibir los datos del cuerpo de la solicitud
type RequestData struct {
	Nombre string `json:"nombre"`
	Fuerte bool   `json:"fuerte"`
}

// manejador para el endpoint "/hello"
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "¡Hola, mundo!")
}

// manejador para el endpoint "/fuerte"
func fuerteHandler(w http.ResponseWriter, r *http.Request) {
	var data RequestData

	// Decodifica el cuerpo de la solicitud JSON en la estructura RequestData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	// Log de los datos recibidos
	log.Printf("Datos recibidos: Nombre=%s, Fuerte=%t", data.Nombre, data.Fuerte)

	// Verifica si el campo "fuerte" es verdadero o falso
	if data.Fuerte {
		// Devuelve el nombre en mayúsculas
		fmt.Fprintf(w, "%s", strings.ToUpper(data.Nombre))
	} else {
		// Devuelve el nombre en minúsculas
		fmt.Fprintf(w, "%s", strings.ToLower(data.Nombre))
	}
}

func main() {
	// Registra los manejadores para los endpoints
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/fuerte", fuerteHandler)

	// Inicia el servidor en el puerto 8080
	fmt.Println("Servidor escuchando en http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
