package main

import (
	"net/http"

	"firstimelang/controllers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	assetHandler := http.FileServer(http.Dir("./assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", assetHandler))

	video := controllers.NewVideo()
	r.HandleFunc("/", video.New).Methods("GET")
	r.HandleFunc("/new", video.New).Methods("GET")
	r.HandleFunc("/new", video.Show).Methods("POST")
	http.ListenAndServe("localhost:8080", r)
}

//http://video.google.com/timedtext?lang=de&v=dL5oGKNlR6I
