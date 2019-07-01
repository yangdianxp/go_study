package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func main()  {
	src := []byte("哈哈哈哈哈")
	sigText := SignatureRSA(src, "private.pem")
	bl := VerifyRSA(src, sigText, "public.pem")
	fmt.Println(bl)
}

// RSA签名
func SignatureRSA(plainText []byte, fileName string)  []byte {
	// 1. 打开磁盘的私钥文件
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	// 2. 将私钥文件中的内容读出
	info, err := file.Stat()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, info.Size())
	file.Read(buf)
	file.Close()
	// 3. 使用pem对数据解码，得到了pem Block结构体变量
	block, _ := pem.Decode(buf)
	// 4. x509将数据解析成私钥结构体， -》 得到私钥
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	// 5. 创建一个哈希对象
	myhash := sha512.New()
	// 6. 给哈希对象添加数据
	myhash.Write(plainText)
	// 7. 计算哈希值
	hashText := myhash.Sum(nil)
	// 8. 使用rsa中的函数对散列值签名
	sigText, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA512, hashText)
	if err != nil {
		panic(err)
	}
	return sigText
}

// RSA签名验证
func VerifyRSA(plainText, sigText []byte, pubFileName string) bool  {
	// 1. 打开公钥文件， 将文件内容读出
	file, err := os.Open(pubFileName)
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
	// 2. 使用pem解码
	block, _ := pem.Decode(buf)
	// 3. x509将数据解析, 得到接口
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	// 4. 进行类型断言
	publickKey := pubInterface.(*rsa.PublicKey)
	// 5. 对原始消息进行哈希运算
	hashText := sha512.Sum512(plainText)
	// 6. 签名认证
	err = rsa.VerifyPKCS1v15(publickKey, crypto.SHA512, hashText[:], sigText)
	if err == nil {
		return true
	} else {
		return false
	}
}
