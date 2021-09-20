package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
	"math/rand"
)

/*
Simple Proof of work:
	1. Take the data from the block
	2. Create a random counter (nonce)
	3. Create the hash o the data plus the nonce
	4. Check the hash to see if it meets a set of requirements, if not, go back to point 2 and start with new nonce

Requirements:
	1. The firsts few bytes must contains 0s
*/

const Difficulty = 12 //By now, the difficulty will be static

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty)) //Left-shift

	pow := &ProofOfWork{b, target}
	return pow
}

func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
			ToHex(int64(nonce)),
			ToHex(int64(Difficulty)),
		},
		[]byte{},
	)
	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte
	nonce := 0
	finish := false

	for !finish {
		nonce = rand.Intn(math.MaxInt64)

		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)

		fmt.Println(hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target) == -1 {
			finish = true //The solution meets the requirements
		}
	}
	fmt.Println()

	return nonce, hash[:]
}

// Aux function
func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}
