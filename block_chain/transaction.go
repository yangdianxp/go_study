package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"os"
)

const reward = 12.5

type Transaction struct {
	//交易ID
	TXID []byte
	//输入
	TXInputs []TXInput
	//输出
	TXOutputs []TXOutput
}

type TXInput struct {
	//所引用输出的交易ID
	TXID []byte
	//所引用output的索引值
	Vout int64
	//解锁脚本，指明可以使用某个output的条件
	ScriptSig string
}

//检查当前的用户能否解开引用的utxo
func (input *TXInput)CanUnlockUTXOWith(unlockData string) bool  {
	return input.ScriptSig == unlockData
}

type TXOutput struct {
	//支付给收款方的金额
	Value float64
	//锁定脚本， 指定收款方的地址
	ScriptPubKey string
}

//检查当前用户是否是这个utxo的所有者
func (output *TXOutput)CanBeUnlockedWith(unlockData string) bool {
	return output.ScriptPubKey == unlockData
}

//设置交易ID，是一个哈希值
func (tx *Transaction)SetTXID() {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(tx)
	CheckErr("SetTXID", err)
	hash := sha256.Sum256(buffer.Bytes())
	tx.TXID = hash[:]
}

//创建coinbase交易，只有收款人，没有付款人，是矿工的奖励交易
func NewCoinbaseTx(address string, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("reward to %s %d btc", address, reward)
	}
	input := TXInput{[]byte{}, -1, data}
	output := TXOutput{reward, address}
	tx := Transaction{[]byte{}, []TXInput{input}, []TXOutput{output}}
	tx.SetTXID()
	return &tx
}

func (tx *Transaction)IsCoinbase() bool {
	if len(tx.TXInputs) == 1{
		if len(tx.TXInputs[0].TXID) == 0 && tx.TXInputs[0].Vout == -1 {
			return true
		}
	}
	return false
}

//创建普通交易send的辅助函数
func NewTransaction(from string, to string, amount float64, bc *BlockChain) *Transaction {
	//map[string][]int64 key ： 交易id, value: 引用output的索引数组
	validUTXOs := make(map[string][]int64)
	var total float64
	validUTXOs/*所需要的，合理的utxo的集合*/, total/*返回utxo的金额总和*/ = bc.FindSuitableUTXOs(from, amount)
	//validUTXOs[0x1111111111] = []int64{1]
	//validUTXOs[0x2222222222] = []int64{0]
	//...
	//validUTXOs[0xnnnnnnnnnnn] = []int64{0, 4, 8}
	if total < amount {
		fmt.Println("Not enough money!")
		os.Exit(1)
	}
	var inputs []TXInput
	var outputs []TXOutput
	///1.创建inputs
	//进行output到input的转换
	//遍历有效utxo的合集
	for txId, outputIndexes := range validUTXOs {
		//遍历所有引用的utxo的索引，每一个索引需要创建五个input
		for _, index := range outputIndexes {
			input := TXInput{[]byte(txId), int64(index), from}
			inputs = append(inputs, input)
		}
	}
	//2. 创建output
	//给对方支付
	output := TXOutput{amount, to}
	outputs = append(outputs, output)
	//找零钱
	if total > amount {
		output := TXOutput{total - amount, from}
		outputs = append(outputs, output)
	}

	tx := Transaction{nil, inputs, outputs}
	tx.SetTXID()
	return &tx
}