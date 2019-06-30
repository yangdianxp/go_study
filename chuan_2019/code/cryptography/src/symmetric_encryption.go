package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
)

func paddingLastGroup(plainText []byte, blockSize int) []byte  {
	padNum := blockSize - len(plainText) % blockSize
	char := []byte{byte(padNum)}
	newPlain := bytes.Repeat(char, padNum)
	return append(plainText, newPlain...)
}

func unPaddingLastGroup(plainText []byte) []byte  {
	length := len(plainText)
	lastChar := plainText[length - 1]
	number := int(lastChar)
	return plainText[:length - number]
}

func desEncrypt(plainText, key[]byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	newText := paddingLastGroup(plainText, block.BlockSize())
	iv := []byte("12345678")
	blockMode := cipher.NewCBCEncrypter(block, iv)
	cipherText := make([]byte, len(newText))
	blockMode.CryptBlocks(cipherText, newText)
	return cipherText
}

func desDecrypt(cipherText, key []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	iv := []byte("12345678")
	blockMode := cipher.NewCBCDecrypter(block, iv)
	blockMode.CryptBlocks(cipherText, cipherText)
	plainText := unPaddingLastGroup(cipherText)
	return plainText
}

func aesEncrypt(plainText, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	iv := []byte("12345678abcdefgh")
	stream := cipher.NewCTR(block, iv)
	cipherText := make([]byte, len(plainText))
	stream.XORKeyStream(cipherText, plainText)
	return cipherText
}

func aesDecrypt(cipherText, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	iv := []byte("12345678abcdefgh")
	stream := cipher.NewCTR(block, iv)
	plainText := make([]byte, len(cipherText))
	stream.XORKeyStream(plainText, cipherText)
	return plainText
}