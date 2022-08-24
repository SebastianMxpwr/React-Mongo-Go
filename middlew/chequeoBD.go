package middlew

import (
	"net/http"

	"github.com/SebastianMxpwr/React-Mongo-Go/bd"
)

func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(w, "Conexion Perdida", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
