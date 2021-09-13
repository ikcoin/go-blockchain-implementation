package main

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
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
	block.DeriveHash()
	return block
}

func main() {

}
