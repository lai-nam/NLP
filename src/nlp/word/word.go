package word

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"

	"db"
)

var (
	tags = []string{"CC", "CD", "DT", "EX", "FW", "IN", "JJ", "JJR", "JJS", "LS", "MD", "NN", "NNS", "NNP", "NNPS", "PDT", "POS", "PRP", "PRP$", "RB", "RBR", "RBS", "RP", "SYM", "TO", "UH", "VB", "VBD", "VBG", "VBN", "VBP", "VBZ", "WDT", "WP", "WP$", "WRB"}
)

type WordController struct {
	Repo db.RepoI
}

type WordStruct struct {
	Word string
	Tag  string
}

func NewWord(config db.Config) *WordController {
	return &WordController{
		db.NewRepoS(config),
	}
}
func (word *WordController) Post(w http.ResponseWriter, r *http.Request) {

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var newWord WordStruct
	err = json.Unmarshal(data, &newWord)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	err = word.Repo.Create(newWord)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	w.Write([]byte("ok"))
}

func (word *WordController) Get(w http.ResponseWriter, r *http.Request) {

	var result []interface{}

	err := word.Repo.List(nil, &result)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	data, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	w.Write(data)
}

func (Word *WordController) Index(w http.ResponseWriter, r *http.Request) {
	index := template.Must(template.ParseFiles("client/html/word.html"))
	index.Execute(w, tags)
}
