package main

import (
	"bytes"
	"encoding/gob"
	"log"
)

type Person struct {
	Name string
	age uint
}

func main()  {
	// 定义一个结构Person
	var xiaoMing Person
	xiaoMing.Name = "xiaoming"
	xiaoMing.age = 18
	var buffer bytes.Buffer
	// 使用gob进行序列化
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(&xiaoMing)
	if err != nil {
		log.Panic("编码错误")
	}

	// 使用gob进行反序列化
}