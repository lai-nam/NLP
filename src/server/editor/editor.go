package editor

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Editor struct {
	Collection *mgo.Collection

	Dir string
}

type Word struct {
	Id          bson.ObjectId `bson:"_id" json:"_id"`
	HeadWord    string        `bson:"headword" json:"headword"`
	WordType    string        `bson:"wordtype" json:"wordtype"`
	Category    string        `bson:"category" json:"category"`
	SubCategory string        `bson:"subcategory" json:"subcategory"`
	Defination  string        `bson:"defination" json:"defination"`
	Tag         string        `bson:"tag" json:"tag"`
	Eng         string        `bson:"eng" json:"eng"`
}

type Pagination struct {
	Page    int
	PerPage int
}

func New(opts ...func(*Editor)) *Editor {
	editor := &Editor{}

	for _, opt := range opts {
		opt(editor)
	}
	return editor
}

func (e *Editor) View(name string) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles(e.Dir + "/" + name))
		t.Execute(w, nil)
	}
}

func (e *Editor) List(pag Pagination) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		var words []Word

		u, err := url.Parse(r.URL.String())

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		values := u.Query()

		page := pag.Page
		perPage := pag.PerPage

		if v, ok := values["page"]; ok {
			page, _ = strconv.Atoi(v[0])
		}

		if v, ok := values["per_page"]; ok {
			perPage, _ = strconv.Atoi(v[0])
		}

		prePage := page - 1

		if prePage < 0 {
			prePage = 0
		}

		fmt.Println("pag :", pag, prePage)
		e.Collection.Find(nil).Skip(prePage * perPage).Limit(perPage).All(&words)

		body, err := json.Marshal(words)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(body)
	}

}

func (e *Editor) Get() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		index := strings.LastIndex(r.URL.String(), "/")
		id := r.URL.String()[index+1:]

		if !bson.IsObjectIdHex(id) {
			http.Error(w, id+"is not objectId", http.StatusBadRequest)
			return
		}
		var newWord Word
		err := e.Collection.FindId(bson.ObjectIdHex(id)).One(&newWord)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		body, _ := json.Marshal(newWord)

		w.WriteHeader(http.StatusOK)
		w.Write(body)
	}
}

func (e *Editor) Put() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		index := strings.LastIndex(r.URL.String(), "/")
		id := r.URL.String()[index+1:]

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var newWord Word

		err = json.Unmarshal(body, &newWord)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println(id, bson.IsObjectIdHex(id), bson.ObjectIdHex(id))
		if !bson.IsObjectIdHex(id) {
			http.Error(w, id+"is not objectid", http.StatusBadRequest)
			return
		}

		err = e.Collection.UpdateId(bson.ObjectIdHex(id), newWord)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(200)
		w.Write(body)
	}
}

func parseBody(r *http.Request, model interface{}) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err.Error())
		return
	}

	err = json.Unmarshal(body, model)
	if err != nil {
		panic(err.Error())
		return
	}
}
func (e *Editor) Create() func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		var newWord Word
		parseBody(r, &newWord)

		err := e.Collection.Insert(newWord)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		body, _ := json.Marshal(newWord)
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	}
}
