package main

import (
	"fmt"
	"strconv"

	"github.com/ikcoin/go-blockchain-implementation/blockchain"
)

func main() {

	chain := blockchain.InitBlockchain()

	chain.AddBlock("First Block")
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")
	fmt.Println()

	fmt.Println("-----BLOCKCHAIN:------")
	for blockNum, block := range chain.Blocks {
		fmt.Println("BLOCK ", blockNum)
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %x\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Println("Pow: ", strconv.FormatBool(pow.Validate()))
		fmt.Println("-----------")
		fmt.Println()
	}

}
