package db

import (
	"log"
)

var db = initWiki()

// FIXME use a real database
func initWiki() map[string][]byte {
	var result = make(map[string][]byte)
	result["home"] = []byte("This is the home page")
	result["test"] = []byte("This is a test page")
	return result
}

func SaveTiddly(name string, body []byte){
	db[name] = body
	log.Printf("Tiddly saved [name: %s]\n", name)
}

func FindTiddlyByName(name string) []byte {
	log.Printf("Finding tiddly [name: %s]",name)
	return db[name]
}


