package bd

import (
	Models "TRABAJO2/Twittor/models"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)



func BuscoPerfil (ID string) (Models.Usuario, error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second *15)

	defer cancel()

	db:= MongoC.Database("twitter")
	col := db.Collection("usuarios")

	var perfil Models.Usuario

	objID, _ := primitive.ObjectIDFromHex(ID)
	fmt.Println(objID)
	condicion := bson.M {
		"_id": objID,
	}



	err:= col.FindOne(ctx,condicion).Decode(&perfil)
	perfil.Password=""

	if err != nil {
		fmt.Println("Registro no encontrado "+err.Error())
		return perfil, err

	}
	return perfil, nil
}