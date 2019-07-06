package main

import (
	"bolt"
	"fmt"
	"log"
)

func main()  {
	fmt.Println("hello")

	// 1. 打开数据库
	db, err := bolt.Open("test.db", 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		// 2. 找到抽屉bucket
		bucket := tx.Bucket([]byte("b1"))
		if bucket == nil{
			//没有抽屉，我们需要创建
			bucket, err = tx.CreateBucket([]byte("b1"))
			if err != nil {
				log.Panic(err)
			}
		}
		// 3. 写数据
		bucket.Put([]byte("11111"), []byte("hello"))
		bucket.Put([]byte("22222"), []byte("world"))
		return nil
	})
	// 4. 读数据
	db.View(func(tx *bolt.Tx) error {
		// 1. 找到抽屉
		bucket := tx.Bucket([]byte("b1"))
		if bucket == nil {
			log.Panic("bucket b1 不应该为空，请检查！！！")
		}
		// 2. 直接读取数据
		v1 := bucket.Get([]byte("11111"))
		v2 := bucket.Get([]byte("22222"))
		fmt.Printf("v1:%s v2:%s", v1, v2)
		return nil
	})
}