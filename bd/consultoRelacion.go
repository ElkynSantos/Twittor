package bd

import (
	"TRABAJO2/Twittor/models"
	Models "TRABAJO2/Twittor/models"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func ConsultoRelacion (t Models.Relacion) (bool,error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("twitter")
	col := db.Collection("relacion")

	condicion := bson.M{
		"usuarioid":t.UsuarioID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}

	var resultado models.Relacion
	fmt.Println(resultado)

	err:=col.FindOne(ctx,condicion).Decode(&resultado)

	if err!= nil {
		fmt.Println(err.Error())
		return false,err
	}

	return true, nil
}