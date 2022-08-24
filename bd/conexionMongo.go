package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Se pone en mayusculas las funciones a exportar mientras que las internas se ponene en minusculas
var MongoCN = conectarBD()

//setar la direcion de la base de datos em Mongo atlas
var clientOptions = options.Client().ApplyURI("mongodb+srv://Admin:12345@mongo-go.b4dx7oi.mongodb.net/?retryWrites=true&w=majority")

func conectarBD() *mongo.Client {
	//contexto es un espacio en memoria que se comparte, ademas sirven para comunicar informacion entre ejecuciones
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexion Exitosa con la BD")
	return client

}

func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
