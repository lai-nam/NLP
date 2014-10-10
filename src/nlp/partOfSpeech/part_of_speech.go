package partofspeech

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"

	"db"
)

type PartOfSpeech struct {
	Tag         string
	Description string
}

type PartOfSpeechController struct {
	Repo db.RepoI
}

func NewPartOfSpeechCtrl(config db.Config) *PartOfSpeechController {

	return &PartOfSpeechController{
		db.NewRepoS(config),
	}
}
func (p *PartOfSpeechController) List(w http.ResponseWriter, r *http.Request) {

	Repo := p.Repo

	var models []interface{}

	err := Repo.List(nil, &models)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := json.Marshal(models)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(data)

}

func (p *PartOfSpeechController) Post(w http.ResponseWriter, r *http.Request) {

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var data PartOfSpeech
	err = json.Unmarshal(bytes, &data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = p.Repo.Create(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("ok"))
}

func (p *PartOfSpeechController) Index(w http.ResponseWriter, r *http.Request) {

	index := template.Must(template.ParseFiles("client/partofspeech.html"))
	index.Execute(w, nil)
}
