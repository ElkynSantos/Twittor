package bd

import (
	Models "TRABAJO2/Twittor/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertoRegistro (u Models.Usuario) (string, bool , error){
	ctx,cancel := context.WithTimeout(context.Background(),15*time.Second)

	defer cancel()
	db := MongoC.Database("Twitter")
	col := db.Collection("usuarios")

	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx,u)

	if err != nil {
		return "",false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(),true, nil
	



}