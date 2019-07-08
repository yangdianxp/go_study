package main

import (
	"fmt"
	"os"
)

//接收命令行参数，控制区块链操作的文件
type CLI struct {
	bc *BlockChain
}

const Usage = `
	addBlock --data DATA     "add data to blockchain"
	printChain 				 "print all blockchain data"
	getBalance --address ADDRESS "获取指定地址的余额"
`

// 接收参数的动作， 我们放到一个函数中
func (cli *CLI) Run()  {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf(Usage)
		return
	}
	// 分析命令
	cmd := args[1]
	switch cmd {
	case "addBlock":
		fmt.Println("添加区块")
		if len(args) == 4 && args[2] == "--data" {
			data := args[3]
			cli.AddBlock(data)
		} else {
			fmt.Printf(Usage)
			return
		}
	case "printChain":
		fmt.Println("打印区块")
		cli.PrintBlockChain()
	case "getBalance":
		fmt.Printf("获取余额\n")
		if len(args) == 4 && args[2] == "--address" {
			addr := args[3]
			cli.GetBalance(addr)
		} else {
			fmt.Printf(Usage)
			return
		}
	default:
		fmt.Printf(Usage)
	}
}
