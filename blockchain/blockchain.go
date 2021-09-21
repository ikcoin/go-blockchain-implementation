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

//create the Blockchain
func InitBlockchain() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}} //Array of blocks with a call to Genesis function
}
