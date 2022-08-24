package handlers

import (
	"log"
	"net/http"
	"os"
	"github.com/SebastianMxpwr/React-Mongo-Go/middlew"
	"github.com/SebastianMxpwr/React-Mongo-Go/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Handlers() {
	//manejador de las peticiones http
	router := mux.NewRouter()
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")


	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	//manejaro de cors
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
