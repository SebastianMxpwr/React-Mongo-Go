package routers

import (
	"encoding/json"
	"net/http"

	"github.com/SebastianMxpwr/React-Mongo-Go/bd"
	"github.com/SebastianMxpwr/React-Mongo-Go/models"
)

//esto es un metodo ya que no devuelve nada
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email es requerido", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "La contraseña debe de ser de almenos 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ChkUsuarioExistente(t.Email)
	if encontrado {
		http.Error(w, "Ya esta usuado el email", 400)
		return
	}

	_, status, err := bd.InsertarRegistro(t)

	if err != nil {
		http.Error(w, "Ya la cagaste mano"+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado insertar el regitro del usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
