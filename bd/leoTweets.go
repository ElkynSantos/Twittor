package bd

import (
	"context"
	"log"
	"time"

	Models "TRABAJO2/Twittor/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func LeoTweets (ID string, pagina int64) ([] *Models.DevuelvoTweets,bool){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("twitter")
	col := db.Collection("tweet")

	var resultados []*Models.DevuelvoTweets

	condicion := bson.M{
		"userid": ID,
	}

	opciones := options.Find()
	opciones.SetLimit(20)
	opciones.SetSort(bson.D{{Key: "fecha",Value: -1}})

	opciones.SetSkip((pagina-1)*20)

	cursor, err := col.Find(ctx,condicion,opciones)

	if err != nil {
		log.Fatal(err.Error())
		return resultados,false
	}

	for cursor.Next(context.TODO()){
		var registro Models.DevuelvoTweets
		err := cursor.Decode(&registro)

		if err != nil {

			return resultados,false
		}
		resultados = append(resultados, &registro)
	}

	return resultados,true
}