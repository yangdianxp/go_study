package main

import (
	"crypto/x509"
	"encoding/pem"
	"os"
)

func main()  {
	
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
}

// RSA签名验证
func VerifyRSA()  {

}
