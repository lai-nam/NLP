package document_test

// import (
// 	"bytes"
// 	"document"
// 	"encoding/json"
// 	"github.com/gorilla/mux"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"gopkg.in/mgo.v2"
// )

// var (
// 	router *mux.Router
// )

// func init() {
// 	router = mux.NewRouter()
// 	db := connectDatabase("127.0.0.1", "nlp_test")

// 	docCtrl := document.NewDocumentController(document.Config{
// 		CollectionName: "document",
// 		Db:             db,
// 	})

// 	router.PathPrefix("/documents/").Methods("GET").HandlerFunc(docCtrl.Index)
// 	router.PathPrefix("/api/documents/").Methods("POST").HandlerFunc(docCtrl.Post)
// 	router.PathPrefix("/api/documents/").Methods("GET").HandlerFunc(docCtrl.Get)
// }

// func connectDatabase(url, dbname string) *mgo.Database {
// 	session, err := mgo.Dial(url)
// 	if err != nil {
// 		panic("Could not connect to database: " + err.Error())
// 	}

// 	return session.DB(dbname)
// }

// func TestDocumentPost(t *testing.T) {
// 	doc := document.Document{
// 		Category: "News",
// 		Content:  "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
// 	}

// 	data, _ := json.Marshal(doc)

// 	body := bytes.NewReader(data)

// 	req, err := http.NewRequest("POST", "/documents/", body)

// 	if err != nil {
// 		t.Errorf("error %v", err)
// 		return
// 	}

// 	w := httptest.NewRecorder()

// 	router.ServeHTTP(w, req)

// 	if w.Code != 200 {
// 		t.Errorf("expected 200 but got %d", w.Code)
// 		return
// 	}

// 	t.Log("not", w.Body.String(), "dont")
// }

// func TestDocumentGet(t *testing.T) {

// }
