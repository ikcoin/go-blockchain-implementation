package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

//Function to create the Hash based on the previous Hash and Data
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

//Function to create a block
func CreateBlock(data string, prevHash []byte) *Block {
	//block := &Block{[]byte{}, []byte(data), prevHash}
	block := new(Block) //constructor
	block.Hash = []byte{}
	block.Data = []byte(data)
	block.PrevHash = prevHash
	block.Nonce = 0 //initial nonce
	//block.DeriveHash()

	pow := NewProof(block)
	block.Nonce, block.Hash = pow.Run()

	return block
}

type Blockchain struct {
	Blocks []*Block
}

//Function to add a Block to the Blockchain
func (chain *Blockchain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)
}

//Function to create the GenesisBlock(first Block) of the Blockchain
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

//create the Blockchain
func InitBlockchain() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}} //Array of blocks with a call to Genesis function
}
