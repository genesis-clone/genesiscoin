package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data string
	hash string
	prevHash string
}

func main() {
	genesisBlock := block{"genesis block", "", ""}
	hash := sha256.Sum256([]byte(genesisBlock.data + genesisBlock.hash))
	hexHash := fmt.Sprintf("%x", hash)
	genesisBlock.hash = hexHash
	fmt.Println(genesisBlock)
}
