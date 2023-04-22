package bd

import (
	Models "TRABAJO2/Twittor/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func ChequeoYaExisteUsuario (email string) (Models.Usuario, bool, string){

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)


	defer cancel()

	db := MongoC.Database("twitter")
	col := db.Collection("usuarios")

	condicion := bson.M{"email":email}
	var resultado Models.Usuario

	err := col.FindOne(ctx,condicion).Decode(&resultado)
	ID:=resultado.ID.Hex()

	if err !=nil {
		return resultado,false, ID
	}

	return resultado, true, ID
}