package main

import (
	"bolt"
	"fmt"
	"os"
)

const dbFile = "blockChain.db"
const blockBucket = "bucket"
const lastHashKey = "key"

type BlockChain struct {
	//数据库操作句柄
	db   *bolt.DB
	tail []byte
}

func InitBlockChain() *BlockChain {
	if isDBExist() {
		fmt.Println("blockchain exist already, no need to create!")
		os.Exit(1)
	}

	db, err := bolt.Open(dbFile, 0600, nil)
	CheckErr("InitBlockChain", err)

	var lastHash []byte
	db.Update(func(tx *bolt.Tx) error {
		genesis := NewGenesisBlock()
		bucket, err := tx.CreateBucket([]byte(blockBucket))
		CheckErr("InitBlockChain1", err)
		err1 := bucket.Put(genesis.Hash, genesis.Serialize()) //TODO
		CheckErr("InitBlockChain2", err1)
		err2 := bucket.Put([]byte(lastHashKey), genesis.Hash)
		CheckErr("InitBlockChain3", err2)
		lastHash = genesis.Hash

		return nil
	})
	return &BlockChain{db, lastHash}
}

func isDBExist() bool  {
	_, err := os.Stat(dbFile)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func GetBlockChainHandler() *BlockChain {
	if !isDBExist() {
		fmt.Println("Please create blockchain first!")
		os.Exit(1)
	}
	db, err := bolt.Open(dbFile, 0600, nil)
	CheckErr("GetBlockChainHandler", err)

	var lastHash []byte
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket != nil {
			lastHash = bucket.Get([]byte(lastHashKey))
		}

		return nil
	})
	return &BlockChain{db, lastHash}
}

func (bc *BlockChain) AddBlock(data string) {
	var prevBlockHash []byte

	bc.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			os.Exit(1)
		}

		prevBlockHash = bucket.Get([]byte(lastHashKey))
		return nil
	})
	block := NewBlock(data, prevBlockHash)
	err3 := bc.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			os.Exit(1)
		}

		err1 := bucket.Put(block.Hash, block.Serialize())
		CheckErr("AddBlock1", err1)
		err2 := bucket.Put([]byte(lastHashKey), block.Hash)
		CheckErr("AddBlock2", err2)
		bc.tail = block.Hash
		return nil
	})
	CheckErr("AddBlock3", err3)
}

//迭代器
type BlockChainIterator struct {
	currHash []byte
	db       *bolt.DB
}

func (bc *BlockChain) NewIterator() *BlockChainIterator {
	return &BlockChainIterator{currHash: bc.tail, db: bc.db}
}

func (it *BlockChainIterator) Next() (block *Block) {
	err := it.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			return nil
		}

		data := bucket.Get(it.currHash)
		block = Deserialize(data)
		it.currHash = block.PrevBlockHash
		return nil
	})
	CheckErr("Next()", err)
	return
}
