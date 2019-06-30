package main

import "fmt"

func main() {
	src := []byte("你好啊")
	key := []byte("helloworld")
	hmac1 := GenerateHmac(src, key)
	bl := VerifyHmac(src, key, hmac1)
	fmt.Println("校验结果： %t\n", bl)
}

/*
{
	fmt.Println("aes 加解密")
	key := []byte("12345678abcdefgh")
	src := []byte("你好啊")
	cipherText := aesEncrypt(src, key)
	plainText := aesDecrypt(cipherText, key)
	fmt.Printf("解密之后的数据:[%s]\n", string(plainText))
}

{
	fmt.Println("des 加解密")
	key := []byte("1234abcd")
	src := []byte("你好啊")
	cipherText := desEncrypt(src, key)
	plainText := desDecrypt(cipherText, key)
	fmt.Printf("解密之后的数据:[%s]\n", string(plainText))
}

 */