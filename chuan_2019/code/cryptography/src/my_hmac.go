package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
)

func main() {
	src := []byte("你好啊")
	key := []byte("helloworld")
	hmac1 := GenerateHmac(src, key)
	bl := VerifyHmac(src, key, hmac1)
	fmt.Println("校验结果： %t\n", bl)
}

// 生成消息认证码
func GenerateHmac(plainText, key []byte)[]byte {
	// 1. 创建哈希接口， 需要指定使用的哈希算法， 和秘钥
	myhash := hmac.New(sha1.New, key)
	// 2. 给哈希对象添加数据
	myhash.Write(plainText)
	// 3. 计算散列值
	hashText := myhash.Sum(nil)
	return hashText
}

// 验证消息认证码
func VerifyHmac(plainText, key, hashText []byte) bool  {
	// 1. 创建哈希接口， 需要指定使用的哈希算法， 和秘钥
	myhash := hmac.New(sha1.New, key)
	// 2. 给哈希对象添加数据
	myhash.Write(plainText)
	// 3. 计算散列值
	hmac1 := myhash.Sum(nil)
	// 4. 对比散列值
	return hmac.Equal(hmac1, hashText)
}
