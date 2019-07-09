package main

import (
	"bolt"
	"fmt"
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
func NewBlockChain(address string) *BlockChain  {
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
				genesisBlock := GenesisBlock(address)
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

func GenesisBlock(address string) *Block {
	coinbase := NewCoinbaseTX(address, "Go一期创世块，老牛逼了！")
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

//6. 添加区块
func (bc *BlockChain)AddBlock(txs []*Transaction) {
	db := bc.db
	lastHash := bc.tail
	// a. 创建新的区块
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			log.Panic("bucket 不应该为空， 请检查")
		}
		block := NewBlock(txs, lastHash)
		bucket.Put(block.Hash, block.Serialize())
		bucket.Put([]byte(lashHashKey), block.Hash)
		bc.tail = block.Hash
		return nil
	})
	// b. 添加到数据库中
}

// 找到指定地址的所有utxo
func (bc *BlockChain)FindUTXOs(address string) []TXOutput {
	var UTXO []TXOutput

	txs := bc.FindUTXOTransactions(address)

	for _, tx := range txs {
		for _, output := range tx.TXOutputs {
			if address == output.PubKeyHash {
				UTXO = append(UTXO, output)
			}
		}
	}

	return UTXO
}

func (bc *BlockChain) FindNeedUTXOs(from string, amount float64) (map[string][]uint64, float64) {
	//找到的合理的utxos集合
	utxos := make(map[string][]uint64)
	//  找到的utxos里面包含钱的总数
	var calc float64

	txs := bc.FindUTXOTransactions(from)

	for _, tx := range txs {
		for i, output := range tx.TXOutputs {
			if from == output.PubKeyHash {
				if calc < amount {
					utxos[string(tx.TXID)] = append(utxos[string(tx.TXID)], uint64(i))
					calc += output.Value

					if calc >= amount {
						fmt.Printf("找到了满足的金额: %f\n", calc)
						return utxos, calc
					}
				}
			}
		}
	}

	return utxos, calc
}

func (bc *BlockChain)FindUTXOTransactions(address string) []*Transaction  {
	var txs []*Transaction  // 存储所有包含utxo交易集合
	// 定义一个map来保存消费过的output, key是这个output的交易id, value 是这个交易中索引的数组
	// map[交易id][]int64
	spentOutputs := make(map[string][]int64)
	// 1. 遍历区块
	// 2. 遍历交易
	// 3. 遍历output, 找到和自己相关的utxo（在添加output之前检查一下是否已经消耗过）
	// 4. 遍历input, 找到自己花费过的utxo的集合（把自己消耗过的标示出来
	// 创建迭代器
	it := bc.NewIterator()
	for {
		block := it.Next()
		for _, tx := range block.Transactions {
			// fmt.Printf("current txid : %x\n", tx.TXID)

		OUTPUT:
			for i, output := range tx.TXOutputs {
				// fmt.Printf("current index : %d\n", i)
				// 在这里做一个过滤，将所有消耗过的outputs和当前的即将添加output对比一下
				// 如果相同，则路过，否则添加
				if spentOutputs[string(tx.TXID)] != nil {
					for _, j := range spentOutputs[string(tx.TXID)] {
						//
						if int64(i) == j {
							// 当前准备添加output已经消耗过了，不要再加了
							continue OUTPUT
						}
					}
				}
				// 这个output和我们目标的地址相同，满足条件，加到返回utxo数组中
				if output.PubKeyHash == address {
					// 返回所有包含我的UTXO的交易的集合
					txs = append(txs, tx)
					break
				}
			}

			// 如果当前交易是挖矿交易的话，那么不做遍历，直接路过
			if !tx.IsCoinbase() {
				for _, input := range tx.TXInputs {
					if input.Sig == address {
						spentOutputs[string(input.TXid)] = append(spentOutputs[string(input.TXid)], input.Index)
					}
				}
			} else {
				// fmt.Printf("这是coinbase,不做input 遍历\n")
			}
		}

		if len(block.PrevHash) == 0 {
			break
			fmt.Printf("区块遍历完成退出！")
		}
	}
	return txs
}