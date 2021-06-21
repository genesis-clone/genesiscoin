package main

import (
	"fmt"
	"github.com/genesis717-clone/genesiscoin/blockchain"
)

func main() {
	chain := blockchain.GetBlockchain()
	fmt.Println(chain)
}
