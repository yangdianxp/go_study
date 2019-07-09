package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"
)

//1. 演示如何使用ecdsa生成公钥私钥
//2. 签名校验

func main()  {
	curve := elliptic.P256()
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	// 生成公钥
	pubKey := privateKey.PublicKey
	data := "hello, world"
	hash := sha256.Sum256([]byte(data))
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		log.Panic(err)
	}
	signature := append(r.Bytes(), s.Bytes()...)

	var r1, s1 big.Int
	r1.SetBytes(signature[:len(signature) / 2])
	s1.SetBytes(signature[len(signature) / 2:])

	// 校验需要三个东西：数据，签名，公钥
	res := ecdsa.Verify(&pubKey, hash[:], &r1, &s1)
	fmt.Printf("校验结果： %v\n", res)
}