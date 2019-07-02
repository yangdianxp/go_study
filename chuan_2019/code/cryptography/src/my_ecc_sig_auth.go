package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
)

func main()  {
	src := []byte("你你你你你")
	rText, sText := EccSignature(src, "eccPrivate.pem")
	bl := EccVerify(src, rText, sText, "eccPublic.pem")
	fmt.Println(bl)
}

// 秘钥对 的生成，并保存到磁盘
func GenerateEccKey()  {
	// 1. 使用ecdsa生成密钥对
	privateKey, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	if err != nil {
		panic(err)
	}
	// 2. 将私钥写入磁盘
	// 使用x509进行序列化
	derText, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		panic(err)
	}
	// 将得到的切片字符串放入pem.Block结构体中
	block := pem.Block{
		Type: "ecdsa private key",
		Bytes: derText,
	}
	// 使用pem编码
	file, err := os.Create("eccPrivate.pem")
	if err != nil {
		panic(err)
	}
	pem.Encode(file, &block)
	file.Close()
	// 3. 将公钥写入磁盘
	// 从私钥中得到公钥
	publicKey := privateKey.PublicKey
	// 使用x509进行序列化
	derText, err = x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}
	// 将得到的切片字符串放入pem.Block结构体中
	block = pem.Block{
		Type: "ecdsa public key",
		Bytes: derText,
	}
	file, err = os.Create("eccPublic.pem")
	pem.Encode(file, &block)
	file.Close()
}



// 使用私钥进行数字签名
func EccSignature(plainText []byte,privName string) (rText, sText []byte) {
	// 1. 打开私钥文件，将内容读出来
	file, err := os.Open(privName)
	if err != nil {
		panic(err)
	}
	info, err := file.Stat()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, info.Size())
	file.Read(buf)
	file.Close()
	// 2. 使用pem进行数据解码   pem.Decode()
	block, _ := pem.Decode(buf)
	// 3. 使用x509，对私钥进行反序列化
	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	// 4. 对原始数据进行哈希运算->散列值
	hashText := sha1.Sum(plainText)
	// 5. 进行数字签名
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashText[:])
	// 签名
	if err != nil {
		panic(err)
	}
	// 得到的r和s不能直接使用， 因为这是指针
	// 将两块内存中的数据进行序列化
	rText, err = r.MarshalText()
	if err != nil {
		panic(err)
	}
	sText, err = s.MarshalText()
	if err != nil {
		panic(err)
	}
	return rText, sText
}



// 使用公钥验证数字签名
func EccVerify(plainText, rText, sText []byte, pubFile string) bool {
	// 1. 打开公钥文件， 将里边的内容读出
	file, err := os.Open(pubFile)
	if err != nil {
		panic(err)
	}
	info, err := file.Stat()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, info.Size())
	file.Read(buf)
	file.Close()
	// 2. pem解码 -> pem.Decode()
	block, _ := pem.Decode(buf)
	// 3. 使用x509对公钥还原
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	// 4. 将接口 -> 公钥
	publicKey := pubInterface.(*ecdsa.PublicKey)
	// 5. 对原始数据进行哈希运算->得到散列值
	hashText := sha1.Sum(plainText)
	// 6. 签名的认证 -> ecdsa
	var r, s big.Int
	err = r.UnmarshalText(rText)
	if err != nil {
		panic(err)
	}
	err = s.UnmarshalText(sText)
	if err != nil {
		panic(err)
	}
	return ecdsa.Verify(publicKey, hashText[:], &r, &s)
}

