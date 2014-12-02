package training

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Training struct {
	Id       bson.ObjectId `json:"_id" bson:"_id"`
	Document string        `json:"document" bson:"document"`
	Class    string        `json:"class" bson:"class"`
}

type TrainingController struct {
	C   *mgo.Collection
	Dir string
}

func New(opts ...func(t *TrainingController)) *TrainingController {
	T := &TrainingController{}

	for _, opt := range opts {
		opt(T)
	}

	return T
}

func (t *TrainingController) List(w http.ResponseWriter, r *http.Request) {

	var trs []Training
	e := t.C.Find(nil).All(&trs)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bytes, e := json.Marshal(trs)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func (t *TrainingController) Get(w http.ResponseWriter, r *http.Request) {
	index := strings.LastIndex(r.URL.String(), "/")

	id := r.URL.String()[index+1:]

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(id + "is not objectId"))
		return
	}

	var tr Training
	e := t.C.FindId(id).One(&tr)
	if e != nil {
		if strings.Contains(e.Error(), "not found") {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bytes, e := json.Marshal(tr)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func (t *TrainingController) Create(w http.ResponseWriter, r *http.Request) {

	bytes, _ := ioutil.ReadAll(r.Body)

	var tr Training

	e := json.Unmarshal(bytes, &tr)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tr.Id = bson.NewObjectId()
	e = t.C.Insert(tr)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(tr.Id.Hex()))
}

func (e *TrainingController) View(name string) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles(e.Dir + "/" + name))
		t.Execute(w, nil)
	}
}
