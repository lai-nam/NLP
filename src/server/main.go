package main

import (
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"

	"db"
	"nlp/document"
	"nlp/partofspeech"
	"nlp/word"
)

var (
	docCtrl      *document.DocumentController
	partOfSpeech *partofspeech.PartOfSpeechController
	newWord      *word.WordController
)

func init() {
	db.Connect("127.0.0.1")

	docCtrl = document.NewDocumentCtrl(db.Config{
		CollectionName: "document",
	})

	partOfSpeech = partofspeech.NewPartOfSpeechCtrl(db.Config{
		CollectionName: "partofspeech",
	})
	newWord = word.NewWord(db.Config{
		CollectionName: "word",
	})
}
func main() {
	router := mux.NewRouter()

	router.PathPrefix("/api/documents/").Methods("POST").HandlerFunc(docCtrl.Post)
	router.PathPrefix("/documents/").Methods("GET").HandlerFunc(docCtrl.Index)
	router.PathPrefix("/api/documents/").Methods("GET").HandlerFunc(docCtrl.Get)

	router.PathPrefix("/api/partofspeechs/").Methods("GET").HandlerFunc(partOfSpeech.List)
	router.PathPrefix("/api/partofspeechs/").Methods("POST").HandlerFunc(partOfSpeech.Post)
	router.PathPrefix("/api/partofspeechs/").Methods("GET").HandlerFunc(partOfSpeech.Index)
	router.PathPrefix("/api/words/").Methods("POST").HandlerFunc(newWord.Post)
	router.PathPrefix("/api/words/").Methods("GET").HandlerFunc(newWord.Get)
	router.PathPrefix("/words/").Methods("GET").HandlerFunc(newWord.Index)

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
