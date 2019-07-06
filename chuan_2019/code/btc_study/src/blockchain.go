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
				bucket.Put(genesisBlock.Hash, genesisBlock.Serialize())
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
	db := bc.db
	lastHash := bc.tail
	// a. 创建新的区块
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			log.Panic("bucket 不应该为空， 请检查")
		}
		block := NewBlock(data, lastHash)
		bucket.Put(block.Hash, block.Serialize())
		bucket.Put([]byte(lashHashKey), block.Hash)
		bc.tail = block.Hash
		return nil
	})
	// b. 添加到数据库中
}