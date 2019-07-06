package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

type ProofOfWork struct {
	block *Block
	target *big.Int
}

func NewProofOfWork(block *Block) *ProofOfWork  {
	pow := ProofOfWork{
		block: block,
	}

	// 我们指定的难度值， 现在是一个string类型，需要进行转换
	targetStr := "000001bf474e8f1f46797bcb4d3454832fd8ef867595889dbfa5b733944d9280"
	tmpInt := big.Int{}
	tmpInt.SetString(targetStr, 16)
	pow.target = &tmpInt
	return &pow
}

func (pow *ProofOfWork)Run() ([]byte, uint64)  {
	//1. 拼装数据 （区块的数据， 还有不断变化的随机数
	// 2. 做哈希运算
	// 3. 与pow中的target进行比较
	// a . 找到了，退出返回
	// b. 断续找，

	var nonce uint64
	block := pow.block
	var hash [32]byte

	fmt.Println("开始挖矿...")
	for {
		//1. 拼装数据 （区块的数据， 还有不断变化的随机数
		tmp := [][]byte{
			Uint64ToByte(block.Version),
			block.PrevHash,
			block.MerkelRoot,
			Uint64ToByte(block.TimeStamp),
			Uint64ToByte(block.Difficulty),
			Uint64ToByte(nonce),
			block.Data,
		}
		blockInfo := bytes.Join(tmp, []byte{})
		// 2. 做哈希运算
		hash = sha256.Sum256(blockInfo)
		// 3. 与pow中的target进行比较
		tmpInt := big.Int{}
		tmpInt.SetBytes(hash[:])
		if tmpInt.Cmp(pow.target) == -1 {
			// a . 找到了，退出返回
			fmt.Printf("挖矿成功！hash:%x nonce:%d\n", hash[:], nonce)
			break
		} else {
			// b. 断续找，
			nonce++
		}
	}

	return hash[:], nonce
}
