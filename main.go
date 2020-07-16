package main

import (
	"fmt"
	"net/http"

	"firstimeLanguage.com/controllers"

	"github.com/gorilla/mux"
)

func main() {
	port := 8080
	r := mux.NewRouter()
	assetHandler := http.FileServer(http.Dir("./assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", assetHandler))
	uiHandler := http.FileServer(http.Dir("./ui/"))
	r.PathPrefix("/ui/").Handler(http.StripPrefix("/ui/", uiHandler))

	s := controllers.NewStatic()
	r.Handle("/", s.Home).Methods("GET")
	r.Handle("/contact", s.Contact).Methods("GET")
	r.Handle("/faq", s.Faq).Methods("GET")

	fmt.Println("Listening on port", port)
	http.ListenAndServe(fmt.Sprintf("localhost:%d", port), r)
}

//http://video.google.com/timedtext?lang=de&v=dL5oGKNlR6I
