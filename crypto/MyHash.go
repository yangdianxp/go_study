package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

// 使用MD5对数据进行哈希运算
func GetMd5Str_1(src []byte) string  {
	// 1. 给哈希算法添加数据
	res := md5.Sum(src)
	// myres := fmt.Sprintf("%x", res)
	myres := hex.EncodeToString(res[:])
	return myres
}

func GetMd5Str_2(src []byte) string {
	// 1. 创建哈希接口
	myHash := md5.New()
	// 2. 添加数据
	// 添加数据的第一种方式
	io.WriteString(myHash, string(src))
	// 添加数据的第二种方式
	//myHash.Write(src)
	// 3. 计算结果
	res := myHash.Sum(nil)
	// 4. 散列值格式化
	myres := hex.EncodeToString(res[:])
	return myres
}