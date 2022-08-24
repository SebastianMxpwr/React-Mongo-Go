package main

import (
	"log"

	"github.com/SebastianMxpwr/React-Mongo-Go/bd"
	"github.com/SebastianMxpwr/React-Mongo-Go/handlers"
)

func main() {

	if bd.CheckConnection() == 0 {
		log.Fatal("Connectin failed")
		return
	}
	handlers.Handlers()

}
