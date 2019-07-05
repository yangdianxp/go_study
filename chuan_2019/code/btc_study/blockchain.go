package main


//4. 引入区块链
type BlockChain struct {
	blocks []*Block
}

//5.  定义一个区块链
func NewBlockChain() *BlockChain  {
	genesisBlock := GenesisBlock()
	return &BlockChain{
		blocks: []*Block{genesisBlock},
	}
}

func GenesisBlock() *Block {
	return NewBlock("Go一期创世块， 老牛逼了！ ", []byte{})
}

//6. 添加区块
func (bc *BlockChain)AddBlock(data string)  {
	// a. 创建新的区块
	lastBlock := bc.blocks[len(bc.blocks) - 1]
	prevHash := lastBlock.Hash
	block := NewBlock(data, prevHash)
	// b. 添加到区块链中
	bc.blocks = append(bc.blocks, block)
}