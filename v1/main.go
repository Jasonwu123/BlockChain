package main

import (
	"crypto/sha256"
	"fmt"
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

// 3. 引入区块链
type BlockChain struct {
	// 定义一个区块链数组
	blocks []*Block
}

// 4. 定义一个区块链
func NewBlockChain() *BlockChain {
	// 创建一个创世块，并作为第一个区块添加到区块链中
	genesisblock := GenesisBlock()

	return &BlockChain{
		blocks: []*Block{
			genesisblock,
		},
	}
}

// 5. 定义一个创世快
func GenesisBlock() *Block {
	return NewBlock("Go创世块，老牛逼了！", []byte{})
}

// 6. 添加区块
func (bc *BlockChain) AddBlock(data string)  {
	// 获取前区块哈希值
	lastBlock := bc.blocks[len(bc.blocks)-1]
	prevHash := lastBlock.Hash

	// 创建新的区块
	block := NewBlock(data, prevHash)

	// 添加到区块链数组中
	bc.blocks = append(bc.blocks, block)
}
// 7. 重构代码

func main()  {
	bc := NewBlockChain()
	bc.AddBlock("班长向班花转了50枚比特币！")
	bc.AddBlock("班长又向班花转了50枚比特币！")

	for i, block := range bc.blocks {
		fmt.Printf("当前区块高度：%d\n", i)
		fmt.Printf("前区块哈希值：%x\n", block.PrevHash)
		fmt.Printf("当前区块哈希值：%x\n", block.Hash)
		fmt.Printf("区块交易数据：%s\n", block.Data)
	}
}