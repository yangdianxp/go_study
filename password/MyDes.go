package main

import "bytes"

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
