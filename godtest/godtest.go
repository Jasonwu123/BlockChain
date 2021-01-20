package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type person struct {
	Name string
	age uint
}

func main() {
	// 定义一个结构Person
	var xiaoMing person
	xiaoMing.Name = "小明"
	xiaoMing.age = 20

	// 编码的数据放到buffer中
	var buffer bytes.Buffer

	// 使用gob进行序列化，得到字节流
	// 1. 定义一个编码器
	encode := gob.NewEncoder(&buffer)

	// 2. 使用编码器进行编码
	err := encode.Encode(&xiaoMing)
	if err != nil {
		log.Panic("编码出错，小明不知去向！！！")
	}
	fmt.Printf("编码后的小明：%v\n", buffer.Bytes())



	// 使用gob进行反序列化，得到Person结构
	// 1. 定义一个解码器
	decode := gob.NewDecoder(bytes.NewReader(buffer.Bytes()))

	// 2. 使用解码器解码
	var daMing person
	err = decode.Decode(&daMing)
	if err != nil {
		log.Panic("解码出错！！！")
	}

	fmt.Printf("解码后的小明：%v\n", daMing)

}
