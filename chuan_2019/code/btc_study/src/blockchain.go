package main

import (
	"bolt"
	"log"
)

//4. 引入区块链
type BlockChain struct {
	// blocks []*Block
	db *bolt.DB
	tail []byte
}

const blockChainDb = "blockChain.db"
const blockBucket = "blockBucket"
const lashHashKey = "LashHashKey"

//5.  定义一个区块链
func NewBlockChain() *BlockChain  {
	db, err := bolt.Open(blockChainDb, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	var lastHash []byte

	db.Update(func(tx *bolt.Tx) error {
		// 2. 找到抽屉bucket
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil{
			//没有抽屉，我们需要创建
			bucket, err = tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				log.Panic(err)
			}
			// 3. 写数据
			genesisBlock := GenesisBlock()
			bucket.Put(genesisBlock.Hash, genesisBlock.toByte())
			bucket.Put([]byte(lashHashKey), genesisBlock.Hash)
		}
		lastHash = bucket.Get([]byte(lashHashKey))
		return nil
	})
	return &BlockChain{
		db, lastHash,
	}
}

func GenesisBlock() *Block {
	return NewBlock("Go一期创世块， 老牛逼了！ ", []byte{})
}

//6. 添加区块
func (bc *BlockChain)AddBlock(data string)  {
	//// a. 创建新的区块
	//lastBlock := bc.blocks[len(bc.blocks) - 1]
	//prevHash := lastBlock.Hash
	//block := NewBlock(data, prevHash)
	//// b. 添加到区块链中
	//bc.blocks = append(bc.blocks, block)
}