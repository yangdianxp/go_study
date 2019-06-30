package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"os"
)

//生成rsa的密钥对， 并且保存到磁盘文件中
func GenerateRsaKey(keySize int)  {
	// 1. 使用rsa中的GenerateKey方法生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		panic(err)
	}
	// 2. 通过x509标准将得到的rsa私钥序列化为ASN.1的DER编码字符串
	derText := x509.MarshalPKCS1PrivateKey(privateKey)
	// 3. 要组织一个pem Block
	block := pem.Block{
		Type: "rsa private key",
		Bytes: derText,
	}
	// 4. pem编码
	file, err := os.Create("private.pem")
	if err != nil {
		panic(err)
	}
	pem.Encode(file, &block)
	file.Close()

	// =========公钥=========
	// 1. 从私钥中取出公钥
	publicKey := privateKey.PublicKey
	// 2. 使用x509标准格式化
	derstream, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}
	// 3. 将得到数据放到pem Block中
	block = pem.Block{
		Type: "rsa public key",
		Bytes: derstream,
	}
	// 4. pem编码
	file, err = os.Create("public.pem")
	if err != nil {
		panic(err)
	}
	pem.Encode(file, &block)
	file.Close()
}

// RSA 加密
func RSAEncrypt(plainText []byte, fileName string) []byte {
	// 1. 打开文件，并且读出文件内容
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, fileInfo.Size())
	file.Read(buf)
	file.Close()
	// 2. pem解码
	block, _ := pem.Decode(buf)
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	pubKey := pubInterface.(*rsa.PublicKey)
	// 3. 使用公钥加密
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, plainText)
	if err != nil {
		panic(err)
	}
	return cipherText
}

// rsa 解密
func RSADecrypt(cipherText []byte, fileName string) []byte {
	// 1. 打开文件，并且读出文件内容
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, fileInfo.Size())
	file.Read(buf)
	file.Close()
	// 2. pem解码
	block, _ := pem.Decode(buf)
	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	// 3. 使用私钥解密
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privKey, cipherText)
	if err != nil {
		panic(err)
	}
	return plainText
}

func main()  {
	GenerateRsaKey(1024)
	src := []byte("我是小崔，如果我死了，肯定不是自杀。。。")
	cipherText := RSAEncrypt(src, "public.pem")
	plainText := RSADecrypt(cipherText, "private.pem")
	fmt.Println(string(plainText))

	myHash()
}

func myHash()  {
	// sha256.Sum256([]byte("hello, go"))

	myHash := sha256.New()
	src := []byte("我是哈哈哈哈")
	myHash.Write(src)
	myHash.Write(src)
	myHash.Write(src)
	res := myHash.Sum(nil)
	myStr := hex.EncodeToString(res)
	fmt.Printf("%s\n", myStr)
}

/*
{
	GenerateRsaKey(1024)
}

 */

