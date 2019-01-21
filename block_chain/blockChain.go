package main

import (
	"bolt"
	"fmt"
	"os"
)

const dbFile = "blockChain.db"
const blockBucket = "bucket"
const lastHashKey = "key"
const genesisInfo = "genesisInfo"

type BlockChain struct {
	//数据库操作句柄
	db   *bolt.DB
	tail []byte
}

func InitBlockChain(address string) *BlockChain {
	if isDBExist() {
		fmt.Println("blockchain exist already, no need to create!")
		os.Exit(1)
	}

	db, err := bolt.Open(dbFile, 0600, nil)
	CheckErr("InitBlockChain", err)

	var lastHash []byte
	db.Update(func(tx *bolt.Tx) error {
		coinbase := NewCoinbaseTx(address, genesisInfo)
		genesis := NewGenesisBlock(coinbase)
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

func (bc *BlockChain) AddBlock(txs []*Transaction) {
	var prevBlockHash []byte

	bc.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			os.Exit(1)
		}

		prevBlockHash = bucket.Get([]byte(lastHashKey))
		return nil
	})
	block := NewBlock(txs, prevBlockHash)
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

//返回指定地址能够支配的utxo的交易集合
func (bc *BlockChain)FindUTXOTransactions(address string) []Transaction  {
	//包含目标utxo的交易集合
	var UTXOTransactions []Transaction
	//存储使用过的utxo的集合 map[交易id] []int64
	//0x111111 : 0, 1 都是给Alice转账
	spentUTXO := make(map[string][]int64)
	it := bc.NewIterator()
	for {
		//遍历区块
		block := it.Next()
		//遍历交易
		//目的： 找到所有能支配utxo
		for _, tx := range block.Transactions{
			//遍历input
			//目的： 找到已经消耗的utxo， 把它们放到一个集合里
			//需要两个字段来标识使用过的utxo: a. 交易ID， b.output的索引
			if !tx.IsCoinbase() {
				for _, input := range tx.TXInputs{
					if input.CanUnlockUTXOWith(address) {
						//map[txid] []int64
						spentUTXO[string(tx.TXID)] = append(spentUTXO[string(tx.TXID)], input.Vout)
					}
				}
			}
		OUTPUTS:
			//遍历output
			//目的：找到所有能支配utxo
			for currIndex, output := range tx.TXOutputs{
				//检查当前的output是否已经被消耗，如果消耗过，那么就进行下一个output检验
				if spentUTXO[string(tx.TXID)] != nil{
					//非空，代表当前交易里面有消耗的utxo
					indexes := spentUTXO[string(tx.TXID)]
					for _, index := range indexes {
						//当前的索引和消耗的索引比较，若相同，表明这个utxo肯定被消耗了，直接跳过，进行下一个output的判断
						if int64(currIndex) == index {
							continue OUTPUTS
						}
					}
				}
				//如果当前地址是这个utxo的所有者，就满足条件
				if output.CanBeUnlockedWith(address){
					UTXOTransactions = append(UTXOTransactions, *tx)
				}
			}
		}

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
	return UTXOTransactions
}

//寻找指定地址能够使用的utxo
func (bc *BlockChain)FindUTXO(address string) []*TXOutput {
	var UTXOs []*TXOutput
	txs := bc.FindUTXOTransactions(address)
	//遍历交易
	for _, tx := range txs {
		//遍历output
		for _, utxo := range tx.TXOutputs {
			//当前地址拥有的utxo
			if utxo.CanBeUnlockedWith(address) {
				UTXOs = append(UTXOs, &utxo)
			}
		}
	}
	return UTXOs
}
