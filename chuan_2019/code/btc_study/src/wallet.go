package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
)

// 这里的钱包是一个结构，每一个钱包保存了公钥，私钥对

type Wallet struct {
	Private *ecdsa.PrivateKey
	// 约定：这里的PubKey不存储原始的公钥，而是存储X和Y拼接的字符串， 在校验端重新拆分
	PubKey []byte
}

// 创建钱包
func NewWallet() *Wallet {
	curve := elliptic.P256()
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	// 生成公钥
	pubKeyOrig := privateKey.PublicKey
	pubKey := append(pubKeyOrig.X.Bytes(), pubKeyOrig.Y.Bytes()...)
	return &Wallet{Private: privateKey, PubKey: pubKey}
}

// 生成地址
