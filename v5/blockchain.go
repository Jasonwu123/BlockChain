package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

// 1. 引入区块链
type BlockChain struct {
	// 定义一个区块链数据库
	db *bolt.DB

	// 存储最后一个区块的哈希
	tail []byte
}

const (
	blockChainDB = "blockchain.db"
	blockBucket = "blockbucket"

)

// 2. 定义一个区块链
func NewBlockChain() *BlockChain {

	// 最后一个区块的哈希，从数据库中读出来的
	var lastHash []byte

	// 1. 打开数据库
	db, err := bolt.Open("blockchain.db", 0600, nil)
	if err != nil {
		log.Panic("打开数据库失败！")
	}

	defer db.Close()

	// 2. 找到抽屉bucket
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("blockbucket"))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte("blockbucket"))
			if err != nil {
				log.Panic("创建bucket(blockbucket)失败")
			}

			// 创建一个创世块，并作为第一个区块添加到区块链中
			genesisblock := GenesisBlock()

			// 3. 写数据
			// hash作为key，block的字节流作为value
			bucket.Put(genesisblock.Hash, genesisblock.Serialize())
			bucket.Put([]byte("lastHashKey"), genesisblock.Hash)
			lastHash = genesisblock.Hash

			// test
			blockBytes := bucket.Get(genesisblock.Hash)
			block := Deserialize(blockBytes)
			fmt.Printf("block info : %v\n", block)

		} else {
			lastHash = bucket.Get([]byte("lastHashKey"))
		}
		return nil
	})

	return &BlockChain{db, lastHash}
}

// 5. 定义一个创世快
func GenesisBlock() *Block {
	return NewBlock("Go创世块，老牛逼了！", []byte{})
}

// 6. 添加区块
func (bc *BlockChain) AddBlock(data string)  {
	/*
	// 获取前区块哈希值
	lastBlock := bc.blocks[len(bc.blocks)-1]
	prevHash := lastBlock.Hash

	// 创建新的区块
	block := NewBlock(data, prevHash)

	// 添加到区块链数组中
	bc.blocks = append(bc.blocks, block)

	 */
}