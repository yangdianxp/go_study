package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"math/big"
)

const reward float64 = 50

// 1. 定义交易结构
type Transaction struct {
	TXID []byte  // 交易ID
	TXInputs []TXInput  //交易输入数组
	TXOutputs []TXOutput //交易输出数组
}
// 定义交易输入
type TXInput struct {
	// 引用的交易ID
	TXid []byte
	// 引用的output的索引值
	Index int64
	// 解锁脚本， 我们用地址来模拟
	// Sig string
	Signature []byte  // 真正的数字签名， 由r, s 拼成的[]byte
	//
	PubKey []byte  // 拼接字符串
}

// 定义交易输出
type TXOutput struct {
	// 转账金额
	Value float64
	// 锁定脚本, 用地址模拟
	// 收款方的公钥哈希
	PubKeyHash []byte
}

// 由于现在存储的字段是地址的公钥哈希，所以无法直接创建TXOutput
func (output *TXOutput)Lock(address string) {
	pubKeyHash := GetPubKeyFromAddress(address)
	output.PubKeyHash = pubKeyHash
}

func NewTXOutput(value float64, address string) *TXOutput  {
	output := TXOutput{
		Value: value,
	}
	output.Lock(address)
	return &output
}

// 设置交易ID
func (tx *Transaction) SetHash()  {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	data := buffer.Bytes()
	hash := sha256.Sum256(data)
	tx.TXID = hash[:]
}
// 判断当前的交易是否为挖矿交易
func (tx *Transaction) IsCoinbase() bool {
	// 1. 交易的input只有一个
	// 2. 交易id为空
	// 3. 交易的index 为 -1
	//if len(tx.TXInputs) == 1 {
	//	input := tx.TXInputs[0]
	//	if bytes.Equal(input.TXid, []byte{}) && input.Index == -1{
	//		return true
	//	}
	//}
	if len(tx.TXInputs) == 1 && len(tx.TXInputs[0].TXid) == 0 && tx.TXInputs[0].Index == -1 {
		return true
	}
	return false
}

// 2. 提供创建交易方法（挖矿交易）
func NewCoinbaseTX(address string, data string) *Transaction  {
	// 挖矿交易的特点：
	// 1. 只有一个input
	// 2. 无需引用交易id
	// 3. 无需引用index
	// 矿工由于挖矿时无需指定签名，所以这个sig字段可以由矿工自由填写数据，一般是填写矿池的名字
	// 签名先填写为空， 后面创建完整交易后，最后做一次签名即可
	input := TXInput{[]byte{}, -1, nil, []byte(data)}
	output := NewTXOutput(reward, address)
	// 对于挖矿交易来说，只有一个input和一个output
	tx := Transaction{[]byte{}, []TXInput{input}, []TXOutput{*output}}
	tx.SetHash()
	return &tx
}
// 3. 创建挖矿交易
// 4. 根据交易调整程序

// 创建普通交易
// 1. 找到最合理UTXO集合 map[string][]uint64
// 2. 将这些UTXO逐一转成input
// 3. 创建outputs
// 4. 如果有零钱，要找零
func NewTransaction(from, to string, amount float64, bc* BlockChain) *Transaction  {
	// 创建交易之后要进行数字签名，所以需要私钥， 打开钱包（NewWallets()）
	// 找到自己的钱包，根据地址返回自己的wallet
	// 得到对应的公钥， 私钥
	ws := NewWallets()
	wallet := ws.WalletsMap[from]
	if wallet == nil {
		fmt.Printf("没有找到该地址的钱包，交易创建失败")
		return nil
	}
	pubKey := wallet.PubKey
	privateKey := wallet.Private  // 稍后再用
	pubKeyHash := HashPubKey(pubKey)

	utxos, resValue := bc.FindNeedUTXOs(pubKeyHash, amount)
	if resValue < amount {
		fmt.Printf("余额不足，交易失败！")
		return nil
	}
	var inputs []TXInput
	var outputs []TXOutput
	for id, indexArray := range utxos {
		for _, i := range indexArray {
			input := TXInput{[]byte(id), int64(i), nil, pubKey}
			inputs = append(inputs, input)
		}
	}
	output := NewTXOutput(amount, to)
	outputs = append(outputs, *output)
	if resValue > amount {
		// 找零
		output = NewTXOutput(resValue - amount, from)
		outputs = append(outputs, *output)
	}
	tx := Transaction{[]byte{}, inputs, outputs}
	tx.SetHash()

	bc.SignTransaction(&tx, privateKey)

	return &tx
}

