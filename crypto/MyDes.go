package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
)

// 填充最后一个分组的函数
// src - 原始数据
// blockSize - 每个分组的数据长度
func paddingText(src []byte, blockSize int) []byte {
	// 1. 求出最后一个分组要填充多少个字节
	padding := blockSize - len(src) % blockSize
	// 2. 创建新的切片，切片的字节数为padding，并初始化，每个字节的值为padding
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	// 3. 将创建出的新切片和原始数据进行连接
	newText := append(src, padText...)
	// 4. 返回新的字符串
	return newText
}

// 删除末尾填充的字节
func unPaddingText(src []byte) []byte  {
	// 1. 求出要处理的切片的长度
	len := len(src)
	// 2. 取出最后一个字符，得到其整型值
	number := int(src[len-1])
	// 3. 将切片末尾的number个字节删除
	newText := src[:len - number]
	return newText
}

// 使用des进行对称加密
func encryptDES(src, key []byte)  []byte {
	// 1. 创建并返回一个D使用ES算法的cipher.Block接口
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// 2 对最后一个明文分组进行数据填充
	src = paddingText(src, block.BlockSize())
	// 3 创建一个密码分组为链接模式的，底层使用DES加密的BlockMode接口
	iv := []byte("aaaabbbb")
	blockMode := cipher.NewCBCEncrypter(block, iv)
	// 4 加密连接的数据块
	blockMode.CryptBlocks(src, src)
	return src
}

// 使用des解密
func decryptDES(src, key []byte) []byte {
	// 1. 创建并返回一个使用DES算法的cipher.Blcok接口
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// 2. 创建一个密码分组为链接模式的，底层使用DES解密的BlockMode接口
	iv := []byte("aaaabbbb")
	blockMode := cipher.NewCBCDecrypter(block, iv)
	// 3. 数据块解密
	blockMode.CryptBlocks(src, src)
	// 4. 去掉最后一组的填充数据
	return unPaddingText(src)
}

// 使用3des加密
func encrypt3DES(src, key []byte)  []byte {
	// 1. 创建并返回一个使用3DES算法的cipher.Block接口
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		panic(err)
	}
	// 2 对最后一个明文分组进行数据填充
	src = paddingText(src, block.BlockSize())
	// 3 创建一个密码分组为链接模式的，底层使用DES加密的BlockMode接口
	blockMode := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])
	// 4 加密连接的数据块
	blockMode.CryptBlocks(src, src)
	return src
}

//使用3des解密
func decrypt3DES(src, key []byte) []byte {
	// 1. 创建并返回一个使用3DES算法的cipher.Block接口
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		panic(err)
	}
	// 2. 创建一个密码分组为链接模式的，底层使用DES解密的BlockMode接口
	blockMode := cipher.NewCBCDecrypter(block, key[:block.BlockSize()])
	// 3. 数据块解密
	blockMode.CryptBlocks(src, src)
	// 4. 去掉最后一组的填充数据
	return unPaddingText(src)
}

// 使用aes进行对称加密
func encryptAES(src, key []byte)  []byte {
	// 1. 创建并返回一个D使用ES算法的cipher.Block接口
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// 2 对最后一个明文分组进行数据填充
	src = paddingText(src, block.BlockSize())
	// 3 创建一个密码分组为链接模式的，底层使用AES加密的BlockMode接口
	blockMode := cipher.NewCBCEncrypter(block, key)
	// 4 加密连接的数据块
	blockMode.CryptBlocks(src, src)
	return src
}

// 使用aes解密
func decryptAES(src, key []byte) []byte {
	// 1. 创建并返回一个使用AES算法的cipher.Blcok接口
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// 2. 创建一个密码分组为链接模式的，底层使用AES解密的BlockMode接口
	blockMode := cipher.NewCBCDecrypter(block, key)
	// 3. 数据块解密
	blockMode.CryptBlocks(src, src)
	// 4. 去掉最后一组的填充数据
	return unPaddingText(src)
}

