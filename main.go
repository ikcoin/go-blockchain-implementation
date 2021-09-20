package main

import (
	"fmt"

	"github.com/ikcoin/go-blockchain-implementation/blockchain"
)

func main() {

	chain := blockchain.InitBlockchain()
	fmt.Println(string(chain.Blocks[0].Data))

	chain.AddBlock("First Block")
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %x\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println("-----------")
	}

}
