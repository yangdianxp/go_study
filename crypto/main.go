package main

import "fmt"

func main()  {
	HashTest()
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

// 测试RSA的加密和解密测试
func RsaTest() {
	err := RsaGenKey(4086)
	fmt.Println("错误信息：", err)

	// 加密
	src := []byte("少壮不努力，老大徒伤悲。")
	data, err := RsaPublicEncrypt(src, "public.pem")
	// 解密
	data, err = RsaPrivateDecrypt(data, "private.pem")
	fmt.Println("解密之后的明文： " + string(data))
}

// 哈希算法测试
func HashTest()  {
	data := []byte("少壮不努力，老大徒伤悲。")
	fmt.Println("md5sum:", GetMd5Str_1(data))
	fmt.Println("md5sum:", GetMd5Str_2(data))
}