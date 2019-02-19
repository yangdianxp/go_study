package main

import "fmt"

func main()  {
	aesTest()
}

// 测试des的加解密
func desTest()  {
	fmt.Println("===== des 加解密 =====")
	src := []byte("少壮不努力，老大徒伤悲。")
	key := []byte("12345678")
	str := encryptDES(src, key)
	str = decryptDES(str, key)
	fmt.Println("解密之后的明文： " + string(str))
}

// 测试3des的加解密
func tripledesTest()  {
	fmt.Println("===== 3des 加解密 =====")
	src := []byte("少壮不努力，老大徒伤悲。")
	key := []byte("123456783232323287654321")
	str := encrypt3DES(src, key)
	fmt.Println("加密之后的密文： " + string(str))
	str = decrypt3DES(str, key)
	fmt.Println("解密之后的明文： " + string(str))
}

// 测试aes的加解密
func aesTest()  {
	fmt.Println("===== aes 加解密 =====")
	src := []byte("少壮不努力，老大徒伤悲。")
	key := []byte("1234567887654321")
	str := encryptAES(src, key)
	fmt.Println("加密之后的密文： " + string(str))
	str = decryptAES(str, key)
	fmt.Println("解密之后的明文： " + string(str))
}