// 签名的具体实现, 参数为：私钥， inputs里面所有引用的交易的结构，map[string]Transaction
func (tx *Transaction) Sign(privateKey *ecdsa.PrivateKey, prevTxs map[string]Transaction)  {
	if tx.IsCoinbase() {
		return
	}
	//具体签名的动作先不管，稍后继续
	// 1. 创建一个当前交易的copy： TrimmedCopy: 要把Signature和PubKey字段设置nil
	txCopy := tx.TrimmedCopy()
	// 2. 循环遍历txCopy的input, 得到这上input索引的output的公钥hash
	for i, input := range txCopy.TXInputs {
		prevTX := prevTxs[string(input.TXid)]
		if len(prevTX.TXID) == 0 {
			log.Panic("引用的交易无效")
		}
		// 不要对
		txCopy.TXInputs[i].PubKey = prevTX.TXOutputs[input.Index].PubKeyHash

		// 所需要的三个数据都具备了， 开始做哈希处理
		txCopy.SetHash()
		txCopy.TXInputs[i].PubKey = nil
		signData := txCopy.TXID
		r, s, err := ecdsa.Sign(rand.Reader, privateKey, signData)
		if err != nil {
			log.Panic(err)
		}
		signature := append(r.Bytes(), s.Bytes()...)
		tx.TXInputs[i].Signature = signature
	}
	// 3. 生成要签名的数据，要签名的数据一定是哈希值
	// a. 我们对每一个input都要签名一次， 签名的数据是由当前input引用的output的哈希+当前的outputs(都承载在当前的txCopy里面)
	// b. 要对这个拼好的txCopy进行哈希处理， SetHash得到TXID，这个TXID就是我们要签名最终
	// 4. 执行签名动作得到 r,s 字节流
	// 5. 放到我们所签名的input的Signature
}

func (tx *Transaction) TrimmedCopy() Transaction {
	var inputs []TXInput
	var outputs []TXOutput
	for _, input := range tx.TXInputs {
		inputs = append(inputs, TXInput{input.TXid, input.Index, nil, nil})
	}
	for _, output := range tx.TXOutputs {
		outputs = append(outputs, output)
	}
	return Transaction{tx.TXID, inputs, outputs}
}

// 分析检验
func (tx *Transaction) Verify(prevTxs map[string]Transaction) bool {
	if tx.IsCoinbase() {
		return true
	}
	// 得到签名的数据
	txCopy := tx.TrimmedCopy()
	for i, input := range tx.TXInputs {
		prevTX := prevTxs[string(input.TXid)]
		if len(prevTX.TXID) == 0 {
			log.Panic("引用的交易无效")
		}
		txCopy.TXInputs[i].PubKey = prevTX.TXOutputs[input.Index].PubKeyHash
		txCopy.SetHash()
		dataHash := txCopy.TXID
		signature := input.Signature
		pubKey := input.PubKey  // 拆，

		var r, s big.Int
		r.SetBytes(signature[:len(signature) / 2])
		s.SetBytes(signature[len(signature) / 2:])

		var X, Y big.Int
		X.SetBytes(pubKey[:len(pubKey) / 2])
		Y.SetBytes(pubKey[len(pubKey) / 2:])

		pubKeyOrigin := ecdsa.PublicKey{elliptic.P256(), &X, &Y}
		if !ecdsa.Verify(&pubKeyOrigin, dataHash, &r, &s) {
			return false
		}
	}
	// 得到Signature, 反推r, s

	// 拆解PubKey, X, Y 得到原生公钥
	// 4. Verify
	return true
}
