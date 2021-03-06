package main

import (
	"flag"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"

	. "server/editor"

	_ "server/lexicons"
	"server/training"
)

var (
	db *mgo.Database

	dir = flag.String("dir", "client", `-dir=???`)
)

func init() {
	db = connectDatabase("127.0.0.1", "nlp")

	// edit = editor.New(
	// 	func(e *editor.Editor) {
	// 		e.Collection = db.
	// 	}
	// )

}
func main() {

	flag.Parse()

	router := mux.NewRouter()
	staticSetup(router)

	// router for editor

	edit := New(
		func(e *Editor) {
			e.Collection = db.C("dictionary")
		},
		func(e *Editor) {
			e.Dir = *dir
		},
	)

	sub := router.PathPrefix("/v1").Subrouter()
	sub.Path("/editor").Methods("GET").HandlerFunc(edit.View("html/editor.html"))
	sub.Path("/datatable").Methods("GET").HandlerFunc(edit.View("html/datatable.html"))
	sub.Path("/editors/").Methods("GET").HandlerFunc(edit.List(Pagination{Page: 0, PerPage: 10}))
	sub.Path("/editor/{id}").Methods("PUT").HandlerFunc(edit.Put())
	sub.Path("/editor/{id}").Methods("GET").HandlerFunc(edit.Get())
	sub.Path("/editor/").Methods("POST").HandlerFunc(edit.Create())

	tr := training.New(func(t *training.TrainingController) {
		t.C = db.C("trainingset")
	},
		func(t *training.TrainingController) {
			t.Dir = *dir
		})

	sub.Path("/training/").Methods("GET").HandlerFunc(tr.List)
	sub.Path("/trainging/{id}").Methods("GET").HandlerFunc(tr.Get)
	sub.Path("/training/").Methods("POST").HandlerFunc(tr.Create)
	sub.Path("/training").Methods("GET").HandlerFunc(tr.View("html/dataset.html"))

	// lex := lexicons.New(
	// 	func(l *lexicons.Editor) {
	// 		l.Collection = db.C("lexicon")
	// 	},
	// 	func(l *lexicons.Editor) {
	// 		l.Dir = *dir
	// 	})

	// sub.Path("/lexicon").Methods("GET").HandlerFunc(lex.View("html/lexicon.html"))
	// sub.Path("/lexicon/").Methods("GET").HandlerFunc(lex.List(lexicons.Pagination{Page: 0, PerPage: 10}))
	// sub.Path("/lexicon/{id}").Methods("PUT").HandlerFunc(lex.Put())
	// sub.Path("/lexicon/{id}").Methods("GET").HandlerFunc(lex.Get())
	// sub.Path("/lexicon/").Methods("POST").HandlerFunc(lex.Create())

	log.Printf("starting server at port %d", 80)
	http.ListenAndServe(":80", router)
}

func staticSetup(router *mux.Router) {
	router.PathPrefix("/js/").Methods("GET").
		Handler(http.StripPrefix("/js/", http.FileServer(http.Dir(*dir+"/js"))))

	router.PathPrefix("/fonts/").Methods("GET").
		Handler(http.StripPrefix("/fonts/", http.FileServer(http.Dir(*dir+"/fonts"))))

	router.PathPrefix("/css/").Methods("GET").
		Handler(http.StripPrefix("/css/", http.FileServer(http.Dir(*dir+"/css"))))

	router.PathPrefix("/images/").Methods("GET").
		Handler(http.StripPrefix("/images/", http.FileServer(http.Dir(*dir+"/images"))))
}

func connectDatabase(url, dbname string) *mgo.Database {
	session, err := mgo.Dial(url)
	if err != nil {
		panic("Could not connect to database: " + err.Error())
	}

	return session.DB(dbname)
}
