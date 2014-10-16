package main

import (
	_ "encoding/json"
	"encoding/xml"
	"flag"
	"io/ioutil"
	"log"

	"gopkg.in/mgo.v2"
)

type Document struct {
	XMLName xml.Name `xml:"Document"`
	Entrys  []Entry  `xml:"Entry"`
}

type Entry struct {
	HeadWord  string
	WordType  string    `xml:"Morphology>WordType"`
	Syntactic Syntactic `xml:"Syntactic"`
	Semantic  string    `xml:"Semantic>def"`
}

type Word struct {
	HeadWord    string
	WordType    string
	Category    string
	SubCategory string
	Defination  string
	Tag         string
}

type Syntactic struct {
	Category    string
	SubCategory string
}

var (
	dir    = flag.String("dirxml", "vcl_xml", ``)
	dbName = flag.String("dbname", "nlp", ``)
	DB     *mgo.Database
)

func main() {

	flag.Parse()

	dirInfo, err := ioutil.ReadDir(*dir)
	if err != nil {
		log.Println(err)
		return
	}

	connectDB("127.0.0.1")

	for i := 0; i < len(dirInfo); i++ {
		fileInfo := dirInfo[i]

		fileName := *dir + "/" + fileInfo.Name()
		handleFileXML(fileName)
	}

}

func handleFileXML(filename string) {
	C := DB.C("dictionary")

	fileInfo, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err.Error())
	}

	var doc Document
	err = xml.Unmarshal(fileInfo, &doc)
	if err != nil {
		log.Println(err.Error())
		return
	}

	for i := 0; i < len(doc.Entrys); i++ {
		entry := doc.Entrys[i]

		word := Word{}
		word.HeadWord = entry.HeadWord
		word.WordType = entry.WordType

		if entry.Syntactic.Category == "A" {
			word.Tag = "JJ"
		}

		word.Category = entry.Syntactic.Category
		word.SubCategory = entry.Syntactic.SubCategory
		word.Defination = entry.Semantic
		C.Insert(word)
	}

}

func connectDB(url string) {
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err.Error())
	}

	if dbName == nil {
		panic("dbname is nil")
	}

	DB = session.DB(*dbName)
}
