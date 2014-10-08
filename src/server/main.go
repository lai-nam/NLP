package main

import (
	"document"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
)

var (
	db = connectDatabase("127.0.0.1", "nlp")

	docCtrl = document.NewDocumentController(document.Config{
		CollectionName: "document",
		Db:             db,
	})
)

func main() {
	router := mux.NewRouter()

	router.PathPrefix("/api/documents/").Methods("POST").HandlerFunc(docCtrl.Post)
	router.PathPrefix("/documents/").Methods("GET").HandlerFunc(docCtrl.Index)
	router.PathPrefix("/api/documents/").Methods("GET").HandlerFunc(docCtrl.Get)

	router.PathPrefix("/js/").Methods("GET").
		Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("client/js"))))

	router.PathPrefix("/fonts/").Methods("GET").
		Handler(http.StripPrefix("/fonts/", http.FileServer(http.Dir("client/fonts"))))

	router.PathPrefix("/css/").Methods("GET").
		Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("client/css"))))

	log.Printf("starting server at port %d", 80)
	http.ListenAndServe(":80", router)
}

func connectDatabase(url, dbname string) *mgo.Database {
	session, err := mgo.Dial(url)
	if err != nil {
		panic("Could not connect to database: " + err.Error())
	}

	return session.DB(dbname)
}
