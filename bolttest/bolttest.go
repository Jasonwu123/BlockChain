package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

func main() {

	// 1. 打开数据库
	db, err := bolt.Open("test.db", 0600, nil)
	if err != nil {
		log.Panic("打开数据库失败！")
	}
	defer db.Close()
	// 2. 找到抽屉bucket
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("b1"))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte("b1"))
			if err != nil {
				log.Panic("创建bucket(b1)失败")
			}
		}

		// 3. 写数据
		bucket.Put([]byte("1111"), []byte("hello"))
		bucket.Put([]byte("2222"), []byte("world"))
		return nil
	})

	// 4. 读数据
	db.View(func(tx *bolt.Tx) error {
		// 找到抽屉，没有就直接报错退出
		bucket := tx.Bucket([]byte("b1"))
		if bucket == nil {
			log.Panic("bucket b1 不应该为空，请检查！！！")
		}

		// 直接读取数据
		v1 := bucket.Get([]byte("1111"))
		v2 := bucket.Get([]byte("2222"))
		fmt.Printf("v1 : %s\n", v1)
		fmt.Printf("v2 : %s\n", v2)

		return nil
	})
}
