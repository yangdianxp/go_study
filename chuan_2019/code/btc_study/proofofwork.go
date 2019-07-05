package main

import "math/big"

type ProofOfWork struct {
	block *Block
	target *big.Int
}

func NewProofOfWork(block *Block) *ProofOfWork  {
	pow := ProofOfWork{
		block: block,
	}

	// 我们指定的难度值， 现在是一个string类型，需要进行转换
	targetStr := "000019bf474e8f1f46797bcb4d3454832fd8ef867595889dbfa5b733944d9280"
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
	for {
		//1. 拼装数据 （区块的数据， 还有不断变化的随机数
		
	}

	return []byte("HelloWorld"), 10
}
