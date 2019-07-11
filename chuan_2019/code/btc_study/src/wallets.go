package main

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"io/ioutil"
	"log"
	"os"
)

const walletFile = "wallet.data"

// 定一个Wallets结构，保存所有的wallet以及它的地址
type Wallets struct {
	WalletsMap map[string] *Wallet
}
// 创建方法
func NewWallets() *Wallets {
	//ws := loadFile()
	//return ws
	var ws Wallets
	ws.WalletsMap = make(map[string]*Wallet)
	ws.loadFile()
	return &ws
}

func (ws *Wallets)Print()  {
	for addr := range ws.WalletsMap {
		fmt.Printf("地址：%s\n", addr)
	}
}

func (ws *Wallets)CreateWallet() string {
	wallet := NewWallet()
	address := wallet.NewAddress()
	ws.WalletsMap[address] = wallet
	ws.saveToFile()
	return address
}



//保存方法，把新建的wallet添加进去
func (ws *Wallets) saveToFile()  {
	var buffer bytes.Buffer
	gob.Register(elliptic.P256())
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(ws)
	if err != nil {
		log.Panic(err)
	}
	ioutil.WriteFile(walletFile, buffer.Bytes(), 0600)
}

// 读取文件方法，把所有的钱包读出来
func (ws *Wallets) loadFile()  {
	_, err := os.Stat(walletFile)
	if os.IsNotExist(err) {
		return
	}

	content, err := ioutil.ReadFile(walletFile)
	if err != nil {
		log.Panic(err)
	}
	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(content))
	var wsLocal Wallets
	err = decoder.Decode(&wsLocal)
	if err != nil {
		log.Panic(err)
	}
	// ws.WalletsMap = wsLocal.WalletsMap
	*ws = wsLocal
}

func (ws *Wallets)GetAllAddresses() []string {
	var addresses []string
	for address := range ws.WalletsMap {
		addresses = append(addresses, address)
	}
	return addresses
}

//通过地址返回公钥hash
func GetPubKeyFromAddress(address string) []byte {
	addressByte := base58.Decode(address)  // 25字节
	pubKeyHash := addressByte[1:len(addressByte)-4]
	return pubKeyHash
}

func IsValidAddress(address string) bool {
	// 1. 解码
	addressByte := base58.Decode(address)
	if len(addressByte) < 4 {
		return false
	}
	payLoad := addressByte[:len(addressByte) - 4]
	checksum1 := addressByte[len(addressByte) - 4:]
	checksum2 := CheckSum(payLoad)
	fmt.Printf("checksum1:%x\n", checksum1)
	fmt.Printf("checksum2:%x\n", checksum2)
	// 2. 取数据
	// 3. 做checksum函数
	// 4. 比较
	return bytes.Equal(checksum1, checksum2)
}