package db

import (
	"gopkg.in/mgo.v2"
)

var (
	DB *mgo.Database
)

type RepoI interface {
	List(query interface{}, models interface{}) error
	Get(id interface{}, model interface{}) error
	Create(model interface{}) error
	Update(id interface{}, model interface{}) error
	Delete(id interface{}) error
}

type Config struct {
	CollectionName string
}

type RepoS struct {
	C *mgo.Collection
}

func Connect(url string) {
	session, err := mgo.Dial(url)

	if err != nil {
		panic(err.Error())
		return
	}

	DB = session.DB("nlp")
}

func NewRepoS(config Config) RepoI {

	return &RepoS{C: DB.C(config.CollectionName)}
}

func (r *RepoS) List(query interface{}, models interface{}) error {
	return r.C.Find(query).All(models)
}

func (r *RepoS) Get(id interface{}, model interface{}) error {

	return r.C.FindId(id).One(&model)
}

func (r *RepoS) Update(id interface{}, model interface{}) error {

	return r.C.UpdateId(id, model)
}

func (r *RepoS) Delete(id interface{}) error {

	return r.C.RemoveId(id)
}

func (r *RepoS) Create(model interface{}) error {
	C := r.C

	return C.Insert(model)
}
