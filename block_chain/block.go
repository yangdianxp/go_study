package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"time"
)

type Block struct {
	//版本
	Version int64
	//前区块的哈希值
	PrevBlockHash []byte
	//当前区块的哈希值，为了简化代码
	Hash []byte
	//梅克尔根
	MerKelRoot []byte
	//时间戳
	TimeStamp int64
	//难度值
	Bits int64
	//随机值
	Nonce int64

	//交易信息
	Transactions []*Transaction
}

func (block *Block) Serialize() []byte {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(block)
	CheckErr("Serialize", err)
	return buffer.Bytes()
}

func Deserialize(data []byte) *Block {
	if len(data) == 0 {
		return nil
	}

	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	CheckErr("Deserialize", err)
	return &block
}

func NewBlock(txs []*Transaction, preBlockHash []byte) *Block {
	var block Block
	block = Block{
		Version:       1,
		PrevBlockHash: preBlockHash,
		MerKelRoot: []byte{},
		TimeStamp:  time.Now().Unix(),
		Bits:       targetBits,
		Nonce:      0,
		Transactions: txs}

	//block.SetHash()
	pow := NewProofOfWork(&block)
	nonce, hash := pow.Run()
	block.Nonce = nonce
	block.Hash = hash
	return &block
}

func NewGenesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

//粗略模拟梅克尔树， 将交易的哈希值进行拼接，生成root hash
func (block *Block)HashTransactions() []byte  {
	var txHashes [][]byte
	txs := block.Transactions
	for _,tx := range txs{
									//[]byte
		txHashes = append(txHashes, tx.TXID)
	}
	//对二维切片进行拼接，生成一维切片
	data := bytes.Join(txHashes, []byte{})
	hash /*[32]byte*/:= sha256.Sum256(data)
	return hash[:]
}