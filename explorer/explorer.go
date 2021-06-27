package explorer

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/genesis-clone/genesiscoin/blockchain"
)

const (
	templatesDir = "explorer/templates/"
)

var (
	templates *template.Template
	port      string
)

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	data := homeData{"Home", blockchain.GetBlockchain().GetAllBlocks()}
	templates.ExecuteTemplate(rw, "home", data)
}

func add(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(rw, "add", nil)
	case "POST":
		r.ParseForm()
		data := r.Form.Get("blockData")
		blockchain.GetBlockchain().AddBlock(data)
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
	}
}

func Start(aPort int) {
	handler := http.NewServeMux()
	port = fmt.Sprintf(":%d", aPort)
	templates = template.Must(template.ParseGlob(templatesDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templatesDir + "partials/*.gohtml"))
	handler.HandleFunc("/", home)
	handler.HandleFunc("/add", add)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, handler))
}
