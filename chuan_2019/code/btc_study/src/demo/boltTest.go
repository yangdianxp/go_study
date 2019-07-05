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
}