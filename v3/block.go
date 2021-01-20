package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
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

	block.SetHash()

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

// 2. 生成哈希
func (block *Block) SetHash()  {
	// var blockInfo []byte

	// 拼装数据
	/*
	blockInfo = append(blockInfo, Uint64ToByte(block.Version)...)
	blockInfo = append(blockInfo, block.PrevHash...)
	blockInfo = append(blockInfo, block.MerkelRoot...)
	blockInfo = append(blockInfo, Uint64ToByte(block.TimeStamp)...)
	blockInfo = append(blockInfo, Uint64ToByte(block.Difficulty)...)
	blockInfo = append(blockInfo, Uint64ToByte(block.Nonce)...)
	blockInfo = append(blockInfo, block.Hash...)
	blockInfo = append(blockInfo, block.Data...)
	*/

	tmp := [][]byte {
		Uint64ToByte(block.Version),
		block.PrevHash,
		block.MerkelRoot,
		Uint64ToByte(block.TimeStamp),
		Uint64ToByte(block.Difficulty),
		Uint64ToByte(block.Nonce),
		block.Hash,
		block.Data,
	}

	// 将二维的切片数组连接起来，返回一个一维的切片
	blockInfo := bytes.Join(tmp, []byte{})

	// sha256
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}

