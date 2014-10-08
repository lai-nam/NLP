package document

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"
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
	Db *mgo.Database
	C  *mgo.Collection
}

func NewDocumentController(config Config) *DocumentController {

	return &DocumentController{
		Db: config.Db,
		C:  config.Db.C(config.CollectionName),
	}
}

func (d *DocumentController) Post(w http.ResponseWriter, r *http.Request) {
	C := d.C

	data, _ := ioutil.ReadAll(r.Body)

	var doc Document
	err := json.Unmarshal(data, &doc)

	if err != nil {
		log.Printf("unmarshal erro %v", err)
		return
	}

	err = C.Insert(doc)

	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	w.Write([]byte("ok"))
}

func (d *DocumentController) Index(w http.ResponseWriter, r *http.Request) {
	index := template.Must(template.ParseFiles("client/index.html"))
	index.Execute(w, nil)
}

func (d *DocumentController) Get(w http.ResponseWriter, r *http.Request) {

	C := d.C

	var docs []Document

	C.Find(nil).All(&docs)

	data, _ := json.Marshal(docs)

	w.Write(data)
}
