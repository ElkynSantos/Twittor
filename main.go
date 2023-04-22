package main

import (
	Bd "TRABAJO2/Twittor/bd"
	handlers "TRABAJO2/Twittor/handlers"
	"log"
)
func main() {


	if Bd.Checkconnection()==0{
		log.Fatal("Sin Conexion a la BD")
		return
	}

	handlers.Manejadores()
}