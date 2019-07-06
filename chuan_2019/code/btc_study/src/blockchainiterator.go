package main

import (
	"bolt"
	"log"
)

type BlockChainIterator struct {
	db *bolt.DB
	// 游标
	currentHashPointer []byte
}

//func NewIterator(bc *BlockChain)  {
//
//}

func (bc *BlockChain) NewIterator() *BlockChainIterator {
	return &BlockChainIterator{
		db: bc.db,
		currentHashPointer: bc.tail,
	}
}

//1. 返回当前的区块
//2. 指针前移
func (it *BlockChainIterator) Next() *Block {
	var block Block
	it.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			log.Panic("迭代器遍历时，bucket不应该为空")
		}
		blockTmp := bucket.Get(it.currentHashPointer)
		block.Deserialize(blockTmp)
		it.currentHashPointer = block.PrevHash
		return nil
	})
	return &block
}