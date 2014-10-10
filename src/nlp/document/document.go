package document

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"

	"db"
)

type Config struct {
	CollectionName string
	Db             *mgo.Database
}

type Document struct {
	Category string `bson:"category" json:"category"`
	Content  string `bson:"content" json:"content"`
}

type DocumentController struct {
	Repo db.RepoI
}

func NewDocumentCtrl(config db.Config) *DocumentController {

	return &DocumentController{
		db.NewRepoS(config),
	}
}

func (d *DocumentController) Post(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)

	var doc Document
	err := json.Unmarshal(data, &doc)

	if err != nil {
		log.Printf("unmarshal error %v", err)
		return
	}

	Repo := d.Repo

	err = Repo.Create(doc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err = json.Marshal(doc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(data)
}

func (d *DocumentController) Index(w http.ResponseWriter, r *http.Request) {
	index := template.Must(template.ParseFiles("client/index.html"))
	index.Execute(w, nil)
}

func (d *DocumentController) Get(w http.ResponseWriter, r *http.Request) {

	var docs []Document

	err := d.Repo.List(nil, &docs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := json.Marshal(docs)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(data)
}
