package main

import (
	"fmt"
	"github.com/genesis717-clone/genesiscoin/blockchain"
	"html/template"
	"log"
	"net/http"
)

const (
	port = ":4000"
	templatesDir = "templates/"
)

var templates *template.Template

type homeData struct {
	PageTitle string
	Blocks []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	data := homeData{"Home", blockchain.GetBlockchain().GetAllBlocks()}
	templates.ExecuteTemplate(rw, "home", data)
}

func main() {
	templates = template.Must(template.ParseGlob(templatesDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templatesDir + "partials/*.gohtml"))
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
