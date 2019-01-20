// proofOfWork
package main

import (
	"bytes"
	"fmt"
	"math"
	"math/big"
	"crypto/sha256"
)

type ProofOfWrok struct {
	block *Block

	//目标值
	target *big.Int
}

const targetBits = 24

func NewProofOfWork(block *Block) *ProofOfWrok {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	pow := ProofOfWrok{block: block, target: target}
	return &pow
}

func (pow *ProofOfWrok) PrepareData(nonce int64) []byte {
	block := pow.block
	temp := [][]byte{
		IntToByte(block.Version),
		block.PrevBlockHash,
		block.MerKelRoot,
		IntToByte(block.TimeStamp),
		IntToByte(targetBits),
		IntToByte(nonce),
		//block.Transactions.TransactionHash() //TODO
	}
	data := bytes.Join(temp, []byte{})
	return data
}

func (pow *ProofOfWrok) Run() (int64, []byte) {
	var nonce int64 = 0
	var hash [32]byte
	var hashInt big.Int

	fmt.Printf("begin mining...\n");
	fmt.Printf("target:%x\n", pow.target)
	for nonce < math.MaxInt64 {
		data := pow.PrepareData(nonce)
		hash = sha256.Sum256(data)

		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.target) == -1 {
			fmt.Printf("found nonce:%d, hash:%x\n", nonce, hash)
			break
		} else {
			//fmt.Printf("not found nonce, current nonce:%d, current hash:%x\n", nonce, hash)
			nonce++
		}
	}
	return nonce, hash[:]
}

func (pow *ProofOfWrok)IsValid() bool  {
	data := pow.PrepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	var hashInt big.Int
	hashInt.SetBytes(hash[:])
	return hashInt.Cmp(pow.target) == -1
}
