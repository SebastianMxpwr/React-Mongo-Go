package bd

import (
	"context"
	"time"

	"github.com/SebastianMxpwr/React-Mongo-Go/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertarRegistro(u models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	//defer son funciones que se hacen al final
	defer cancel()

	db := MongoCN.Database("twistor")
	col := db.Collection("users")

	u.Password, _ = encriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
