package main

import (
	"flag"
	"fmt"
	"os"
)

const usage = `
	createChain --address ADDRESS	"create a blockchain"
	send --from FROM --to TO --amount AMOUNT	"send coin from FROM to TO"
	getBalance --address ADRESS		"get balance of the addressd"
	printChain						"print all blocks"
`

const PrintChainCmdString = "printChain"
const CreateChainCmdString = "createChain"
const GetBalanceCmdString = "getBalance"
const SendCmdString = "send"

type CLI struct {
	//bc *BlockChain
}

func (cli *CLI)printUsage()  {
	fmt.Println(usage)
	os.Exit(1)
}

func (cli *CLI)parameterCheck()  {
	if len(os.Args) < 2 {
		fmt.Println("invalid input!")
		cli.printUsage()
	}
}

func (cli *CLI)Run()  {
	cli.parameterCheck()
	createChainCmd := flag.NewFlagSet(CreateChainCmdString, flag.ExitOnError)
	getBalanceCmd := flag.NewFlagSet(GetBalanceCmdString, flag.ExitOnError)
	printChainCmd := flag.NewFlagSet(PrintChainCmdString, flag.ExitOnError)
	sendCmd := flag.NewFlagSet(SendCmdString, flag.ExitOnError)

	//创建区块链相关参数
	createChainCmdPara := createChainCmd.String("address", "", "address info!")
	//余额相关参数
	getBalanceCmdPara := getBalanceCmd.String("address", "", "address info!")
	//send相关参数
	fromPara := sendCmd.String("from", "", "sender address info!")
	toPara := sendCmd.String("to", "", "to address info!")
	amountPara := sendCmd.Float64("amount", 0, "amount info!")

	switch os.Args[1] {
	case CreateChainCmdString:
		err := createChainCmd.Parse(os.Args[2:])
		CheckErr("Run1()", err)
		if createChainCmd.Parsed() {
			if *createChainCmdPara == "" {
				cli.printUsage()
			}
			cli.CreateChain(*createChainCmdPara)
		}
	case SendCmdString:
		//发送交易
		err := sendCmd.Parse(os.Args[2:])
		CheckErr("Run4()", err)
		if sendCmd.Parsed() {
			if *fromPara == "" || *toPara == "" || *amountPara == 0 {
				fmt.Println("send cmd parameters invalid!!")
				cli.printUsage()
			}
			cli.Send(*fromPara, *toPara, *amountPara)
		}
	case GetBalanceCmdString:
		err := getBalanceCmd.Parse(os.Args[2:])
		CheckErr("Run4()", err)
		if getBalanceCmd.Parsed() {
			if *getBalanceCmdPara == "" {
				cli.printUsage()
			}
			cli.GetBalance(*getBalanceCmdPara)
		}
	case PrintChainCmdString:
		err := printChainCmd.Parse(os.Args[2:])
		CheckErr("Run3()", err)
		if printChainCmd.Parsed() {
			cli.PrintChain()
		}
	default:
		cli.printUsage()
	}
}