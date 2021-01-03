package routes

import (
	"fmt"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"io/ioutil"
	"log"
	"net/http"
	"wiki/db"
)

// https://www.golangprograms.com/example-to-handle-get-and-post-request-in-golang.html
func HandlePostWiki(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	name:= r.URL.Path[len("/wiki/"):]
	db.SaveTiddly(name, getBody(r))
	w.WriteHeader(http.StatusOK)
	printRequest(w,r)
}

func printRequest(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
	fmt.Fprintf(w, "Form values: %v\n", r.Form)
	//fmt.Fprintf(w, "Body: %v", getBodyAsString(r))
}

func getBody(r *http.Request) []byte {
	body, _ := ioutil.ReadAll(r.Body)
	return body
}

func HandleGetWiki(w http.ResponseWriter, r *http.Request){
	name:= r.URL.Path[len("/wiki/"):]
	log.Printf("Wiki called, name:%s\n", name)
	body:= db.FindTiddlyByName(name)
	if body == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Accept","application/json")
	fmt.Fprintf(w, "%s", parseMarkdown(body))
}

// TODO consider remove it from here
// https://github.com/gomarkdown/markdown
func parseMarkdown(body []byte) []byte{
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)
	return markdown.ToHTML(body,parser,nil)
}