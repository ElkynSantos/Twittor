package bd

import (
	"context"
	"time"

	Models "TRABAJO2/Twittor/models"
)

func BorroRelacion(t Models.Relacion)(bool, error){

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("twitter")
	col := db.Collection("relacion")

	_, err:= col.DeleteOne(ctx,t)
	
	if err == nil {
		return false, nil
	}

	return true , nil;
	
}