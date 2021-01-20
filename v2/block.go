package main

import (
	"crypto/sha256"
)

// 0. 定义结构
type Block struct {
	//  前区块哈希
	PrevHash []byte

	// 当前区块哈希
	Hash []byte

	// 交易数据
	Data []byte
}

// 1. 创建区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		PrevHash: prevBlockHash,
		Hash:     []byte{},
		Data:     []byte(data),
	}

	block.SetHash()

	return &block
}


// 2. 生成哈希
func (block *Block) SetHash()  {
	// 拼装数据
	blockInfo := append(block.PrevHash, block.Data...)
	// sha256
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}