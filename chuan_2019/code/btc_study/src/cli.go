package main

import (
	"fmt"
	"os"
	"strconv"
)

//接收命令行参数，控制区块链操作的文件
type CLI struct {
	bc *BlockChain
}

const Usage = `
	printChain 				 "print all blockchain data"
	getBalance --address ADDRESS "获取指定地址的余额"
	send FROM TO AMOUNT MINER DATA "由FROM转AMOUNT给TO， 由MINER探矿，同时写入DATA"
	newWallet 		"创建一个新的钱包（私钥公钥对）"
	listAddresses  "列举所有的钱包地址"
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
	case "send":
		if len(args) != 7 {
			fmt.Printf("参数个数错误\n")
			fmt.Printf(Usage)
			return
		}
		fmt.Printf("转账开始。。。\n")
		from := args[2]
		to := args[3]
		amount, _ := strconv.ParseFloat(args[4], 64)
		miner := args[5]
		data := args[6]
		cli.Send(from, to, amount, miner, data)
	case "newWallet":
		fmt.Printf("创建新的钱包...\n")
		cli.NewWallet()
	case "listAddresses":
		fmt.Printf("列举所有的钱包地址\n")
		cli.List()
	default:
		fmt.Printf(Usage)
	}
}
