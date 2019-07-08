package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type Person struct {
	Name string
	Age uint
}

func main()  {
	// 定义一个结构Person
	var xiaoMing Person
	xiaoMing.Name = "xiaoming"
	xiaoMing.Age = 18
	var buffer bytes.Buffer
	// 使用gob进行序列化
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(&xiaoMing)
	if err != nil {
		log.Panic("编码错误")
	}
	//fmt.Printf("编码后的小明：%v \n", buffer.Bytes())
	// 使用gob进行反序列化
	var daMing Person
	decoder := gob.NewDecoder(&buffer)
	err = decoder.Decode(&daMing)
	if err != nil {
		log.Panic("解码出错")
	}
	fmt.Printf("%s, %d\n", daMing.Name, daMing.Age)
}