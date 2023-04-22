package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoC=ConectarBD()
var clientOptions=options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0.ycbzwzk.mongodb.net/test")

/* ConectarBD ES LA FUNCION QUE CONECTA BASE DE DATOS*/
func ConectarBD() *mongo.Client {
	client,err:=mongo.Connect(context.TODO(),clientOptions)

	if err != nil {
		log.Fatal(err)
		return client
	}
	err = client.Ping(context.TODO(),nil)
	if err != nil {
		log.Fatal(err)
		return client
	}

	log.Println("Conexion valida")
return client
}
/* Checkconnection chequea conexion*/
func Checkconnection() int{
	err := MongoC.Ping(context.TODO(),nil)

	if err != nil{

		return 0
	}
return 1
}