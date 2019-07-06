package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"log"
	"time"
)

// 0. 定义结构
type Block struct {
	// 1. 版本号
	Version uint64
	// 2. 前区块哈希
	PrevHash []byte
	// 3. Merkel根 （梅克尔根， 这就是一个哈希值）
	MerkelRoot []byte
	// 4. 时间戳
	TimeStamp uint64
	// 5. 难度值
	Difficulty uint64
	// 6. 随机数，探矿要找的数据
	Nonce uint64

	// a. 当前区块哈希  正常比特币区块中没有当前区块的哈希，我们为了实现方便，简化
	Hash []byte
	// b. 数据
	Data []byte
}

//

// 2. 创建区块
func NewBlock(data string, prevBlockHash []byte) *Block  {
	block := Block{
		Version: 00,
		PrevHash: prevBlockHash,
		MerkelRoot: []byte{},
		TimeStamp: uint64(time.Now().Unix()),
		Difficulty: 0,
		Nonce: 0,
		Hash: []byte{},		// 先填空，后面再计算
		Data: []byte(data),
	}
	pow := NewProofOfWork(&block)
	hash, nonce := pow.Run()
	block.Hash = hash
	block.Nonce = nonce
	return &block
}

func Uint64ToByte(num uint64) []byte  {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		panic(err)
	}
	return buffer.Bytes()
}

//序列化
func (block *Block) Serialize() []byte  {
	var buffer bytes.Buffer
	// 使用gob进行序列化
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(&block)
	if err != nil {
		log.Panic("编码错误")
	}
	return buffer.Bytes()
}
//反序列化
func (block *Block) Deserialize(data []byte) error {
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(block)
	if err != nil {
		log.Panic("解码出错")
	}
	return err
}

func (block *Block) Print()  {
	fmt.Printf("=================================================\n")
	fmt.Printf("版本号： %d\n", block.Version)
	fmt.Printf("前区块哈希值: %x\n", block.PrevHash)
	fmt.Printf("梅克尔根: %x\n", block.MerkelRoot)
	fmt.Printf("时间戳: %d\n", block.TimeStamp)
	fmt.Printf("难度值: %d\n", block.Difficulty)
	fmt.Printf("随机值：%d\n", block.Nonce)
	fmt.Printf("当前区块哈希值: %x\n", block.Hash)
	fmt.Printf("交易数据: %s\n", block.Data)
	fmt.Printf("=================================================\n")
}

//// 3. 生成哈希
//func (block *Block)SetHash()  {
//	tmp := [][]byte{
//		Uint64ToByte(block.Version),
//		block.PrevHash,
//		block.MerkelRoot,
//		Uint64ToByte(block.TimeStamp),
//		Uint64ToByte(block.Difficulty),
//		Uint64ToByte(block.Nonce),
//		block.Data,
//	}
//	blockInfo := bytes.Join(tmp, []byte{})
//	hash := sha256.Sum256(blockInfo)
//	block.Hash = hash[:]
//
//
//	hashBlock := sha256.New()
//	hashBlock.Write(Uint64ToByte(block.Version))
//	hashBlock.Write(block.PrevHash)
//	hashBlock.Write(block.MerkelRoot)
//	hashBlock.Write(Uint64ToByte(block.TimeStamp))
//	hashBlock.Write(Uint64ToByte(block.Difficulty))
//	hashBlock.Write(Uint64ToByte(block.Nonce))
//	hashBlock.Write(block.Data)
//	// 2. sha256
//	block.Hash = hashBlock.Sum(nil)
//
//}
