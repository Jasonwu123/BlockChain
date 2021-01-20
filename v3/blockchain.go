package main

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