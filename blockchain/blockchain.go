package blockchain

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
