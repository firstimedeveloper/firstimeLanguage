package main

import (
	"fmt"
	"net/http"

	"firstimelang/controllers"

	"github.com/gorilla/mux"
)

func main() {
	port := 8080
	r := mux.NewRouter()
	assetHandler := http.FileServer(http.Dir("./assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", assetHandler))

	video := controllers.NewVideo()
	r.HandleFunc("/", video.New).Methods("GET")
	r.HandleFunc("/new", video.New).Methods("GET")
	r.HandleFunc("/new", video.Show).Methods("POST")
	fmt.Println("Listening on port", port)
	http.ListenAndServe(fmt.Sprintf("localhost:%d", port), r)
}

//http://video.google.com/timedtext?lang=de&v=dL5oGKNlR6I
