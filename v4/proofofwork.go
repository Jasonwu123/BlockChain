package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

/* 1. 定义一个工作量证明结构ProofOfWork
*   a. block
*   b. 目标值
*  2. 提供创建POW函数
*  3. 提供计算不断计算hash的函数
*  4. 提供一个检验函数
*/


// 1. 定义一个工作量证明结构ProofOfWork
type ProofOfWork struct {
	block *Block
	target *big.Int
}

// 2. 提供创建POW函数
func NewProofOfWork(block *Block) *ProofOfWork {
	pow := ProofOfWork{
		block:  block,
	}

	// 指定的难度值，现在是一个string类型，需要进行转换
	targetStr := "01"

	// 引入辅助变量，目的是将上面的难度值转换成big.Int
	tmpInt := big.Int{}

	// 将难度值赋值给big.Int，指定16进制格式
	tmpInt.SetString(targetStr, 16)

	pow.target = &tmpInt
	return &pow
}

// 3. 提供计算不断计算hash的函数
func (pow *ProofOfWork) Run() ([]byte, uint64) {
	// 拼装数据
	// 做哈希运算
	// 与pow中的target进行比较
	    // a. 找到了，退出返回
	    // b. 没找到， 继续找，随机数+1
	var nonce uint64
	var hash [32]byte
	block := pow.block

	for {
		tmp := [][]byte {
			Uint64ToByte(block.Version),
			block.PrevHash,
			block.MerkelRoot,
			Uint64ToByte(block.TimeStamp),
			Uint64ToByte(block.Difficulty),
			Uint64ToByte(nonce),
			block.Data,
		}

		blockInfo := bytes.Join(tmp, []byte{})

		hash = sha256.Sum256(blockInfo)

		tmpInt := big.Int{}
		tmpInt.SetBytes(hash[:])

		if tmpInt.Cmp(pow.target) == -1 {
			fmt.Printf("挖矿成功！hash：%x, nonce: %d\n", hash, nonce)
			return hash[:], nonce
		} else {
			nonce++
		}
	}

}