package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
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
	block.Nonce, block.Hash = pow.Mining()

	return block
}

//Function to create the GenesisBlock(first Block) of the Blockchain
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)

	if err != nil {
		log.Panic(err)
	}

	return res.Bytes()
}

func Deserialize(data []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)

	if err != nil {
		log.Panic(err)
	}
	return &block
}
