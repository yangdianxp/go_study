package main

import (
	"bolt"
)

const dbFile = "blockChain.db"
const blockBucket = "bucket"
const lastHashKey = "key"

type BlockChain struct {
	//数据库操作句柄
	db *bolt.DB
	tail []byte
}

func NewBlockChain() *BlockChain {
	db, err := bolt.Open(dbFile, mode:0x0600, nil)
	CheckErr("NewBlockChain1", err)
	
	var lastHash []byte
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket != nil {
			lastHash = bucket.Get(byte[](lastHashKey))
		} else {
			genesis := NewGenesisBlock()
			bucket, err := tx.CreateBucket([]byte(blockBucket))
			CheckErr("NewBlockChain2", err)
			err1 := bucket.Put(genesis.Hash, genesis.Serialize())//TODO
			CheckErr("NewBlockChain3", err1)
			err2 := bucket.Put([]byte(lastHashKey), genesis.Hash)
			CheckErr("NewBlockChain4", err2)
			lastHash = genesis.Hash
		}
		
		return nil
	})
	return &BlockChain{db, lastHash}
}

func (bc *BlockChain) AddBlock(data string) {
	prevBlockHash := bc.blocks[len(bc.blocks)-1].Hash
	block := NewBlock(data, prevBlockHash)
	bc.blocks = append(bc.blocks, block)
}

//迭代器
//type BlockChainIterator struct {
//	currHash []byte
//	db *bolt
//}
