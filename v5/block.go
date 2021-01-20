package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"log"
	"time"
)

/*
*  补充区块字段
*  更新计算哈希函数
*  优化代码
*/

// 0. 定义结构
type Block struct {
	//  版本号
	Version uint64

	//  前区块哈希
	PrevHash []byte

	// Merkel根(梅克尔根，一个哈希值)
	MerkelRoot []byte

	// 时间戳
	TimeStamp uint64

	// 难度值
	Difficulty uint64

	// 随机数
	Nonce uint64

	// 当前区块哈希
	Hash []byte

	// 交易数据
	Data []byte
}

// 1. 创建区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		Version: 00,
		PrevHash: prevBlockHash,
		MerkelRoot: []byte{},
		TimeStamp: uint64(time.Now().Unix()),
		Difficulty: 0,
		Nonce: 0,
		Hash:     []byte{},
		Data:     []byte(data),
	}

	//block.SetHash()
	// 创建一个pow对象
	pow := NewProofOfWork(&block)

	// 查找随机数，不停的进行哈希运算
	hash, nonce := pow.Run()

	// 根据挖矿结果对区块数据进行更新
	block.Hash = hash
	block.Nonce = nonce

	return &block
}

// 实现一个辅助函数，功能是将uint64转成[]byte
func Uint64ToByte(num uint64) []byte {
	var buffer bytes.Buffer

	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buffer.Bytes()
}

// 序列化
func (block *Block) Serialize() []byte {
	var buffer bytes.Buffer
	encode := gob.NewEncoder(&buffer)
	err := encode.Encode(&block)
	if err != nil {
		log.Panic("编码出错！！！")
	}
	return buffer.Bytes()
}


// 反序列化
func Deserialize(data []byte) Block {
	var block Block
	decode := gob.NewDecoder(bytes.NewReader(data))
	err := decode.Decode(&block)
	if err != nil {
		log.Panic("解码出错！！！")
	}
	return block
}