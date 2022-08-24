package bd

import (
	"context"
	"time"

	"github.com/SebastianMxpwr/React-Mongo-Go/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ChkUsuarioExistente(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twistor")
	col := db.Collection("users")

	condicion := bson.M{"email": email}

	var resultado models.Usuario
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}

	return resultado, true, ID

}
