package main

import "bytes"

func paddingLastGroup(plainText []byte, blockSize int) []byte  {
	padNum := blockSize - len(plainText) % blockSize
	char := []byte{byte(padNum)}
	newPlain := bytes.Repeat(char, padNum)
	return append(plainText, newPlain...)
}
