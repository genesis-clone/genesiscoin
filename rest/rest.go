package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/genesis-clone/genesiscoin/blockchain"
	"github.com/genesis-clone/genesiscoin/utils"
)

var port string

type url string

func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type urlDescription struct {
	URL         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []urlDescription{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         url("/blocks"),
			Method:      "POST",
			Description: "Add A Block",
			Payload:     "data:string",
		},
	}
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)
}

type addBlockBody struct {
	Message string
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(blockchain.GetBlockchain().GetAllBlocks())
	case "POST":
		var addBlockBody addBlockBody
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
		blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated)
	}
}

func Start(aPort int) {
	handler := http.NewServeMux()
	port = fmt.Sprintf(":%d", aPort)
	handler.HandleFunc("/", documentation)
	handler.HandleFunc("/blocks", blocks)
	fmt.Printf("Listening on http://localhost%s\n", port)
	http.ListenAndServe(port, handler)
}
