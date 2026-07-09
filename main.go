package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"os"
	"strings"
)

type Peticion struct {
	Texto string `json:"Texto"`
	Opcion string `json:"Opcion"`
}

type Respuesta struct {
	Resultado string `json:"resultado"`
}


func main() {
	// Se define la ruta
	http.HandleFunc("/conversor-minuscula-mayuscula", func(w http.ResponseWriter, r *http.Request) {

		// Validamos el método
		if r.Method != http.MethodPost {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
			return
		}

		// se deben crear las variables
		var p Peticion
		var res Respuesta

		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}

		//En Go, el código ejecutable (como un switch, asignaciones, o llamadas a funciones) obligatoriamente debe vivir dentro de una función. Por ejemplo dentro de un HandleFunc
		switch p.Opcion {
			case "mayuscula":
				res.Resultado = strings.ToUpper(p.Texto)
			case "minuscula":
				res.Resultado = strings.ToLower(p.Texto)
			default:
				http.Error(w, "Opción no válida. Usa 'mayuscula' o 'minuscula'", http.StatusBadRequest)
				return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081" // Puerto por defecto en caso de falla
	}

	fmt.Println("Servidor corriendo en el puerto " + port + "...")
	err := http.ListenAndServe(":"+port, nil)

	// Si ocurre un error al levantar el servidor, se imprime
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
