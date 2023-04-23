package handlers

import (
	Middlew "TRABAJO2/Twittor/middleW"
	"TRABAJO2/Twittor/routers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Manejadores() {

    router := mux.NewRouter()

    
    router.HandleFunc("/registro", Middlew.ChequeoBD(routers.Registro)).Methods("POST")
    router.HandleFunc("/login", Middlew.ChequeoBD(routers.Login)).Methods("POST")
    router.HandleFunc("/verperfil", Middlew.ChequeoBD(Middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")
    router.HandleFunc("/modificarPerfil", Middlew.ChequeoBD(Middlew.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
    router.HandleFunc("/graboTweet", Middlew.ChequeoBD(Middlew.ValidoJWT(routers.GraboTweet))).Methods("POST")
    router.HandleFunc("/leoTweets", Middlew.ChequeoBD(Middlew.ValidoJWT(routers.LeoTweets))).Methods("GET")
    router.HandleFunc("/eliminarTweet", Middlew.ChequeoBD(Middlew.ValidoJWT(routers.EliminarTweet))).Methods("DELETE")
    
    router.HandleFunc("/subirAvatar", Middlew.ChequeoBD(Middlew.ValidoJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/obtenerAvatar", Middlew.ChequeoBD(routers.ObtenerAvatar)).Methods("GET")
	router.HandleFunc("/subirBanner", Middlew.ChequeoBD(Middlew.ValidoJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtenerBanner", Middlew.ChequeoBD(routers.ObtenerBanner)).Methods("GET")
    
    
    
    PORT := os.Getenv("PORT")

    if PORT == "" {
        PORT = "8080"
    }

    handler := cors.Default().Handler(router)

    log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
